package excel

import (
	"fmt"
	"testing"
)

func TestImport(t *testing.T) {
	file := "./export.xlsx"
	data, err := Import(file)
	fmt.Println(data, err)
}
