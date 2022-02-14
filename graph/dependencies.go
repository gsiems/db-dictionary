package graph

import (
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"

	m "github.com/gsiems/db-dictionary/model"
)

var typeColors = map[string]string{
	"TABLE":             "#FFFFE0", // LightYellow
	"BASE TABLE":        "#FFFFE0", // LightYellow
	"MATERIALIZED VIEW": "#FFD700", // Gold
	"VIEW":              "#DDA0DD", // Plum
	"FUNCTION":          "#7FFFD4", // Aquamarine
	"PACKAGE":           "#E0FFFF", // LightCyan
	"PROCEDURE":         "#87CEFA", // LightSkyBlue
	"SEQUENCE":          "#FFFF00", // Yellow
	"DEFAULT":           "#F5F5F5", // Grey
}

type DependencyNode struct {
	ID           int
	ObjectSchema string
	ObjectName   string
	ObjectType   string
	Color        string
}

type DependencyEdge struct {
	Node1 DependencyNode
	Node2 DependencyNode
	// direction? or can this be inferred?
}

type SMap map[string]map[string]DependencyNode
type EMap map[int]map[int]DependencyEdge

type DependencyGraph struct {
	id            int
	graphviz      string
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	OutputDir     string
	SchemaName    string
	SchemaComment string
	SchemaNodes   SMap
	OtherNodes    SMap
	Edges         EMap
}

func MakeDepenencyGraphs(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		g := NewDependencyGraph(&vs, md)

		dependencies := md.FindDependencies(vs.Name, "")
		if len(dependencies) > 0 {
			md.SortDependencies(dependencies)
			for _, v := range dependencies {
				g.AddDependency(&v)
			}
		}

		dependents := md.FindDependents(vs.Name, "")
		if len(dependents) > 0 {
			md.SortDependencies(dependents)
			for _, v := range dependents {
				g.AddDependent(&v)
			}
		}

		// generate the graph
		err = g.RenderDotGraph()
		if err != nil {
			return err
		}

	}
	return err
}

func NewDependencyGraph(vs *m.Schema, md *m.MetaData) *DependencyGraph {

	g := DependencyGraph{
		Title: "Dependencies for " + md.Alias + "." + vs.Name,
		//TmspGenerated: md.TmspGenerated,
		DBMSVersion:   md.Version,
		DBName:        md.Name,
		DBComment:     md.Comment,
		OutputDir:     md.OutputDir,
		SchemaName:    vs.Name,
		SchemaComment: vs.Comment,
	}

	if !md.Cfg.NoGraphviz {
		g.graphviz = md.Cfg.GraphvizCmd
	}

	// TODO Add the legend and  title block

	g.id = 0

	g.SchemaNodes = make(SMap)
	g.OtherNodes = make(SMap)
	g.Edges = make(EMap)

	return &g
}

func (g *DependencyGraph) RenderDotGraph() (err error) {

	if len(g.SchemaNodes) == 0 && len(g.OtherNodes) == 0 {
		return err
	}

	ft := `
digraph {
    layout="fdp";
    overlap="false";
    ranksep=4;
    clusterrank=local
    fontname="Helvetica"
    fontnames="Helvetica,sans-Serif"
    stylesheet="../css/svg.css"
    node [style="rounded,filled"; fontname="Helvetica"; fontnames="Helvetica,sans-Serif"]
    {{ range $sn, $sv := .SchemaNodes }}{{ range $nn, $nv := $sv }}"{{.ID}}" [label="{{.ObjectName}}"; fillcolor="{{.Color}}"; shape="rect"]
    {{ end }}{{ end }}{{ range $sn, $sv := .OtherNodes }}
    subgraph cluster_{{$sn}}{
        label="{{$sn}}"
        bgcolor="#FCFCFC"{{ range $nn, $nv := $sv }}
        "{{.ID}}" [label="{{.ObjectName}}"; color="{{.Color}}"; shape="rect"]{{ end }}
    }{{ end }}
    {{ range $i, $ix := .Edges }}{{ range $j, $jx := $ix }}"{{ $i }}" -> "{{ $j }}"
    {{ end }}{{ end }}
}
`
	// parse the template
	templates, err := template.New("doc").Parse(ft)
	if err != nil {
		return err
	}

	// ensure that the file directory exists
	dirName := path.Join(g.OutputDir, g.SchemaName)
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0745)
		if err != nil {
			return err
		}
	}

	// TODO: graphviz can be very slow so do we want/need to first check to
	// see if there was a previous dependencies.gv file and if so, has the
	// contents of the file changed-- no change then no need to re-run graphviz

	// create the file
	outFileName := path.Join(dirName, "dependencies.gv")
	outfile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	// render and write the file
	err = templates.Lookup("doc").Execute(outfile, g)
	if err != nil {
		return err
	}

	if g.graphviz != "" {
		// attempt to run graphviz
		svgFileName := path.Join(dirName, "dependencies.svg")
		cmd := exec.Command(g.graphviz, "-Tsvg", "-o", svgFileName, outFileName)
		cerr := cmd.Run()
		if cerr != nil {
			log.Printf("could not run Graphviz (%s): %s", g.graphviz, cerr)
		}
	}

	return err
}

func (g *DependencyGraph) AddDependency(d *m.Dependency) {

	var n1 DependencyNode
	var n2 DependencyNode
	var ok bool

	if d.ObjectSchema == g.SchemaName {
		n1, ok = g.SchemaNodes[d.ObjectSchema][d.ObjectName]
	} else {
		n1, ok = g.OtherNodes[d.ObjectSchema][d.ObjectName]
	}

	if !ok {
		g.id++
		n1 = DependencyNode{
			ID:           g.id,
			ObjectSchema: d.ObjectSchema,
			ObjectName:   d.ObjectName,
			ObjectType:   d.ObjectType,
			Color:        typeColor(d.ObjectType),
		}
		g.AddNode(n1)
	}

	if d.DepObjectSchema == g.SchemaName {
		n2, ok = g.SchemaNodes[d.DepObjectSchema][d.DepObjectName]
	} else {
		n2, ok = g.OtherNodes[d.DepObjectSchema][d.DepObjectName]
	}

	if !ok {
		g.id++
		n2 = DependencyNode{
			ID:           g.id,
			ObjectSchema: d.DepObjectSchema,
			ObjectName:   d.DepObjectName,
			ObjectType:   d.DepObjectType,
			Color:        typeColor(d.DepObjectType),
		}
		g.AddNode(n2)
	}
	g.AddEdge(n1, n2)
}

func (g *DependencyGraph) AddDependent(d *m.Dependency) {

	var n1 DependencyNode
	var n2 DependencyNode
	var ok bool

	if d.ObjectSchema == g.SchemaName {
		n1, ok = g.SchemaNodes[d.ObjectSchema][d.ObjectName]
	} else {
		n1, ok = g.OtherNodes[d.ObjectSchema][d.ObjectName]
	}

	if !ok {
		g.id++
		n1 = DependencyNode{
			ID:           g.id,
			ObjectSchema: d.ObjectSchema,
			ObjectName:   d.ObjectName,
			ObjectType:   d.ObjectType,
			Color:        typeColor(d.ObjectType),
		}
		g.AddNode(n1)
	}

	if d.DepObjectSchema == g.SchemaName {
		n2, ok = g.SchemaNodes[d.DepObjectSchema][d.DepObjectName]
	} else {
		n2, ok = g.OtherNodes[d.DepObjectSchema][d.DepObjectName]
	}

	if !ok {
		g.id++
		n2 = DependencyNode{
			ID:           g.id,
			ObjectSchema: d.DepObjectSchema,
			ObjectName:   d.DepObjectName,
			ObjectType:   d.DepObjectType,
			Color:        typeColor(d.DepObjectType),
		}
		g.AddNode(n2)
	}
	g.AddEdge(n1, n2)
}

func (g *DependencyGraph) AddNode(n DependencyNode) {

	if n.ObjectSchema == g.SchemaName {
		if g.SchemaNodes[n.ObjectSchema] == nil {
			g.SchemaNodes[n.ObjectSchema] = make(map[string]DependencyNode)
		}
		g.SchemaNodes[n.ObjectSchema][n.ObjectName] = n
	} else {
		if g.OtherNodes[n.ObjectSchema] == nil {
			g.OtherNodes[n.ObjectSchema] = make(map[string]DependencyNode)
		}
		g.OtherNodes[n.ObjectSchema][n.ObjectName] = n
	}
}

func (g *DependencyGraph) AddEdge(n1, n2 DependencyNode) {

	if g.Edges[n1.ID] == nil {
		g.Edges[n1.ID] = make(map[int]DependencyEdge)
	}

	_, ok := g.Edges[n1.ID][n2.ID]
	if !ok {
		g.Edges[n1.ID][n2.ID] = DependencyEdge{
			Node1: n1,
			Node2: n2,
		}
	}
}

func typeColor(s string) string {
	color, ok := typeColors[s]
	if ok {
		return color
	}
	return "#F5F5F5"
}
