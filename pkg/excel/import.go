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
	// Get all the rows in the Sheet1.
	return f.GetRows("Sheet1")
}
