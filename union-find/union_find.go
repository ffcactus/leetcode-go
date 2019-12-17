package union_find

// union find 描述了这样一类问题：
// 输入是若个数据对，每个数据对中的2个元素由p,q表示。一个数据对p和q是相连的。
// [1,2], [2,3], 则这3个数都在一个域里。
// 求出一个有多少个域

// element 表示union find问题中的数据对中的元素。
type element int

// domain 表示union find问题中的域
type domain int

type unionFind interface {
	// NewUnionFind 初始化一个union find的解决方案，最多有n个域。
	NewUnionFind(n int) unionFind

	// Union 连接元素p, q。
	Union(p, q element)

	// Find 找出元素p所在的domain。
	Find(p int) domain

	// Connected 判断p, q两个元素是否在同一个域中。
	Connected(p, q element) bool

	// Count 输出一共有多少个域。
	Count() int
}

type solutionTemplate struct {
	unionFind
	count int
}

func (impl *solutionTemplate) NewUnionFind(n int) unionFind {

}
