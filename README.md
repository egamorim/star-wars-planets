# Star wars API

It is a service built upon GoLang and MongoDB with CRUD operations for Planets which appears in Star Wars Movies

## Requirements
- Go
- MongoDB

## Installation and run

Simple Go application
```bash
go build -o star-wars cmd/server/main.go
./star-wars
```
Using Docker

```
docker build . -t star-wars 
docker-compose up
```

## Usage

- The base URL id : http://localhost:8000/planets
- The authentication and authorization mechanisms aren't scope of that sample
 
### Insert
Inserts a new planet, the user must provide name, terrain and climate then the service will call the [swapi](https://swapi.dev/about) api to get information about how many movies the provided planet appears, this information will be saved at the database.

See an example of request bellow:
```
curl --request POST \
  --url http://localhost:8000/planets \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Alderaan",
	"terrain" : "desert",
	"climate" : "arid"
}'
```
In case of the provided planet name was found in swapi api, it will return an error of `Planet not found in swapi`

### Get all Planets
Return a paginated list of all the planets from the database.

See an example of request bellow:
```
curl --request GET \
  --url 'http://localhost:8000/planets?offset=0&limit=5'
```

### Get a planet by ID
Return a planet by the provided ID

See an example of request bellow:
```
curl --request GET \
  --url http://localhost:8000/planets/601d387524016ab66116cd85

```

### Delete a planet 
Delete a planet from the database by its provided ID

See an example of request bellow:
```
curl --request DELETE \
  --url http://localhost:8000/planets/601d387524016ab66116cd85
```

### Get by name 
Returns a planet with the given name. It is considered part of the name as well

See an example of request bellow:
```
curl --request GET \
  --url http://localhost:8000/planets/findByName/aaaaa
```

### Note
If you have any questions or comments feel free to reach out me by my e-mail: egamorim@gmail.com