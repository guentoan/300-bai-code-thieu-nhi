## 345. Reverse Vowels of a String

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/reverse-vowels-of-a-string/description/)

### Mục Tiêu
Cho một chuỗi ký tự `s`, hãy đảo ngược tất cả các nguyên âm trong chuỗi và trả về kết quả.

Các nguyên âm bao gồm: `'a', 'e', 'i', 'o', và 'u'` (bao gồm cả chữ thường và chữ hoa, và có thể xuất hiện nhiều lần).

### Ý Tưởng
Có thể tạo một map để lưu các ký tự nguyên âm.
Sau đó, duyệt chuỗi ký tự từ hai đầu, nếu gặp ký tự là nguyên âm thì dừng lại và swap. Duyệt đến khi nào `left >= right` thì dừng lại.
