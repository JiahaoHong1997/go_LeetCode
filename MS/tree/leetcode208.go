/*
 * @Descripttion: 实现 Trie (前缀树)
 Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：
1.Trie() 初始化前缀树对象。
2.void insert(String word) 向前缀树中插入字符串 word 。
3.boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
4.boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
 

示例：

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
 

提示：
1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成
insert、search 和 startsWith 调用次数 总计 不超过 3 * 10^4 次
 * @Solution:
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-15 13:55:43
 * @LastEditors: 洪笳淏
 */
package tree

type Trie struct {
	layer     int		// 记录当前层级
	isWord    bool		// 记录该节点是否能够形成单词
	letterMap []*Trie	// 子节点
}

func Constructor() Trie {

	m := make([]*Trie, 26)
	return Trie{
		layer:     0,
		isWord:    false,
		letterMap: m,
	}
}

func (this *Trie) Insert(word string) {
	p := this
	for i := 0; i < len(word); i++ {
		x := word[i]
		if p.letterMap[int(x)-'a'] != nil {
			p = p.letterMap[int(x)-'a']
		} else {
			m := make([]*Trie, 26)
			p.letterMap[int(x)-'a'] = &Trie{layer: p.layer + 1, letterMap: m}
			p = p.letterMap[int(x)-'a']
		}
	}
	p.isWord = true
}

func (this *Trie) Search(word string) bool {
	p := this
	for i := 0; i < len(word); i++ {
		x := word[i]
		if p.letterMap[int(x)-'a'] != nil {
			p = p.letterMap[int(x)-'a']
		} else {
			return false
		}
	}

	if p.isWord {
		return true
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for i := 0; i < len(prefix); i++ {
		x := prefix[i]
		if p.letterMap[int(x)-'a'] != nil {
			p = p.letterMap[int(x)-'a']
		} else {
			return false
		}
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
