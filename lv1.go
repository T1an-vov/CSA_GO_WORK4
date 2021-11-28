package main

import (
	"github.com/gin-gonic/gin"
)
var m = make(map[string]string)//保存已注册的用户
var in = make(map[string]string)//记录信息
func main() {
	r:=gin.Default()
	r.GET("/users/register", func(c *gin.Context) {
		if in,ok:=c.GetQueryMap("reg"); ok==true{
			for x,y:=range in{
				if _,flag:=m[x];flag==true{
					c.String(200,x+" has been registered!")
				}else {
					m[x]=y
					c.String(200,"register success!")
				}
			}//注册用户
		}else {
			c.String(200,"wrong")
		}
	})
	r.GET("/users/login", func(c *gin.Context) {
		if in,ok:=c.GetQueryMap("login"); ok==true{
			for x,y:=range in {
				if _, ok = m[x]; ok == true {
					if m[x] == y {
						c.String(200, x+" login in!")
					} else {
						c.String(200, x+" password wrong!")
					}
				}else {
					c.String(200,"not registered!")
				}
			}//登录用户
		}else {
			c.String(200,"wrong")
		}
	})
	r.Run(":8080")
}

