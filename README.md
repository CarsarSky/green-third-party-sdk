# 背景：
## 阿里绿网没有官方提供golang SDK，在工作中用到小踩了点坑，就贴了出来，后续如果有改动不一定会及时更新。
## 非必填字段均未添加

# 使用方法
## 1. 测试过程在main.go，配置好相应服务的path、scenes,并将accessKeyId accessKeySecret替换为自己阿里云账户的密钥 （[官方网址](https://ak-console.aliyun.com/#/accesskey)）
## 2. 本示例调用得是绿网的文本检测服务
## 3. 在绿网API 最复杂的就是生产Header的签名逻辑，直接复制sign.go到项目中，其余的Http服务一般自己公司封装好了，不用参考，作为测试就可以了

