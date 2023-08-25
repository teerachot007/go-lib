package excel

import "github.com/xuri/excelize/v2"

func (e *Excel) OpenExcel(filepath string) (*excelize.File, error) {
	f, err := excelize.OpenFile(filepath)
	e.File = f
	return f, err
}

// "Sheet1" "Sheet2"
func (e *Excel) GetRowExcel(getBysheet string) ([][]string, error) {
	rows, err := e.File.GetRows(getBysheet)
	return rows, err
}
