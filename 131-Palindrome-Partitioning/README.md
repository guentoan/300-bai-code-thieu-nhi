## 131. Palindrome Partitioning

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/palindrome-partitioning/description/)

### Đề bài
Cho một chuỗi `s`, yêu cầu phân chia chuỗi sao cho mỗi phần tử trong phân chia là một palindrome.


### Phân tích
Do bài toán yêu cầu trả về tất các cách phân chia, nên có thể sử dụng phương pháp `backtracking`.
duyệt từ phần tử thứ `0`, đến phần tử thứ `n-1` của chuỗi.
Tại mỗi vị trí ta sẽ tiến hành kiểm tra xem một chuỗi con bắt đầu từ chỉ số `i` đến chỉ số `j` có phải là palindrome không.
Nếu là palindrome, thêm nó vào phân chia hiện tại và tiếp tục đệ quy với phần còn lại của chuỗi.
Khi phân chia toàn bộ chuỗi thành công, thêm phân chia hiện tại vào kết quả.

