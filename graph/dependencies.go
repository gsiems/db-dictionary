package graph

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"text/template"
	"time"

	m "github.com/gsiems/db-dictionary/model"
)

const (
	colLabelH        = 15.0
	defaultTextWidth = 191.94140817732347
	hTextPadding     = 16.0
	minNodeHeight    = 30.0
	minNodeWidth     = 224.47070408866173
	nodeLabelH       = 29.1264648438
	vSpacing         = 20.0
	// Font Names
	Helvetica = "Helvetica"
	Dialog    = "Dialog"
	// Font Styles
	Bold   = "bold"
	Normal = "normal"
	// Shapes
	Ellipse        = "ellipse"
	Parallelogram  = "parallelogram"
	Rectangle      = "rectangle"
	RoundRectangle = "roundrectangle"
	// Colours
	Aquamarine    = "#7FFFD4"
	Black         = "#000000"
	Cornsilk      = "#FFF8DC"
	GoldenRod     = "#DAA520"
	Grey          = "#808080"
	LightGreen    = "#90EE90"
	LightGrey     = "#D3D3D3"
	LightSkyBlue  = "#87CEFA"
	NavajoWhite   = "#FFDEAD"
	Orchid        = "#DA70D6"
	PaleTurquoise = "#AFEEEE"
	Tan           = "#D2B48C"
	Thistle       = "#D8BFD8"
	Turquoise     = "#40E0D0"
	PaleGreen     = "#98FB98"
	DarkSlateGrey = "#2F4F4F"

	// AliceBlue            = "#F0F8FF"
	// AntiqueWhite         = "#FAEBD7"
	// Aqua                 = "#00FFFF"
	// Azure                = "#F0FFFF"
	// Beige                = "#F5F5DC"
	// Bisque               = "#FFE4C4"
	// BlanchedAlmond       = "#FFEBCD"
	// Blue                 = "#0000FF"
	// BlueViolet           = "#8A2BE2"
	// Brown                = "#A52A2A"
	// BurlyWood            = "#DEB887"
	// CadetBlue            = "#5F9EA0"
	// Chartreuse           = "#7FFF00"
	// Chocolate            = "#D2691E"
	// Coral                = "#FF7F50"
	// CornflowerBlue       = "#6495ED"
	// Crimson              = "#DC143C"
	// Cyan                 = "#00FFFF"
	// DarkBlue             = "#00008B"
	// DarkCyan             = "#008B8B"
	// DarkGoldenRod        = "#B8860B"
	// DarkGray             = "#A9A9A9"
	// DarkGrey             = "#A9A9A9"
	// DarkGreen            = "#006400"
	// DarkKhaki            = "#BDB76B"
	// DarkMagenta          = "#8B008B"
	// DarkOliveGreen       = "#556B2F"
	// DarkOrange           = "#FF8C00"
	// DarkOrchid           = "#9932CC"
	// DarkRed              = "#8B0000"
	// DarkSalmon           = "#E9967A"
	// DarkSeaGreen         = "#8FBC8F"
	// DarkSlateBlue        = "#483D8B"
	// DarkSlateGray        = "#2F4F4F"
	// DarkTurquoise        = "#00CED1"
	// DarkViolet           = "#9400D3"
	// DeepPink             = "#FF1493"
	// DeepSkyBlue          = "#00BFFF"
	// DimGray              = "#696969"
	// DimGrey              = "#696969"
	// DodgerBlue           = "#1E90FF"
	// FireBrick            = "#B22222"
	// FloralWhite          = "#FFFAF0"
	// ForestGreen          = "#228B22"
	// Fuchsia              = "#FF00FF"
	// Gainsboro            = "#DCDCDC"
	// GhostWhite           = "#F8F8FF"
	// Gold                 = "#FFD700"
	// Gray                 = "#808080"
	// Green                = "#008000"
	// GreenYellow          = "#ADFF2F"
	// HoneyDew             = "#F0FFF0"
	// HotPink              = "#FF69B4"
	// IndianRed            = "#CD5C5C"
	// Indigo               = "#4B0082"
	// Ivory                = "#FFFFF0"
	// Khaki                = "#F0E68C"
	// Lavender             = "#E6E6FA"
	// LavenderBlush        = "#FFF0F5"
	// LawnGreen            = "#7CFC00"
	// LemonChiffon         = "#FFFACD"
	// LightBlue            = "#ADD8E6"
	// LightCoral           = "#F08080"
	// LightCyan            = "#E0FFFF"
	// LightGoldenRodYellow = "#FAFAD2"
	// LightGray            = "#D3D3D3"
	// LightPink            = "#FFB6C1"
	// LightSalmon          = "#FFA07A"
	// LightSeaGreen        = "#20B2AA"
	// LightSlateGray       = "#778899"
	// LightSlateGrey       = "#778899"
	// LightSteelBlue       = "#B0C4DE"
	// LightYellow          = "#FFFFE0"
	// Lime                 = "#00FF00"
	// LimeGreen            = "#32CD32"
	// Linen                = "#FAF0E6"
	// Magenta              = "#FF00FF"
	// Maroon               = "#800000"
	// MediumAquaMarine     = "#66CDAA"
	// MediumBlue           = "#0000CD"
	// MediumOrchid         = "#BA55D3"
	// MediumPurple         = "#9370DB"
	// MediumSeaGreen       = "#3CB371"
	// MediumSlateBlue      = "#7B68EE"
	// MediumSpringGreen    = "#00FA9A"
	// MediumTurquoise      = "#48D1CC"
	// MediumVioletRed      = "#C71585"
	// MidnightBlue         = "#191970"
	// MintCream            = "#F5FFFA"
	// MistyRose            = "#FFE4E1"
	// Moccasin             = "#FFE4B5"
	// Navy                 = "#000080"
	// OldLace              = "#FDF5E6"
	// Olive                = "#808000"
	// OliveDrab            = "#6B8E23"
	// Orange               = "#FFA500"
	// OrangeRed            = "#FF4500"
	// PaleGoldenRod        = "#EEE8AA"
	// PaleVioletRed        = "#DB7093"
	// PapayaWhip           = "#FFEFD5"
	// PeachPuff            = "#FFDAB9"
	// Peru                 = "#CD853F"
	// Pink                 = "#FFC0CB"
	// Plum                 = "#DDA0DD"
	// PowderBlue           = "#B0E0E6"
	// Purple               = "#800080"
	// RebeccaPurple        = "#663399"
	// Red                  = "#FF0000"
	// RosyBrown            = "#BC8F8F"
	// RoyalBlue            = "#4169E1"
	// SaddleBrown          = "#8B4513"
	// Salmon               = "#FA8072"
	// SandyBrown           = "#F4A460"
	// SeaGreen             = "#2E8B57"
	// SeaShell             = "#FFF5EE"
	// Sienna               = "#A0522D"
	// Silver               = "#C0C0C0"
	// SkyBlue              = "#87CEEB"
	// SlateBlue            = "#6A5ACD"
	// SlateGray            = "#708090"
	// SlateGrey            = "#708090"
	// Snow                 = "#FFFAFA"
	// SpringGreen          = "#00FF7F"
	// SteelBlue            = "#4682B4"
	// Teal                 = "#008080"
	// Tomato               = "#FF6347"
	// Violet               = "#EE82EE"
	// Wheat                = "#F5DEB3"
	// White                = "#FFFFFF"
	// WhiteSmoke           = "#F5F5F5"
	// Yellow               = "#FFFF00"
	// YellowGreen          = "#9ACD32"
)

var nodeTypes = []string{"TABLE", "BASE TABLE", "FOREIGN TABLE", "MATERIALIZED VIEW", "VIEW", "FUNCTION", "PACKAGE", "PROCEDURE", "SEQUENCE"}

var nodeColors = map[string]string{
	"TABLE":             Cornsilk,
	"BASE TABLE":        NavajoWhite,
	"FOREIGN TABLE":     Tan,
	"MATERIALIZED VIEW": Orchid,
	"VIEW":              Thistle,
	"FUNCTION":          Aquamarine,
	"PACKAGE":           PaleTurquoise,
	"PROCEDURE":         Turquoise,
	"SEQUENCE":          LightGreen,
}

var nodeShapes = map[string]string{
	"TABLE":             Rectangle,
	"BASE TABLE":        Rectangle,
	"FOREIGN TABLE":     Rectangle,
	"MATERIALIZED VIEW": Rectangle,
	"VIEW":              Rectangle,
	"FUNCTION":          Parallelogram,
	"PACKAGE":           Parallelogram,
	"PROCEDURE":         Parallelogram,
	"SEQUENCE":          Ellipse,
}

type nodeColumn struct {
	Name            string
	DataType        string
	IsPK            bool
	IsNullable      bool
	OrdinalPosition int32
	Y               float64
}

type graphNode struct {
	ID           int
	ObjectSchema string
	ObjectName   string
	ObjectType   string
	Color        string  // to simplify the dot file creation
	H            float64 // gml
	W            float64 // gml
	X            float64 // gml
	Y            float64 // gml
	dtOffset     float64 // gml
	Columns      []nodeColumn
}

type graphSchema struct {
	ID         int
	SchemaName string
	Color      string
}

type graphEdge struct {
	Node1 graphNode
	Node2 graphNode
}

type nodeMap map[string]map[string]graphNode // [schema name][object name]
type edgeMap map[int]map[int]graphEdge       // [dependent node id][depends on node id]

type dependencyGraph struct {
	id            int
	graphviz      string
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	OutputDir     string
	SchemaName    string
	SchemaComment string
	gmlTitleBlock string
	gmlLegend     string
	nodeTypes     []string
	OtherSchemas  []graphSchema
	SchemaNodes   nodeMap
	OtherNodes    nodeMap
	Edges         edgeMap
}

func (g *dependencyGraph) AddDependency(d *m.Dependency, md *m.MetaData) {

	var n1 graphNode
	var n2 graphNode
	var ok bool
	var addCols bool

	if d.ObjectSchema == g.SchemaName {
		n1, ok = g.SchemaNodes[d.ObjectSchema][d.ObjectName]
		addCols = true
	} else {
		n1, ok = g.OtherNodes[d.ObjectSchema][d.ObjectName]
		addCols = false
	}

	if !ok {
		g.id++
		n1 = mkNode(g.id, d.ObjectSchema, d.ObjectName, d.ObjectType, addCols, md)
		g.AddNode(n1)
	}

	if d.DepObjectSchema == g.SchemaName {
		n2, ok = g.SchemaNodes[d.DepObjectSchema][d.DepObjectName]
		addCols = true
	} else {
		addCols = false
		n2, ok = g.OtherNodes[d.DepObjectSchema][d.DepObjectName]
	}

	if !ok {
		g.id++
		n2 = mkNode(g.id, d.DepObjectSchema, d.DepObjectName, d.DepObjectType, addCols, md)
		g.AddNode(n2)
	}
	g.AddEdge(n1, n2)
}

func (g *dependencyGraph) AddEdge(n1, n2 graphNode) {

	if g.Edges[n1.ID] == nil {
		g.Edges[n1.ID] = make(map[int]graphEdge)
	}

	_, ok := g.Edges[n1.ID][n2.ID]
	if !ok {
		g.Edges[n1.ID][n2.ID] = graphEdge{
			Node1: n1,
			Node2: n2,
		}
	}
}

func (g *dependencyGraph) AddGMLLegend() {

	x0 := 0.0
	y0 := 300.0
	width := 250.0
	x := x0 + width/2.0
	y := y0

	var legend []string

	g.id++

	gid := g.id

	itemFmt := `	node
	[
		id	%d
		label	"%s"
		graphics
		[
			x	%f
			y	%f
			w	%f
			h	%f
			type	"%s"
			fill	"%s"
			outline	"%s"
		]
		LabelGraphics
		[
			text	"%s"
			fontSize	13
			fontName	"Dialog"
			anchor	"c"
		]
		gid	%d
	]`

	for _, nodeType := range g.nodeTypes {
		g.id++
		y = y0 + float64(g.id-gid)*(minNodeHeight+vSpacing)
		fillColor := nodeColor(nodeType)
		shape := nodeShape(nodeType)

		legend = append(legend, fmt.Sprintf(itemFmt, g.id, nodeType, x, y, defaultTextWidth, minNodeHeight, shape, fillColor, Black, nodeType, gid))
	}

	yg := (y + y0) / 2
	h := float64(g.id-gid)*(minNodeHeight+vSpacing) + vSpacing

	lFmt := `	node
	[
		id	%d
		label	"Legend"
		graphics
		[
			x	%f
			y	%f
			w	%f
			h	%f
			type	"roundrectangle"
			fill	"%s"
			outline	"%s"
			outlineStyle	"dotted"
		]
		LabelGraphics
		[
			text	"Legend"
			fill	"%s"
			fontSize	15
			fontName	"Dialog"
			autoSizePolicy	"node_width"
			anchor	"t"
		]
		isGroup	1
	]
%s`

	g.gmlLegend = fmt.Sprintf(lFmt, gid, x, yg, width, h, LightGrey, Black, LightSkyBlue, strings.Join(legend, "\n"))
}

func (g *dependencyGraph) AddGMLTitleBlock() {

	x0 := 0.0
	y0 := 0.0
	width := 900.0

	xl := x0 + 20.0
	xv := x0 + 200.0
	yl := y0 + 75.0
	height := yl + 6.0*vSpacing

	x := x0 + width/2.0
	y := y0 + height/2.0

	label := "Database Dependency Graph"
	var ta []string

	tbFmt := `	node
	[
		id	%d
		label	"%s"
		graphics
		[
			x	%f
			y	%f
			w	%f
			h	%f
			type	"%s"
			fill	"%s"
			outline	"%s"
		]
		LabelGraphics
		[
			text	"%s"
			color	"%s"
			fontSize	36
			fontStyle	"bold"
			fontName	"Dialog"
			anchor	"t"
		]`

	itemFmt := `		LabelGraphics
		[
			text	"%s"
			color	"%s"
			fontSize	14
			fontStyle	"bold"
			fontName	"Helvetica"
			x	%f
			y	%f
		]
		LabelGraphics
		[
			text	"%s"
			color	"%s"
			fontSize	14
			fontStyle	"normal"
			fontName	"Helvetica"
			x	%f
			y	%f
		]`

	ta = append(ta, fmt.Sprintf(tbFmt, g.id, label, x, y, width, height, Rectangle, LightGrey, Black, label, Black))

	t := time.Now()
	ta = append(ta, fmt.Sprintf(itemFmt, "Created", Black, xl, yl, t.Format("2006-01-02 15:04:05 MST"), Black, xv, yl))
	yl += vSpacing

	ta = append(ta, fmt.Sprintf(itemFmt, "Database", Black, xl, yl, g.DBName, Black, xv, yl))
	yl += vSpacing

	ta = append(ta, fmt.Sprintf(itemFmt, "Database Version", Black, xl, yl, g.DBMSVersion, Black, xv, yl))
	yl += vSpacing

	ta = append(ta, fmt.Sprintf(itemFmt, "Database Comment", Black, xl, yl, g.DBComment, Black, xv, yl))
	yl += vSpacing

	ta = append(ta, fmt.Sprintf(itemFmt, "Schema", Black, xl, yl, g.SchemaName, Black, xv, yl))
	yl += vSpacing

	ta = append(ta, fmt.Sprintf(itemFmt, "Schema Comment", Black, xl, yl, g.SchemaComment, Black, xv, yl))
	yl += vSpacing

	ta = append(ta, "\t]")
	g.gmlTitleBlock = strings.Join(ta, "\n")

}

func (g *dependencyGraph) AddNode(n graphNode) {

	if n.ObjectSchema == g.SchemaName {
		if g.SchemaNodes[n.ObjectSchema] == nil {
			g.SchemaNodes[n.ObjectSchema] = make(map[string]graphNode)
		}
		g.SchemaNodes[n.ObjectSchema][n.ObjectName] = n
	} else {
		if g.OtherNodes[n.ObjectSchema] == nil {
			g.OtherNodes[n.ObjectSchema] = make(map[string]graphNode)
		}
		g.OtherNodes[n.ObjectSchema][n.ObjectName] = n
	}
}

func (g *dependencyGraph) RenderDotGraph() (err error) {

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

func (g *dependencyGraph) RenderGMLGraph(includeCols bool) (err error) {

	x0 := 300.0
	y0 := 300.0
	y := y0
	yMin := 16.0 + y0
	yMax := yMin

	schemaFmt := `	node
	[
		id	%d
		label	"%s"
		graphics
		[
			x	533.0
			y	%f
			w	%f
			h	%f
			type	"roundrectangle"
			fill	"%s"
			outline	"%s"
			outlineStyle	"dotted"
		]
		LabelGraphics
		[
			text	"%s"
			fill	"%s"
			fontSize	15
			fontName	"Dialog"
			autoSizePolicy	"node_width"
			anchor	"t"
		]
		isGroup	1
	]`

	objFmt := `	node
	[
		id	%d
		label	"%s"
		graphics
		[
			x	%f
			y	%f
			w	%f
			h	%f
			type	"%s"
			fill	"%s"
			outline	"%s"
		]
		LabelGraphics
		[
			text	"%s"
			fontSize	13
			fontStyle	"bold"
			fontName	"Dialog"
			anchor	"t"
		]`
	// Helvetica
	pkFmt := `		LabelGraphics
		[
			text	"&#xd83d;&#xdd11;"
			color	"%s"
			fontSize	10
			fontStyle	"bold"
			fontName	"Dialog"
			x	%f
			y	%f
		]`

	// Helvetica
	colFmt := `		LabelGraphics
		[
			text	"%s"
			fontSize	10
			fontStyle	"%s"
			fontName	"Dialog"
			x	%f
			y	%f
		]
		LabelGraphics
		[
			text	"%s"
			fontSize	10
			fontStyle	"%s"
			fontName	"Dialog"
			x	%f
			y	%f
		]`

	othObjFmt := `	node
	[
		id	%d
		label	"%s"
		graphics
		[
			x	533.0
			y	%f
			w	%f
			h	%f
			type	"%s"
			fill	"%s"
			outline	"%s"
		]
		LabelGraphics
		[
			text	"%s"
			fontSize	13
			fontStyle	"bold"
			fontName	"Dialog"
			anchor	"t"
		]
		gid	%d
	]`

	var ta []string

	//////////////////////////////////////////////////////////////////////
	// Start the graph
	ta = append(ta, `Creator	"dep_graph"
graph
[
	label	""
	directed	1`)

	ta = append(ta, g.gmlTitleBlock)
	ta = append(ta, g.gmlLegend)

	//////////////////////////////////////////////////////////////////////
	// Add the "Other" schemas
	osy := make(map[string]float64)
	msh := 0.0
	for _, vs := range g.OtherSchemas {

		osy[vs.SchemaName] = y

		msw := minNodeWidth // max width of the current "Other" schema
		tCount := 0
		for _, obj := range g.OtherNodes[vs.SchemaName] {
			w := textWidth(obj.ObjectName, 15.0, Dialog, true) + (2.0 * hTextPadding)
			if w > msw {
				msw = w
			}
			tCount++
		}
		msh = ((minNodeHeight + vSpacing) * float64(tCount)) + 45.3143245

		ta = append(ta, fmt.Sprintf(schemaFmt, vs.ID, vs.SchemaName, y, msw, msh, PaleGreen, DarkSlateGrey, vs.SchemaName, LightSkyBlue))
		y += msh + (2.0 * vSpacing)
	}

	//////////////////////////////////////////////////////////////////////
	// Add the objects for the "Other" schemas
	objH := minNodeHeight

	for _, vs := range g.OtherSchemas {
		objY := osy[vs.SchemaName]

		// sort the "Other" schema objects by name
		var objs []graphNode
		for oName := range g.OtherNodes[vs.SchemaName] {
			objs = append(objs, g.OtherNodes[vs.SchemaName][oName])
		}
		sort.Slice(objs, func(i, j int) bool { return objs[j].ObjectName > objs[i].ObjectName })

		for _, obj := range objs {

			color := nodeColor(obj.ObjectType)
			shape := nodeShape(obj.ObjectType)
			objY += vSpacing + objH
			objW := textWidth(obj.ObjectName, 13.0, Dialog, true) + (2.0 * hTextPadding)
			if defaultTextWidth > objW {
				objW = defaultTextWidth
			}

			ta = append(ta, fmt.Sprintf(othObjFmt, obj.ID, obj.ObjectName, objY, objW, objH, shape, color, Black, obj.ObjectName, vs.ID))
			if objY > yMax {
				yMax = objY
			}

		}
	}

	//////////////////////////////////////////////////////////////////////
	// Add the schema objects

	y = yMin

	for schemaName, _ := range g.SchemaNodes {

		// sort the schema objects by name
		var objs []graphNode
		for oName := range g.SchemaNodes[schemaName] {
			objs = append(objs, g.SchemaNodes[schemaName][oName])
		}
		sort.Slice(objs, func(i, j int) bool { return objs[j].ObjectName > objs[i].ObjectName })

		// Determine the layout of the schema objects. For the minimal graph (no column names) we can
		// layout the objects in a circular fashion (and generate an svg file from that).
		// For the extended graph we just lay out the objects in a vertical line.

		iCount := len(objs)

		// For the minimal graph, determine a radius of the layout circle.
		// If the calculated radius is under a minimum size then use the minimum size,
		// otherwise use the calculated radius.
		r := (minNodeHeight + vSpacing) * float64(iCount) / 2.0
		rMin := 300.0

		if r < rMin {
			r = rMin
		}

		// Having a radius, specify the X, Y coordinate of the circle center
		cY := (yMin + yMax) / 2.0
		cX := x0 + 700.0 + r

		// For the minimal graph, determine a ratio to stretch the Y-axis by.
		// This is to help with object crowding at the top and bottom of the circle (now oval) and to also
		// scale the schema objects layout to the other schemas objects (if any)
		yRatio := (yMax - yMin - 4.0*minNodeHeight) / (2.0 * r)
		switch {
		case yRatio < 1.0:
			yRatio = 1.0
		case yRatio > 2.0:
			yRatio = 2.0
		}

		cRad := (2.0 * math.Pi)
		aInc := cRad / float64(iCount)

		// Determine the object placements
		for i, obj := range objs {

			objW := textWidth(obj.ObjectName, 13.0, Dialog, true) + (2.0 * hTextPadding)

			if defaultTextWidth > objW {
				objW = defaultTextWidth
			}

			if includeCols {

				obj.X = x0 + 1000.0
				obj.Y = y + (obj.H / 2.0)
				y += obj.H + vSpacing

				if objW > obj.W {
					obj.W = objW
				}

			} else {

				aI := aInc * float64(i)
				dX := (r * math.Sin(aI))
				dY := (r * math.Cos(aI) * yRatio)

				// TODO: If we're near the top/bottom of the "circle" then we should adjust the y ratio
				// somewhat

				obj.X = cX + dX
				obj.Y = cY - dY
				obj.H = minNodeHeight
				obj.W = objW
			}
			objs[i] = obj
		}

		// Add the objects to the graph
		for _, obj := range objs {

			color := nodeColor(obj.ObjectType)
			shape := nodeShape(obj.ObjectType)

			ta = append(ta, fmt.Sprintf(objFmt, obj.ID, obj.ObjectName, obj.X, obj.Y, obj.W, obj.H, shape, color, Black, obj.ObjectName))

			if includeCols {

				xPk := obj.X - (obj.W / 2.0)
				xl := obj.X - (obj.W / 2.0) + hTextPadding
				xd := xl + obj.dtOffset

				var columns []string

				for _, col := range obj.Columns {

					colY := obj.Y + col.Y - (obj.H / 2.0)

					if col.IsPK {
						columns = append(columns, fmt.Sprintf(pkFmt, GoldenRod, xPk, colY))
					}

					var fontStyle string
					if col.IsNullable {
						fontStyle = "normal"
					} else {
						fontStyle = "bold"
					}

					columns = append(columns, fmt.Sprintf(colFmt, col.Name, fontStyle, xl, colY, col.DataType, fontStyle, xd, colY))

				}
				ta = append(ta, strings.Join(columns, "\n"))
			}
			ta = append(ta, "	]")
		}
	}

	//////////////////////////////////////////////////////////////////////
	// Add the edges
	edgeFmt := `	edge
	[
		source	%d
		target	%d
		graphics
		[
			fill	"%s"
			targetArrow	"standard"
		]
		edgeAnchor
		[
			ySource	-1.0
			yTarget	1.0
		]
	]`

	for i, v := range g.Edges {
		for j, _ := range v {
			ta = append(ta, fmt.Sprintf(edgeFmt, i, j, Black))
		}
	}

	//////////////////////////////////////////////////////////////////////
	// Close the graph
	ta = append(ta, "]")

	res := strings.Join(ta, "\n")

	// ensure that the file directory exists
	dirName := path.Join(g.OutputDir, g.SchemaName)
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0745)
		if err != nil {
			return err
		}
	}

	var outFileName string
	if includeCols {
		outFileName = path.Join(dirName, "dependencies.gml")
	} else {
		outFileName = path.Join(dirName, "dependencies-min.gml")
	}
	err = os.WriteFile(outFileName, []byte(res), 0644)

	return err
}

func MakeDepenencyGraphs(vs *m.Schema, md *m.MetaData) (err error) {

	g := NewDependencyGraph(vs, md)

	if len(g.SchemaNodes) == 0 && len(g.OtherNodes) == 0 {
		return err
	}

	// generate the DOT graph
	err = g.RenderDotGraph()
	if err != nil {
		return err
	}

	// generate the GML graph with just the node names
	err = g.RenderGMLGraph(false)
	if err != nil {
		return err
	}

	// generate the GML graph that includes the column definitions
	err = g.RenderGMLGraph(true)
	if err != nil {
		return err
	}

	return err
}

func NewDependencyGraph(vs *m.Schema, md *m.MetaData) *dependencyGraph {

	g := dependencyGraph{
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

	g.id = 0

	////////////////////////////////////////////////////////////////
	// Filter the node types to just those found in the specific graph
	// Determine the "other" schemas
	var tt = make(map[string]bool)
	var ts = make(map[string]bool)
	dependencies := md.FindDependencies(g.SchemaName, "")
	if len(dependencies) > 0 {
		for _, v := range dependencies {
			_, ok := tt[v.ObjectType]
			if !ok {
				tt[v.ObjectType] = true
			}
			_, ok = ts[v.ObjectSchema]
			if !ok {
				ts[v.ObjectSchema] = true
			}
			_, ok = tt[v.DepObjectType]
			if !ok {
				tt[v.DepObjectType] = true
			}
			_, ok = ts[v.DepObjectSchema]
			if !ok {
				ts[v.DepObjectSchema] = true
			}
		}
	}
	dependents := md.FindDependents(g.SchemaName, "")
	if len(dependents) > 0 {
		for _, v := range dependents {
			_, ok := tt[v.ObjectType]
			if !ok {
				tt[v.ObjectType] = true
			}
			_, ok = ts[v.ObjectSchema]
			if !ok {
				ts[v.ObjectSchema] = true
			}
			_, ok = tt[v.DepObjectType]
			if !ok {
				tt[v.DepObjectType] = true
			}
			_, ok = ts[v.DepObjectSchema]
			if !ok {
				ts[v.DepObjectSchema] = true
			}
		}
	}

	for _, nodeType := range nodeTypes {
		_, ok := tt[nodeType]
		if ok {
			g.nodeTypes = append(g.nodeTypes, nodeType)
		}
	}

	////////////////////////////////////////////////////////////////
	g.AddGMLTitleBlock()
	g.AddGMLLegend()

	////////////////////////////////////////////////////////////////
	// Add the "other" schemas
	var os []string
	for k := range ts {
		os = append(os, k)
	}
	sort.Slice(os, func(i, j int) bool { return os[j] > os[i] })

	for _, v := range os {
		if g.SchemaName != v {
			g.id++
			g.OtherSchemas = append(g.OtherSchemas, graphSchema{
				ID:         g.id,
				SchemaName: v,
			})
		}
	}

	////////////////////////////////////////////////////////////////
	// Add dependencies information
	g.SchemaNodes = make(nodeMap)
	g.OtherNodes = make(nodeMap)
	g.Edges = make(edgeMap)

	//dependencies = md.FindDependencies(vs.Name, "")
	if len(dependencies) > 0 {
		md.SortDependencies(dependencies)
		for _, v := range dependencies {
			g.AddDependency(&v, md)
		}
	}

	//dependents = md.FindDependents(vs.Name, "")
	if len(dependents) > 0 {
		md.SortDependencies(dependents)
		for _, v := range dependents {
			g.AddDependency(&v, md)
		}
	}

	return &g
}

func mkNode(nodeId int, objectSchema, objectName, objectType string, addCols bool, md *m.MetaData) (n graphNode) {

	nodeHeight := nodeLabelH

	// determine the width for the node using the greater of
	// - the default width,
	// - the node label width, and
	// - the max column width
	nodeWidth := defaultTextWidth + (2 * hTextPadding)
	dtOffset := 0.0
	tw := textWidth(objectName, 13.0, Dialog, true) + (2.0 * hTextPadding)
	if tw > nodeWidth {
		nodeWidth = tw
	}

	var cols []nodeColumn

	if addCols {
		cols = mkNodeColumns(objectSchema, objectName, objectType, md)
		if len(cols) > 0 {
			for _, c := range cols {
				colWidth := textWidth(c.Name, 10.0, Helvetica, c.IsNullable)
				dtWidth := textWidth(c.DataType, 10.0, Helvetica, c.IsNullable)
				tw = colWidth + dtWidth + (4.0 * hTextPadding)
				if tw > nodeWidth {
					nodeWidth = tw
				}

				if colWidth+(2.0*hTextPadding) > dtOffset {
					dtOffset = colWidth + (2.0 * hTextPadding)
				}
			}

			// adjust the node height based on the number of columns for the node
			nodeHeight += (colLabelH * float64(len(cols)))
		}
	}

	return graphNode{
		ID:           nodeId,
		ObjectSchema: objectSchema,
		ObjectName:   objectName,
		ObjectType:   objectType,
		H:            nodeHeight,
		W:            nodeWidth,
		dtOffset:     dtOffset,
		Columns:      cols,
		Color:        nodeColor(objectType),
	}

}

func mkNodeColumns(objectSchema, objectName, objectType string, md *m.MetaData) (cols []nodeColumn) {

	switch strings.ToUpper(objectType) {
	case "TABLE", "BASE TABLE", "FOREIGN TABLE", "MATERIALIZED VIEW", "VIEW":
		c := md.FindColumns(objectSchema, objectName)

		if len(c) > 0 {

			var pkc = make(map[string]int)
			pks := md.FindPrimaryKeys(objectSchema, objectName)

			if len(pks) > 0 {
				p := strings.Split(pks[0].Columns, ", ")
				for _, pkCol := range p {
					pkc[pkCol] = 1
				}
			}

			for _, v := range c {
				nullable := v.IsNullable == "YES"
				_, isPk := pkc[v.Name]
				y := (colLabelH * float64(v.OrdinalPosition)) + colLabelH/2.0

				cols = append(cols, nodeColumn{
					Name:            v.Name,
					DataType:        v.DataType,
					OrdinalPosition: v.OrdinalPosition,
					IsPK:            isPk,
					IsNullable:      nullable,
					Y:               y,
				})
			}
		}
	}

	return cols
}

func nodeColor(nodeType string) string {
	color, ok := nodeColors[nodeType]
	if ok {
		return color
	}
	return Grey
}

func nodeShape(nodeType string) string {
	shape, ok := nodeShapes[nodeType]
	if ok {
		return shape
	}
	return Rectangle
}

// textWidth performs a rough estimate of the width of a string in units where
// 1 equals the width of the majority of characters. This does not make any
// consideration of differing font faces, kerning, or any other such thing.
func textWidth(s string, pts float64, fontFace string, isBold bool) (w float64) {

	// TODO consider font face.

	dependencyFontWidthFactor := 0.6
	fontBoldFactor := 1.09738
	w = 0.0

	for _, v := range []byte(s) {
		switch string(v) {
		case "i", "j", "l", "'":
			w += 0.44703
		case "I", "f", "r", "t":
			w += 0.58656
		case " ", ",", "(", ")", "[", "]", "-", "`", "\"":
			w += 0.58656
		case ".":
			w += 0.62791
		case "J", "L", "a", "b", "c", "d", "e", "g", "h", "k", "n", "o", "p", "q", "s", "u", "v", "x", "y", "z":
			w += 1
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_":
			w += 1.0
		case "F", "T", "Z":
			w += 1.13953
		case "A", "B", "C", "D", "E", "H", "K", "N", "P", "R", "S", "U", "V", "X", "Y", "w":
			w += 1.27907
		case "G", "O", "Q":
			w += 1.41602
		case "M", "m":
			w += 1.55556
		case "W":
			w += 1.69251
		default:
			w += 1.07456
		}
	}

	w *= pts * dependencyFontWidthFactor
	if isBold {
		w *= fontBoldFactor
	}
	return w
}
