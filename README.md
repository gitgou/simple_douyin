# Simple_douyin README

# Simple_douyin README

## 项目架构：
### 服务简介
该项目实现一款极简版抖音应用后台功能，主要采用了微服务架构对服务进行解耦。主要实现了以下功能：
1. 用户注册及登陆鉴权。
2. 视频流推送。
3. 用户点赞及评论，查看评论信息等互动功能。
4. 用户互相关注，用户关系管理等相关功能。
5. 好友间进行即时通信，消息缓存及在线推送等功能。 
   
项目的解耦后的7个微服务功能与技术特征见下表：

| 服务名称      | 服务功能              |  代码框架    |数据传输协议  | 服务路径      |  IDL             |
|--------------|----------------------|-------------|-------------|-------------|-----------------|
| Api          | HTTP interface       | hertz | http   | cmd/api     | ------  |
| User         | 用户信息管理           | kitex/gorm    | protobuf3      | cmd/user    | idl/user.proto |
| Redis        | 缓存热点数据           |  kitex/gorm   | protobuf3      | cmd/redis   | idl/redis.proto |
| Relation     | 管理用户关系 | kitex/gorm  | protobuf3      | cmd/relation    | idl/relaton.proto |
| Chat    | 用户通信服务 | kitex/gorm  | protobuf3      | cmd/chat    | idl/chat.proto |
| Interaction    | 视频互动服务 | kitex/gorm  | protobuf3      | cmd/interaction    | idl/interaction.proto |
| Video   | 视频流推送 | kitex/gorm  | protobuf3      | cmd/video    | idl/video.proto |



![easy-note-arch](./images/easy-note-arch.png)

## 应用技术点：

kitex: 微服务

hertz: API

gorm： 数据库API 

redis 缓存热点数据、减少 I/O

jaeger:  tracer

docker: 环境配置

jwt: 登录认证



## 项目结构：

主要文件如下：

|                             Dir                              |                         Introduction                         |
| :----------------------------------------------------------: | :----------------------------------------------------------: |
| [**cmd**](https://github.com/gitgou/simple_douyin/tree/master/cmd) |                   API 网关、微服务代码文件                   |
| **[data](https://github.com/gitgou/simple_douyin/tree/master/data)** |                   保存用户上传的视频源文件                   |
| [**idl**](https://github.com/gitgou/simple_douyin/tree/master/idl) |    proto、 thrift 文件，用于生成微服务 API、数据传输协议     |
| **[kitex_gen](https://github.com/gitgou/simple_douyin/tree/master/kitex_gen)** |                  kitex 工具生成的微服务 API                  |
| [**minio**](https://github.com/gitgou/simple_douyin/tree/master/minio) |            记录对象存储meta数据 ：视频转化为 url             |
| [**pkg**](https://github.com/gitgou/simple_douyin/tree/master/pkg) | 数据库创建表格 sql 文件、中间件 middleware、trace、常量 constants、错误机制 errno、工具函数 utils |
| [**tools**](https://github.com/gitgou/simple_douyin/tree/master/tools) |    保存项目中用到的第三方工具及配置文件，比如 redis.conf     |
| **[docker-compose.yml](https://github.com/gitgou/simple_douyin/blob/master/docker-compose.yml)** |                       docker 启动文件                        |

![image-20230222133426973](/Users/yinliliang/Library/Application Support/typora-user-images/image-20230222133426973.png)

### 微服务代码结构， 以 [user](https://github.com/gitgou/simple_douyin/tree/master/cmd/user) 微服务为例 （其它微服务结构类似）：

|                             File                             |                         Introduction                         |
| :----------------------------------------------------------: | :----------------------------------------------------------: |
| **[Makefile](https://github.com/gitgou/simple_douyin/blob/master/cmd/user/Makefile)** |                记录 kitex 生成微服务接口命令                 |
| **[cache ](https://github.com/gitgou/simple_douyin/tree/master/cmd/user/cache)** | 登陆用户数据记录于内存，数据定时入库，用户数据修改前先查找cache层，再查数据库曾，减少数据库 I/O次数； |
| **[dal](https://github.com/gitgou/simple_douyin/tree/master/cmd/user/dal)** |              数据库操作接口，用于 user 增删改查              |
| **[handler.go](https://github.com/gitgou/simple_douyin/blob/master/cmd/user/handler.go)** |                        user 服务接口                         |
| **[main.go](https://github.com/gitgou/simple_douyin/blob/master/cmd/user/main.go)** |                    user 服务主函数入口；                     |
| **[pack](https://github.com/gitgou/simple_douyin/tree/master/cmd/user/pack)** |  用于数据打包，转化；比如数据库访问 model 转化为 proto格式   |
| **[rpc](https://github.com/gitgou/simple_douyin/tree/master/cmd/user/rpc)** |                      调用其它微服务接口                      |
| **[service](https://github.com/gitgou/simple_douyin/tree/master/cmd/user/service)** |                         服务逻辑处理                         |

<img src="/Users/yinliliang/Library/Application Support/typora-user-images/image-20230222141214539.png" alt="image-20230222141214539" style="zoom:50%;" />



## Quick Start：

### 1. Setup Basic Dependence:

```sh
docker-compose up # 执行前，注意把 tools 里面的 redis.conf 文件放到 docker-compose.yml 中配置的redis配置文件路径下
```

### 2. Run Redis RPC Service :

```sh
cd cmd/redis
sh build.sh
sh output/bootstrap.sh
```

### 3. Run User RPC Server :

```sh
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 4. Run Video RPC Server:

```sh
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```

### 5. Run Relation RPC Server:

```sh
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

### 6. Run Chat RPC Server:

```sh
cd cmd/chat
sh build.sh
sh output/bootstrap.sh
```

### 7. Run Server:

```sh
cd cmd/...
sh build.sh
sh output/bootstrap.sh
```

### 8. Run API Server

```sh
cd cmd/api
./run.sh
```

### 9.Jaeger

visit `http://127.0.0.1:16686/` on browser.

![image-20230222164952177](/Users/yinliliang/Library/Application Support/typora-user-images/image-20230222164952177.png)



## 测试 API:



## Future:

- **数据库分库设计，**用多数据库保存服务数据，解决高并发场景问题：

  - 大量请求阻塞

    在高并发场景下，大量请求都需要操作数据库，导致连接数不够了，请求处于阻塞状态。

  - SQL 操作变慢

    如果数据库中存在一张上亿数据量的表，一条 SQL 没有命中索引会全表扫描，这个查询耗时会非常久。

  - 存储出现问题

    业务量剧增，单库数据量越来越大，给存储造成巨大压力。

- **redis 连接池，**redis集群部署，保证高可用

- **负载均衡机制**（一致性哈希），各微服务进行横向扩展，提高并发率

- **消息队列：**各服务间添加消息队列，各服务间解耦，减少数据丢失
