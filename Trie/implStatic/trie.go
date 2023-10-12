package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* 静态数组实现 推荐 */

const Max = 150001

// 前缀树上的每个节点
var tree = [Max][26]int{}

// 每个节点的经过次数
var pass = [Max]int{}

// 每个节点作为结尾的次数
var end = [Max]int{}

// 已经使用的位置
var cnt = 1

// Insert 插入单词
func Insert(word string) {
	cur := 1
	pass[cur]++
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		// 查找路径
		if tree[cur][path] == 0 {
			// 路径还不存在
			tree[cur][path] = cnt + 1
			cnt++
		}
		cur = tree[cur][path]
		pass[cur]++
	}
	end[cur]++
}

// Search 查找单词是否在前缀树中
func Search(word string) bool {
	cur := 1
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		// 查找路径
		if tree[cur][path] == 0 {
			// 路径还不存在
			return false
		}
		cur = tree[cur][path]
	}
	return end[cur] > 0
}

// 查看指定前缀出现的次数
func countWordsStartingWith(prefix string) int {
	cur := 1
	for i := 0; i < len(prefix); i++ {
		path := prefix[i] - 'a'
		// 查找路径
		if tree[cur][path] == 0 {
			// 路径还不存在
			return 0
		}
		cur = tree[cur][path]
	}
	return pass[cur]
}

// Delete 删除单词
func Delete(word string) {
	if !Search(word) {
		return
	}
	cur := 1
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		// 查找路径
		pass[tree[cur][path]]--
		if pass[tree[cur][path]] == 0 {
			tree[cur][path] = 0
			return
		}
		cur = tree[cur][path]
	}
	end[cur]--
}

// 清理静态空间
func clear() {
	for i := 1; i <= cnt; i++ {
		tree[i] = [26]int{}
	}

}

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func read() (byte, string) {
	op, _ := in.ReadByte()
	_, _ = in.ReadByte()
	word := make([]byte, 0, 20)
	for c, _ := in.ReadByte(); c >= 'a'; c, _ = in.ReadByte() {
		word = append(word, c)
	}
	return op, string(word)
}

var m int
var op byte
var word string

func main() {
	fmt.Scanln(&m)
	for i := 0; i < m; i++ {
		op, word = read()
		switch op {
		case '1':
			Insert(word)
		case '2':
			Delete(word)
		case '3':
			if Search(word) {
				out.WriteString("YES\n")
			} else {
				out.WriteString("NO\n")
			}
		case '4':
			out.WriteString(strconv.Itoa(countWordsStartingWith(word)) + "\n")
		}
	}
	out.Flush()
	clear()

}
