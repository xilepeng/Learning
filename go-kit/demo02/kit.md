 Go Kit 的三层架构
 
## 1. Transport
主要负责与 HTTP gRPC thrift 等相关的逻辑
## 2. Endpoint
定义 Request 和 Response 格式，并可以使用装饰器包装函数，以此来实现各种中间件嵌套
## 3. Service
实现业务类、接口

