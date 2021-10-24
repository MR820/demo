/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/22
 * Time 10:07 上午
 */

package config

import (
	"encoding/json"
	"os"
)

type Database struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}

type Config struct {
	Database *Database `json:"database`
}

var GlobalConfigSetting = &Config{}

func init() {
	filePtr, err := os.Open("/Users/xingzhiwei/go/src/learngo/gin-gorm/config/bms.json")
	if err != nil {
		panic(err.Error())
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(GlobalConfigSetting)
	DatabaseSetting = GlobalConfigSetting.Database
}

var DatabaseSetting = &Database{}
