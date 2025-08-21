package main

import (
	"fmt"
	"sync"
	"github.com/emirpasic/gods/sets/treeset"
)


func reader(y chan<- int, limit int, wg *sync.WaitGroup) { // passed the sending end of the channel

	defer wg.Done()

	for j:=0;j<limit;j++{

		// read input 
		var num int;
		fmt.Scan(&num);
		y <- num
	}
	
	close(y); // reading work done close channel now

}

func writer( y <-chan int, wg *sync.WaitGroup, set *treeset.Set, mu *sync.Mutex) { // the feceiving end of the channel
	
	defer wg.Done()

	// val, ok := <-y  // this is wong  
	//read all values from y untill it is closed 

	for val := range y {
		// spawn 2 go routines now one to find them and other to save them into the array 
		// simpler implementation currently

		var isp []int;
		for j := 0;j<=val;j++ {
			isp = append(isp, 1); // consider all as prime for now 
		}

		for j := 2; j*j <= val; j++ {
			for i := j; i*j <= val; i++ {
				isp[i*j] = 0;
			} 
    	}

		fmt.Println(" RESULTS : ");
		for j:=2;j<=val;j++{
			if(isp[j] == 1) {

				mu.Lock()
				set.Add(j)
				mu.Unlock()
				fmt.Print(j," ");
			}
		}
		fmt.Println(" Bye ")
	}



	

}

func main() {

	fmt.Println(" enter the number of testcases upt which you will want to find the prime numbers ");
	var n int;
	
	fmt.Scan(&n);
	fmt.Println("============== you have entered N : ",n," ====== CALCULATIONS START NOW =============");
	
	
	var wg sync.WaitGroup
	ch := make(chan int);
	set := treeset.NewWithIntComparator();
	var mu sync.Mutex


	wg.Add(2);

	go reader(ch,n,&wg)
	go writer(ch,&wg,set, &mu)

	wg.Wait()

	fmt.Println("============== COMPUTATIONS DONE ================")
	fmt.Println(" Overall primes recorded : ");
	fmt.Print(set);


	// reader must works n times only then





}