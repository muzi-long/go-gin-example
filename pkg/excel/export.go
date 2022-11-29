package excel

import (
	"bytes"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Export 导出excel表格
func Export(data [][]interface{}) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	streamWriter, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		return nil, err
	}
	// 设置单元格的值
	for i, row := range data {
		rowNum := i + 1
		err := streamWriter.SetRow("A"+strconv.Itoa(rowNum), row)
		if err != nil {
			return nil, err
		}
	}

	err = streamWriter.Flush()
	if err != nil {
		return nil, err
	}
	bufferFile, err := streamWriter.File.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	// 根据指定路径保存文件
	return bufferFile, nil
}
