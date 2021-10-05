## 思路
1. 收集窗口 [0,k-1) 的信息
2. right 从 k-1 开始，将 right 加入窗口
3. 判断是否要将 left 移出窗口
4. 取当前窗口的值

## Template

````
for right := k-1; right < n; right++ {
    addRight()
    removeLeft()
    calculateWindow()
}
````