# Domain 领域层，最核心最纯粹的业务实体及其规则的抽象定义

domain 层不得依赖 adapter/app/infra 层的具体实现

- entity: 业务实体
- service: 领域服务
- repo: 领域服务依赖的资源库定义，理论上只包含 interface 接口定义，通过依赖反转注入