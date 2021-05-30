package d2txt

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testdata() (*DataDictionary, error) {
	data, err := ioutil.ReadFile("./testdata/testdata.txt")
	if err != nil {
		return nil, err
	}

	return Load(data)
}

func Test_Load(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/testdata.txt")
	if err != nil {
		t.Error(err)
	}

	dict, err := Load(data)
	if err != nil {
		t.Error(fmt.Errorf("error loading data dictionary: %w", err))
	}

	lookup := map[string]int{
		"Name":     0,
		"Age":      1,
		"Employee": 2,
		"Hobby":    3,
	}

	assert.Equal(t, lookup, dict.lookup, "unexpected lookup table")
}

func Test_String(t *testing.T) {
	dict, err := testdata()
	if err != nil {
		t.Error(err)
	}

	names := []string{"Rob", "Steven", "Mark"}

	i := 0

	for ; dict.Next(); i++ {
		assert.Equal(t, names[i], dict.String("Name"), "Unexpected result returned by dict.String()")
	}

	if i != len(names) {
		t.Fatal("unexpected number of records read")
	}
}

func Test_Number(t *testing.T) {
	dict, err := testdata()
	if err != nil {
		t.Error(err)
	}

	ages := []int{35, 20, 73}

	i := 0

	for ; dict.Next(); i++ {
		assert.Equal(t, ages[i], dict.Number("Age"), "Unexpected result returned by dict.Number()")
	}

	if i != len(ages) {
		t.Fatal("unexpected number of records read")
	}
}

func Test_List(t *testing.T) {
	dict, err := testdata()
	if err != nil {
		t.Error(err)
	}

	hobbies := [][]string{
		{"Swimming", "Programming"},
		{"Horse riding"},
		{"books reading", "dancing"},
	}

	i := 0

	for ; dict.Next(); i++ {
		assert.Equal(t, hobbies[i], dict.List("Hobby"), "Unexpected result returned by dict.List()")
	}

	if i != len(hobbies) {
		t.Fatal("unexpected number of records read")
	}
}

func Test_Bool(t *testing.T) {
	dict, err := testdata()
	if err != nil {
		t.Error(err)
	}

	employee := []bool{true, false, false}

	i := 0

	for ; dict.Next(); i++ {
		assert.Equal(t, employee[i], dict.Bool("Employee"), "Unexpected result returned by dict.Bool()")
	}

	if i != len(employee) {
		t.Fatal("unexpected number of records read")
	}
}
