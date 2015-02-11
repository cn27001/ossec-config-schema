package ossecconf

import "encoding/xml"

type RegistryIgnore struct {
	XMLName xml.Name `xml:"registry_ignore"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
