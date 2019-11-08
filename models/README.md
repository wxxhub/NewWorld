# models

## redis
1. 用户, 使用哈希表(user:user_id name, pwd, image);
2. 消息, 消息列表和消息
    消息列表(message_list:user_id messages_id...)
    消息哈希表(message:message_id user_id, text, time, image).  
3. 评论列表: (commits:message_id commits)
4. 点赞集合: (praise_set:message_id praises)
5. 关注, 集合(user_concern:user_id user_id).  
6. 消息计数器, message_counter, 每添加一次消息, 作为消息ID,自增一次.  