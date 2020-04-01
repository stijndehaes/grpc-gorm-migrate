package user

import (
	"github.com/gin-gonic/gin"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/errors"
)

func AddRoutes(r *gin.Engine) {
	v1Users := r.Group("api/v1/users")
	{
		v1Users.GET("/", GetUsers)
		v1Users.PUT("/", CreateUser)
	}
}

func GetUsers(c *gin.Context) {
	users, err := ListUsers()
	if err != nil {
		errors.SomethingWentWrong(c, err)
		return
	}
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {

}

