package main

import (
	goList "github.com/stucamp/golist"
)

func main() {
	lst := goList.List{}

	for i := 0; i <= 100; i += 10 {
		lst.Insert(i)
	}

	lst.Reverse()

	lst.Display()
}
