# Migrations
Require
migrate https://github.com/gemnasium/migrate

## Postgres Migration
Configuration Store database
Usage from Terminal
```
# create new migration file in path
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ create alarm

# apply all available migrations
migrate -url "postgres://amoeba:password@localhost:35432/?sslmode=disable" -path .\database\migration\ up

# roll back all migrations
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ down

# roll back the most recently applied migration, then run it again.
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ redo

# run down and then up command
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ reset

# show the current migration version
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ version

# apply the next n migrations
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ migrate +1
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ migrate +2
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ migrate +n

# roll back the previous n migrations
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ migrate -1
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ migrate -2
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ migrate -n

# go to specific migration
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ goto 1
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ goto 10
migrate -url "postgres://postgres@localhost:5432/postgres?sslmode=disable" -path .\database\migration\ goto v
```
