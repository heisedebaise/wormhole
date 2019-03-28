# 临时文件

## 复制

将已存在的文件复制一份到临时目录，并通过设置自动清理规则定期清理过期文件。

复制处理URI默认为`/whtemp/copy`，但可通过修改[文件配置](../conf/temporary.json)文件的`Root`修改为其他URI地址。

### 参数

|名称|类型|必须|说明|
|---|---|---|---|
|uri|string|是|要复制的源文件URI地址。|

> 复制接口可设置[IP白名单](../conf/whitelist.json)，或提供[参数签名](sign.md)验证。

### 返回

如果复制成功则返回文件`URI`地址，否则返回`404`。

## 获取文件

获取文件直接使用文件`URI`地址即可，如：`http://wormhole/path/name.json`。

## 配置文件

[temporary.json](../conf/temporary.json)
