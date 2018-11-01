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
<<<<<<< HEAD
- [√]regist         register a new user
- remove         To remove Participator from the meeting


## Golang debug in vscode

### 配置lauch.json

点击 VSCODE 左侧的 DEBUG 图标，进入debug panel, 点击launch, 一开始默认生成的langch.json文件。
```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${filedir}",
            "env": {},
            "args": [],             
        }
    ]
}
```

- name: 调试文件的名字

- request: 请求，一般都是

- mode， 有以下几个字段可以设置
    
    - auto: 根据你设置的 `program` 来设置相应的方式。  

    - debug: 把包里面的文件编译成
    
    - test // todo, 没试过
    
    - exec: 调试一个已经编译好的可执行文件(?)，需要和program 一起使用，在program 中指定

    - 

- program: (必须指定的一个参数，不能为空)，这个参数的有以下方式 
    
    - 指定一个包文件进行debug

    - 需要使用绝对路径,不支持相对路径。

    - vscode 本身定义一些变量，用于指定路径。
        
        - ${workspaceFolder} : 当前的工作目录，使用 vscode 中打开的文件

        - ? ${file}：指定当前文件
        
        - ${fileDirname} 当前文件所在的包中

- env: // todo， 没用过

- args: 指定调试程序时的命令参数

在本次试验中，可以用来指定对应的命令，以一个列表的形式，元素是字符串，每个空格隔开作为一个元素
``` json
"args": ["creat", "-m", "aliceMeeting", "-s", "2000-02-02/00:00", "-e", "2000-02-03/00:00", "-p", "a"], 
```

### 设置断点

设置断点的方式, 

1. 普通断点的设置: 直接在行号的左边点击，出现红点就是断点

2. 条件断点的设置：在普通断点处点击右键` edit the break point`， 输入需要需要的条件, 当条件为`True`断点生效

3. 


### 
=======
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
>>>>>>> f3d5e5d73e0bc8cb43d40b14e19b637ca47484ef
