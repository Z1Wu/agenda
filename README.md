# 服务计算第三次作业

## 命名规范

- 文件名字：使用小写字母的加下划线

- 函数/模块/变量：均使用驼峰式命名，根据是否暴露给外部确定首字母是否大写

## 应用架构

- 文件结构:
    - cmd => 表示层 // todo，命令的创建
    
        -  
    - entity 
----------------service => 编写业务逻辑，包括会议添加/删除，用户注册等----------------
        
        - user_logic.go // 用户相关逻辑（包括：用户添加，登陆等）

        - meetting_logic.go // 会议相关逻辑（会议创建，成员加入等）
  
-----------------storage => 处理存储数据------------

        - storage_hanlder.go => 处理数据的读写
        
        - date.go => 定义date
        
        - user.go => 定义user
        
        - meeting.go => 定义meeting
    
    - data => 储存用户数据和会议数据
    
        - meetings.json
    
        - users.json
    
        - cache.json
            - 记录当前的登陆信息，在用户登陆之后，可以会把登陆的用户的用户名写进该文件中。
            - 另外，由于如果没有手动调用logout，之后登陆的信息不会被消除
    main.go // cobra 文件


## 命令详情

// 用户相关

- register

    - 添加新用户，检查是否存在用户之间的重复
    
    - 选项
        - -n | --name: 名字
        - -k | --key: 密码
        - -e | --email：邮箱
        - -p | --phone: 电话
- login

- logout

// 会议相关

- deleteAllMeeting
 
- createMeeting

- addUser

- queryUser

- queryMeeting

- quitFromMeeting

- removeUserFromMeeting

// 其他

- help