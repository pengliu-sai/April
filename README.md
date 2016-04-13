#April

![image](https://github.com/pengliu-sai/April/blob/master/logo.jpeg?raw=true)


###描述
使用Golang语言开发的游戏服务器, 设计中使用了pipeline的思考模式


###设置环境变量
- export GOPATH=/Users/liupeng/April
- export APRIL_PATH=/Users/liupeng/April/
- export PATH=$PATH:/opt/redis-3.0.7/src

###特性
* 流水线的处理模式! 从用户发出请求到处理, 不采用transProxy代理模式, 每条请求相互独立
* 入口单一, 无论是用户的登陆, 还是游戏服务器的管理, 对于客户端来说, 都是指向同一个tcp连接.
* 仅支持json格式的文件配置, json格式的描述性强, 且配置简单.
* 依赖的库单一, 本工程核型依赖link框架 + protobuf前端协议(长连使用) + json协议(短连使用)

###设计图

![image](https://github.com/pengliu-sai/April/blob/master/April.png?raw=true)

###参考项目
* [link](https://github.com/funny/link)
* [nsq](https://github.com/nsqio/nsq)
* [GOGAMESERVER](https://github.com/yicaoyimuys/GoGameServer)


###项目起源
* 在做客户端开发时, 经常性要和服务器进行接口的对接, 因双方对项目的认知不同, 以及技术水平的差异, 经常会出现扯皮, 互相拉扯的现象.
    在与策划进行需求PK时, 策划往往以先做做看看效果, 以往的数据已经证明这个需求一定可行等借口打发客户端, 即便是数据有作弊嫌疑, 但
    因客户端对线上数据不了解, 造成PK过程当中往往落于下风.
    So, 为了可以重新领跑, 有必要证明 服务器也是可以搞定的, 数据分析也是没问题的.
    单挑, who 怕 who!


###项目特点
####1.单一入口
* 对于客户端来说, 永远只和一个地址进行tcp通讯, 不再区分登陆服务器地址和游戏服务器地址.

  好处:
  - 降低客户端心智负担,
  - 对于每个大区都有一个登陆服务器, 解决登陆负载问题, 通过AdminServer内部通讯, 解决跨区登陆问题
  - 对外隐藏了GameServer的网络地址

####2.GameServer负载
* 玩家在登陆成功之后, 由AdminServer分配GameServer, AdminServer记录了当前用户和GameServer之间的管理
    所以, 以后可自由定制负载方案.

####3.代码风格统一
* 在设计过程当中, 有一些常规的约定.
  - 对于用户的主动登陆和登出操作, 采用Login 和 Exit 命名
  - 对于被GameServer或者AdminServer提下线的操作, 采用Online 和 Offline 命名
  - 从客户端 发送到 服务器, 协议采用 C2S => Client -> Server 结尾命令 (奇数)
  - 从服务器 发到到 客户端, 协议采用 S2C => Server -> Client 结尾命名 (偶数)
  - 等等

####4.测试用例简单, 并且覆盖完成
* 请看代码



###TODO
* GameServer redis缓存
* 公告, 签到, 排行榜, 邮件, 抽奖, 兑换码, 充值, 计费点, 通知, 商城等通用模块的开发
* game 与 world 通信时, 每一个用户都会开启一个world tcp连接, 我是想之后能不能 只开几个?
* 压力测试脚本
* 等等等等等等