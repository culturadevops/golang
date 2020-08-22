package libs

import (
	"html/template"
	"time"
)

//时间戳转换为日期
func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StrToHtml(content string) template.HTML {
	return template.HTML(content)
}
