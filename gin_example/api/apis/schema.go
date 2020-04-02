package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luxiaotong/go_practice/gin_example/api/models"
)

//GetDatabaseList is a function to list all databases
func GetDatabaseList(c *gin.Context) {
	result := models.Databases()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}
