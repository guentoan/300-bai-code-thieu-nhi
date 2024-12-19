package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	result := make([]byte, 0, n)
	cycle := 2 * (numRows - 1)

	for row := 0; row < numRows; row++ {
		for j := row; j < n; j += cycle {
			result = append(result, s[j])

			// Xử lý ký tự xen giữa (nếu không phải hàng đầu tiên hoặc cuối cùng)
			if row != 0 && row != numRows-1 {
				midIndex := j + cycle - 2*row
				if midIndex < n {
					result = append(result, s[midIndex])
				}
			}
		}
	}

	return string(result)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) //PAHNAPLSIIGYIR
	fmt.Println(convert("PAYPALISHIRING", 4)) //PINALSIGYAHRPI
}
