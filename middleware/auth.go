package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, password, ok := c.Request.BasicAuth()

		if (user == "" || password == "") && ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"result": "Username atau Password tidak boleh kosong",
			})
			c.Abort()
			return
		}

		if ok {
			if (user == "editor" && password == "secret") || (user == "admin" && password == "password") {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"result": "Username atau Password salah",
				})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"result": "Unauthorized",
			})
			c.Abort()
		}
	}
}
