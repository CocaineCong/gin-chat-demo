# gin-chat-demo
gin+websocket+mongodb实现 IM 即时聊天系统

> 
> 这个项目是基于WebSocket + MongoDB + MySQL + Redis。
> 业务逻辑很简单，只是两人的聊天。

- `MySQL` 用来存储用户基本信息
- `MongoDB` 用来存放用户聊天信息
- `Redis` 用来存储处理过期信息

# 项目教程
B站：https://www.bilibili.com/video/BV1BP4y1H7gV

博客：https://blog.csdn.net/weixin_45304503/article/details/121787022


# 项目结构

```
gin-chat-demo/
├── cache
├── conf
├── e
├── model
├── router
└── service
```

- cache : 放置redis配置
- conf : 放置配置文件 
- model : 数据库模型
- pkg : 防止一些错误码
- router ： 路由模块
- service ：服务模块

# 项目功能

- 两人通信
- 在线、不在线应答
- 查看历史聊天记录

# 配置文件
- conf/config.ini

```ini
#debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000 
# 运行端口号 3000端口

[mysql]
Db = mysql
DbHost = "" 
# mysql的ip地址
DbPort = ""
# mysql的端口号,默认3306
DbUser = ""
# mysql user
DbPassWord = ""
# mysql password
DbName = ""
# 数据库名字

[redis]
RedisDb = ""
# redis 名字
RedisAddr = ""
# redis 地址
RedisPw = ""
# redis 密码
RedisDbName = ""
# redis 数据库名

[MongoDB]
MongoDBName =  ""
MongoDBAddr = ""
MongoDBPwd = ""
MongoDBPort = ""
```

# 项目运行

- 下载依赖

```go
go mod tidy
```

- 执行

```go
go run main.go
```

# 演示
- 测试http连接

![在这里插入图片描述](https://img-blog.csdnimg.cn/7ddfc3253d3d4df48460eea2772ede60.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5bCP55Sf5Yeh5LiA,size_20,color_FFFFFF,t_70,g_se,x_16)
- 进行ws连接，连接服务器

![在这里插入图片描述](https://img-blog.csdnimg.cn/cb69e11c10d341abb4a421d304c3c6e1.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5bCP55Sf5Yeh5LiA,size_20,color_FFFFFF,t_70,g_se,x_16)


- 当id=1上线，但是id=2没上线的时候发送消息

![在这里插入图片描述](https://img-blog.csdnimg.cn/b52f92d1437b4d4891c47c2187ffdd65.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5bCP55Sf5Yeh5LiA,size_20,color_FFFFFF,t_70,g_se,x_16)

- 当id=2上线之后

![在这里插入图片描述](https://img-blog.csdnimg.cn/c2a2fd17956846d6be96745e452987fd.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5bCP55Sf5Yeh5LiA,size_20,color_FFFFFF,t_70,g_se,x_16)

- 再次发消息，就是在线应答了

![在这里插入图片描述](https://img-blog.csdnimg.cn/715dee6ffb224f77a788a80e17539cb3.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5bCP55Sf5Yeh5LiA,size_20,color_FFFFFF,t_70,g_se,x_16)

- 这边就实时接受到消息了

![在这里插入图片描述](https://img-blog.csdnimg.cn/4f0495ef968940a28b49e0c5b3e9f346.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5bCP55Sf5Yeh5LiA,size_20,color_FFFFFF,t_70,g_se,x_16)

- 获取历史信息

![在这里插入图片描述](https://img-blog.csdnimg.cn/2e88c30fa5ce47ce94b681a5ffeafb88.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5bCP55Sf5Yeh5LiA,size_20,color_FFFFFF,t_70,g_se,x_16)