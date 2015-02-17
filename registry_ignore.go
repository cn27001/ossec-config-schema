package ossecconf

import "encoding/xml"

type RegistryIgnore struct {
	XMLName xml.Name `xml:"registry_ignore"`
	Type    string   `xml:"type,attr,omitempty"`
	Value   string   `xml:",chardata"`
}
