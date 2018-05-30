package xi18n

import (
	"encoding/xml"
)

// XML xliff
type XML struct {
	Version string  `xml:"version,attr"`
	Xmlns   string  `xml:"xmlns,attr"`
	File    XMLFile `xml:"file"`
}

// XMLFile file
type XMLFile struct {
	Language string `xml:"source-language,attr"`
	Datatype string `xml:"datatype,attr"`
	Original string `xml:"original,attr"`
}

// NewXML .
func NewXML(b []byte) (rs *XML, e error) {
	var root XML
	e = xml.Unmarshal(b, &root)
	if e != nil {
		return
	}
	rs = &root
	return
}
