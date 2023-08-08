# go-mysqlping
Ping MySQL database with [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

## Installation

```
$ go install github.com/utgwkk/go-mysqlping
```

## Usage

```
Usage:
  go-mysqlping [options]

Application Options:
  -u, --user=user_name       MySQL user name (default: root)
  -p, --password=password    MySQL user password
  -h, --host=host_name       Host to connect to the MySQL server (default: 127.0.0.1)
  -P, --port=port_num        Port used for the connection (default: 3306)
      --timeout=timeout      Timeout seconds (default: 60)
      --help                 Show this help
      --version              Show this version
```
