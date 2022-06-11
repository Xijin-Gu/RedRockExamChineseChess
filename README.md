# 数据表：

## 用户表

id
name	
password
email
sendactivation
createtime
salt

id为主键，name为唯一字段

## 房间表

id
create_name
participate_name
create_state
participate_state

id为主键

## 游戏表

id

house_id

mover 

waiter

winlose

checkerboard

id为主键

|棋子名|字段名|码号|
|---|---|---|
|无|-|0|
|将帅|king|1|
|士|guard|2|
|象|bishop|3|
|马|knight|4|
|车|rook|5|
|炮|cannon|6|
|兵|pawn|7|







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
游戏初始化：游戏表中创建一个字段，house_id为房间id，mover字段和waiter字段初始化为create_name和participate_name字段。创建两个棋子表，绑定游戏id和用户名

# 游戏相关

### `/game/:id` 

- websocket协议

- id为`:id`的游戏|

- 参数必须使用param格式传入

|请求参数|类型|说明|
|---|---|---|
|name|必选|用户名|
|start|可选|棋子初始位置|
|end|可选|棋子目标位置|

**该接口的说明放在文档后方**



# 一些说明

## 关于游戏表checkerboard字段

游戏棋盘共有10行9列。

可以用数组切片存储

数据库中没有数组切片结构。

那么定义，一个字符串，用go语言解析。

共有90个位置，棋子类型共有8种（包含无棋子），又红黑两色。

每一个位置可以用一个长度为2的字符串表示棋子状态。

故，可以用一个180位的字符串来进行存储。

例如，初始状态：

|位置|存在棋子|编码|
|---|---|---|
|1,1|黑车|15|
|1,2|黑马|14|
|1,3|黑象|13|
|1,4|黑士|12|
|1,5|黑将|11|
|1,6|黑士|12|
|1,7|黑象|13|
|1,8|黑马|14|
|1,9|黑车|15|
|3,2|黑炮|16|
|3,8|黑炮|16|
|4,1|黑卒|17|
|4,3|黑卒|17|
|4,5|黑卒|17|
|4,7|黑卒|17|
|4,9|黑卒|17|
|7,1|红卒|27|
|7,3|红卒|27|
|7,5|红卒|27|
|7,7|红卒|27|
|7,9|红卒|27|
|8,2|红炮|26|
|8,8|红炮|26|
|10,1|红车|25|
|10,2|红马|24|
|10,3|红象|23|
|10,4|红士|22|
|10,5|红将|21|
|10,6|红士|22|
|10,7|红象|23|
|10,8|红马|24|
|10,9|红车|25|

棋盘初始状态编码为

```go
151413121112131415
000000000000000000
001600000000001600
170017001700170017
000000000000000000
000000000000000000
270027002700270027
002600000000002600
000000000000000000
252423222122232425
```

## 关于一些游戏界面的一些后端流程

基于websocket，的实现。

棋盘信息每次改变，服务端都会将信息更新到数据库中，并发送给用户。

### 准备阶段

提取用户发送的信息，和game表中的mover和waiter字段对照检验用户是否合法。

如果用户参数不合法直接退出

解析game表中的checkerboard字段，将其切分成一个10-9的字符串型切片

解析用户提供的参数，如果name和move不同r则跳过走棋阶段

如果合法：

- 提取目标位置数据end，起始位置数据star,将两个数值利用`Atoi`方法转化为`int`型

- 进行数据处理：`si = startCode%10,sj=startCode-si*10`

​								`ei = endCode%10,ei = endCode-ei*10`

- 确定棋盘信息，存储到\[10\]\[9\]checkerboard中

### 走棋阶段

#### 判断棋子类型

读取起始阶段的位置信息

对照编码确认棋子类型。

#### 走棋逻辑

##### 将

- 满足`3<j<7`,
- 编码为11，满足`1<=i<=3`
- 编码为21，满足`8<=i<=10`
- 满足`|si+sj-ei-ej|==1`

##### 士

- 满足`3<i<7`,
- 编码为12，满足`1<=i<=3`
- 编码为22，满足`8<=i<=10`
- 满足`|si*10+sj-ei*10+ej|==11`

##### 象

- 编码为13，满足`1<=i<=5`
- 编码为23，满足`6<=i<=10`
- 满足`|si*10+sj-ei*10-ej|==22`
- 且`(start+end)/2处值为00`

##### 马

- |end-start|== 12 || 21
- 马腿为，|end-start|== 12,(si,sj+1)不能由棋子；|end-start|== 21，（si+1,sj)处不能由棋子。



##### 车，炮

- i,j只有一个能动。
- 开始结束之间不能有棋子

##### 卒

- `ej>=sj`

- 编码为17，j<=5,si=ei
- 编码为27，j>=5,si=ei
- `|end-start| == 1 || 10`

#### 目的位置有己方棋子

如果表达式为真，进行下一项判断，否则跳过走棋阶段。

#### 吃子逻辑

##### 炮

- ei,ej处编码和si,sj处不同
- 遍历中间的值，存在且仅存在一个非00值

##### 非炮

- 直接替代就行了。

### 返回阶段

每隔十秒刷新一次棋盘信息。





























