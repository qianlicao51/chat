# chat
TCP协程聊天

### 问题
```
for _, v := range loginResMes.UserIds {
   if v != userId { //不显示当前自己
      fmt.Println("在线用户ID:", v)
      //初始化客户端维护的 在线列表
      user := &message.User{
         UserID:     v,
         UserStatus: message.UserOnLice,
      }
      fmt.Println("--------",userId,user)
      onlineUser[v] = user//TODO 此处导致 查询在线用户不匹配|
   }
}
```

![](pic/2020-02-25_170846.png)

