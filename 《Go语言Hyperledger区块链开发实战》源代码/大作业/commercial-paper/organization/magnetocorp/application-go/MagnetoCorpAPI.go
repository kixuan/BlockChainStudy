package main 
import (
    "github.com/gin-gonic/gin"
    "fmt"
    "log"
    "net/http"
)
 
func main() {
    r := gin.Default()
    r.Use(CorsMiddleware()) //解决跨域
    r.GET("/add_to_wallet", func(c *gin.Context) {
        code,err := AddtoWallet()
        message := ""
        if code == -1 {
            message = fmt.Sprintf("设置环境变量报错:%v", err)  
        }
        if code == -2 {
            message = fmt.Sprintf("创建钱包报错:%v", err)  
        }
        if code == -3 {
            message = fmt.Sprintf("构造钱包内容报错:%v", err)  
        }
        if code == -4 {
            message = fmt.Sprintf("读取证书内容报错:%v", err)
        }
        if code == -5 {
            message = fmt.Sprintf("读取私钥的内容报错:%v", err)
        }
        c.JSON(200, gin.H{
            "errcode": code,
            "message": message,
        })
    })

    r.GET("/issue", func(c *gin.Context) {
    //获取接收到的参数
    faceValue:=c.Query("price")
    issuer:=c.Query("issuer")
    paperNumber:=c.Query("papernumber")
    issueDateTime:=c.Query("issuedatetime")   
    maturityDateTime:=c.Query("maturityDateTime")    
    log.Printf("issuer:%v",issuer)    
    log.Printf("issueDateTime:%v",issueDateTime)    
    log.Printf("paperNumber:%v",paperNumber)    
    log.Printf("maturityDateTime:%v",maturityDateTime)    
    log.Printf("faceValue:%v",faceValue)    

    code, err := issue(issuer, paperNumber, issueDateTime, maturityDateTime, faceValue)
    log.Printf("code:%v",code)    

    message := ""
     if code == 0 {
        message = fmt.Sprintf("交易执行成功！")
    }
    if code == -1 {
        message = fmt.Sprintf("创建钱包报错:%v",err)
    }
    if code == -2 {
        message = fmt.Sprintf("连接到网关报错:%$v", err)
    }
    if code == -3 {
        message = fmt.Sprintf("获取通道的Network对象报错:", err);  
    }
    if code == -4 {
        message = fmt.Sprintf("提交交易报错:", err)  
    }        
    log.Printf(message)
    c.JSON(200, gin.H{
        "errcode": code,
        "message": message,
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
