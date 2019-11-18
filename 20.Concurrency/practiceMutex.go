package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

var wg sync.WaitGroup

func main() {

	vS := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	square(&vS)
	fmt.Println(vS)

}

func square(vi *[]int) {
	vx := []int{}
	vsquared := 0
	for _, v := range *vi {
		fmt.Println(vsquared)
		wg.Add(1)
		mu.Lock()
		go func() {
			vsquared = v * v
			fmt.Println(v, vsquared)
			vx = append(vx, vsquared)
			mu.Unlock()
			wg.Done()
		}()
		wg.Wait()
	}
	*vi = vx
}
