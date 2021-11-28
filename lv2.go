package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func Middleare() gin.HandlerFunc {
	return func(c *gin.Context) {
		flag:=false
		cs := c.Request.Cookies()
		for i:=0;i<len(cs);i++{
			if cs[i].Value=="cookied" {
				c.JSON(http.StatusOK,gin.H{
					"code":http.StatusOK,
					"message":cs[i].Name+"你好！",
				})
				flag=true
			}
		}
		if flag==false {
			c.JSON(http.StatusOK, gin.H{
				"code":http.StatusOK,
				"message":"游客你好！",
			})
		}
	}
}
func main() {
	router:=gin.Default()
	router.GET("/login", func(c *gin.Context) {
		name:= c.Query("username")
		_= c.Query("password")
		cookie := &http.Cookie{
			Name: name,
			Value: "cookied",
			Path: "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.String(http.StatusOK, "Login successful")
	})
	Middleare()
	router.GET("/test",Middleare())//测试所有已注册用户接口
	router.Run(":8080")
}

