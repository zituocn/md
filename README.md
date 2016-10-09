# MarkDown for beego
### 介绍 
markdown在beego上的使用   

### 第三方包  
`github.com/russross/blackfriday`  

### 使用
把md文件放在`file`目录下的某个地方，可用  
`http://ip:port/md/xxx.md`  
来访问


### 安装
`go get github.com/zituocn/md`

### 运行
`./md`

### 例子
访问 `http://127.0.0.1:8080/md/sam/22v.md`

### 自行修改
可自行修改 `views/_readfile.html`和`/static/css/`中的样式文件，达到你想要的效果

