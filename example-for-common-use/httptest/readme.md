# httptest

## 基本测试使用(Demo1)
```
req, err := http.NewRequest("GET", "/health-check", nil)    //第二个参数无效
rr := httptest.NewRecorder()
//直接使用HealthCheckHandler，传入参数rr,req
HealthCheckHandler(rr, req)
直接测试handler是否有效,不用管请求路径
```

## 使用newServer方法(Demo2)
创建一个http服务器,使用httptest的newServer方法做测试
http内置的server没有实现handler接口，错误，无法利用newServer方法进行测试