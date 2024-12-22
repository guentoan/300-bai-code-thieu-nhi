## 125. Valid Palindrome

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/valid-palindrome/description)

### Đề bài
Một cụm từ được gọi là palindrome nếu, sau khi chuyển tất cả các chữ cái in hoa thành chữ cái in thường và loại bỏ tất cả các ký tự không phải chữ cái và số, nó đọc giống nhau từ trái sang phải và từ phải sang trái. Các ký tự chữ cái và số được gọi là ký tự alphanumeric.

Cho một chuỗi `s`, hãy trả về `true` nếu chuỗi đó là palindrome, hoặc `false` nếu không phải.

### Phân tích
sử dụng phương pháp **two-pointer**. Bỏ qua các ký tự đặc biệt. 