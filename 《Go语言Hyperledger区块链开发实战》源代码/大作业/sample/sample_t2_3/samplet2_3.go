package main 
import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
 )

func main() {
    r := gin.Default()
    r.Use(CorsMiddleware()) //解决跨域
    r.POST("/adduser", func(c *gin.Context) {

     json := make(map[string]interface{}) //注意该结构接受的内容
     c.ShouldBind(&json)
     log.Printf("%v",&json)
        username := json["username"]
        pwd := json["pwd"]

        c.JSON(200, gin.H{
            "username":username, "pwd":pwd,
        })
    })
    r.Run() 
}

// CorsMiddleware 解决跨域的方法
func CorsMiddleware() gin.HandlerFunc {
    return func(context *gin.Context) {
        method := context.Request.Method
        
        context.Header("Access-Control-Allow-Origin", "*")
        context.Header("Access-Control-Allow-Credentials", "true")
        context.Header("Access-Control-Allow-Headers", "*")
        context.Header("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,DELETE,OPTIONS")
        context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
        
        if method == "OPTIONS" {
            context.AbortWithStatus(http.StatusNoContent)
        }
        context.Next()
    }
}


