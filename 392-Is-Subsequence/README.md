## 392. Is Subsequence

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/is-subsequence/description)

### Đề bài

Cho hai chuỗi `s` và `t`. Hãy trả về `true` nếu `s` là một chuỗi con (subsequence) của `t`, ngược lại trả về `false`.

Một chuỗi con của một chuỗi là một chuỗi mới được tạo từ chuỗi ban đầu bằng cách xóa một số (hoặc không xóa) ký tự mà **không làm xáo trộn thứ tự tương đối của các ký tự còn lại**. Ví dụ: `"ace"` là một chuỗi con của `"abcde"` nhưng `"aec"` thì không.

### Phân tích
Có thể dùng hai con trỏ `i`, `j` để duyệt qua chuỗi `s` và `t`.
mục đích là kiểm tra mọi ký tự của `s` có xuất hiện tuần tự trong `t` hay không.
