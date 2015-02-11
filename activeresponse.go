package ossecconf

import "encoding/xml"

type ActiveResponse struct {
	XMLName  xml.Name `xml:"active-response"`
	Disabled string   `xml:"disabled"`
}
