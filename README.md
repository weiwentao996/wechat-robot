# WECHAT ROBOT
## 项目目录
```
|-- wechatrobot
  |-- .gitignore
  |-- banner # 启动banner
  |-- config.yaml.example # 配置文件实例，使用时去掉.example后缀名
  |-- go.mod
  |-- go.sum
  |-- main.go
  |-- README.md
  |-- storage.json
  |-- src
      |-- config
      |   |-- config.go
      |-- lib
          |-- banner.go
          |-- gpt.go
          |-- wechat.go
```
## 配置文件
> 修改配置后需重启程序
```yaml
gpt_token: 'xxxxxxxxxxxxxxxx'
work_mode: 1
start_key_word: '开启自动回复'
end_key_word: '关闭自动回复'
prefix_word: 'robot'
```
### gpt_token
gpt机器人token，需到openapi官网自行申请
### work_mode
* 为1时代表自动回复模式，好友需通过输入 [start_key_word] 所配置关键字来启动机器人自动回复，[start_key_word] 所配置关键字为停止自动回复命令。该模式在私聊窗口与群组均可开启，但群组中开启只能监听到好友下发信息。
* 为2时代表前缀触发模式，好友需通过输入 [prefix_word] 所配置关键字来触发机器人自动回复。该模式私聊与群组均可开启，但群组中开启只能监听到好友下发信息。
### start_key_word 
自动模式下启动关键字
### end_key_word
自动模式下关闭关键字
### prefix_word
触发模式下前缀关键字
