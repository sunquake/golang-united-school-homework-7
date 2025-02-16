package coverage

import (
	"os"
	"testing"
	"time"
	"reflect"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T){

	var p People
	tNow := time.Now()
	p = append(p, Person{firstName: "Joe1", lastName: "Trump1", birthDay: tNow})
	p = append(p, Person{firstName: "Joe2", lastName: "Trump2", birthDay: tNow.Add(time.Hour * 24)})
	if p.Len() != 2 {
		t.Errorf("Expected 3, got %d", p.Len())
	}

}

func TestLess(t *testing.T){
	var p People
	tNow := time.Now()
	p = append(p, Person{firstName: "Joe", lastName: "Trump", birthDay: tNow})
	p = append(p, Person{firstName: "Donald", lastName: "Biden", birthDay: tNow})
	p = append(p, Person{firstName: "Joe", lastName: "Trump", birthDay: tNow.Add(time.Hour * 24)})
	p = append(p, Person{firstName: "Donald", lastName: "Biden", birthDay: tNow.Add(time.Hour * 48)})
	p = append(p, Person{firstName: "Donald", lastName: "Putin", birthDay: tNow.Add(time.Hour * 48)})
	
	ler := p.Less(1,0)
	if !ler{
		t.Errorf("Expected true, got %v", ler)
	}
	ler = p.Less(3,2)
	if !ler{
		t.Errorf("Expected true, got %v", ler)
	}
	ler = p.Less(3,4)
	if !ler{
		t.Errorf("Expected true, got %v", ler)
	}
}

func TestSwap(t *testing.T){
	var p People
	p = append(p, Person{firstName: "Joe", lastName: "Trump", birthDay: time.Now()})
	p = append(p, Person{firstName: "Donald", lastName: "Biden", birthDay: time.Now().Add(time.Hour * 24)})

	p0 := p[0]

	p.Swap(0,1)

	if p0 == p[0]{

		t.Error("Didn't swap", p[0])
	}
}

func TestNew(t *testing.T){
	
	_, err := New("1 2 3 4 5 6 7 8 9")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	_, err = New(" 1 ")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	_, err = New("1 2 3\n0 5 9")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestCols(t *testing.T){

	m, err := New("1 2 3 4 5 6 7 8 9")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	var cols0 [][]int
	for i := 1; i < 10; i++{
		cols0 = append(cols0, []int{i})
	} 

	cols := m.Cols()

	b := reflect.DeepEqual(cols, cols0)
	if !b {
		t.Error("Error")
	}

}

func TestRows(t *testing.T){

	m, err := New("1 2 3 4 5 6 7 8 9")
	if err != nil {
		t.Error("Error")
	}


	s := []int{1,2,3,4,5,6,7,8,9}

	var rows0 [][]int
		rows0 = append(rows0, s)

	rows := m.Rows()

	b0 := reflect.DeepEqual(rows, rows0)
	if !b0 {
		t.Error("Error")
	}
}

func TestSet(t *testing.T){

	m, err := New("1 2 3 4 5 6 7 8 9")
	if err != nil {
		t.Error("Error")
	}

	n := m.Set(0, 1, 5)

	n0 := m.Set(-1,4,9)

	if !n || m.data[1] != 5 {
		t.Error("Error set")
	}

	if n0 {
		t.Error("No out of bounds check")
	}

}
