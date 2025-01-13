package main

import (
	"errors"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetLocalIP() (ipv4 string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet
		isIpNet bool
	)
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}

	for _, addr = range addrs {
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String()
				return
			}
		}
	}
	err = errors.New("No local IP found!")
	return
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("views/*.html")
	engine.GET("/", func(ctx *gin.Context) {
		name := ctx.Query("name")       //クエリパラメータの取得
		project := os.Getenv("PROJECT") //環境変数の取得
		ip, _ := GetLocalIP()
		ctx.HTML(http.StatusOK, "home.html", gin.H{"name": name, "project": project, "ip": ip})
	})
	port := os.Getenv("PORT")
	engine.Run(":" + port)
}

//docker build -t blog:v1.0.0 .
//docker run -p 80:8080 --name blog blog:v1.0.0
