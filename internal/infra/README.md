# infra 基础设施层，各种外部依赖，组件的衔接，以及 domain/repo 的具体实现

- cache 内层所依赖缓存的实现，可以是 redis、memcached 等
- client 各种中间件的 client 初始化，如 db、redis 等
- config 配置实现
- database 内层所需持久化的实现，可以是 mysql、mongo 等
- log 日志实现，在此接入第三方日志库，避免对内层的污染
- mq 内层所需的消息队列实现
- rpc 广义上的第三方服务的访问实现，可以通过 HTTP、grpc 等