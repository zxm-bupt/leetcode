package day

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	sheet map[string][]int
}

func Constructor_3484(rows int) Spreadsheet {
	sheet := make(map[string][]int, 26)
	for i := 0; i < 26; i++ {
		row := make([]int, rows)
		sheet[string(i+'A')] = row
	}

	return Spreadsheet{
		sheet: sheet,
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	chars := []rune(cell)
	col := string(chars[0])
	row, _ := strconv.Atoi(string(chars[1:]))
	this.sheet[col][row-1] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.SetCell(cell, 0)
}

func (this *Spreadsheet) GetValue(formula string) int {
	fixFormula := []rune(formula)
	fixFormula = fixFormula[1:]
	parameters := strings.Split(string(fixFormula), "+")
	x := parameters[0]
	y := parameters[1]
	return this.Parser(x) + this.Parser(y)
}

func (this *Spreadsheet) Parser(formula string) int {
	fixFormula := []rune(formula)
	if fixFormula[0] < 'A' || 'Z' < fixFormula[0] {
		param, _ := strconv.Atoi(formula)
		return param
	} else {
		col := string(fixFormula[0])
		row, _ := strconv.Atoi(string(fixFormula[1:]))
		return this.sheet[col][row-1]
	}
}
