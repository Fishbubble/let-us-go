package baseop

import (
	"fmt"
)

const PtrSize = 4 << (^uintptr(0) >> 63)

func BaseOp() {
	//op1()
	//op10()

	fmt.Println(PtrSize)
	s := []int{1, 2}
	b := []int{4, 5, 6}
	s = append(s, b...)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func printSlice(m []int) {
	fmt.Printf("len:%d cap:%d val:%v addr:%p \n", len(m), cap(m), m, m)
}

func printSliceNoVal(m []int) {
	fmt.Printf("len:%d cap:%d addr:%p \n", len(m), cap(m), m)
}

//len:2 cap:2 val:[1 2] addr:0xc0000140b0
//len:2 cap:5 val:[0 0] addr:0xc000016150
//len:3 cap:4 val:[1 2 1] addr:0xc00001a180
func op1() {
	m := []int{1, 2}
	printSlice(m)

	m1 := make([]int, 2, 5)
	printSlice(m1)

	m1 = append(m, 1)
	printSlice(m1)
}

func op2() {
	m := []int{1, 2}
	printSlice(m)

	m1 := make([]int, 2, 5)
	printSlice(m1)

	m1 = append(m, 1)
	printSlice(m1)
}

//len:0 cap:0 addr:0x11a6c30
//len:1 cap:1 addr:0xc0000140c0
//len:2 cap:2 addr:0xc0000140e0
//len:3 cap:4 addr:0xc00001a180
//len:5 cap:8 addr:0xc000018100
//len:9 cap:16 addr:0xc00001c080
//len:17 cap:32 addr:0xc00007a000
//len:33 cap:64 addr:0xc000022200
//len:65 cap:128 addr:0xc00007c000
//len:129 cap:256 addr:0xc00007e000
//len:257 cap:512 addr:0xc000100000
//len:513 cap:1024 addr:0xc000102000
//len:1025 cap:1280 addr:0xc000104000
//len:1281 cap:1696 addr:0xc00010e000
//len:1697 cap:2304 addr:0xc000118000
//len:2305 cap:3072 addr:0xc00012a000
//len:3073 cap:4096 addr:0xc000130000
//len:4097 cap:5120 addr:0xc000138000
//len:5121 cap:7168 addr:0xc000142000

// runtime.slice.go  growslice()
func op10() {
	m := make([]int, 0)
	printSliceNoVal(m)
	var oldcap int
	for i := 0; i < 6000; i++ {
		oldcap = cap(m)
		m = append(m, i)
		if cap(m) != oldcap {
			printSliceNoVal(m)
			fmt.Println(cap(m)-oldcap, float64(cap(m))/float64(oldcap))
		}
	}
}
