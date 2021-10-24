/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/8/27
 * Time 4:48 下午
 */

package main

import (
	"flag"
	"learngo/api/token"
	"learngo/api/user"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	//_ "github.com/mkevac/debugcharts"
)

// 第一步: 开一个全局变量
//var zkTracer opentracing.Tracer

var (
	listenAddr       = flag.String("web.listen-port", "9000", "An port to listen on for web interface and telemetry.")
	metricsPath      = flag.String("web.telemetry-path", "/metrics", "A path under which to expose metrics.")
	metricsNamespace = flag.String("metric.namespace", "ECSDATA", "Prometheus metrics namespace, as the prefix of metrics name")
)

func main() {

	// 第二步: 初始化 tracer
	//{
	//	reporter := zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
	//	defer reporter.Close()
	//	endpoint, err := zipkin.NewEndpoint("main3", "localhost:8888")
	//	if err != nil {
	//		log.Fatalf("unable to create local endpoint: %+v\n", err)
	//	}
	//	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	//	if err != nil {
	//		log.Fatalf("unable to create tracer: %+v\n", err)
	//	}
	//	zkTracer = zkOt.Wrap(nativeTracer)
	//	opentracing.SetGlobalTracer(zkTracer)
	//}

	//go func() {
	//	flag.Parse()
	//	metrics := metric.NewMetrics(*metricsNamespace)
	//	registry := prometheus.NewRegistry()
	//	registry.MustRegister(metrics)
	//	http.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	//	log.Println(http.ListenAndServe(":9100", nil))
	//}()
	Router()
}

func Router() {
	router := gin.Default()

	// 第三步: 添加一个 middleWare, 为每一个请求添加span
	//router.Use(func(c *gin.Context) {
	//	span := zkTracer.StartSpan(c.FullPath())
	//	defer span.Finish()
	//	c.Next()
	//})

	{
		router.GET("/login", func(c *gin.Context) {
			t, e := token.GenToken("xzw")
			if e != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 1001,
					"msg":  "token创建失败",
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"code":  2000,
				"token": t,
			})
		})

		//router.Use(token.JWTAuthMiddleware())

		router.GET("/", func(c *gin.Context) {
			username, ok := c.Get("username")
			if ok == false {
				c.JSON(http.StatusOK, gin.H{
					"code": 3001,
					"msg":  "username获取失败",
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 3000,
				"msg":  username,
			})
		})
		router.GET("/user", user.InitPage)
		router.GET("/user/list", user.ListUser)
		router.POST("/user/create", user.CreateUser)
		router.GET("/user/detail/:id", user.DetailUser)
		router.PATCH("/user/:id", user.UpdateUser)
		router.DELETE("/user/:id", user.DeleteUser)
		//log.Fatal()
	}

	router.Run(":8888")
}
