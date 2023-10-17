package implHash

/* 类的方式实现前缀树 以hash表存储下级路径 */

type Trie struct {
	root *TrieNode
}
type TrieNode struct {
	// 路过当前字符的次数
	pass int
	// 以当前字符作为结尾的次数
	end int
	// 下级节点的路径
	nexts map[byte]*TrieNode
}

// 构造前缀树

func Constructor() *Trie {
	trie := &Trie{
		&TrieNode{
			nexts: make(map[byte]*TrieNode),
		},
	}
	return trie
}

// 插入单词

func (t *Trie) Insert(word string) {
	node := t.root
	node.pass++
	for i := 0; i < len(word); i++ {
		if _, ok := node.nexts[word[i]]; !ok {
			// 新建
			newNode := &TrieNode{nexts: make(map[byte]*TrieNode)}
			node.nexts[word[i]] = newNode
		}
		node = node.nexts[word[i]]
		node.pass++
	}
	node.end++
}

// 搜索单词是否在前缀树中

func (t *Trie) Search(word string) bool {
	node := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := node.nexts[word[i]]; !ok {
			return false
		}
		node = node.nexts[word[i]]
	}
	return node.end > 0

}

// 判断前面的字符串中是否有以此为前缀的

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		if _, ok := node.nexts[prefix[i]]; !ok {
			return false
		}
		node = node.nexts[prefix[i]]
	}
	return true
}

// 查找单词在前缀树中出现多少次

func (t *Trie) countWordsEqualTo(word string) int {
	node := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := node.nexts[word[i]]; !ok {
			return 0
		}
		node = node.nexts[word[i]]
	}
	return node.end
}

// 查看指定前缀在前缀树中出现了几次

func (t *Trie) countWordsStartingWith(prefix string) int {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		if _, ok := node.nexts[prefix[i]]; !ok {
			return 0
		}
		node = node.nexts[prefix[i]]
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
		if node.nexts[word[i]].pass == 1 {
			node.nexts[word[i]] = nil
			return
		}
		node = node.nexts[word[i]]
		node.pass--
	}
	node.end--
}
