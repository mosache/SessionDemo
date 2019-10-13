package session

import (
	"github.com/gin-gonic/gin"
)

//GinSessionWare ...

const (
	ginSessionKey = "_gin_session_id"
	//GinContextSessionKey ...
	GinContextSessionKey = "_gin_context_session_key"
)

//GinSessionWare ...
func GinSessionWare() gin.HandlerFunc {
	if manager == nil {
		panic("manger is not init!")
	}
	return func(c *gin.Context) {
		//session
		var session *Data

		//1.Cookie 取值
		cookie, err := c.Cookie(ginSessionKey)
		if err != nil {
			//没有cookie，是新用户
			session = manager.Create()
		}
		//有值，则判断cookie是否存在
		session, err = manager.GetSession(cookie)
		if err != nil {
			//session不存在
			session = manager.Create()
		}
		c.SetCookie(ginSessionKey, session.ID, 36000, "/", c.Request.URL.Hostname(), false, true)
		c.Set("session", &session)
		c.Next()

	}
}
