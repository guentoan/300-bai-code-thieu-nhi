## 13. Roman to Integer

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/roman-to-integer)

### Đề bài
Số La Mã được biểu diễn bằng 7 ký hiệu khác nhau:

| **Ký hiệu** | **Giá trị** |
|-------------|-------------|
| I           | 1           |
| V           | 5           |
| X           | 10          |
| L           | 50          |
| C           | 100         |
| D           | 500         |
| M           | 1000        |

Ví dụ:
- Số **2** được viết là `II` trong số La Mã, là tổng của hai số 1 (I + I).
- Số **12** được viết là `XII`, là tổng của X (10) và II (2).
- Số **27** được viết là `XXVII`, là tổng của XX (20), V (5), và II (2).

Số La Mã thường được viết từ **lớn đến nhỏ**, từ trái sang phải. Tuy nhiên, có một số trường hợp **phép trừ** được sử dụng thay vì phép cộng.

Ví dụ:
- Số **4** không được viết là `IIII`, mà được viết là `IV`. Vì chữ I đứng trước chữ V (5), nên ta **trừ 1 cho 5** để được 4.
- Tương tự, số **9** được viết là `IX` (trừ 1 cho 10).

1. `I` có thể đứng trước `V` (5) và `X` (10) để tạo thành **4** và **9**.
2. `X` có thể đứng trước `L` (50) và `C` (100) để tạo thành **40** và **90**.
3. `C` có thể đứng trước `D` (500) và `M` (1000) để tạo thành **400** và **900**.

Cho một số La Mã, hãy chuyển đổi nó thành số nguyên.

### Ví dụ
- Input: `"III"`  
  Output: `3`

- Input: `"IV"`  
  Output: `4`

- Input: `"IX"`  
  Output: `9`

- Input: `"LVIII"`  
  Output: `58`  
  **Giải thích**: L = 50, V = 5, III = 3. Tổng: 50 + 5 + 3 = 58.

- Input: `"MCMXCIV"`  
  Output: `1994`  
  **Giải thích**:
    - M = 1000
    - CM = 900
    - XC = 90
    - IV = 4  
      Tổng: 1000 + 900 + 90 + 4 = 1994.

### Phân tích
Ta có thể giải bài toán này bằng cách duyệt qua các ký tự của chuỗi đầu vào.
Xác định giá trị số nguyên của ký tự hiện tại và ký tự tiếp theo (nếu tồn tại).
Nếu ký tự hiện tại nhỏ hơn ký tự tiếp theo (ví dụ `I` trước `V`, hoặc `X` trước `L`), ta thực hiện phép trừ (ví dụ: `IV = 5 - 1 = 4`)
Nếu ký tự hiện tại lớn hơn hoặc bằng ký tự tiếp theo, ta thực hiện phép cộng.
Cộng dồn lại ta sẽ được kết quả.
