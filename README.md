# Quick start
```shell
go mod download

# 分别启动 user 和 hello 服务
kratos run

# 增加用户
curl -X POST "http://localhost:8000/api/users" -d '{"name": "Alice", "age": 25}'

# 查询用户
curl -X GET "http://localhost:8000/api/users/1"   

# 用户列表
curl -X GET "http://localhost:8000/api/users"        
```

# debug
```shell
xiaolatiao@xiaolatiaodeMacBook-Pro GolandProjects % curl -X GET "http://localhost:8000/api/users/1"
{"code":500,"reason":"","message":"failed FindByID user: ent: user not found","metadata":{}}%                                                                                           
xiaolatiao@xiaolatiaodeMacBook-Pro GolandProjects % curl -X POST "http://localhost:8000/api/users" -d '{"name": "Alice", "age": 25}'

{"user":{"id":"1","name":"Alice","age":"25"}}%                                                                                                                                          
xiaolatiao@xiaolatiaodeMacBook-Pro GolandProjects % curl -X GET "http://localhost:8000/api/users/1"                                 
{"user":{"id":"1","name":"Alice","age":"25"}}%                                                                                                                                          
xiaolatiao@xiaolatiaodeMacBook-Pro GolandProjects % 
```

# ref
- https://go-kratos.dev/docs/component/transport/http