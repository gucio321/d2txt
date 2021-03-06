package d2txt

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// DataDictionary represents a data file (Excel)
type DataDictionary struct {
	lookup   map[string]int
	position int
	Records  [][]string
}

// Load loads the contents of a spreadsheet style txt file
func Load(buf []byte) (*DataDictionary, error) {
	cr := csv.NewReader(bytes.NewReader(buf))
	cr.Comma = '\t'
	cr.ReuseRecord = true

	lines, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading file lines: %w", err)
	}

	fieldNames := lines[0]

	data := &DataDictionary{
		lookup:   make(map[string]int, len(fieldNames)),
		Records:  lines[0:],
		position: 0,
	}

	for i, name := range fieldNames {
		data.lookup[name] = i
	}

	return data, nil
}

// Encode encodes data dictionary into a byte slice
func (d *DataDictionary) Encode() ([]byte, error) {
	var result bytes.Buffer

	c := csv.NewWriter(&result)

	c.Comma = '\t'

	err := c.WriteAll(d.Records)
	if err != nil {
		return nil, fmt.Errorf("error occurred while writing data: %w", err)
	}

	if err := c.Error(); err != nil {
		return nil, fmt.Errorf("error occurred while writing data: %w", err)
	}

	return result.Bytes(), nil
}

// Next reads the next row, skips Expansion lines or
// returns false when the end of a file is reached or an error occurred
func (d *DataDictionary) Next() (isntLast bool) {
	d.position++

	if d.position == len(d.Records) {
		d.position = 0
		return false
	}

	if d.Records[d.position][0] == "Expansion" {
		return d.Next()
	}

	return true
}

// String gets a string from the given column
func (d *DataDictionary) String(field string) string {
	return d.Records[d.position][d.lookup[field]]
}

// Number gets a number for the given column
func (d *DataDictionary) Number(field string) int {
	n, err := strconv.Atoi(d.String(field))
	if err != nil {
		return 0
	}

	return n
}

// List splits a delimited list from the given column
func (d *DataDictionary) List(field string) []string {
	str := d.String(field)
	return strings.Split(str, ",")
}

// Bool gets a bool value for the given column
func (d *DataDictionary) Bool(field string) bool {
	n := d.Number(field)
	if n > 1 {
		log.Panic("Bool on non-bool field ", field)
	}

	return n == 1
}

// Reset sets position to 0
func (d *DataDictionary) Reset() {
	d.position = 0
}
