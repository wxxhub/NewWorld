# NewWorld
![](static/img/small_logo.png)

## API 
|API |作用 |参数 |
| - | - | - |
|/login |登录 |user_id(用户ID), pwd(密码)  |
|/logon |注册 |user_id(用户ID), name(用户昵称), image(头像), email(邮箱), pwd(密码)  |
|/commit |评论 |message_id(要评论的消息ID), commit(评论内容)  |
|/praise |点赞 |message_id(要评论的消息ID), praise(true表示点赞, false表示取消) |
|/add_message |发布消息 |text(内容), image(消息中的图片) |
|/concern |关注 |goal_user_id(要关注的ID), concern(true表示关注, false表示取消关注) |
|/self_message |获取自己的消息 |start(获取起点,第几条消息), end(获取结束,第几条消息) |
|/concern_message |获取关注的消息 |size(获取的数据量)  |
|/hot_message |获取热点数据 |不需要参数 |
|/test_message |获取测试数据 |直接GET就可  |
|/api(已关闭) |获取需要的数据 |name(获取的数据名, 获取测试数据"name=test_data"))  |

## 返回码
| | |
| - | - |
|401 | 身份信息过期 |

## 示例
[javascript使用api及解析数据](views/test/api_test.html)