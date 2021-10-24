/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/10
 * Time 上午9:42
 */

package ziface

/*
封包数据和拆包数据
*/
type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
