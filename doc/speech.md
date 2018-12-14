# 演示

演示服务为`1`个生产者对应`n`个消费者的服务，当接收到生产者数据时，将自动推送到所有消费者端。

演示使用[WebSocket](websocket.md)服务。

需要先添加[认证](auth.md)后才能使用。

## 生产

```json
{
    "auth": "认证Ticket",
    "operation": "speech.produce",
    "unique": "消息唯一值",
    "type": "类型",
    "content": "消息内容"
}
```

## 注册消费者

```json
{
    "auth": "认证Ticket",
    "operation": "speech.consumer"
}
```

## 消费

```json
{
    "operation": "speech.consume",
    "unique": "消息唯一值",
    "type": "类型",
    "content": "消息内容"
}
```

## 拉取

```json
{
    "operation": "speech.pull",
    "unique": "消息唯一值范围",
    "type": "类型"
}
```

> `unique`格式为`开始:结束`，如果为`空`则表示全部；如果`开始`为`空`则表示从第一份开始；如果`结束`为`空`则表示到最后一份。

## 结束

```json
{
    "operation": "speech.finish",
    "unique": "消息唯一值"
}
```

> 如果未推送结束消息，则默认`8h`未收到推送消息后自动结束。

## 概要

HTTP POST: /whspeech/outline

参数

|名称|类型|必须|说明|
|---|---|---|---|
|auth|string|是|演示ID。|

返回值

```json
{
    "create": "创建时间戳，单位：秒",
    "modify": "更新时间戳，单位：秒",
    "unique": "最后unique值。",
    "types": [{
        "name": "类型",
        "count": "数量"
    }]
}
```

> 未结束演示时每分钟更新一次。

## Unique集

HTTP POST: /whspeech/uniques

参数

|名称|类型|必须|说明|
|---|---|---|---|
|auth|string|是|演示ID。|

返回值

```string
type:unique
```

> 每行一对`type:unique`数据。

## 轨迹

HTTP POST: /whspeech/track

参数

|名称|类型|必须|说明|
|---|---|---|---|
|ticket|string|是|认证Ticket。|
|type|string|否|消息type值。|
|start|string|否|开始unique，为空则不限制。|
|end|string|否|结束unique，为空则不限制。|

返回值

```string
type:unique
```

## 配置文件

[speech.json](../conf/speech.json)
