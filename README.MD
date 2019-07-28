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

DB Setup in Config/Database.go
Change DBName, Username, Password and host, port as well if running on podunction environment 

## APIs

* `GET` : Get all todos
* `POST` : Create a todo
* `GET` : Get a todo
* `PUT` : Update a todo
* `DELETE` : Delete a todo

#Post Params
```
{
	"name": "Groceries shopping",
	"description": "Biscuits, Tea, Milk, Soap, Powder",
}
```
