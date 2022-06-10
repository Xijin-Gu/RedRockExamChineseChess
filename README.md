数据表：
用户表
id
name
password
email
sendactivation
createtime
salt
# 基本接口

## 用户相关

### `user/regist/sendactivation` POST

- `application/x-www.form-urlencoded`

- 用户注册界面

| 请求参数 | 类型 | 说明       |
| -------- | ---- | ---------- |
| Username | 必选 | 注册用户名 |
|Password|必选|注册用户密码|
|Email|必选|用户注册邮箱|

|返回参数|说明|
|---|---|
|ActicationCode|激活码|


流程：
用户传入用户名，密码，注册邮箱。
检验用户传入参数是否合法，用户名和密码不为空就行，邮箱是否符合格式
参数合法，则发送邮箱激活码，将用户信息和激活码存储在用户表中
激活码有效期为10分钟。将当前时间加10分钟存储在createtime字段
过期未激活用户信息会被从表中删除

### `user/regist/active` POST

- `application/x-www.form-urlencoded`

- 用户激活界面

| 请求参数 | 类型 | 说明       |
| -------- | ---- | ---------- |
| UserName | 必选 | 注册用户名 |
|ActivationCode|必选|激活码|

流程，通过用户名读取数据库信息
和激活码进行匹配。
激活码过期或者为空，返回激活码不存在或已经失效
如果激活码正确，激活码字段变为0，表示用户已激活。
当用户激活以后，密码进行加盐保护




### `user/login`





### `user/login/email`







### `user/login/forget`











# 扩展项