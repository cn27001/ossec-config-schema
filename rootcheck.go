package ossecconf

import "encoding/xml"

type RootCheck struct {
	XMLName        xml.Name `xml:"rootcheck"`
	WindowsAudit   string   `xml:"windows_audit"`
	WindowsApps    string   `xml:"windows_apps"`
	WindowsMalware string   `xml:"windows_malware"`
}
