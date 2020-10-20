package opml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Body struct {
	Outlines []Outline `xml:"outline"`
}

type OPML struct {
	// Version string `xml:"version,attr"`
	// Head    Head   `xml:"head"`
	Body Body `xml:"body"`
}

type Outline struct {
	Outlines []Outline `xml:"outline"`
	Text     string    `xml:"text,attr"`
	Type     string    `xml:"type,attr"`
	XMLURL   string    `xml:"xmlUrl,attr"`
}

// type Head struct {
// }

func Parse(byteData []byte) (*OPML, error) {
	var opml OPML
	err := xml.Unmarshal(byteData, &opml)
	if err != nil {
		return nil, err
	}

	return &opml, nil
}

func Load(name string) (*OPML, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteData, _ := ioutil.ReadAll(file)

	return Parse(byteData)
}
