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


命令集合
// 
- add            To add Participators of the meeting    
- cancel         Cancel a meeting named MeetingName     
- clear          Cancel all the meeting created by the current user
- creat          Create a meeting
- delete         Delete a user
- help           Help about any command
- listAllMeeting List all meetings the sponsor created
- listAllUser    List all users' name
- login          Login to the meeting system.
- logout         Logout the meeting system
- query          To query all the meeting have attended during [StartTime] and [EndTime]
- quit           quit the meeting with the name [MeetingName]
- regist         register a new user
- remove         To remove Participator from the meeting

命令具体使用方法
//
- 需要登录才能使用的功能：
    - ./agenda add -n [加入参与者的姓名] -m [加入会议的名字]
    - ./agenda cancel -m [取消的会议的名字] （在已经登录的用户发起的会议中取消指定会议）
    - ./agenda clear
    - ./agenda creat -m [创建的会议的名字] -s [开始时间，格式为：xxxx-xx-xx/xx:xx] -e [结束时间，格式为：xxxx-xx-xx/xx:xx] -p [参会者，不能为发起人]
    - ./agenda listAllMeeting
    - ./agenda listAllUser
    - ./agenda query
    - ./agenda quit -m [想要推出的会议的名字]
    - ./agedna remove -p [想要移除的参会者的名字] -m [从中移除参会者的会议名字]
 - 不需要登录即可使用的功能：
    - ./agenda regist -n [创建的用户名] -c [创建的密码] -e [邮箱] -t [电话号码]
    - ./agenda login -n [用户名]  -c [密码]
    - ./agenda logout
    - ./agenda delete -n [要删除的用户的用户名] -c [要删除的用户的密码]
    - ./agenda --help [产看所有的指令]
