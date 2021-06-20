# geek-go-camp

【Week02 作业题目】

Q:我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
A:dao层报错，wrap后抛给上层，由service(业务了逻辑）处理。

【Week03 作业题目】
基于 errgroup 实现一个 http server 的启动和关闭 ，
以及 linux signal 信号的注册和处理，
要保证能够一个退出，全部注销退出。

【Week04 作业题目】
按照自己的构想，写一个项目满足基本的目录结构和工程，
代码需要包含对数据层、业务层、API 注册，
以及 main 函数对于服务的注册和启动，信号处理，
使用 Wire 构建依赖。可以使用自己熟悉的框架。