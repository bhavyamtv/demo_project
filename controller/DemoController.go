package controller

import (
	"demo_project/models"

	"github.com/gin-gonic/gin"
)

func ContactUs(c *gin.Context) { // hio
	data := make(map[string]string)
	data["name"] = c.PostForm("name")
	data["email"] = c.PostForm("email")
	data["mobile"] = c.PostForm("mobile")
	data["msg"] = c.PostForm("msg")

	if data["name"] != "" && data["email"] != "" && data["mobile"] != "" && data["msg"] != "" {
		Flag, _ := models.ContactUs(data)

		if Flag == 1 {
			c.JSON(200, gin.H{"code": 1, "result": "success!!"})
			return
		} else if Flag == 2 {
			c.JSON(200, gin.H{"code": 1, "result": "success!!"})
			return
		} else {
			c.JSON(200, gin.H{"code": 1, "result": "failed!!"})
			return
		}

	} else {
		c.JSON(200, gin.H{"code": 0, "result": "invalid input"})
		return
	}

}
