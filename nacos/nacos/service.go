/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/1
 * Time 下午3:55
 */

package nacos

import (
	"log"
	"net"
	"net/http"

	"github.com/nacos-group/nacos-sdk-go/util"

	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"

	"github.com/gin-gonic/gin"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var namingClient naming_client.INamingClient

func init() {
	clientConnfig := constant.ClientConfig{
		TimeoutMs:           5000, // 超时时间
		NamespaceId:         "95d1f5d2-cc16-4b65-86e0-9ac0357d623d",
		NotLoadCacheAtStart: true,
		CacheDir:            "./cache",
		LogDir:              "./log",
		RotateTime:          "1h",    // 日志轮转周期,比如：30m, 1h, 24h, 默认是24h
		MaxAge:              3,       // 日志最大文件数，默认3
		LogLevel:            "debug", // 日志默认级别，值必须是：debug,info,warn,error，默认值是info
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "172.16.213.134",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	client, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConnfig,
	})
	if err != nil {
		panic(err)
	}
	namingClient = client
}

func GetIntranetIp() []string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(err)
	}

	ips := make([]string, 0)

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return ips
}

func RegisterInstance(c *gin.Context) {
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "172.16.231.1",
		Port:        8080,
		ServiceName: "user-service",
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a",
		GroupName:   "group-a",
	})
	if err != nil {
		panic(err)
	}
	//c.Next()
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": success, "msg": "register"})
}

func DeregisterInstance(c *gin.Context) {
	success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "172.16.231.1",
		Port:        8080,
		ServiceName: "user-service",
		Ephemeral:   true,
		Cluster:     "cluster-a",
		GroupName:   "group-a",
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": success, "msg": "deregister"})
}

func GetService(c *gin.Context) {
	services, err := namingClient.GetService(vo.GetServiceParam{
		Clusters:    []string{"cluster-a"},
		ServiceName: "user-service",
		GroupName:   "group-a",
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": services})
}

// SelectAllInstance可以返回全部实例列表,包括healthy=false,enable=false,weight<=0
func SelectAllInstances(c *gin.Context) {
	instances, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		Clusters:    []string{"cluster-a"},
		ServiceName: "user-service",
		GroupName:   "group_a",
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": instances})
}

// SelectInstances 只返回满足这些条件的实例列表：healthy=${HealthyOnly},enable=true 和weight>0
func SelectInstances(c *gin.Context) {
	instances, err := namingClient.SelectInstances(vo.SelectInstancesParam{
		Clusters:    []string{"cluster-a"},
		ServiceName: "user-service",
		GroupName:   "group-a",
		HealthyOnly: true,
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": instances})
}

// SelectOneHealthyInstance将会按加权随机轮询的负载均衡策略返回一个健康的实例
// 实例必须满足的条件：health=true,enable=true and weight>0
func SelectOneHealthyInstance(c *gin.Context) {
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		Clusters:    []string{"cluster-a"},
		ServiceName: "user-service",
		GroupName:   "group-a",
	})
	if err != nil {
		panic(err)
	}
	c.Set("serviceIp", instance.Ip)
	c.Set("servicePort", instance.Port)
	c.Next()
	//c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": instance})
}

// Subscribe key=serviceName+groupName+cluster
// 注意:我们可以在相同的key添加多个SubscribeCallback.
func Subscribe(c *gin.Context) {
	err := namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: "user-service",
		GroupName:   "group-a",
		Clusters:    []string{"cluster-a"},
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			if err != nil {
				panic(err)
			}
			//fmt.Println(services)
			log.Printf("\n\n subscribe callback return services:%s \n\n", util.ToJsonString(services))
		},
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func Unsubscribe(c *gin.Context) {
	err := namingClient.Unsubscribe(&vo.SubscribeParam{
		ServiceName: "user-service",
		Clusters:    []string{"cluster-a"},
		GroupName:   "group-a",
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			if err != nil {
				panic(err)
			}
			log.Printf("\n\n unsubscribe callback return services:%s \n\n", util.ToJsonString(services))
		},
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func GetAllServicesInfo(c *gin.Context) {
	serviceInfos, err := namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: "95d1f5d2-cc16-4b65-86e0-9ac0357d623d",
		GroupName: "group-a",
		PageNo:    1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": serviceInfos})
}
