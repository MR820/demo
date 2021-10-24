/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/10
 * Time 上午10:07
 */

package ziface

/*
   将请求的一个消息封装到message中，定义抽象层接口
*/
type IMessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetData([]byte)
	SetDataLen(uint32)
}
