package ossecconf

import "encoding/xml"

type SysCheck struct {
	XMLName         xml.Name          `xml:"syscheck"`
	Frequency       int               `xml:"frequency"`
	Disabled        string            `xml:"disabled"`
	Directories     []Directory       `xml:"directories"`
	Ignore          []Ignore          `xml:"ignore"`
	WindowsRegistry []WindowsRegistry `xml:"windows_registry"`
	RegistryIgnore  []RegistryIgnore  `xml:"registry_ignore"`
}
