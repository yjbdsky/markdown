package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	s, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("get clipboard err:", err)
		return
	}
	var reslut string
	maxs, ret := getKeywords(s)
	if maxs == nil || ret == nil {
		fmt.Println("can not format clipboard")
		return
	}

	for _, line := range ret {
		reslut = fmt.Sprintf("%s| %s |\n", reslut, strings.Join(line, "|"))
	}
	err = clipboard.WriteAll(reslut)
	if err != nil {
		fmt.Println("write clipboard err:", err)
	}
}

func getSecondLine(i int) []string {
	r := make([]string, i)
	for j := range r {
		r[j] = "---"
	}
	return r
}

func getKeywords(s string) ([]int32, [][]string) {
	var ret [][]string
	lines := strings.Split(s, "\n")
	if len(lines) <= 1 {
		return nil, nil
	}
	cloum := len(strings.Split(lines[0], "\t"))
	row := len(lines)
	if cloum <= 1 {
		return nil, nil
	}

	for i, line := range lines {
		ks := make([]string, cloum)
		keys := strings.Split(line, "\t")
		copy(ks, keys)

		ret = append(ret, ks)
		if i == 0 {
			ret = append(ret, getSecondLine(cloum))
			row = row + 1
		}
	}
	fmt.Println(ret)
	fmt.Println(cloum, row, len(ret))
	maxCloumLen := make([]int32, cloum)
	for j := 0; j < cloum; j++ {
		for i := 0; i < row; i++ {
			fmt.Println(i, j)
			keyLen := int32(len(ret[i][j]))
			maxCloumLen[j] = Max(maxCloumLen[j], keyLen)
		}
	}

	for j := 0; j < cloum; j++ {
		for i := 0; i < row; i++ {
			keyLen := int32(len(ret[i][j]))
			maxLen := maxCloumLen[j]
			fmt.Println(maxLen, keyLen)
			if keyLen < maxLen {
				for x := int32(0); x < maxLen-keyLen; x++ {
					ret[i][j] = ret[i][j] + " "
				}
			}
		}
	}
	return maxCloumLen, ret
}

func Max(x, y int32) int32 {
	if x < y {
		return y
	}
	return x
}
