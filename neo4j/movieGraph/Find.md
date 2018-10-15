# 找到名字为"Tom Hanks"的演员
```
MATCH (tom {name: "Tom Hanks"}) RETURN tom
```

# 找到标题为"Cloud Atlas"的电影
```
MATCH (cloudAtlas {title: "Cloud Atlas"}) RETURN cloudAtlas
```

# 找10个人
```
MATCH (people:Person) RETURN people.name LIMIT 10
```

# 找到1990s的电影
```
MATCH (nineties:Movie) WHERE nineties.released >= 1990 AND nineties.released < 2000 RETURN nineties.title
```

