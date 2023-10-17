package __implArray

/* 类的方式实现前缀树 以数组存储下级路径 */

type Trie struct {
	root *TrieNode
}
type TrieNode struct {
	// 路过当前字符的次数
	pass int
	// 以当前字符作为结尾的次数
	end int
	// 下级节点的路径
	nexts [26]*TrieNode
}

// 构造前缀树

func Constructor() *Trie {
	trie := &Trie{
		&TrieNode{},
	}
	return trie
}

// 插入单词

func (t *Trie) Insert(word string) {
	node := t.root
	node.pass++
	for i := 0; i < len(word); i++ {
		// 下级路径索引
		path := word[i] - 'a'
		if node.nexts[path] == nil {
			// 新建
			newNode := &TrieNode{}
			node.nexts[path] = newNode
		}
		node = node.nexts[path]
		node.pass++
	}
	node.end++
}

// 搜索单词是否在前缀树中

func (t *Trie) Search(word string) bool {
	node := t.root
	for i := 0; i < len(word); i++ {
		// 下级路径索引
		path := word[i] - 'a'
		if node.nexts[path] == nil {
			// 不存在此路径
			return false
		}
		node = node.nexts[path]
	}
	return node.end > 0

}

// 判断前面的字符串中是否有以此为前缀的

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		// 下级路径索引
		path := prefix[i] - 'a'
		if node.nexts[path] == nil {
			// 不存在此路径
			return false
		}
		node = node.nexts[path]
	}
	return true
}

// 查找单词在前缀树中出现多少次

func (t *Trie) countWordsEqualTo(word string) int {
	node := t.root
	for i := 0; i < len(word); i++ {
		// 下级路径索引
		path := word[i] - 'a'
		if node.nexts[path] == nil {
			// 不存在此路径
			return 0
		}
		node = node.nexts[path]
	}
	return node.end
}

// 查看指定前缀在前缀树中出现了几次

func (t *Trie) countWordsStartingWith(prefix string) int {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		// 下级路径索引
		path := prefix[i] - 'a'
		if node.nexts[path] == nil {
			// 不存在此路径
			return 0
		}
		node = node.nexts[path]
	}
	return node.pass
}

// 删除单词

func (t *Trie) Delete(word string) {
	// 前缀树中不存在
	if !t.Search(word) {
		return
	}
	node := t.root
	node.pass--
	for i := 0; i < len(word); i++ {
		// 下级路径索引
		path := word[i] - 'a'
		// 当前节点删除后往后节点无用了
		if node.nexts[path].pass == 1 {
			node.nexts[path] = nil
			return
		}
		node = node.nexts[path]
		node.pass--
	}
	node.end--
}
