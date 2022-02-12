# db-dictionary

Use relational database meta-data to generate data dictionaries.

## Features/roadmap

 - [x] Generates static HTML pages

 - [x] Dictionary information for:

    - [x] Schemas
    - [x] Tables
    - [x] Views
    - [x] Materialized views
    - [x] Constraints
    - [x] Domains
    - [x] Indexes
    - [x] Relationhips

 - [x] Ability to filter tabular data on pages

 - [x] Ability to sort tabular data on pages

    - single column sort only

 - [ ] Generate dependency (network) graphs (WIP)

    - [ ] GraphML file (WIP)
    - [x] Graphviz file
    - [x] Svg (requires Graphviz)

 - [ ] Generate relationship (network) graphs (WIP)

    - [ ] GraphML file (TODO)
    - [ ] Graphviz file (TODO)
    - [ ] Svg (TODO)

 - [x] Ability to specify different CSS/image files for theming

 - [x] Ability to specify additional javascript files

 - [x] Reasonably fast

 - [x] PostgreSQL support

 - [x] Sqlite support

 - [x] MySQL/Mariadb support

 - [ ] MS SQL-server support (WIP)

 - [x] Oracle support

## Configuration

Configurations can be specified through a combination of configuration file, environment variables, and command line arguments.

 * Command line arguments take precedence over both configuration file parameters and environment variables.
 * Environment variables take precedence over configuration file parameters.

| Parameter      | Command line | Environment/Config file | Description |
| -------------- | ------------ | ----------------------- | ----------- |
| CommentsFormat | -f           | comments_format         | The formatter to use for rendering comments {none, markdown} (default: none) |
| ConfigFile     | -c           |                         | The configurations file to read, if any |
| CSSFiles       | -css         | css_files               | The comma-separated list of CSS files to use in place of the default (default: none) |
| DbComment      | -comment     | db_comment              | The comment to use for the database (for those databases that do not support ```COMMENT ON DATABASE ...```) (default: none) |
| DbEngine       | -dbms        | dbms                    | The dbms to generate the dictionary for {oracle, postgresql, mariadb, mysql, sqlite} |
| DbName         | -db          | db_name                 | The name of the database to connect to |
| DSN            |              | dsn                     | The DSN to use for connecting to the database (will attempt to create one based on DbName, Host, Port, etc. if not specified) |
| ExcludeSchemas | -x           | exclude_schemas         | The comma-separated list of schemas to exclude (default: none) |
| File           | -file        | file                    | (Sqlite) The database file to read |
| GraphvizCmd    | -gv          | graphviz_cmd            | The Graphviz command to run (default: fdp) |
| HideSQL        | -nosql       | hide_sql                | Do not show the queries used for views and materialized views (default is to show queries) |
| Host           | -host        | host                    | The database host to connect to (default: localhost) |
| ImgFiles       | -img         | img_files               | The comma-separated list of image files to include (for use with custom CSS) (default: none) |
| IncludeSchemas | -s           | include_schemas         | The comma-separated list of schemas to include. Takes precedence over ExcludeSchemas. If neither are specified than all non-system schemas are included. |
| JSFiles        | -js          | js_files                | The comma-separated list of javascript files to include (default: none) |
| Minify         | -minify      | minify                  | Indicates if the output should be minified to reduce files size (default: false) |
| NoGraphviz     | -nogv        | no_graphviz             | Do not (attempt to) run Graphviz  (default is to run Graphviz) |
| OutputDir      | -out         | output_dir              | The directory to write the output files to (defaults to the current directory) |
| Port           | -port        | port                    | The port number to connect to (default depends on the database engine) |
| SSLMode        | -sslmode     | ssl_mode                | (Postgresql) Set the SSL mode to use {disable, require, verify-ca, verify-full} (default: require) |
| Username       | -user        | username                | The username to connect as (defaults to the current OS user) |
| UserPass       |              | user_pass               | The password to use to connect as |
| Verbose        | -v           | verbose                 | Indicates if additional feedback should be printed to STDOUT (default: false) |

----

See [db-dictionary-example](https://github.com/gsiems/db-dictionary-example) for some sample data dictionaries.
