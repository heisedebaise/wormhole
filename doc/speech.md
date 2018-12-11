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

## 