## 题目描述
721. 账户合并

给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元素是 emails 表示该账户的邮箱地址。

现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。

合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是 按字符 ASCII 顺序排列 的邮箱地址。账户本身可以以 任意顺序 返回。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/accounts-merge
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## 思路
1. 要根据条件将数据合并，可以使用并查集两两判断
2. 这里的条件是邮箱相同就合并，我们可以使用一个 map emailToInt 来存放邮箱第一次出现的位置
3. 遍历所有邮箱，如果当前账户 x 的 邮箱已在 emailToInt y 中，则可以将账户 x 和 y 合并
4. 遍历所有邮箱，将 root 相同的邮箱并到同一个组
5. 根据题目要求将每个组的邮箱排序，并在最前面加上账户名
6. 将所有组组合成结果数组

## Template for UnionFind

````
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

func (uf *UnionFind) toGroups(a int) map[int][]int {
   var ans = make(map[int][]int)
   for k, root := range uf.parents {
      ans[root] = append(ans[root], k)
   }
   return ans
}
````