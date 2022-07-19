package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"gitlab.com/ishangoyal/goapi/models"
	//jwtAuth "gitlab.com/ishangoyal/goapi/services"
)

type todo struct {
	ID          string `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
}

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	// routes
	router.GET("/task", GetTasks)
	//router.GET("/task/:id", GetTaskById)
	router.POST("/task", PostTasks)
	router.POST("/signup", jwtAuth.SignUp)
	router.DELETE("/task/:id", DeleteTask)
	router.PUT("/task/:id", UpdateTask)

	return router
}

var DB = models.Init()

// get all tasks
func GetTasks(c *gin.Context) {
	var allTask []todo
	out, err := DB.Query(`SELECT * from tasks`)
	checkErr(err)

	for out.Next() {
		var id string
		var task string
		var description string
		err = out.Scan(&task, &description, &id)
		checkErr(err)

		allTask = append(allTask, todo{ID: id, Task: task, Description: description})

	}
	c.IndentedJSON(http.StatusOK, allTask)
}

// get task by id
/*
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	var allTask []todo

	sqlStmnt := `SELECT * from tasks WHERE ID IN ($1)`
	out, err := DB.Query(sqlStmnt, id)
	checkErr(err)

	for out.Next() {
		var id string
		var task string
		var description string
		err = out.Scan(&task, &description, &id)
		checkErr(err)
		allTask = append(allTask, todo{ID: id, Task: task, Description: description})

	}
	c.IndentedJSON(http.StatusOK, allTask)
}*/

// post a new task
func PostTasks(c *gin.Context) {
	var newtodo todo

	if err := c.BindJSON(&newtodo); err != nil {
		return
	}

	sqlStmnt := `INSERT INTO tasks (id,task,description) VALUES ($1,$2,$3)`
	_, err := DB.Exec(sqlStmnt, newtodo.ID, newtodo.Task, newtodo.Description)
	checkErr(err)
	c.IndentedJSON(http.StatusCreated, newtodo)
}

// update a task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var update todo

	if err := c.BindJSON(&update); err != nil {
		return
	}

	sqlStmnt := `UPDATE tasks SET task=$1, description=$2 WHERE id=$3`
	_, err := DB.Exec(sqlStmnt, update.Task, update.Description, id)
	checkErr(err)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

// delete a task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	sqlStatement := `DELETE FROM tasks WHERE id = $1;`
	_, err := DB.Exec(sqlStatement, id)
	checkErr(err)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
}

// error handling
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
