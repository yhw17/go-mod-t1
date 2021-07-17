package t1

var onlyMe = "专用内容，以小写字母开始"

var Version = "v0.1 公共内容，以大写字母开始"

func notOpen(num int) (value int) {
	return num + 1
}

func Open(num int) (str string) {
	return string(num)
}
