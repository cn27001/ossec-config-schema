package ossecconf

type Client struct {
	ServerIp      Ip  `xml:"server-ip"`
	NotifyTime    int `xml:"notify_time,chardata"`
	TimeReconnect int `xml:"time-reconnect,chardata"`
}

type Ip struct {
	Value string `xml:",chardata"`
}
