package lib

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
	"os"
	"wechatrobot/src/config"
)

var GPT *gogpt.Client

func init() {
	if config.GlobalConfig.GptToken == "" {
		log.Println("机器人未配置,无法回复信息！")
		os.Exit(0)
	}
	GPT = gogpt.NewClient(config.GlobalConfig.GptToken)
	log.Println("机器人启动成功！")
}

func Q(q string) string {
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 4000,
		Prompt:    q,
	}
	resp, err := GPT.CreateCompletion(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return "抱歉，我回答不了该问题！"
	}
	return resp.Choices[0].Text
}
