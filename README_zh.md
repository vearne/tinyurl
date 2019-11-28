* [English README](https://github.com/vearne/tinyurl/blob/master/README.md)
## 概览
这是一个简单可用的短地址服务。 

[在线测试](http://tool.vearne.cc/#/tinyurl)

## 快速开始
### 1. 创建数据库
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
### 2. 开发调试

#### 2.1 修改配置文件
```
vim ./config/config.yaml
```
修改域名和数据库地址

#### 2.2 启动服务
```
go run main.go web
```

#### 2.3 测试
1）生成短地址
```
curl -XPOST http://localhost:8080/api/tinyurl -d '{"url":"http://vearne.cc/archives/39217"}'
```
2）测试效果
```
curl -iv http://st.vearne.cc/2h7
curl -L http://st.vearne.cc/2h7
```
也可以直接在浏览器中测试

![seq chart](https://raw.githubusercontent.com/vearne/tinyurl/master/seq.png)

### 3. 生产环境
#### 3.1 编译
```
go build -ldflags "-s -w" -o tinyurl
```
#### 3.2 启动服务
```
./tinyurl web --config ./config/config.yaml
```

### 4. 感谢
代码实现受到了 [国内有哪些靠谱的短链接服务？](https://www.zhihu.com/question/20188969)的启发。

### 参考资料
1. [国内有哪些靠谱的短链接服务？](https://www.zhihu.com/question/20188969)

