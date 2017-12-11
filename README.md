# Agenda

[![Build Status](https://travis-ci.org/tpisntgod/Agenda.svg?branch=master)](https://travis-ci.org/tpisntgod/Agenda)

## Team member
- 胡子昂 15331111 
- 侯培中 15331105
- 柯永基 15331135

## 项目完成情况
实现了github上多人协作开发；使用 json 存储 User 和 Meeting 实体，当前用户信息存储在 curUser.txt 中，实现了数据持久化；使用travis进行项目持续集成；添加 log 服务，记录用户的操作过程，以及关键的输出。

## 项目分工

- 胡子昂
    1. 整体项目框架的初构建
    2. entity/user包的实现 
    3. log日志的实现
- 侯培中
    1. 整体项目框架的改进  
    2. entity/meeting包的实现  
    3. meeting的test实现  
    4. 使用travis进行项目持续集成
- 柯永基
  1. cobra项目的创建 
  2. 逻辑层的实现 
  3. 项目总管理

## 项目功能

1. 用户注册，登录登出，用户查找，用户删除
2. 用户创建会议，取消会议，修改会议人员，显示所有要参加的会议
3. 用户和会议储存

## 安装和运行

使用go get 安装后，在$GOPATH中的bin文件夹直接运行Agenda。

```shell
~$ go get github.com/tpisntgod/Agenda
~$ $GOPATH/bin/Agenda
```

## Usage
- 使用$GOPATH/bin/Agenda -h调用总体帮助界面

  ```shell
  ~$ $GOPATH/bin/Agenda -h
  Usage:
    agenda [command]

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

  Use "agenda [command] --help" for more information about a command.
  ```

- 对每个指令，使用$GOPATH/bin/Agenda XX -h查看帮助页面，如下。

  其中包含：

  ​	Example：具体实例（$GOPATH/bin/Agenda ap -ttitle -pPeter -pMarry，意为添加Peter和Marry到名字叫title的会议中）

  ​	Flags： 每个flag的使用方法和作用

  ```shell
  ~$ $GOPATH/bin/Agenda ap -h
  to add some participators to a meeting with 
  	the title of the meeting and the name of the new participators.
  	 For example:

  ./app ap -ttitle -pPeter -pMarry

  Usage:
    Agenda-GO ap [flags]

  Flags:
    -h, --help                help for ap
    -p, --parti stringArray   name of the participators you want to add 
    -t, --title string        title of the meeting

  Global Flags:
        --config string   config file (default is $HOME/.Agenda-GO.yaml)
  ```
