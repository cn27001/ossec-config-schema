package ossecconf

import "encoding/xml"

type Directory struct {
	XMLName  xml.Name `xml:"directories"`
	CheckAll string   `xml:"check_all,attr"`
	Value    string   `xml:",chardata"`
}
