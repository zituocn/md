package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"

	"io/ioutil"
	"os"
	"strings"
	"time"
)

type MdHandle struct {
	beego.Controller
}

var (
	dRoot = "file" //文档根目录
)

type FileInfo struct {
	Path     string
	Name     string
	FileSize int64
	Modtime  time.Time
	Isdir    bool
}

const DEFAULT_TITLE = ""

//字串处理
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//获取上一级目录
func getParentDirectory(dirctory string) string {
	if strings.Index(dirctory, "/") > -1 {
		return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
	}
	return ""
}

// @title 读取file目录下的文件列表或目录在页面上打印出来
// @router /md/ [get]
func (this *MdHandle) GetFileList() {
	var (
		path       string
		parentPath string
	)
	path = this.GetString("path")
	if len(path) == 0 {
		path = dRoot
	}
	if strings.Index(path, ".md") > -1 {
		title, content := getMarkDown(path)
		this.Data["content"] = content
		this.Data["title"] = title
		this.TplName = "_readfile.html"
	} else {
		files, err := WalkDir(path)
		if err != nil {
			beego.Info(err.Error())
		}
		parentPath = getParentDirectory(path)
		this.Data["parentPath"] = parentPath

		this.Data["files"] = files
		this.Data["path"] = path
		this.TplName = "_index.html"
	}
}

func WalkDir(path string) (files []*FileInfo, err error) {
	if path == "" {
		path = dRoot
	}
	file, err := ioutil.ReadDir(path)
	if err != nil {
		beego.Info(err)
	}
	for _, fi := range file {
		if fi.IsDir() {
			files = append(files, &FileInfo{path, fi.Name(), fi.Size(), fi.ModTime(), true})
		} else {
			files = append(files, &FileInfo{path, fi.Name(), fi.Size(), fi.ModTime(), false})
		}
	}
	return files, nil
}

func getMarkDown(path string) (title string, content string) {
	file, err := os.Open(path)
	if err != nil {
		beego.Info(err.Error())
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		beego.Info(err.Error())
	}

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	title = ""
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
	return title, content
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

func DateT(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}
