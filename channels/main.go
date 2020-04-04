// +build js, wasm

package main


func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {

	s :=[]int{1, 2, 3,4, 5, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <- c // receive from c

	println(x, y, x+y)
}