package main

// solution
// đặc điểm của chuỗi đối xứng:
// 1. luôn có một điểm là tâm của chuỗi đối xứng. tâm đối xứng có thể là một ký tự với chuỗi có kích thước lẻ, và hai ký tự với chuỗi có kích thước chẵn.
// 2. các giá trị ở hai bên tâm luôn có giá trị bằng nhau.
// với hai đặc điểm trên, ta có thể duyệt qua từng ký tự của chuỗi `s`, và gọi điểm đó là tâm. sau đó duyệt về hai bên của tâm đó để tìm ra chuỗi đối xứng dài nhất.
func solution(s string) string {
	// hàm này dùng để tính kích thước của chuỗi đối xứng có tâm đầu vào là left, right.
	// trường hợp là chuỗi lẻ thì left = right
	// trường hợp là chuỗi chẵn thì right = left + 1
	expand := func(left, right int) int {
		// duyệt về hai phía của tâm đối xứng.
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		// trả về độ dài của chuỗi đối xứng
		return right - left - 1
	}

	// hai biến để lưu vị trí hai đầu của chuỗi đối xứng dài nhất.
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// duyệt qua từng ký tự của chuỗi s, và tìm kích thước chuỗi đối xứng tại vị trí i
		// tại mỗi vị trí đều cần phải kiểm tra kích thước của chuỗi chẵn và chuỗi lẻ
		o := expand(i, i)
		e := expand(i, i+1)
		maxLength := e
		if o > e {
			maxLength = o
		}

		// nếu kích thước của chuỗi đối xứng tại tâm i lớn hơn chuỗi dài nhất thì ta sẽ tính lại vị trí của chuỗi mới.
		if maxLength > (end - start) {
			start = i - (maxLength-1)/2
			end = i + maxLength/2
		}
	}

	return s[start : end+1]
}
