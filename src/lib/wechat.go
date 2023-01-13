package lib

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
	"log"
	"strings"
	"sync"
	"wechatrobot/src/config"
)

func init() {
	hub = ClientHub{Group: map[string]bool{}, User: map[string]bool{}}
}

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}

func Login(initCallback chan bool) {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式
	// 注册消息处理函数
	bot.MessageHandler = msgCallback
	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")
	// 登陆
	if err := bot.HotLogin(reloadStorage); err != nil {
		fmt.Println(err)
		// 注册登陆二维码回调
		bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
		//bot.UUIDCallback = ConsoleQrCode
		// 登陆
		if err := bot.Login(); err != nil {
			fmt.Println(err)
			initCallback <- false
			return
		}
	}
	log.Println("微信启动成功！")
	initCallback <- true
	bot.Block()
}

func msgCallback(msg *openwechat.Message) {
	switch config.GlobalConfig.WorkMode {
	case config.PrefixModel:
		prefixModel(msg)
	case config.AutoAnswerModel:
		autoAnswerModel(msg)
	}

}

func prefixModel(msg *openwechat.Message) {
	if msg.IsText() && !msg.IsSendByGroup() && strings.HasPrefix(strings.ToUpper(msg.Content), strings.ToUpper(config.GlobalConfig.PrefixWord)) {
		q := Q(strings.TrimPrefix(strings.ToUpper(msg.Content), strings.ToUpper(config.GlobalConfig.PrefixWord)))
		msg.ReplyText(q)
		msg.ReplyText(fmt.Sprintf("🤖以上为机器人回复！"))
	}
	if msg.IsText() && msg.IsSendByGroup() && strings.HasPrefix(strings.ToUpper(msg.Content), strings.ToUpper(config.GlobalConfig.PrefixWord)) {
		q := Q(strings.TrimPrefix(strings.ToUpper(msg.Content), strings.ToUpper(config.GlobalConfig.PrefixWord)))
		sender, _ := msg.Sender()
		group := openwechat.Group{User: sender}
		group.SendText(q)
		_, err := group.SendText("🤖以上为机器人回复！")
		if err != nil {
			fmt.Println(err)
		}
	}
}

type ClientHub struct {
	User  map[string]bool
	Group map[string]bool
	mux   sync.Mutex
}

var hub ClientHub

func autoAnswerModel(msg *openwechat.Message) {
	if msg.IsText() && strings.ToUpper(msg.Content) == config.GlobalConfig.StartKeyWord {
		if msg.IsSendByGroup() {
			sender, _ := msg.Sender()
			group := openwechat.Group{User: sender}
			hub.mux.Lock()
			hub.Group[group.NickName] = true
			hub.mux.Unlock()
			group.SendText(fmt.Sprintf("🤖自动回复已经开启！"))
		} else {
			sender, _ := msg.Sender()
			hub.mux.Lock()
			hub.User[sender.NickName] = true
			hub.mux.Unlock()
			msg.ReplyText(fmt.Sprintf("🤖自动回复已经开启！"))
		}
		return
	}
	if msg.IsText() && strings.ToUpper(msg.Content) == config.GlobalConfig.EndKeyWord {
		if msg.IsSendByGroup() {
			sender, _ := msg.Sender()
			group := openwechat.Group{User: sender}
			hub.mux.Lock()
			delete(hub.Group, group.NickName)
			hub.mux.Unlock()
			group.SendText(fmt.Sprintf("🤖自动回复已经关闭！"))
		} else {
			sender, _ := msg.Sender()
			hub.mux.Lock()
			delete(hub.User, sender.NickName)
			hub.mux.Unlock()
			msg.ReplyText(fmt.Sprintf("🤖自动回复已经关闭！"))
		}
		return
	}

	if msg.IsSendByGroup() {
		sender, _ := msg.Sender()
		group := openwechat.Group{User: sender}
		hub.mux.Lock()
		_, ok := hub.Group[group.NickName]
		hub.mux.Unlock()
		if !ok {
			return
		}
		q := Q(msg.Content)
		group.SendText(q)
		_, err := group.SendText("🤖以上为机器人回复！")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		sender, _ := msg.Sender()
		a := hub
		_ = a
		hub.mux.Lock()
		_, ok := hub.User[sender.NickName]
		hub.mux.Unlock()
		if !ok {
			return
		}
		q := Q(msg.Content)
		msg.ReplyText(q)
		msg.ReplyText(fmt.Sprintf("🤖以上为机器人回复！"))
	}
}
