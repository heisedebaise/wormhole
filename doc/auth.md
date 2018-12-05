# 认证

Auth提供认证服务，由可信任（IP白名单或参数签名）的业务服务提交认证信息。

提交认证信息时，将参数以`json`格式通过`body`提交。

## 添加生产者认证

### URI

```
/whauth/producer
```

### 参数

|名称|类型|必须|说明|
|---|---|---|---|
|token|string|是|认证Token。|
|ticket|string|是|认证Ticket。|

## 添加消费者认证

### URI

```
/whauth/consumer
```

### 参数

|名称|类型|必须|说明|
|---|---|---|---|
|token|string|是|认证Token。|
|ticket|string|是|认证Ticket。|

## 配置

[auth.json](../conf/auth.json)