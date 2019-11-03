# models

## redis
1. 用户, 使用哈希表(user:user_id name, pwd, head);
2. 消息, 消息列表和消息
    消息列表(message_list:user_id messages_id...)
    消息哈希表(message:message_id user_id, text, time, praise, commit[], image)
3. 关注, 集合(user_concern:user_id user_id)