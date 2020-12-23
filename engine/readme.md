# engine

This is where the queries and connection creation for various databases
reside, one subdirectory per database engine.

For dictionary creation of any database to work there needs to be:
 1. a go driver for that database engine. See https://github.com/golang/go/wiki/SQLDrivers
 1. some means of querying the database metadata (tables/views/function calls?)
 1. enough understanding of the metadata to write the appropriate queries

Table of top contenders:
 * Ranking: The ranking from https://db-engines.com/en/ranking/relational+dbms as of 2020-12-22
 * Name: The name of the database engine
 * Avail: Is there a readily available database to test with?
 * Metadata: Are there tables/views/function calls available for obtaining desired metadata?
 * Driver: A/the database driver to use




| Ranking | Name                                | Avail | Metadata  | Driver                            | License |
| ------- | ----------------------------------- | ----- | --------- | --------------------------------- | ------- |
|         | ODBC driver                         |       |           | github.com/alexbrainman/odbc      | BSD     |
| 1       | Oracle                              | Y     | Y         | github.com/godror/godror          | Apache  |
| 2       | MySQL                               | Y     | Y         | github.com/go-sql-driver/mysql    | MPL     |
| 3       | Microsoft SQL Server                | ?     | Y         | github.com/denisenkom/go-mssqldb  | BSD     |
| 4       | PostgreSQL                          | Y     | Y         | github.com/lib/pq                 | Pg?     |
| 5       | IBM Db2                             |       |           | github.com/ibmdb/go_ibm_db        | BSD     |
| 6       | SQLite                              | Y     |           | github.com/mattn/go-sqlite3       | MIT     |
| 7       | Microsoft Access                    | ?     |           | github.com/bennof/accessDBwE      | BSD     |
|         |                                     |       |           | github.com/mattn/go-adodb         | MIT     |
| 8       | MariaDB                             | Y     | Y         | github.com/go-sql-driver/mysql    | MPL     |
| 13      | SAP HANA                            |       |           | github.com/SAP/go-hdb/driver      | Apache? |
| 15      | Google BigQuery                     |       |           | github.com/viant/bgc              | Apache  |
| 16      | Firebird                            |       |           | github.com/nakagami/firebirdsql   | MIT     |

The core? package *should* define an interface for the querying. It
would be nice to have a script/utility for creating the shell function
calls for each database engine that ensures that a bare-minimum
interface is implemented for each engine.
