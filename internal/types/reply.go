package types

type Reply struct {
	Chat    int64  `json:"chat"`
	ReplyTo int    `json:"reply_to"`
	Text    string `json:"text"`
}
