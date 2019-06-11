package router

import (
	"config"
	"model"
	"github.com/gin-gonic/gin"
)

var router =gin.Default()
func RouterInit() {

	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")

	//api接口设置router
	router_api:=configMap["router_api"]
	v1 :=router.Group(router_api)
	{
		v1.POST("/", model.PostDepwit)
	}

	//路径设置
	http_port:=configMap["http_port"]

	router.Run(http_port)
}
