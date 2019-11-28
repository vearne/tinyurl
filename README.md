* [中文 README](https://github.com/vearne/tinyurl/blob/master/README_zh.md)

## Overview
This is a simple, available short-address service.

[Online Test](http://tool.vearne.cc/#/tinyurl)

## Quick Start
### 1. configuration database 
```
CREATE database tinyurl;
USE tinyurl;
CREATE TABLE `tinyurl` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `url` varchar(255) DEFAULT NULL COMMENT 'URL',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
```
### 2.  develop and debug

#### 2.1  Modify configuration file 
```
vim ./config/config.yaml
```
**Notice** Please modify the domainand database address

#### 2.2 start
```
go run main.go web
```

#### 2.3 test
1）generate short address
```
curl -XPOST http://localhost:8080/api/tinyurl -d '{"url":"http://vearne.cc/archives/39217"}'
```
2）watch result
```
curl -iv http://st.vearne.cc/2h7
curl -L http://st.vearne.cc/2h7
```
You can also test directly in a browse.

![seq chart](https://raw.githubusercontent.com/vearne/tinyurl/master/seq.png)

### 3. Production Environment
#### 3.1 compile
```
go build -ldflags "-s -w" -o tinyurl
```
#### 3.2 start
```
./tinyurl web --config ./config/config.yaml
```

### 4. Thanks
The project inspired by  [国内有哪些靠谱的短链接服务？](https://www.zhihu.com/question/20188969) 



