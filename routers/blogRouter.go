package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishangoyal13/blogPost/controllers"
	"github.com/ishangoyal13/blogPost/middleware"
	"github.com/ishangoyal13/blogPost/models"
	//jwtAuth "gitlab.com/ishangoyal/goapi/services"
)

type CreateTaskInput struct {
	ID      uint   `gorm:"primary key:autoIncrement" json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	// routes
	router.GET("/blog", GetTasks)
	//router.GET("/task/:id", GetTaskById)
	router.POST("/blog", PostTasks)
	// router.POST("/signup", jwtAuth.SignUp)
	router.DELETE("/blog/:id", DeleteTask)
	//router.PUT("/task/:id", UpdateTask)
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}

// get all tasks
func GetTasks(c *gin.Context) {
	var allBlog []models.Blog
	models.DB.Find(&allBlog)

	c.IndentedJSON(http.StatusOK, gin.H{"data": allBlog})
}

// post a new task
func PostTasks(c *gin.Context) {
	var newtodo CreateTaskInput

	if err := c.ShouldBindJSON(&newtodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tod := models.Blog{Author: &newtodo.Author, Title: &newtodo.Title, Content: &newtodo.Content}
	models.DB.Create(&tod)
	c.IndentedJSON(http.StatusCreated, gin.H{"data": tod})
}

// delete a task
func DeleteTask(c *gin.Context) {
	var delTodo models.Blog
	if err := models.DB.Where("id = ?", c.Param("id")).First(&delTodo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&delTodo)

	c.JSON(http.StatusOK, gin.H{"data": true})

}

// extra code

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

// update a task
/*
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
}*/
