package api_vo

type SMTPVO struct {
	Enable    bool   `json:"enable"`
	Hostname  string `json:"hostname"`
	Mode      int    `json:"mode"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FromName  string `json:"fromName"`
	FromEmail string `json:"fromEmail"`
}
