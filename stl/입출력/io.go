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

var input = `1 2`
var sc = bufio.NewScanner(strings.NewReader(input))

func main() {
	sc.Split(bufio.ScanWords)
	//--- Main Code Start----

	a := nextInt()
	b := nextInt()
	fmt.Println(a + b)
}

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
