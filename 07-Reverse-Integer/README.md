## 7. Reverse Integer

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/reverse-integer/)

### Đề bài
Cho một số nguyên 32-bit có dấu `x`, hãy trả về `x` với các chữ số của nó bị đảo ngược. Nếu việc đảo ngược `x` khiến giá trị vượt ra ngoài phạm vi của số nguyên 32-bit có dấu `[-2^31, 2^31 - 1]`, thì trả về `0`.

**Giả sử môi trường không cho phép bạn lưu trữ số nguyên 64-bit (có dấu hoặc không dấu).**

### Phân tích
su dung cong thuc `res = res*10 + x%10`.