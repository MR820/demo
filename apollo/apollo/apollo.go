/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/26
 * Time 下午2:53
 */

package apollo

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache/memory"
	"github.com/apolloconfig/agollo/v4/component/log"
	"github.com/apolloconfig/agollo/v4/env/config"
)

var client *agollo.Client

func init() {
	// apollo配置
	c := &config.AppConfig{
		AppID:          "G02-go-test",
		Cluster:        "default",
		NamespaceName:  "application",
		IP:             "http://d-apollo.yundasys.com:30224",
		IsBackupConfig: false,
	}
	agollo.SetCache(&memory.DefaultCacheFactory{})
	agollo.SetLogger(&log.DefaultLogger{})
	// 创建客户端
	cli, err := agollo.StartWithConfig(func() (appConfig *config.AppConfig, err error) {
		return c, nil
	})
	if err != nil {
		panic(err)
	}
	client = cli
}

func GetClient() *agollo.Client {
	return client
}
