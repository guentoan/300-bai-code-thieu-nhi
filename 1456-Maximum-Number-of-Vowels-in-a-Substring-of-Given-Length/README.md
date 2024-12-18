## 1456. Maximum Number of Vowels in a Substring of Given Length

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description)

### Đề bài

Bạn được cho một chuỗi s và một số nguyên k. Hãy trả về số lượng nguyên âm tối đa trong bất kỳ chuỗi con nào của s có độ dài k.

Các nguyên âm trong tiếng Anh là 'a', 'e', 'i', 'o' và 'u'.

### Ví dụ

**Input:**
- s = "abciiidef"
- k = 3

**Output:**
- 3

**Giải thích:**
- Chuỗi con có độ dài k=3 trong s là "abc", "bci", "cii", "iii", "iie", "def".
- Chuỗi "iii" chứa 3 nguyên âm, là chuỗi con có số lượng nguyên âm tối đa.

### Phân tích
Bài toán này có thể dùng kỹ thuật slice window để giải quyết.
Lúc đầu tạo một chuỗi con `s[0:k]`, tính số lượng nguyên âm có trong chuỗi này.
Sau đó, ta sẽ di chuyển cửa sổ này qua một vị trí một lần, tức là loại bỏ ký tự ở đầu cửa sổ và thêm ký tự mới vào cuối cửa sổ, đồng thời cập nhật lại số lượng nguyên âm.
Ta sẽ lưu lại số lượng nguyên âm tối đa trong các cửa sổ mà ta đã quét qua.
