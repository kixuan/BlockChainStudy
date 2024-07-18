package main 
import (
  "github.com/gin-gonic/gin"
  "net/http"
)
 
func main() {
    r := gin.Default()
    r.Use(CorsMiddleware()) //解决跨域
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "hello world",
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