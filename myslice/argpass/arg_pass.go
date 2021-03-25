// 测试切片当参数传递,在函数内修改值,append,对比函数修改前后的变化
// todo 结论:

package argpass

import (
	"fmt"
)

func AppendFunc() {
	m := make([]int, 3, 10)
	m1 := m[:4]
	//m = append(m, 1)
	//m = append(m, 2)
	m[0] = 1
	m[1] = 2
	fmt.Println("len", len(m), "cap", cap(m))
	fmt.Println("len", len(m1), "cap", cap(m1))

	fmt.Println()
	fmt.Println(m)
	changeMem(m)
	fmt.Println(" after ChangeMem", m, m1)

	fmt.Println()
	fmt.Println(m)
	appendLessCap(m)
	fmt.Println(" after appendLessCap", m, m1)

	fmt.Println()
	fmt.Println(m)
	chgAppendLessCap(m)
	fmt.Println(" after chgAppendLessCap", m, m1)

	fmt.Println()
	fmt.Println(m)
	appendMoreCap(m)
	fmt.Println(" after AppendMoreCap", m, m1)
}

func changeMem(m []int) {
	for i := 0; i < len(m); i++ {
		m[i] *= 2
	}
	fmt.Println("ChangeMem", m)
}

func appendLessCap(m []int) {
	m = append(m, -1)
	fmt.Println("appendLessCap", m)
}

func chgAppendLessCap(m []int) {
	for i := 0; i < len(m); i++ {
		m[i] *= 2
	}
	m = append(m, -1)
	fmt.Println("chgAppendLessCap", m)
}

func appendMoreCap(m []int) {
	c := cap(m)
	for len(m) <= c {
		m = append(m, -1)
	}

	fmt.Println("AppendMoreCap", m)
}

// 结构体切片

type str struct {
	A int
}

func (s *str) string() string {
	return fmt.Sprintf("v:%d", s.A)
}

func PrintSlice(name string, s interface{}) {
	ss := s.([]str)
	fmt.Println(name, "len", len(ss), "cap", cap(ss), "val", ss)
}

func AppendStructFunc() {
	m := make([]str, 3, 10)
	m1 := m[:4]
	m[0] = str{1}
	m[1] = str{2}

	PrintSlice("m", m)
	PrintSlice("m1", m1)
	changeStructMem(m)
	PrintSlice("m", m)
	PrintSlice("m1", m1)

	appendStructLessCap(m)
	PrintSlice("m", m)
	PrintSlice("m1", m1)

	chgAppendStructLessCap(m)
	PrintSlice("m", m)
	PrintSlice("m1", m1)

	appendStructMoreCap(m)
	PrintSlice("m", m)
	PrintSlice("m1", m1)
}

func changeStructMem(m []str) {
	for i := 0; i < len(m); i++ {
		m[i].A *= 2
	}
	PrintSlice("ChangeStructMem", m)
}

func appendStructLessCap(m []str) {
	m = append(m, str{-1})
	PrintSlice("appendLessCap", m)
}

func chgAppendStructLessCap(m []str) {
	for i := 0; i < len(m); i++ {
		m[i].A *= 2
	}
	m = append(m, str{-1})
	PrintSlice("chgAppendLessCap", m)
}

func appendStructMoreCap(m []str) {
	c := cap(m)
	for len(m) <= c {
		m = append(m, str{-1})
	}

	PrintSlice("AppendMoreCap", m)
}
