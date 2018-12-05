# WebSocket

WebSocket默认跟随HTTP启动，数据格式描述如下：

```json
{
    "auth": "认证Ticket",
    "operation": "操作",
    "unique": "消息唯一值",
    "content": "消息内容"
}
```

> 添加认证参考[auth](auth.md)。

WebSocket请求URI默认为`/whws`，可修改[websocket.json](../conf/websocket.json)配置文件更改请求URI。

## 配置

[websocket.json](../conf/websocket.json)