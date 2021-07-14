# iot-video-p2p-key-demo 简介

此模块提供加密key供p2p加密流使用，客户可参考此模块代码或接口定义，自行开发或封装加密模块。

## 配置文件定义

参考配置config.yml.example并新建配置config.yml，配置说明如下
```yaml
service:
  addr: 0.0.0.0:8080 # 服务绑定地址
redis:
  addr: 127.0.0.1:6379 # redis实例地址
  pass: password # redis访问密码
  expire: 86400 # key过期时间，单位为秒
```

## 服务编译及运行

生成iot-video-p2p-key-demo可执行文件

```
go build -o iot-video-p2p-key-demo
```

指定配置文件，运行iot-video-p2p-key-demo
```
iot-video-p2p-key-demo -config=config.yml
```

## 接口定义

### 生成加密key

请求示例
```
POST /code
```

返回示例
```js
{
    "code":"5eaa68f4-91e9-457c-9f02-c9b4ed5db234", // key对应的code
    "expire":1625795932, // key过期时间，unix时间戳
    "key":"2c8dad1f044b644d01260217e563132d5cae95f07ae01d23f1da887df0563a62" // key至少32位字节长度
}
```

### 查询加密key

请求示例
```
GET /code/5eaa68f4-91e9-457c-9f02-c9b4ed5db234
```

返回示例
```js
{
    "key":"2c8dad1f044b644d01260217e563132d5cae95f07ae01d23f1da887df0563a62"
}
```
