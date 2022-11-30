package excel

import (
	"github.com/xuri/excelize/v2"
)

// Import 导入，读取excel表格数据
func Import(file string) ([][]string, error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	// 数据量大时，使用迭代 if rows.Next() { fmt.Println(rows.Columns()) }
	// rows, err := f.Rows()
	return f.GetRows("Sheet1")
}
