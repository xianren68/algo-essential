package unionFind

// 牛客版并查集（包含小挂大）
type unionFind struct {
	// 指向前面的节点(索引代表当前值，值代表父节点指针)
	father []int
	// 每个并查集的大小
	size []int
	// 记录寻找代表节点时沿途遇到的节点（帮助扁平化）
	stack []int
}

func NewUnionFind(l int) *unionFind {
	res := &unionFind{
		father: make([]int, l),
		size:   make([]int, l),
	}
	for i := 0; i < l; i++ {
		res.father[i] = i
		res.size[i] = 1
	}
	return res
}

// 查找某个值所在集合的代表节点
func (n *unionFind) find(value int) int {
	for n.father[value] != value {
		// 将沿途节点添加到栈中
		n.stack = append(n.stack, value)
		// 进入父节点
		value = n.father[value]
	}
	// 扁平化
	for i := 0; i < len(n.stack); i++ {
		n.father[n.stack[i]] = value
	}
	// 将栈置空
	n.stack = []int{}
	return value
}

// 判断两个值是否在同一集合
func (n *unionFind) isSameSet(v1, v2 int) bool {
	return n.find(v1) == n.find(v2)
}

// 合并两个元素所在的集合
func (n *unionFind) union(v1, v2 int) {
	f1, f2 := n.find(v1), n.find(v2)
	// 小挂大
	if n.size[f1] > n.size[f2] {
		n.father[f2] = f1
	} else {
		n.father[f1] = f2
	}

}
