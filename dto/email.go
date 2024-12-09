package dto

// EmailBodyType メールのタイプ(HTML, TEXT)
type EmailBodyType string

// EmailBodyType　の種類
const (
	HTML EmailBodyType = "HTML"
	TEXT EmailBodyType = "TEXT"
)

// EmailInfo メール送信のエンティティー
type EmailInfo struct {
	From    string        // 送信元
	To      []string      // 宛先
	Cc      []string      // CC
	Bcc     []string      // BCC
	Subject string        // タイトル
	Body    string        // 本文
	Type    EmailBodyType // HTML　または TEXT
}
