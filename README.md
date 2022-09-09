# 项目信息

项目`github`链接：https://github.com/CNA1949/wendaxitong.git

API测试文档链接：https://console-docs.apipost.cn/preview/420f53fda19e0517/1ee1d4387cf0f8cd



# 需求

**主题：实现一个类似知乎的问答社区**

- 多用户登录/注册
- 用户可以发起一个问题
- 所有用户都可以回答问题
- 所用用户都可以对**问题的回答**进行评论
- 用户可以获取自己的所有问题、所有回答、所有评论（回答和评论需要能定位出处，比如是哪个问题的回答，或者是哪个回答的评论）
- 用户可以删除或修改自己发布的问题或回答
- 实现点赞功能
- 实现关注功能
- 使用rpc完成网关和服务发现功能实现HTTP调用
- 使用redis作为缓存

	





# 系统设计

## 功能模块设计

### 基础功能

用户注册（已完成）

用户登录（已完成）

用户注销（已完成）

退出登录（已完成）

修改个人信息（已完成）

用户查看其他用户信息（已完成）

### 关注与取消关注

关注/取消关注用户（已完成）

查看所有已关注用户（已完成）

查看所有粉丝（已完成）



### 互动

创建话题（待完成）



删除话题（待完成）



获取话题具体信息（待完成）



评论话题（待完成）



删除评论（待完成）



点赞话题（待完成）



点赞评论（待完成）



## 数据库设计

### 用户信息表

|    字段名称     | 数据类型 |   null   | 唯一性 |  键  |                说明                |
| :-------------: | :------: | :------: | :----: | :--: | :--------------------------------: |
|     user_id     |   int    | not null | unique | 主键 |               用户id               |
|    user_name    | varchar  | not null | unique |      |               用户名               |
|      phone      | varchar  |   null   |        |      |             用户手机号             |
|      email      | varchar  |   null   |        |      |              用户邮箱              |
|    num_fans     |   int    |   null   |        |      |    用户粉丝数（即被关注的数量）    |
|    num_idols    |   int    |   null   |        |      | 用户关注数（即用户关注的用户数量） |
|    password     | varchar  | not null |        |      |              用户密码              |
|   fans_names    | varchar  |   null   |        |      |              用户粉丝              |
|   idols_names   | varchar  |   null   |        |      |             关注的用户             |
|   created_at    | varchar  | not null |        |      |            用户注册时间            |
|   updated_at    | varchar  | not null |        |      |          用户信息更新时间          |
| last_login_time | varchar  | not null |        |      |          用户最近登录时间          |





### 关注信息表





## 系统功能实现

项目`github`链接：https://github.com/CNA1949/wendaxitong.git



# 系统测试

API测试文档链接：https://console-docs.apipost.cn/preview/420f53fda19e0517/1ee1d4387cf0f8cd

