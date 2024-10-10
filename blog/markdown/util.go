package markdown

import (
	"regexp"
	"strings"
)

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re := regexp.MustCompile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re = regexp.MustCompile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re = regexp.MustCompile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re = regexp.MustCompile(`<[\S\s]+?>`)
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re = regexp.MustCompile(`\s{2,}`)
	src = re.ReplaceAllString(src, "")
	return strings.TrimSpace(src)
}
