package ossecconf

import "encoding/xml"

type Ossec_config struct {
	XMLName         xml.Name       `xml:"ossec_config"`
	ServerIp        string         `xml:"client>server-ip"`
	NotifyTime      int            `xml:"client>notify_time"`
	TimeReconnect   int            `xml:"client>time-reconnect"`
	LocalFiles      []LocalFile    `xml:"localfile"`
	Rootchecks      RootCheck      `xml:"rootcheck"`
	Syschecks       SysCheck       `xml:"syscheck"`
	ActiveResponses ActiveResponse `xml:"active-response"`
}
