package handler

import (
	voicechat "reading-articles/pkg/voiceChat"

	"github.com/bwmarrin/discordgo"
)

type IHandler interface {
	OnAnyoneJoinVoiceChannel(s *discordgo.Session, m *discordgo.VoiceStateUpdate)
}

type Handler struct {
	vcm voicechat.IVoiceChatManager
}

func (h Handler) OnAnyoneJoinVoiceChannel(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	if m.Member.User.Bot {
		return
	}
	vc := voicechat.NewVoiceChat()
	vc.Connect(s, m.ChannelID)
	h.vcm.AddVoiceChat(m.ChannelID, &vc)
}

func NewHandlers(vcm *voicechat.IVoiceChatManager) IHandler {

	return &Handler{
		vcm: *vcm,
	}
}
