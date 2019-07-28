# Api go simple
It is a just simple REST API with Go using:
1. **Gin Framework**
2. **Gorm** 

## Installation & Run
```bash
# Download this project
$ go get github.com/go-sql-driver/mysql

# Download Gin Framework
$ go get github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm
```

## DB Setup
1. Go to Config/Database.go
2. Update DBName, User, Password and Host, Port according to your database configuration

## API list

* `GET` : Get all todos
* `POST` : Create a todo
* `GET` : Get a todo
* `PUT` : Update a todo
* `DELETE` : Delete a todo

## Post Params
```
{
	"name": "Groceries shopping",
	"description": "Biscuits, Tea, Milk, Soap, Powder",
}
```
