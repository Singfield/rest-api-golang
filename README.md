# Go Rest api

Hi ! This is my first little project in golang.
I'm learning this super language by making a fully operational rest api wich retrieve users comments

## techs

- gorm
- gorill/mux
- logrus
- go 1.17

- docker


## instruction

docker :

```
docker run --name rest-api-golang-comments -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```


powershell env variables :
```
    $env:DB_USERNAME = "postgres"
    $env:DB_PASSWORD = "postgres"
    $env:DB_TABLE = "postgres"
    $env:DB_PORT = "5432"
    $env:DB_HOST ="localhost"
```
for bash : 

```
export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_TABLE=postgres
export DB_PORT=5432
export DB_HOST=localhost

```

## launch

``` go run cmd/server/main.go```

## test endpoint

**api/comments** (Post)

```curl --header "Content-Type: application/json" -d "{Slug:"/", Body:"first commit",Author:"test",Created:"User"}" http://localhost:8000/api/comments -v```