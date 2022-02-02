package main

import (
	"fmt"
	"github.com/GingJan/redpacket/src"
	"time"
)

func main() {
	totalAmount := 2000
	totalNum := 6
	seed := time.Now().UnixNano()
	fmt.Println(src.Avg2Times(totalAmount, totalNum, seed))
	fmt.Println(src.SplitLine(totalAmount, totalNum, seed))
	fmt.Println(src.AvgSplit(totalAmount, totalNum, seed))
}