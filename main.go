package main

import (
	"svc-user/services"
)

// @title           User system documentation
// @version         1.0
// @license.name  	Apache 2.0
// @description     Management和Initialized类为管理后台接口，普通用户无权访问。

// @host      192.168.88.56:81/
// @BasePath  api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	services.App.Run()
}
