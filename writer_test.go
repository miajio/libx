package libx_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/miajio/libx/write"
)

func TestWriter(t *testing.T) {
	resp, err := http.Get("https://miajio.oss-cn-hangzhou.aliyuncs.com/group1/M00/01/3B/hello.xml")
	if err != nil {
		fmt.Printf("get failed, err: %v", err)
		return
	}
	defer resp.Body.Close()
	totle, err := write.WriteIO("hello.txt", resp.Body)
	if err != nil {
		fmt.Printf("writer failed, err: %v\n", err)
		return
	}
	fmt.Printf("file totle is: %d\n", totle)
}

func TestWriteByte(t *testing.T) {
	err := write.WriteByte("hello.txt", []byte("红鲤鱼与绿鲤鱼与驴与鱼"))
	if err != nil {
		fmt.Printf("writer failed, err: %v\n", err)
		return
	}
}
