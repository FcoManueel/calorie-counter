# REST API for calorie counter

## Installation steps:

1. Get dependencies:
    - `go get github.com/tanel/dbmigrate`
    - `go get github.com/lib/pq`
    - `go get goji.io`
    - `go get github.com/satori/go.uuid`
    - `go get gopkg.in/pg.v3`
    
2. Create database user/table (check instructions at the top of db/migrate/001_initial.sql)
3. Run the server with `go run main.go`