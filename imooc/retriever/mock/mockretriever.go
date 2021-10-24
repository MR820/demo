/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/12
 * Time 上午10:07
 */

package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
