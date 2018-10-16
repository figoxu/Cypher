# 解决问题 ( solve )
## 找到4度（关系/跳）内的电影和演员
```
MATCH (bacon:Person {name:"Kevin Bacon"})-[*1..4]-(hollywood)
RETURN DISTINCT hollywood
```

## 寻找Meg Ryan和 Kevin Bacon的最短路径
```
MATCH p=shortestPath(
  (bacon:Person {name:"Kevin Bacon"})-[*]-(meg:Person {name:"Meg Ryan"})
)
RETURN p
```


