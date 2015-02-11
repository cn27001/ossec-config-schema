package ossecconf

import "encoding/xml"

type LocalFile struct {
	XMLName   xml.Name `xml:"localfile"`
	Location  string   `xml:"location"`
	LogFormat string   `xml:"log_format"`
}
