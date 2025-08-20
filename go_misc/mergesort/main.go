package main

import (
	"fmt"
	"sync"
	"time"

	// "time"
	"github.com/emirpasic/gods/sets/treeset"
)

func merge(x *[]int, stpf int, enpf int, stps int, enps int) {
	var temp []int

	j := stpf
	k := stps

	for j <= enpf && k <= enps {
		if (*x)[j] <= (*x)[k] {
			temp = append(temp, (*x)[j])
			j++
		} else {
			temp = append(temp, (*x)[k])
			k++
		}
	}

	for ; j <= enpf; j++ {
		temp = append(temp, (*x)[j])
	}

	for ; k <= enps; k++ {
		temp = append(temp, (*x)[k])
	}

	k = 0
	for j = stpf; j <= enps; j++ {
		(*x)[j] = temp[k]
		k++
	}

}

func mergeSort(x *[]int, st_index int, en_index int, wg *sync.WaitGroup, depth int) {

	if st_index < en_index {
		var mid int = (st_index + en_index) / 2
		if depth < 3 {
			var localWg sync.WaitGroup
			localWg.Add(2)

			wg.Add(1)
			go func() {
				defer wg.Done()
				defer localWg.Done()
				mergeSort(x, st_index, mid, wg, depth+1)
			}()
			
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer localWg.Done()
				mergeSort(x, mid+1, en_index, wg, depth+1)
			}()
			localWg.Wait()
		} else {
			singleSort(x, st_index, mid)
			singleSort(x, mid+1, en_index)
		}
		merge(x, st_index, mid, mid+1, en_index)
	}
}

func singleSort(x *[]int, st_index int, en_index int) {

	if st_index < en_index {
		var mid int = (st_index + en_index) / 2
		singleSort(x, st_index, mid)
		singleSort(x, mid+1, en_index)
		merge(x, st_index, mid, mid+1, en_index)
	}

}

func main() {

	fmt.Println("hello world")

	var wg sync.WaitGroup
	set := treeset.NewWithIntComparator()


	fmt.Println("finally the set has : ", set.Size())
	fmt.Println("Set values:", set.Values())

	var arr []int
	var brr []int
	for n := 1000000; n > 0; n-- {
		arr = append(arr, n)
		brr = append(brr, n)
	}

	t1 := time.Now()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mergeSort(&arr, 0, 1000000-1, &wg, 1)
	}()

	wg.Wait()
	t2 := time.Now()
	fmt.Println(" array is now sorted : ", t2.Sub(t1))

	t1 = time.Now()
	singleSort(&brr, 0, 1000000-1)
	t2 = time.Now()
	fmt.Println(" barray is now sorted : ", t2.Sub(t1))

	// fmt.Println(arr);
}
