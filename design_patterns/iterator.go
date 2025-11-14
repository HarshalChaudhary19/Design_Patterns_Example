package designpatterns

import "fmt"

//Basically struct mein data rakh kr iterate krna
//Bss iterate krne ka alag se function use krna hai

type IntIterator struct {
	data []int
	cur  int
}

func (it *IntIterator) Next() (int, bool) {
	if it.cur >= len(it.data) {
		return 0, false
	}
	v := it.data[it.cur]
	it.cur++
	return v, true
}

func Iterator() {
	it := IntIterator{data: []int{1, 2, 3, 4, 5, 6, 7}, cur: 0}
	it.Next()
	fmt.Println("Done")
}
