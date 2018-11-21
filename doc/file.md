# 文件服务

## 上传文件

上传图片的处理URI默认为`/whfile/save`，但可通过修改[文件配置](../conf/file.json)文件的`Save`修改为其他URI地址。

### 参数

|名称|类型|说明|
|---|---|---|
|path|string|文件保存的目录，未提供则保存到根目录。|
|name|string|文件保存的文件名，未提供则使用文件的`MD5值+后缀`作为文件名。|
|file|file|文件。|
> 上传接口可设置[IP白名单](../conf/whitelist.json)，或提供[参数签名](sign.md)验证。

### 返回

如果上传成功则返回文件`URI`地址，否则返回`404`。

## 获取文件

获取文件直接使用文件`URI`地址即可，如：`http://wormhole/path/name.json`。

## 配置文件

[file.json](../conf/file.json)
