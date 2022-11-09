package main

import "fmt"

type celsius float64
type fahrenheit float64

//method for converting temperature celsius to fahrenheit
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

//method for converting temperature fahrenheit to celsius
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================"
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

type getRowFn func(row int) (string, string)
type replay func(a int) int

//creates a table with two columns
func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)                 //"======================" тут выводим переменнную line
	fmt.Printf(rowFormat, hdr1, hdr2) //переменна rowFormat  с двумЯ строками hdr1, hdr2"| %8s | %8s |\n"
	fmt.Println(line)                 //"======================" тут выводим переменнную line
	for row := 0; row < rows; row++ { //в цикле
		cell1, cell2 := getRow(row) //ctof(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}
func intPrint(r replay) {
	a := iterInt(5)
	fmt.Println("Vse ok", a)

}
func iterInt(a int) int {
	a = a + a
	return a
}
func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}
func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}
func main() {
	drawTable("°C", "°F", 29, ctof)
	fmt.Println()
	drawTable("°F", "°C", 29, ftoc)
	a := iterInt(5)
	intPrint(iterInt)
	fmt.Println(a)
}
