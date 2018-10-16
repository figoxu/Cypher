#Neo4j Praticse
## Install On MacOs
```
brew install neo4j
brew services list
brew services start neo4j
brew services stop neo4j
```
默认安装的是3.4.5版本

## 实战
### 资料地址
http://neo4j.com.cn/public/cypher/neo4j_cql_create_node.html

### Neo4j地址
http://localhost:7474

连接地址：bolt://localhost:7687
默认账号:neo4j
默认密码：neo4j
开发密码: 1234567890 (修改后)

输入 :help 在网站顶部的输入框里，可以得到命令帮助

### 官方范例
[  数据插入 ](./movieGraph/Insert.md)
[  查找 ](./movieGraph/Find.md)
[  搜索 ](./movieGraph/Query.md)
[  解题 ](./movieGraph/Solve.md)
[  推荐 ](./movieGraph/Recommend.md)
[  清理数据 ](./movieGraph/Del.md)

### Cypher文档
[ Neo4j开发文档 ](https://neo4j.com/docs/developer-manual/current/cypher/#cypher-intro)