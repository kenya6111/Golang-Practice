package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionSet(c *gin.Context, userID uint64) {
	session := sessions.Default(c)
	// var idInterface interface{} = userID
	session.Set("id", userID)
	session.Save()
}
func SessionGet(c *gin.Context, userID uint64) uint64 {

	session := sessions.Default(c)
	return session.Get("id").(uint64)
}
func SessionClear(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

}
