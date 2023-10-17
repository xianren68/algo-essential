package wordSearchII

const MAX = 10001

var trie = [MAX][26]int{}
var pass = [MAX]int{}

// 记录每个单词
var end = [MAX]string{}
var cnt int

// Insert 将单词加入前缀树
func Insert(word string) {
	cur := 1
	for i := 0; i < len(word); i++ {
		pass[cur]++
		path := word[i] - 'a'
		if trie[cur][path] == 0 {
			trie[cur][path] = cnt + 1
			cnt++
		}
		cur = trie[cur][path]
	}
	pass[cur]++
	end[cur] = word
}

// 构建前缀树
func build(words []string) {
	cnt = 1
	for _, val := range words {
		Insert(val)
	}
}
func findWords(board [][]byte, words []string) []string {
	build(words)
	res := make([]string, 0, len(words))
	for i, val := range board {
		for j, _ := range val {
			res, _ = dfs(board, res, i, j, 1)
		}
	}
	clear()
	return res
}

// 递归上下左右
// t前缀树所在节点
func dfs(board [][]byte, res []string, i, j, t int) ([]string, int) {
	// 判断越界或当前值已被访问
	if i == len(board) || i < 0 || j == len(board[i]) || j < 0 || board[i][j] == 0 {
		return res, 0
	}
	// 记录当前值，便于后续还原
	tmp := board[i][j]
	// 判断当前值是否在前缀树中
	path := tmp - 'a'
	t = trie[t][path]
	// if pass[t] == 0 || trie[t][path] == 0
	// 前者表示这个单词已经被访问过，后者表示没有当前路
	// 可以合并，因为pass[0] == 0
	if pass[t] == 0 {
		return res, 0
	}
	// 记录有多少单词从当前节点找出
	fix := 0
	if end[t] != "" {
		res = append(res, end[t])
		end[t] = ""
		fix++
	}
	var f int
	board[i][j] = 0
	res, f = dfs(board, res, i+1, j, t)
	fix += f
	res, f = dfs(board, res, i-1, j, t)
	fix += f
	res, f = dfs(board, res, i, j+1, t)
	fix += f
	res, f = dfs(board, res, i, j-1, t)
	fix += f
	// 防止已经找过的单词再找一次
	pass[t] -= fix
	// 还原单词
	board[i][j] = tmp
	return res, fix
}
func clear() {
	for i := 1; i <= cnt; i++ {
		trie[i] = [26]int{}
		pass[i] = 0
		end[i] = ""
	}
}
