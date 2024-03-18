package voicechat

type IVoiceChatManager interface {
	AddVoiceChat(
		id string,
		vc *IVoiceChat,
	)
	RemoveVoiceChat()
}
type VoiceChatManager struct {
	connections map[string]*IVoiceChat
}

func (vcm *VoiceChatManager) AddVoiceChat(id string, vc *IVoiceChat) {
	vcm.connections[id] = vc
}

func (vcm *VoiceChatManager) RemoveVoiceChat() {

}
func NewVoiceChatManager() IVoiceChatManager {
	return &VoiceChatManager{
		connections: make(map[string]*IVoiceChat),
	}
}
