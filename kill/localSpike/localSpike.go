/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/30
 * Time 9:52 上午
 */

package localSpike

type LocalSpike struct {
	LocalInStock     int64
	LocalSalesVolume int64
}

func (spike *LocalSpike) LocalDeductionStock() bool {
	spike.LocalSalesVolume = spike.LocalSalesVolume + 1
	return spike.LocalSalesVolume < spike.LocalInStock
}
