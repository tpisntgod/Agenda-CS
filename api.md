- [Post]   v1/users  -- 创建用户
- [Get]   v1/user/login{?name,password}  -- 登录用户
- [Get]   v1/user/logout  -- 登出用户
- [Get]   v1/users -- 显示所有用户
- [Delete]   v1/users/{id}  -- 删除用户  **？**
- [Get]   v1/user  -- 显示当前用户
- ​



- [Get]   v1/users/{id}/all-meetings   -- 显示所有会议
- [Delete]   v1/users/{id}/quit-meeting/{title}   -- 退出会议
- [Delete]   v1/users/{id}/cancel-meeting/{title}   -- 取消会议
- [Delete]   v1/users/{id}/cancel-all-meeting   -- 取消所有会议
- [Get]   v1/meetings/{id}  -- 显示用户所有会议  **?**
- [Put]   v1/meeting/{title}/add-participators   -- 会议创建参与者
- [Delete]   v1/meeting/{title}/delete-participators   -- 会议删除参与者