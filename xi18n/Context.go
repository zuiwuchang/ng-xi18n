package xi18n

import (
	"fmt"
	"io"
	"regexp"
)

var matchTransUnit = regexp.MustCompile(`([ \t]*)<trans-unit.+id="(\S+)".*>(.|\s)*?</trans-unit>`)
var matchTarget = regexp.MustCompile(`<target>((.|\s)*?)</target>`)

// Context ng xi18n 的內存映射
type Context struct {
	Keys  map[string]*Node
	Items []*Node
	XML   *XML
}

// NewContext .
func NewContext(b []byte) (context *Context, e error) {
	var xml *XML
	xml, e = NewXML(b)
	if e != nil {
		return
	}

	c := &Context{
		Keys: make(map[string]*Node),
		XML:  xml,
	}
	strs := matchTransUnit.FindAll(b, -1)
	c.Items = make([]*Node, len(strs))
	for i, str := range strs {
		id := string(matchTransUnit.ReplaceAll(str, []byte("$1$2")))
		val := matchTarget.Find(str)
		if val != nil {
			val = matchTarget.ReplaceAll(val, []byte("$1"))
		}

		node := &Node{
			ID:  id,
			Tag: str,
			Val: val,
		}

		c.Items[i] = node
		if val != nil {
			c.Keys[id] = node
		}
	}
	context = c
	return
}

// Marshal .
func (c *Context) Marshal(w io.Writer, locale string) (e error) {
	_, e = w.Write([]byte(`<?xml version="1.0" encoding="UTF-8" ?>` + "\n"))
	if e != nil {
		return
	}

	_, e = w.Write(
		[]byte(
			fmt.Sprintf(`<xliff version="%s" xmlns="%s">`,
				c.XML.Version, c.XML.Xmlns,
			) + "\n",
		),
	)
	if e != nil {
		return
	}
	{
		_, e = w.Write(
			[]byte(
				fmt.Sprintf(`  <file source-language="%s" datatype="%s" original="%s">`,
					locale, c.XML.File.Datatype, c.XML.File.Original,
				) + "\n",
			),
		)
		if e != nil {
			return
		}
		{
			_, e = w.Write([]byte(fmt.Sprintf(`    <body>` + "\n")))
			if e != nil {
				return
			}
			for _, item := range c.Items {
				_, e = w.Write(item.Tag)
				if e != nil {
					return
				}
				_, e = w.Write([]byte("\n"))
				if e != nil {
					return
				}
			}
			_, e = w.Write([]byte(fmt.Sprintf(`    </body>` + "\n")))
			if e != nil {
				return
			}
		}
		_, e = w.Write([]byte(fmt.Sprintf(`  </file>` + "\n")))
		if e != nil {
			return
		}
	}
	_, e = w.Write([]byte(fmt.Sprintf(`</xliff>` + "\n")))
	if e != nil {
		return
	}
	return
}
