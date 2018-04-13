# 图片服务

## 上传图片

上传图片的处理URI默认为`/save`，但可通过修改[图片配置](../conf/image.json)文件的`SaveUri`修改为其他URI地址。

### 参数

|名称|类型|说明|
|---|---|---|
|path|string|图片保存的目录，未提供则保存到根目录。|
|name|string|图片保存的文件名，未提供则使用图片文件的`MD5值+后缀`作为文件名。|
|file|file|图片文件。|
> 上传接口需提供[参数签名](sign.md)验证。

### 返回

如果上传成功则返回图片`URI`地址，否则返回`404`。

## 获取图片

获取图片直接使用图片`URI`地址即可，如：`http://wormhole/path/name.jpg`。

### JPEG图片缩放与压缩

通过在`JPEG`图片名与后缀间添加参数可对图片进行压缩与缩放：
```
http://wormhole/path/name.$scale.$quality.jpg
```
|名称|类型|说明|
|---|---|---|
|scale|int|缩放比例，百分比，必须大于0；小于100表示缩小；等于100则不缩放；大于100表示放大。|
|quality|int|图片质量，1-100之间的数值。|

如：`http://wormhole/path/name.50.75.jpg`表示输出图片缩小为`50%`，图片质量`75`。

## 配置文件

[image.json](../conf/image.json)
