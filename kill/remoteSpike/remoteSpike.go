/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/30
 * Time 9:54 上午
 */

package remoteSpike

import "github.com/gomodule/redigo/redis"

type RemoteSpikeKeys struct {
	SpikeOrderHashKey  string
	TotalInventoryKey  string
	QuantityOfOrderKey string
}

const LuaScript = `
	local ticket_key = KEYS[1]
	local ticket_total_key = ARGV[1]
	local ticket_sold_key = ARGV[2]
	local ticket_total_nums = tonumber(redis.call('HGET',ticket_key,ticket_total_key))
	local ticket_sold_nums = tonumber(redis.call('HGET',ticket_key,ticket_sole_key))
	-- 查看是否还有余票,增加订单数量,返回结果值
	if(ticket_total_nums >= ticket_sold_nums) then
		return redis.call('HINCRBY',ticket_key,ticket_sold_key,1)
	end
	return 0
`

func (RemoteSpikeKeys *RemoteSpikeKeys) RemoteDeductionStock(conn redis.Conn) bool {
	lua := redis.NewScript(1, LuaScript)
	result, err := redis.Int(lua.Do(conn, RemoteSpikeKeys.SpikeOrderHashKey, RemoteSpikeKeys.TotalInventoryKey, RemoteSpikeKeys.QuantityOfOrderKey))
	if err != nil {
		return false
	}
	return result != 0
}
