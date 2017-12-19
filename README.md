# Agenda-CS

[![Build Status](https://travis-ci.org/tpisntgod/Agenda.svg?branch=master)](https://travis-ci.org/tpisntgod/Agenda)

## Team member
- 胡子昂 15331111
- 侯培中 15331105
- 柯永基 15331135

## 项目分工

- 胡子昂
    1. 整体项目框架的初构建
    2. entity/user包的实现
    3. dockerfile
- 侯培中
    1. 整体项目框架的改进  
    2. entity/meeting包的实现  
    3. meeting的test实现  
    4. 使用travis进行项目持续集成
- 柯永基  
    1. cli 端代码实现  
    2. 项目DockerHub管理  
    3. README编写  

## 项目功能

    1. 用户注册，登录登出，用户查找，用户删除
    2. 用户创建会议，取消会议，修改会议人员，显示所有要参加的会议
    3. 用户和会议储存


## API设计
See [https://agendacs.docs.apiary.io](https://agendacs.docs.apiary.io)

## 安装和运行

1.从Docker Hub上拉取镜像

```
~$ sudo docker pull yokyj/agenda-cs
```
2.运行服务端

```
~$ sudo docker run -p 8080:8080 --name agenda-service -d yokyj/agenda-cs
```

3.docker运行客户端

```
~$ sudo docker run -it --rm --name agenda-cli --net host yokyj/agenda-cs "sh"
```
4.shell安装客户端

```
~$ go get github.com/bilibiliChangKai/Agenda-CS/cli
```

## Usage
- 使用cli -h调用总体帮助界面

  ```shell
  ~$ cli -h
  Usage:
    cli [command]

  Available Commands:
    help        Help about any command
    register    to create a new account
    login       to sign in
    logout      to sign out
    usrSch      to list all the users
    usrDel      to delete the current user
    mc          to create a new meeting
    ap          to add some participators to a meeting
    dp          to delete some partipators from a meeting
    ms          to list all meetings you take part in according to the time you provide
    mq          to quit a meeting
    mc          to cancel a meeting
    mclr        to clear all the meetings you host

  Flags:
    -h, --help   help for agenda

  Use "cli [command] --help" for more information about a command.
  ```

- 对每个指令，使用cli XX -h查看帮助页面，如下。

  其中包含：

  ​	Example：具体实例（cli ap -ttitle -pPeter -pMarry，意为添加Peter和Marry到名字叫title的会议中）

  ​	Flags： 每个flag的使用方法和作用

  ```shell
  ~$ cli ap -h
  to add some participators to a meeting with
  	the title of the meeting and the name of the new participators.
  	 For example:

  cli ap -ttitle -pPeter -pMarry

  Usage:
    cli ap [flags]

  Flags:
    -h, --help                help for ap
    -p, --parti stringArray   name of the participators you want to add
    -t, --title string        title of the meeting

  Global Flags:
        --config string   config file (default is $HOME/.Agenda-GO.yaml)
  ```
  
## 功能展示
  
  ### register
```
~$ cli register -ukyj -p123 -e123@163.com -n123
register called
a new account is registered named by kyj
```

  ### login
```
~$ cli login -ukyj -p123
kyj is logined
```

  ### logout
```
~$ cli logout
logout called
logout successfully
```

  ### usrSch
```
~$ cli usrSch
usrSch called
[
    {
        "Email": "123@163.com",
        "Name": "kyj",
        "Phone": "123"
    },
    {
        "Email": "123@163.com",
        "Name": "jzy",
        "Phone": "123"
    },
    {
        "Email": "123@163.com",
        "Name": "hza",
        "Phone": "123"
    },
    {
        "Email": "123@163.com",
        "Name": "hpz",
        "Phone": "123"
    }
]
```
  ### usrDel
```
~$ cli usrDel
usrDel called
user is canceled successfully.
```

  ### mc
```
 ~$ cli mc -ttest -phza -phpz -s"2017-10-28 09:30:00" -e"2017-10-28 10:30:00"
mc called
create meeting test successfully
```
当标题含空格时
```
~$ cli mc -t"a long title" -phza -phpz -s"2017-10-28 07:30:00" -e"2017-10-28 08:30:00"
```
  ### ap
```
~$ cli ap -ttest -pjzy
ap called
meeting:test add participators successfully
```

  ### ms
```
~$ cli ms -s"2017-10-28 06:30:00" -e"2017-10-28 11:30:00" 
ms called
指定时间范围内找到的所有会议安排
会议主题： 起始时间：      终止时间：      发起者：  参与者：      
test    2017-10-28 09:30:00  2017-10-28 10:30:00  kyj    hza;hpz;jzy
a long title    2017-10-28 07:30:00  2017-10-28 08:30:00  kyj    hza;hpz
```

  ### mcc
```
~$ cli mcc -ttest
mcc called
cancel meeting test successfully
```
