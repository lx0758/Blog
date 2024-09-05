package api_vo

type SMTP struct {
	Enable    bool
	Hostname  string
	Port      int
	Ssl       bool
	Username  string
	Password  string
	FromName  string
	FromEmail string
}
