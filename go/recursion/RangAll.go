package recursion

import "fmt"

//实习一组数据集合的全排列
type RangeType struct {
	val []interface{}
}

func NewRangeArray(n int) *RangeType {
	return &RangeType{
		make([]interface{}, n),
	}
}

func (slice *RangeType) RangeAll(start int) {
	len := len(slice.val)

	if start == len-1 {
		//如果已经是最后位置，直接把数组合并输出
		fmt.Println(slice.val)
	}

	for i := start; i < len; i++ {
		//i = start 时输出自己
		//如果i和start的值相同就没必要交换
		if i == start || slice.val[i] != slice.val[start] {
			//交换当前这个与后面的位置
			slice.val[i], slice.val[start] = slice.val[start], slice.val[i]
			//递归处理索引+1
			slice.RangeAll(start + 1)
			//换回来，因为是递归，如果不换回来会影响后面的操作，并且出现重复
			slice.val[i], slice.val[start] = slice.val[start], slice.val[i]
		}
	}
}
