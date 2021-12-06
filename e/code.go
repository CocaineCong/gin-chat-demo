package e

type Code int

const (
	//通信信号
	WebsocketSuccessMessage  = 50001
	WebsocketSuccess = 50002
	WebsocketEnd = 50003
	WebsocketOnlineReply = 50004
	WebsocketOfflineReply = 50005
	WebsocketLimit = 50006
	//WebsocketMsg = 60001
	//WebsocketRead = 60002
	//WebsocketUnread = 60003
)

func (c Code) Msg() string {
	return codeMsg[c]
}
