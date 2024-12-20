## 28. Find the Index of the First Occurrence in a String

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description)

### Đề bài

Cho hai chuỗi `needle` và `haystack`, hãy trả về **chỉ số của lần xuất hiện đầu tiên** của chuỗi `needle` trong chuỗi `haystack`.  
Nếu `needle` không xuất hiện trong `haystack`, trả về `-1`.

### Phân tích

**Cách 1**: đơn giản và dễ hiểu nhất là cắt chuỗi từ vị trí thứ `i` đến vị trí thứ `i+n` của chuỗi `haystack` (`n` là độ dài của chuỗi `needle`), sau đó so sánh hai chuỗi với nhau.

**Cách 2**: Sử dụng thuật toán **Knuth-Morris-Pratt(KMP)**.

### Thuật toán Knuth-Morris-Pratt
Thuật toán KMP sẽ xây dựng một bảng **LPS (Longest Prefix Suffix)**, sau đó duyệt qua hai chuỗi `haystack` và `needle` để tìm vị trí phù hợp.

1. Vậy câu hỏi đặt ra **LPS** là gì?

**LPS** là một mảng có độ dài bằng với độ dài chuỗi `needle`, mà giá trị phần tử thứ `i` của mảng này biểu thị độ dài của chuỗi con dài nhất từ đầu chuỗi (prefix) mà cũng là chuỗi con kết thúc tại vị trí `i` (suffix).
Câu này có vẻ hơi khó hiểu, để dễ hiểu hơn ta sẽ lấy ví dụ cụ thể. Giả sử `needle = "ababaca"`.  
Ta xây dựng bảng **LPS** như sau:

| Vị trí i | Ký tự | Tiền tố và hậu tố trùng nhau | Độ dài LPS[i] | Giải thích giá trị LPS[i] |
|----------|-------|-----------------------------|---------------|----------------------------|
| 0        | `a`   | Không có                    | 0             | Không có tiền tố và hậu tố, nên LPS[0] = 0. |
| 1        | `b`   | Không có                    | 0             | Không có tiền tố và hậu tố trùng nhau, nên LPS[1] = 0. |
| 2        | `a`   | `a`                         | 1             | Tiền tố và hậu tố là `a`, độ dài là 1, nên LPS[2] = 1. |
| 3        | `b`   | `ab`                        | 2             | Tiền tố và hậu tố là `ab`, độ dài là 2, nên LPS[3] = 2. |
| 4        | `a`   | `aba`                       | 3             | Tiền tố và hậu tố là `aba`, độ dài là 3, nên LPS[4] = 3. |
| 5        | `c`   | Không có                    | 0             | Không có tiền tố và hậu tố trùng nhau, nên LPS[5] = 0. |
| 6        | `a`   | `a`                         | 1             | Tiền tố và hậu tố là `a`, độ dài là 1, nên LPS[6] = 1. |

**LPS array:** `[0, 0, 1, 2, 3, 0, 1]`

2. Tại sao cần bảng **LPS**?

Trong KMP, khi có một ký tự không khớp giữa chuỗi `haystack` và `needle`, bảng **LPS** giúp chúng ta nhảy qua các ký tự không cần so sánh lại. Điều này làm giảm thời gian tính toán.

Ví dụ: Nếu `needle = "ababaca"` và `haystack = "ababcababac"`, khi ký tự tại vị trí `haystack[4]` không khớp, thay vì quay lại so sánh từ đầu của needle, chúng sử dụng bảng **LPS** để nhảy đến vị trí `LPS[4-1] = 2`, mà ko cần phải quay lại vị trí `0`.

3. Thuật toán tính bảng LPS
 - Khởi tạo một mảng `lps` có kích thước bằng độ dài của chuỗi `needle`, với giá trị đầu tiên là 0 (vì không có tiền tố và hậu tố nào trùng nhau ở ký tự đầu tiên).
 - Khởi tạo một biến `length` để lưu trữ độ dài của tiền tố dài nhất cũng là hậu tố. Ban đầu, `length = 0`.
 - Duyệt qua các ký tự trong `needle`.
   - Bắt đầu từ vị trí thứ 1 (vị trí thứ 0 đã được xử lý trước).
   - Nếu ký tự tại vị trí `i` của `needle` bằng ký tự tại vị trí `length` (đã có tiền tố trùng với hậu tố), tăng giá trị của `length` lên và lưu giá trị này vào `lps[i]`.
   - Nếu không khớp, nếu `length > 0`, cập nhật `length` thành `lps[length - 1]` để kiểm tra tiền tố dài hơn.
   - Nếu `length = 0`, đặt `lps[i] = 0`.
 - Tiếp tục lặp lại quá trình này cho đến khi duyệt hết chuỗi `needle`.

4. Thuật toán KMP
 - Xây dựng bảng **LPS**.(như trên)
 - **Duyệt qua `haystack`**: Tiến hành so sánh chuỗi `needle` với `haystack`. Nếu có không khớp, sử dụng bảng LPS để "nhảy" qua các phần không cần so sánh lại.

5. Độ phức tạp

**Time complex:** `O(n + m),` với `n` là độ dài chuỗi `haystack` và `m` là độ dài chuỗi `needle`. Việc tính bảng **LPS** mất `O(m)` và duyệt qua chuỗi `haystack` mất `O(n)`.
