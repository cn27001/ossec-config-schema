package ossecconf

import "encoding/xml"

type WindowsRegistry struct {
	XMLName xml.Name `xml:"windows_registry"`
	Value   string   `xml:",chardata"`
}
