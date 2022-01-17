# db-dictionary-core

Core libraries for using relational database meta-data to generate data dictionaries.

This is intended as the next generation of https://github.com/gsiems/DataDict with the following differences:
 * Use go instead of perl
 * Just data-dictionaries
 * Default output is static html files

Supported databases are:

 * Postgresql
 * Sqlite
 * MySQL/MariaDB (TODO)
 * Oracle (TODO)
 * MS SQL-Server (TODO)

See [db-dictionary-example](https://github.com/gsiems/db-dictionary-example) for some sample data dictionaries.



## Configuration

| Parameter      | Command line | Environment/Config file | Description |
| -------------- | ------------ | ----------------------- | ----------- |
| CommentsFormat | -f           | comments_format         | The formatter to use for rendering comments {none, markdown} (default: none) |
| ConfigFile     | -c           |                         | The configurations file to read, if any. |
| CSSFiles       | -css         | css_files               | The comma-separated list of CSS files to use in place of the default (default: none) |
| DbComment      | -comment     | db_comment              | The comment to use for the database (for those databases that do not support ```COMMENT ON DATABASE ...```) (default: none) |
| DbName         | -db          | db_name                 | The name of the database to connect to |
| DSN            |              | dsn                     | The DSN to use for connecting to the database (will attempt to create one based on DbName, Host, Port, etc. if not specified) |
| ExcludeSchemas | -x           | exclude_schemas         | The comma-separated list of schemas to exclude (default: none) |
| File           | -file        | file                    | The database file to read (sqlite) |
| Host           | -host        | host                    | The database host to connect to (default: localhost) |
| ImgFiles       | -img         | img_files               | The comma-separated list of image files to include (for use with custom CSS) (default: none) |
| IncludeSchemas | -s           | include_schemas         | The comma-separated list of schemas to include. Takes precedence over ExcludeSchemas. If neither are specified than all non-system schemas are included. |
| JSFiles        | -js          | js_files                | The comma-separated list of javascript files to include (default: none) |
| Minify         | -minify      | minify                  | Indicates if the output should be minified to reduce files size (default: false)
| OutputDir      | -out         | output_dir              | The directory to write the output files to (defaults to the current directory) |
| Port           | -port        | port                    | The port number to connect to (default depends on the database engine) |
| Username       | -user        | username                | The username to connect as (defaults to the current OS user) |
| UserPass       |              | user_pass               | The password to use to connect as |
| Verbose        | -v           | verbose                 | Indicates if additional feedback should be printed to STDOUT (default: false) |
