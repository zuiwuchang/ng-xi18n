package cmdupdate

import (
	"fmt"
	"github.com/zuiwuchang/ng-xi18n/xi18n"
	"io/ioutil"
	"os"
)

// Update .
func (c *Context) Update() (e error) {
	_, e = os.Stat(c.Dist)
	if e != nil {
		if os.IsNotExist(e) {
			e = c.copyFile()
		}
		return
	}

	//dist
	var b []byte
	b, e = ioutil.ReadFile(c.Dist)
	if e != nil {
		return
	}
	var merge *xi18n.Context
	merge, e = xi18n.NewContext(b)
	if e != nil {
		return
	}
	if c.Locale == "" {
		c.Locale = merge.XML.File.Language
	} else {
		if merge.XML.File.Language != c.Locale {
			e = fmt.Errorf("locale not mathc %s != %s", merge.XML.File.Language, c.Locale)
			return
		}
	}

	//src
	b, e = ioutil.ReadFile(c.Src)
	if e != nil {
		return
	}
	var src *xi18n.Context
	src, e = xi18n.NewContext(b)
	if e != nil {
		return
	}

	for i := 0; i < len(src.Items); i++ {
		id := src.Items[i].ID
		if find, ok := merge.Keys[id]; ok {
			src.Items[i].UpdateTag(find.Val)
		}
	}

	// write file
	var f *os.File
	f, e = os.Create(c.Dist)
	if e != nil {
		return
	}
	e = src.Marshal(f, c.Locale)
	f.Close()
	return
}
func (c *Context) copyFile() (e error) {
	var b []byte
	b, e = ioutil.ReadFile(c.Src)
	if e != nil {
		return
	}
	var dist *xi18n.Context
	dist, e = xi18n.NewContext(b)
	if e != nil {
		return
	}

	var f *os.File
	f, e = os.Create(c.Dist)
	if e != nil {
		return
	}
	e = dist.Marshal(f, c.Locale)
	f.Close()
	return
}
