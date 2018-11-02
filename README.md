# 服务计算第三次作业

[![Build Status](https://travis-ci.com/Z1Wu/agenda.svg?branch=master)](https://travis-ci.com/Z1Wu/agenda)

## 命名规范

- 文件名字：使用小写字母的加下划线

- 函数/模块/变量：均使用驼峰式命名，根据是否暴露给外部确定首字母是否大写

## 应用架构

- 文件结构:
    - cmd => 表示层 //命令的创建  
        
        - add.go            To add Participators of the meeting            
        
        - cancel.go         Cancel a meeting named MeetingName     
        
        - clear.go          Cancel all the meeting created by the current user
        
        - creat.go          Create a meeting
        
        - delete.go         Delete a user
        
        - help.go           Help about any command
        
        - listAllMeeting.go List all meetings the sponsor created
        
        - listAllUser.go    List all users' name
        
        - login.go          Login to the meeting system.
        
        - logout.go         Logout the meeting system
        
        - query.go          To query all the meeting have attended during [StartTime] and [EndTime]
        
        - quit.go           quit the meeting with the name [MeetingName]
        
        - regist.go         register a new user
        
        - remove.go         // 删除用户
        
        - root.go           cobra 文件
    - entity =>  服务层和存储层文件  
        
        - use r_logic.go // 用户相关逻辑（包括：用户添加，登陆等）

        - meetting_logic.go // 会议相关逻辑（会议创建，成员加入等）
  
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
    
    - 实验截图 =》简单的检测命令功能是否正常 
        - ...

    - main.go // cobra
    - travis.yml // CI


## 命令具体使用方法

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


## 添加Travis CI

### what is Travis CI

首先需要了解一下什么 `CI`， `CI` 是 `Continuous Integration`，个人理解，简单来说就是不停的把小的变化集成到项目中，而不是等到最后有了一个很大的变再把项目集成，关键在于每次集成的时候需要保证新添加入的变化不会影响已经有的功能，或者应该有的功能是否达到预期效果。`Travis CI` 一个平台，每次你把代码 `commit and push` 到 github 上的时候, 会按照你的要求常见一个虚拟环境，在虚拟环境中使用你已经编写好的测试文件来测试代码。 


### How to use

1. 使用 repo owner 的 github 账户登陆 travis， 并授权给 travis。

2. 在对应的项目处，点击 `activate` 开启对这个仓库的CI

3. 编写测试文件，在Go语言中，测试文件的命名有特殊的要求，以 `_test.go`结尾的文件都是测试文件。

4. 在项目的根目录中写好 `travis.yml`指定测试环境和测试文件。



## Golang debug in vscode（记录项目过程中的调试方法）

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

下面就是本次实验的调试所用的 `launch.json`

``` json
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
            "mode": "debug",
            "program": "${workspaceFolder}",
            "env": {},
            // "args": ["listAllUser"], // arguments for list all users
            // "args": ["login", "-n", "alice", "-c", "aaa"], // arguments for login  
            // test logout
            // "args": ["logout"]
            // test register function
            // "args": ["regist", "-n", "bob", "-c", "aaa", "-e", "bob@mail.com", "-t", "110"], // arguments for login  
            // "args": ["regist", "-n", "alice", "-c", "aaa", "-e", "alice@mail.com", "-t", "110"], // arguments for login             
            // test create a meeting
            "args": ["creat", "-m", "aliceMeeting", "-s", "2000-02-02/00:00", "-e", "2000-02-03/00:00", "-p", "a"], 
            // test for list all meeting
            // "args": ["listAllMeeting"],  test pass           
            // test delte a user, logined user alice delete bob's account
            // "args": ["delete", "-n", "bob", "-c", "aaa"],  //test pass           
            // "args": ["", "-m", "aliceMeeting", "-s", "2000-02-02/00:00", "-e", "2000-02-03/00:00", "-p", "a"], 
            // arguments for list all meeting
            // "args": ["creat", "-m", "aliceMeeting", "-s", "2000-02-02/00:00", "-e", "2000-02-03/00:00", "-p", "a"],
        }
    ]
}
```

### 设置断点

设置断点的方式

1. 普通断点的设置: 直接在行号的左边点击，出现红点就是断点

2. 条件断点的设置：在普通断点处点击右键` edit the break point`， 输入需要需要的条件, 当条件为`True`断点生效

### debuger pannel

调试的界面和其他的调试工具大同小异，包括查看调用栈等

