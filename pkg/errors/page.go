package errors

import "github.com/gin-gonic/gin"

func SomethingWentWrong(c *gin.Context, err error) {
	c.String(500, "Oops! Please retry.")
	// A custom error page with HTML templates can be shown by c.HTML()
	_ = c.Error(err)
	c.Abort()
}

