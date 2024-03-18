package main

import (
	"fmt"
	"reading-articles/pkg/bot"
	"reading-articles/pkg/config"
	"reading-articles/pkg/handler"
	voicechat "reading-articles/pkg/voiceChat"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config := config.NewConfig()
	vcm := voicechat.NewVoiceChatManager()
	handlers := handler.NewHandlers(&vcm)

	bot := bot.NewBot(*config, &handlers)
	bot.InitializeBot()
	// dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN")) //トークンは環境変数にしてログイン
	// if err != nil {                                                   //エラー発生時
	// 	fmt.Println("error creating Discord session,", err)
	// 	return
	// }
	// fmt.Print(onAnyoneJoinVoiceChannel)
	// dg.AddHandler(onAnyoneJoinVoiceChannel) //メッセージが投稿されたら実行
	// err = dg.Open()                         //セッション開始
	// if err != nil {                         //エラーが起きたら
	// 	fmt.Println("error opening connection,", err)
	// 	return
	// }
	// // 終了時にちゃんとクローズするように
	// defer dg.Close()

	// fmt.Println("ログイン成功!")
	// stopBot := make(chan os.Signal, 1)
	// signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-stopBot
}

func onAnyoneJoinVoiceChannel(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	fmt.Print(m.UserID)
	if m.Member.User.Bot || m.ChannelID == "" {
		return
	}
	vc := voicechat.NewVoiceChat()
	vc.Connect(s, m.ChannelID)
}

// func onVoiceChannelIsEmpty(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
// 	fmt.Print(m.UserID)
// 	if m.Member.User.Bot || m.ChannelID == "" {
// 		return
// 	}
// 	vc := voicechat.NewVoiceChat()
// 	vc.Disconnect()
// }
