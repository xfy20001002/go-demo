# httptest

## 基本测试使用(Demo1)
```
req, err := http.NewRequest("GET", "/health-check", nil)    //第二个参数无效
rr := httptest.NewRecorder()
//直接使用HealthCheckHandler，传入参数rr,req
HealthCheckHandler(rr, req)
直接测试handler是否有效,不用管请求路径
```