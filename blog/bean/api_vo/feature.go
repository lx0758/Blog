package api_vo

type SMTPVO struct {
	Enable    bool   `json:"enable"`
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
	Ssl       bool   `json:"ssl"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FromName  string `json:"fromName"`
	FromEmail string `json:"fromEmail"`
}
