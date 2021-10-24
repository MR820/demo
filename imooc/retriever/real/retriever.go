/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/12
 * Time 上午10:11
 */

package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}
	return string(result)
}
