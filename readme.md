# 主题
Database Suite for developing ,testing,debugging,monitoring and managing databases.


# Roadmap

## 阶段1：基本支持
- [ ] A.原生查询
  - [ ] one
    - [ ] first
    - [ ] last
  - [ ] all 
  - [ ] execute 执行
  - [ ] bindNames 绑定参数
- [ ] B.SQL builder
  - [ ] select
    - [ ] fields 表达式
  - [ ] insert
  - [ ] update
  - [ ] delete
  - [ ] expression  for where and having
    - [ ] like
    - [ ] in
    - [ ] between
    - [ ] and
    - [ ] or
    - [ ] is null & is not null
    - [ ] exists & not exists
    - [X] raw expression
  - [ ] having
  - [ ] where
  - [ ] group by
  - [ ] order by
    - [ ] 按自定字段值顺序排序,如：`ORDER BY FIELD(id, 1, 2, 3)`
    - [ ] 纯文本排序语句，如：`ORDER BY id DESC, name ASC`
    - [ ] 随机排序，如：`ORDER BY RAND()`
  - [ ] limit
    - [ ] 分页插件
  - [ ] offset
    - [ ] 分页插件 
  - [ ] join
    - [ ] 各种关联
  - [ ] union & union all
  - [ ] subquery
  - [ ] jsontable
  - [ ] one
  - [ ] all
  - [ ] exec
  - [ ] get raw sql & executed sql
  - [ ] 其他
- [ ] C.ActiveRecord Orm & Model
- [X] 事务
  - [X] 单次事务
  - [X] 事务嵌套-通过计数器
  - [ ] 事务嵌套-通过savepoint
- [ ] 关联关系（for sql buider and ActiveRecord Orm）
  - [ ] 一对一
  - [ ] 一对多
  - [ ] 多对多
  - [ ] 多态关联
  - [ ] with方法，即直接leftjoin出数据

## 阶段2  
- [ ] Sql语句优化
- [ ] 钩子、事件
  - [ ] command bus
- [ ] 验证器
- [ ] 聚合查询，clone方式
- [ ] Model生成
  - [ ] Web ui 交互式生成（选表名、设置关联关系、设置验证规则）
  - [ ] 命令行生成
  - [ ] 自动跟线上服务器比对表结构
- [ ] 代码生成
- [ ] 错误追踪、context
- [ ] schema操作
- [ ] 多数据源，主从，权重
- [ ] 多数据库，mysql，sqlite，mssql，oracle
- [ ] migration、auto migration
- [ ] 多语言
- [ ] 缓存
  - [ ] 缓存schema结构信息
  - [ ] 缓存查询结果
- [ ] 高性能线程池
  - [ ] 通用接口定义
  - [ ] 默认接入ants
- [ ] 多日志
- [ ] 插件机制
- [ ] 代码覆盖率
- [ ] 表名、字段名转义，格式化
- [ ] SQL格式化
- [ ] 内存、cpu、性能优化。
- [ ] 生成假数据Fake data generation
  - 假数据生成器
  - 假数据类型
- [ ] SQL简易审计  


# 感谢
    早期版本的代码更多是来自：go-ozzo/ozzo-dbx