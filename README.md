# libx
Method toolkit based on golang

Package libx encapsulates tools based on file, read/write stream, array, map and other operations to simplify developers' repeated implementation of the same type of functions

### read
read function is open file read or net file read tool, it also supports io reading

io read demo:
```
func TestIORead(t *testing.T) {
	resp, err := http.Get("https://miajio.oss-cn-hangzhou.aliyuncs.com/group1/M00/01/3B/hello.xml")
	if err != nil {
		fmt.Printf("get failed, err: %v", err)
		return
	}
	defer resp.Body.Close()
	res, err := libx.NewRead().IORead(resp.Body).Read()

	if err != nil {
		fmt.Printf("path read fail: %v\n", err)
	}
	fmt.Println(string(res))
}
```

net read demo:
```
func TestNetRead(t *testing.T) {
	res, err := libx.NewRead().NetRead("https://miajio.oss-cn-hangzhou.aliyuncs.com/group1/M00/01/3B/hello.xml").Read()
	if err != nil {
		fmt.Printf("net read fail: %v\n", err)
	}
	fmt.Println(string(res))
}
```

path read demo:
```
func TestPathRead(t *testing.T) {
	base, _ := os.Getwd()
	filePath := base + "/hello.xml"
	res, err := libx.NewRead().PathRead(filePath).Read()
	if err != nil {
		fmt.Printf("path read fail: %v\n", err)
	}
	fmt.Println(string(res))
}
```

### writer
