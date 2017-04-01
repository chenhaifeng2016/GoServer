package message

type MsgHeader struct {
	version int32
	msgBoyeLen int32
}

type Message struct {
	msgHeader MsgHeader
	msgBody []byte
}

