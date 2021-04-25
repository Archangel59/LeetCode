## 线段树

```
用处: 对一个区间点进行 update 或者 query 操作，时间复杂度都是O(log2(n)).
  
原理: 将[1,n]分解成若干特定的子区间(数量不超过4*n),然后，将每个区间[L,R]都分解为少量特定的子区间  
     通过对这些少量子区间的修改或者统计，来实现快速对[L,R]的修改或者统计.
```

###Code
[cpp](https://github.com/Archangel59/LeetCode/blob/main/algorithm/segment-tree/segment-tree.cpp)