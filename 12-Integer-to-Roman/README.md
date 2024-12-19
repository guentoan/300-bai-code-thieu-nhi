## 12. Integer to Roman

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/integer-to-roman/description)

### Đề bài
Số La Mã được biểu diễn bằng 7 ký hiệu khác nhau với các giá trị tương ứng như sau:

| **Ký hiệu** | **Giá trị** |
|-------------|-------------|
| I           | 1           |
| V           | 5           |
| X           | 10          |
| L           | 50          |
| C           | 100         |
| D           | 500         |
| M           | 1000        |

Số La Mã được hình thành bằng cách **thêm dần các giá trị thập phân từ cao đến thấp** theo các quy tắc sau:

1. **Trường hợp giá trị không bắt đầu bằng 4 hoặc 9**:
    - Chọn ký hiệu có giá trị lớn nhất có thể trừ được từ số hiện tại.
    - Thêm ký hiệu đó vào kết quả.
    - Trừ giá trị tương ứng của ký hiệu đã chọn khỏi số hiện tại.
    - Tiếp tục chuyển đổi phần còn lại.

2. **Trường hợp giá trị bắt đầu bằng 4 hoặc 9**:
    - Sử dụng **dạng trừ** (subtractive form), nghĩa là ký hiệu của một giá trị nhỏ hơn được đặt trước ký hiệu lớn hơn để biểu thị phép trừ.
    - Ví dụ:
        - **4** được viết là `IV` (1 [I] nhỏ hơn 5 [V]).
        - **9** được viết là `IX` (1 [I] nhỏ hơn 10 [X]).
    - Các dạng trừ hợp lệ là:
        - 4 (IV), 9 (IX),
        - 40 (XL), 90 (XC),
        - 400 (CD), 900 (CM).

3. **Chỉ các lũy thừa của 10 (I, X, C, M) có thể được lặp lại nhiều nhất 3 lần liên tiếp**:
    - Ví dụ: `III` biểu diễn 3.
    - Các ký hiệu 5 (V), 50 (L), và 500 (D) không được lặp lại nhiều lần.
    - Nếu cần biểu diễn ký hiệu lặp lại **4 lần**, sử dụng dạng trừ.

Cho một số nguyên, hãy chuyển đổi nó thành số La Mã.

Điều kiện: 1 <= num <= 3999

### Phân tích
Có thể tạo một bảng ánh xạ theo các số thập phân hàng nghìn, hàng trăm, hàng chục, hàng đơn vị như sau:

| hàng   | Giá trị ánh xạ                                         |
|--------|-------------------------------------------------------|
| Hàng nghìn | `["", "M", "MM", "MMM"]`                              |
| Hàng trăm | `["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"]` |
| Hàng chục | `["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"]` |
| Hàng đơn vị | `["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"]` |

Chia số đầu vào cho 1000, 100, 10 lấy phần nguyên là có thể suy ra được kết quả.
