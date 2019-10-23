package my_calender_iii

import (
	"container/list"
	"fmt"
	"math"
)

// 思路：
// 把原题目想象成在一个数轴上画线段，如果线段有重合，则重合的部分颜色深度加一。
// 每一次book就是画线段。

// point 表示所画线段的点（起始或结束）
type point struct {
	index int // 该点在数轴上的位置。
	deep  int // 该点之后的颜色深度。
}

// String 输出点的信息。
func (p point) String() string {
	return fmt.Sprintf("[%d %d]", p.index, p.deep)
}

// MyCalendarThree 定义了Calendar。
type MyCalendarThree struct {
	points *list.List
}

// Constructor 创建Calendar。
func Constructor() MyCalendarThree {
	ret := MyCalendarThree{points: list.New()}

	// 预制2个点，一个0点，一个最大点。
	ret.points.PushFront(&point{index: 0, deep: 0})
	ret.points.PushBack(&point{index: math.MaxInt64, deep: 0})
	return ret
}

// Book 订阅。
func (c *MyCalendarThree) Book(start int, end int) int {
	var (
		startElement, endElement *list.Element
	)
	// 插入起始点
	for e := c.points.Front(); e != nil; e = e.Next() {
		p := e.Value.(*point)
		// 避免重复，如果该点已经存在了就不新建了。
		if start == p.index {
			startElement = e
			break
		}
		// 新建一个点，注意新建点的颜色深度 暂时 和它前面的点的颜色深度一致。
		if start > p.index && start < e.Next().Value.(*point).index {
			startElement = c.points.InsertAfter(&point{index: start, deep: p.deep}, e)
			break
		}
	}

	// 插入结束点。
	for e := c.points.Back(); e != nil; e = e.Prev() {
		p := e.Value.(*point)
		if end == p.index {
			endElement = e
			break
		}
		if end < p.index && end > e.Prev().Value.(*point).index {
			endElement = c.points.InsertBefore(&point{index: end, deep: e.Prev().Value.(*point).deep}, e)
			break
		}
	}

	// 对于起始和结束点之间的所有点，颜色深度都加一。
	for e := startElement; e != endElement; e = e.Next() {
		p := e.Value.(*point)
		p.deep++
	}

	// find out the max
	k := 0
	for e := c.points.Front(); e != nil; e = e.Next() {
		if e.Value.(*point).deep > k {
			k = e.Value.(*point).deep
		}
	}
	return k
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
