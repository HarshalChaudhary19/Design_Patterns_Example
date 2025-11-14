package designpatterns

import "fmt"

//Swap algorithms/behaviors at runtime via interchangable objects
//Basically Same task krne ke multiple different ways

//

type SortStrategy interface {
	Sort([]int)
}

type BubbleSort struct{}
type QuickSort struct{}

type Sorter struct {
	strategy SortStrategy
}

func (BubbleSort) Sort(arr []int) {
	fmt.Println("Sorted array BubbleSort", arr)
}

func (QuickSort) Sort(arr []int) {
	fmt.Println("Sorted array QuickSort", arr)
}

func (s *Sorter) SetStrategy(st SortStrategy) {
	s.strategy = st
}

func (s *Sorter) DoSort(a []int) {
	s.strategy.Sort(a)
}

func Strategy() {
	s := &Sorter{}
	s.SetStrategy(QuickSort{})
	arr := []int{1, 2, 3, 4, 5, 6}
	s.DoSort(arr)
	fmt.Println(arr)
}
