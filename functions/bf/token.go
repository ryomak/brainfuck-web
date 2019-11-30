package bf

import (
	"regexp"
	"strings"
)

//commands is list of bf command
type Commands struct {
	NEXT  string
	PREV  string
	INC   string
	DEC   string
	WRITE string
	OPEN  string
	CLOSE string
}

// Ops is getting opelation for regexp
func (c *Commands) Ops() string {
	//for regexp
	return strings.Join(c.list(), "|")
}

// list is getting match list
func (c *Commands) list() []string {
	return []string{
		regexp.QuoteMeta(c.NEXT),
		regexp.QuoteMeta(c.PREV),
		regexp.QuoteMeta(c.INC),
		regexp.QuoteMeta(c.DEC),
		regexp.QuoteMeta(c.WRITE),
		regexp.QuoteMeta(c.OPEN),
		regexp.QuoteMeta(c.CLOSE),
	}
}

func (c Commands) Parse(str string) string {
	str = strings.Replace(str, ">", c.NEXT, -1)
	str = strings.Replace(str, "<", c.PREV, -1)
	str = strings.Replace(str, "+", c.INC, -1)
	str = strings.Replace(str, "-", c.DEC, -1)
	str = strings.Replace(str, ".", c.WRITE, -1)
	str = strings.Replace(str, "[", c.OPEN, -1)
	str = strings.Replace(str, "]", c.CLOSE, -1)
	return str
}
