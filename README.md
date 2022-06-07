# How to structure Golang applications
This is a simple go structured project for CRUD operations in Restful api's

## How to run the application
You don't need much to run this application. It requires the following installed on your computer:
- Golang 1.15+

Build the docker container for postgreSQL:
```
docker-compose up -d
```
To start the server, run following commands:
```
cd cmd/server
task
```
## How to use the application
Hit the following endpoints from postman. First create the Author and Category and give the id in Article body.

### To Create an Author

Method: POST
Url: http://localhost:3000/authors
Body:
```
{
"name": "new author",
"age": 21,
"email": "example@gmail.com"
}
```

### To Create a new Category

Method: POST
Url: http://localhost:3000/category
Body:
```
{
"name": "comedy",
}
```

### To Create a new Article

Method: POST
Url: http://localhost:3000/articles
Body:
```
{
    "name": "new article",
    "description": "description new",
    "author_id": 1,
    "category_id": 1
}
```

## Test with GET endpoints
```
Method: GET
http://localhost:3000/articles
http://localhost:3000/authors
http://localhost:3000/category
```

Similarly you can test other CRUD operations with different methods
```
Method: GET, DELETE, PUT
http://localhost:3000/articles/{article_id}
```
## License
MIT
