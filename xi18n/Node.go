package xi18n

import (
	"regexp"
	"strings"
)

// Node 對應一個 trans-unit
type Node struct {
	// 唯一 標識
	ID string
	// 翻譯 結果
	Val []byte
	// trans-unit 完整 數據
	Tag []byte
}

// UpdateTag 將 trans-unit 的 Val 設置為 指定值 並更新 Val Tag 屬性
func (n *Node) UpdateTag(b []byte) {
	val := string(b)
	tag := string(n.Tag)
	var str string
	if matchTarget.Match(n.Tag) {
		str = matchTarget.ReplaceAllString(tag, "")
	} else {
		return
	}
	pos := strings.Index(tag, "</source>")
	if pos == -1 {
		return
	}
	size := len("</source>") + pos
	left := str[:size] + "\n        <target>" + val + "</target>\n        "
	right := str[size:]

	var matchSpace = regexp.MustCompile(`\n\s*\n`)
	right = matchSpace.ReplaceAllString(right, "")

	str = left + strings.TrimSpace(right)

	n.Val = []byte(val)
	n.Tag = []byte(str)
}
