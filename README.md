数据表：
用户表
id
name
password
email
sendactivation
createtime
salt
房间表
id
create_name
participate_name
create_state
participate_state
# 基本接口

## 用户相关

### `user/regist/sendactivation` POST

- `application/x-www.form-urlencoded`

- 用户注册接口

| 请求参数 | 类型 | 说明       |
| -------- | ---- | ---------- |
| name | 必选 | 注册用户名 |
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
| Name | 必选 | 注册用户名 |
|ActivationCode|必选|激活码|

流程：
通过用户名读取数据库信息
和激活码进行匹配。
激活码过期或者为空，返回激活码不存在或已经失效
如果激活码正确，激活码字段变为0，表示用户已激活。
当用户激活以后，密码进行加盐保护




### `user/login`

- `application/x-www.form-urlencoded`

- 登录接口

|请求参数|类型|说明|
|--- |--- |---|
|Name|必选|用户名|
|Password|必选|用户密码|

流程：
用户检验用户输入参数是否合法，Name字段不能大于15且不为0，Password字段长度大于6
通过用户名读取数据库，判断是否拥有该用户
如果用户存在，判断ActivatiionCode字段是否为”0“，不为”0“，则账号未激活。
用户存在且已激活，通过盐值将用户输入密码与数据匹配。
正确返回token，带有用户信息。有效时间为48个小时

## 房间相关

### `/house/create`	`POST`

- `application/x-www.form-urlencoded`

- 创建房间界面

|请求参数|类型|说明|
|---|---|---|
|token|必选|令牌|

流程：
后端提取jwt信息。
如果jwt有效，生成一个房间
将房间ID返回





### `/house/:id` `GET`

- 加入id为`:id`的房间页

|请求参数|类型|说明|
|---|---|---|
|token|必选|令牌|
|state|可选|用户准备情况|


流程：
后端提取，jwt信息
如果jwt有效，后端通过id查询房间信息。
如果房间存在，提示已加入房间，并更新房间信息。
如果房间信息中，create_state和participate_state字段都为1，则创建一个游戏界面
返回游戏id，并初始化游戏


### 游戏相关













# 扩展项