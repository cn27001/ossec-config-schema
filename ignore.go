package ossecconf

import "encoding/xml"

type Ignore struct {
	XMLName xml.Name `xml:"ignore"`
	Type    string   `xml:"type,attr,omitempty"`
	Value   string   `xml:",chardata"`
}
