package voicechat

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type IVoiceChat interface {
	Connect(s *discordgo.Session, channelID string)
	Disconnect()
	SendVoiceMessage(s *discordgo.Session, channelID string, message string)
}
type VoiceChat struct {
	Connection *discordgo.VoiceConnection
}

func (vc *VoiceChat) Connect(s *discordgo.Session, channelID string) {
	var err error
	vc.Connection, err = s.ChannelVoiceJoin(s.State.Guilds[0].ID, channelID, false, true)
	if err != nil {
		fmt.Printf("Error connecting to voice channel: %v", err)
		return
	}
}

func (vc *VoiceChat) SendVoiceMessage(s *discordgo.Session, channelID string, message string) {
	s.ChannelMessageSend(channelID, message)
}

func (vc *VoiceChat) Disconnect() {
	vc.Connection.Disconnect()
}

func NewVoiceChat() IVoiceChat {
	return &VoiceChat{
		Connection: nil,
	}
}
