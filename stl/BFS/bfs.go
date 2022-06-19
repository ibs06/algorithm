package main

//------------
// BOJ
//------------

// import (
// 	"bufio"
// 	"fmt"
//   "strconv"
// 	"os"
// )

// var sc = bufio.NewScanner(os.Stdin)

//------------
// LOCAL
//------------

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var input = `9
8 6
7
1 2
1 3
2 7
2 8
2 9
4 5
4 6`
var sc = bufio.NewScanner(strings.NewReader(input))

func main() {
	sc.Split(bufio.ScanWords)
	//--- Main Code Start----

	n := nextInt()

	start := nextInt()
	end := nextInt()
	m := nextInt()
	fmt.Println(n, start, end, m)
	graph := map[int][]int{}

	for i := 0; i < m; i++ {
		x := nextInt()
		y := nextInt()
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	type node struct {
		val int
		dis int
	}

	queue := make([]node, 0)
	chk := make([]bool, 101)

	queue = append(queue, node{val: start, dis: 0})
	chk[start] = true

	fmt.Println(graph)
	fmt.Println(queue)
	fmt.Println(chk)

	rst := -1

	for {
		if len(queue) > 0 {
			n := queue[0]
			queue = queue[1:] // pop
			fmt.Println("queue start n: ", n)
			if n.val == end {
				rst = n.dis
				break
			}
			if _, ok := graph[n.val]; ok {
				for _, child := range graph[n.val] {
					fmt.Println("child n: ", child)
					if chk[child] == false {
						queue = append(queue, node{val: child, dis: n.dis + 1})
						chk[child] = true
					}
				}
			}
		} else {
			break
		}
	}

	fmt.Println("rst:", rst)
}

/*
	queue := make([]int, 0)
	// Push to the queue
	queue = append(queue, 1)
	// Top (just get next element, don't remove it)
	x := queue[0]
	// Discard top element
	queue = queue[1:]
	// Is empty ?
	if len(queue) == 0 {
		fmt.Println("Queue is empty !")
	}
*/

//------------
// IO Helper
//------------
func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
func nextInt64() int64 {
	sc.Scan()
	v, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return v
}
func next() string {
	sc.Scan()
	return sc.Text()
}
