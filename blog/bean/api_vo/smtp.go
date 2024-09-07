package api_vo

type SMTPVO struct {
	Enable    bool
	Hostname  string
	Port      int
	Ssl       bool
	Username  string
	Password  string
	FromName  string
	FromEmail string
}
