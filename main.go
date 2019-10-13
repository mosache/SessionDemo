package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SessionDemo/models"
	"SessionDemo/session"
)

func main() {
	r := gin.Default()

	r.Use(session.GinSessionWare())

	r.LoadHTMLGlob("./templates/*")

	r.Any("/login", login)

	r.GET("/index", index)

	r.GET("/vip", vip)

	r.Run()
}

func authMidWare(c *gin.Context) {
	tempSD, ok := c.Get(session.GinContextSessionKey)

	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	sd, ok := tempSD.(*session.Data)

	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	value, err := sd.Get("isLogin")

	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	isLogin, ok := value.(bool)

	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	if !isLogin {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.Next()
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func login(c *gin.Context) {

	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "0",
				"msg":    "paras error!",
			})
			return
		}

		if user.UserName == "admin" && user.Pwd == "123" {

			tempSD, ok := c.Get(session.GinContextSessionKey)

			if !ok {
				panic("session midware")
			}

			sd := tempSD.(*session.Data)

			sd.Set("isLogin", true)

			c.HTML(http.StatusOK, "index.html", nil)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "0",
				"msg":    "username or pwd error",
			})
		}
	}

}

func vip(c *gin.Context) {
	c.HTML(http.StatusOK, "vip.html", nil)
}
