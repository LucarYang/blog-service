package routers

import (
	_ "blog-service/docs"
	"blog-service/global"
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api"
	v1 "blog-service/internal/routers/api/v1"
	"blog-service/pkg/limiter"
	"blog-service/pkg/zapLogger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

/*
路由
*/

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	if global.ServerSetting.RunMode == "debug1" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		//r.Use(middleware.AccessLog())
		//r.Use(middleware.Recovery())
		//middleware.InitLogger(middleware.LogConfig{
		//	Filename: "",
		//
		//})
		//r.Use(middleware.GinLogger())
		//r.Use(middleware.GinRecovery(true))
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations()) //注册validator国际化处理中间件

	r.Use(zapLogger.GinLogger(), zapLogger.GinRecovery(true)) //注册zap中间件
	r.Use(middleware.Tracing())

	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.POST("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", v1.Tag{}.Create)
		apiv1.DELETE("/tags/:id", v1.Tag{}.Delete)
		apiv1.PUT("/tags/:id", v1.Tag{}.Update)
		apiv1.PATCH("/tags/:id/state", v1.Tag{}.Update)
		apiv1.GET("/tags", v1.Tag{}.List)

		apiv1.POST("/articles", v1.Article{}.Create)
		apiv1.DELETE("/articles/:id", v1.Article{}.Delete)
		apiv1.PUT("/articles/:id", v1.Article{}.Update)
		apiv1.PATCH("/articles/:id/state", v1.Article{}.Update)
		apiv1.GET("/articles/:id", v1.Article{}.Get)
		apiv1.GET("/articles", v1.Article{}.List)
	}
	return r
}
