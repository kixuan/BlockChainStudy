package main 
import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "log"
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

    r.GET("/buy", func(c *gin.Context) {
        //获取接收到的参数
    issuer:=c.Query("issuer")
    paperNumber:=c.Query("papernumber")
    currentOwner:=c.Query("currentOwner")   
    newOwner:=c.Query("newOwner")   
    price:=c.Query("price")
    purchaseDateTime:=c.Query("purchaseDateTime")    
    log.Printf("issuer:%v",issuer)    
    log.Printf("paperNumber:%v",paperNumber)    
    log.Printf("currentOwner:%v",currentOwner)    
    log.Printf("newOwner:%v",newOwner)    
    log.Printf("price:%v",price)    
    log.Printf("purchaseDateTime:%v",purchaseDateTime)    

        code, err := buy(issuer, paperNumber, currentOwner, newOwner, price,purchaseDateTime)           
        message := ""
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
        if code == 0 {
            message = "购买成功!"  
        }      
        c.JSON(200, gin.H{
            "errcode": code,
            "message": message,
        })
    })

    r.GET("/redeem", func(c *gin.Context) {
        //获取接收到的参数
    issuer:=c.Query("issuer")
    paperNumber:=c.Query("paperNumber")
    redeemingOwner:=c.Query("redeemingOwner")   
    redeenDateTime:=c.Query("redeenDateTime")    
    log.Printf("issuer:%v",issuer)    
    log.Printf("paperNumber:%v",paperNumber)    
    log.Printf("redeemingOwner:%v",redeemingOwner)    
    log.Printf("redeenDateTime:%v",redeenDateTime)    
 

        code, err := redeem(issuer, paperNumber, redeemingOwner, redeenDateTime)
        message := ""
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
        if code == 0 {
            message = "兑换成功!"  
        }      
        c.JSON(200, gin.H{
            "errcode": code,
            "message": message,
        })
    })
    r.Run("0.0.0.0:8081") 
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