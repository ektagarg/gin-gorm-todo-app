package Controllers

import (
	"../Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todo []Models.Todo
	err := Models.GetAllTodo(&todo)
	if err != nil {
		c.String(http.StatusNotFound, todo)
	} else {
		c.String(http.StatusOk, todo)
	}
}

func CreateATodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)
	err := Models.CreateATodo(&todo)
	if err != nil {
		c.String(http.StatusNotFound, todo)
	} else {
		c.String(http.StatusOk, todo)
	}
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Models.Todo
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.String(http.StatusNotFound, todo)
	} else {
		c.String(http.StatusOk, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.String(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = Models.UpdateATodo(&todo, id)
	if err != nil {
		c.String(http.StatusNotFound, todo)
	} else {
		c.String(http.StatusOk, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.DeleteATodo(&todo, id)
	if err != nil {
		c.String(http.StatusNotFound, todo)
	} else {
		c.String(http.StatusOk, todo)
	}
}
