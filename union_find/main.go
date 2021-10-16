package union_find

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	uf := NewUnionFind(n)
	emailToInt := make(map[string]int)
	for i := 0; i < n; i++ {
		account := accounts[i]
		for j := 1; j < len(account); j++ {
			email := account[j]
			if v, ok := emailToInt[email]; !ok {
				emailToInt[email] = i
			} else {
				uf.union(v, i)
			}
		}
	}
	var used = make(map[string]bool)
	var intToEmails = make(map[int][]string)

	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			email := account[j]
			if !used[email] {
				used[email] = true
				root := uf.find(i)
				intToEmails[root] = append(intToEmails[root], email)
			}
		}
	}
	var ans [][]string
	for idx, emails := range intToEmails {
		var item []string
		item = append(item, accounts[idx][0])

		sort.Slice(emails, func(i, j int) bool {
			return emails[i] < emails[j]
		})
		item = append(item, emails...)
		ans = append(ans, item)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][1] < ans[j][1]
	})
	return ans
}

type UnionFind struct {
	parents map[int]int
}

func NewUnionFind(n int) *UnionFind {
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = i
	}
	return &UnionFind{
		parents: m,
	}
}

func (uf *UnionFind) union(a1, a2 int) {
	f1 := uf.find(a1)
	f2 := uf.find(a2)
	uf.parents[f2] = f1
}

func (uf *UnionFind) find(a int) int {
	f := uf.parents[a]
	if f != a {
		ff := uf.find(f)
		uf.parents[a] = ff
		return ff
	}
	return f
}
