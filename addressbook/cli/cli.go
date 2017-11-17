package main

import (
	"fmt"
	"os"
	"sort"
)

// Some something
type Some interface {
	doSomething()
}

// MyStruct something
type MyStruct struct {
	myVar      int
	myOtherVar string
}

// OtherStruct something
type OtherStruct struct {
	myVar      int
	myOtherVar string
}

func (myS *MyStruct) doSomething() {
	fmt.Fprintf(os.Stdout, "Hello myStruct %v\n", myS)
}

func (other *OtherStruct) doSomething() {
	fmt.Fprintf(os.Stdout, "Hello otherStruct %v\n", other)
}

type myStr string

func (myS *myStr) String() string {
	return fmt.Sprintf("MyStruct: %s", string(*myS))
}

// SuperInt is the bomb
type SuperInt int

func (sInt SuperInt) doSomething() {
	fmt.Printf("Value of super int is %v\n", sInt)
}

func doSomethingWithInt(myInt int) {
	SuperInt(myInt).doSomething()
}

// Sequence rules
type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) String() string {
	sort.Sort(s)
	str := "["

	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

// Stringer is
type Stringer interface {
	String() string
}

func main() {
	var value interface{} = "Hola"

	var s Stringer
	var s2 string

	switch str := value.(type) {
	case string:
		s2 = str
	case Stringer:
		s = str
	}

	fmt.Printf("S is %v, S2 is %v\n", s, s2)
	var seq Sequence = []int{1, 2, 3, 4}

	fmt.Printf("Seq is %v\n", seq)
}
