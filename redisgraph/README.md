#Redis Graph Praticse
## 通过Docker来运行Redis Graph
```
docker pull redislabs/redisgraph
brew services stop redis
docker run -p 6379:6379 -it --rm redislabs/redisgraph
```

## Cypher
### 增加节点和关系
```
GRAPH.QUERY MotoGP "CREATE (:Rider {name:'Valentino Rossi'})-[:rides]->(:Team {name:'Yamaha'}), (:Rider {name:'Dani Pedrosa'})-[:rides]->(:Team {name:'Honda'}), (:Rider {name:'Andrea Dovizioso'})-[:rides]->(:Team {name:'Ducati'})"
```
### 查询骑过“Yamaha”的骑手和团队
```
GRAPH.QUERY MotoGP "MATCH (r:Rider)-[:rides]->(t:Team) WHERE t.name = 'Yamaha' RETURN r,t"
```
### 查询“Ducati”下的骑手
```
 GRAPH.QUERY MotoGP "MATCH (r:Rider)-[:rides]->(t:Team {name:'Ducati'}) RETURN count(r)"
```

## 关停Docker服务
```
docker ps
docker stop a95f089bc9c6
```
