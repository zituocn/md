package controllers

import (
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"strings"
)

type MdHandle struct {
	beego.Controller
}

const DEFAULT_TITLE = ""

func (this *MdHandle) Get() {
	var (
		dRoot   = "file"
		link    string
		content string
	)

	link = this.GetString(":splat")
	beego.Info(link)

	if len(link) == 0 {
		this.Abort("404")
	}

	//文件的绝对地址
	link = dRoot + "/" + link
	file, err := os.Open(link)
	if err != nil {
		beego.Debug(err.Error())
		this.Abort("404")
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		beego.Debug(err.Error())

	}

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	title := ""
	htmlFlags := 0
	//htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	//htmlFlags |= blackfriday.HTML_COMPLETE_PAGE
	title = getTitle(data)
	htmlFlags |= blackfriday.HTML_OMIT_CONTENTS
	//htmlFlags |= blackfriday.HTML_TOC

	var renderer blackfriday.Renderer
	renderer = blackfriday.HtmlRenderer(htmlFlags, "title", "")
	var output []byte
	if len(data) > 0 {
		output = blackfriday.Markdown(data, renderer, extensions)
		content = string(output)
	} else {
		content = "空白文章..."
	}
	this.Data["content"] = content
	this.Data["title"] = title
	this.TplName = "_readfile.html"
}

//获取文章标题
func getTitle(input []byte) string {
	i := 0

	// skip blank lines
	for i < len(input) && (input[i] == '\n' || input[i] == '\r') {
		i++
	}
	if i >= len(input) {
		return DEFAULT_TITLE
	}
	if input[i] == '\r' && i+1 < len(input) && input[i+1] == '\n' {
		i++
	}

	// find the first line
	start := i
	for i < len(input) && input[i] != '\n' && input[i] != '\r' {
		i++
	}
	line1 := input[start:i]
	if input[i] == '\r' && i+1 < len(input) && input[i+1] == '\n' {
		i++
	}
	i++

	// check for a prefix header
	if len(line1) >= 3 && line1[0] == '#' && (line1[1] == ' ' || line1[1] == '\t') {
		return strings.TrimSpace(string(line1[2:]))
	}

	// check for an underlined header
	if i >= len(input) || input[i] != '=' {
		return DEFAULT_TITLE
	}
	for i < len(input) && input[i] == '=' {
		i++
	}
	for i < len(input) && (input[i] == ' ' || input[i] == '\t') {
		i++
	}
	if i >= len(input) || (input[i] != '\n' && input[i] != '\r') {
		return DEFAULT_TITLE
	}

	return strings.TrimSpace(string(line1))
}
