  
package main

import "fmt"

func main(){
	var n int
    var total int = 0
	fmt.Scan(&n)
    s := make([]int, 0, n)
	for i:=0; i<n; i++{
        var aux int
        fmt.Scan(&aux)
        total += aux
        s = append(s, aux)
	}

	fmt.Print(total)
}