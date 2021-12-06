package e

var codeMsg = map[Code]string{
	WebsocketSuccessMessage : "解析content内容信息",
	WebsocketSuccess :"发送信息，请求历史纪录操作成功",
	WebsocketEnd: "请求历史纪录，但没有更多记录了",
	WebsocketOnlineReply :"针对回复信息在线应答成功",
	WebsocketOfflineReply :"针对回复信息离线回答成功",
	WebsocketLimit :"请求收到限制",
	//WebsocketMsg :"历史纪录-对方消息",
	//WebsocketRead :"历史纪录-已读消息",
	//WebsocketUnread :"历史纪录-未读消息",
}
