package excel

import (
	"fmt"
	"os"
	"testing"
)

func TestExport(t *testing.T) {
	data := make([][]interface{}, 0)
	data = append(data, []interface{}{
		"姓名", "年龄", "电话",
	})
	data = append(data, []interface{}{
		"a", "18", "1899999999",
	})
	data = append(data, []interface{}{
		"b", "22", "17712345678",
	})
	export, err := Export(data)
	if err != nil {
		return
	}
	f, _ := os.OpenFile("./export.xlsx", os.O_RDWR|os.O_CREATE, 777)
	write, err := f.Write(export.Bytes())
	if err != nil {
		return
	}
	fmt.Println(write)
	err = f.Close()
	if err != nil {
		return
	}
}
