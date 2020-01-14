package router

import (
	"Lv/internal/bootstrap"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"strings"
)

func RouteProvider(r *gin.Engine) {
	LoadAPIRouter(r)

	if !strings.Contains(bootstrap.GetAppConf().App.Env, "prod") {
		url := ginSwagger.URL("http://localhost:8000/swagger/swagger.json")
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,url))
	}
}
