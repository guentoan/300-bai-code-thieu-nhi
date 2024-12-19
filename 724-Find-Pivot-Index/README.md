## 724. Find Pivot Index

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/find-pivot-index)

### Đề bài
Cho một mảng số nguyên `nums`, hãy tính **pivot index** của mảng này.

**pivot index** là một chỉ số mà tổng của tất cả các phần tử ở bên **trái** của chỉ số đó bằng tổng của tất cả các phần tử ở bên **phải** của chỉ số đó.

- Nếu chỉ số nằm ở cạnh trái của mảng, thì tổng bên trái sẽ là `0` (vì không có phần tử nào ở bên trái).
- Quy tắc tương tự áp dụng cho cạnh phải của mảng: tổng bên phải sẽ là `0` nếu chỉ số nằm ở cạnh phải.

Hãy trả về **pivot index nằm ở vị trí bên trái nhất**. Nếu không tồn tại chỉ số trục nào, trả về `-1`.

### Ví dụ

**Ví dụ 1:**
- **Input**: `nums = [1, 7, 3, 6, 5, 6]`
- **Output**: `3`
- **Giải thích**:
    - Tổng các phần tử bên trái của chỉ số `3`: `1 + 7 + 3 = 11`
    - Tổng các phần tử bên phải của chỉ số `3`: `5 + 6 = 11`

**Ví dụ 2:**
- **Input**: `nums = [1, 2, 3]`
- **Output**: `-1`
- **Giải thích**:
    - Không có chỉ số nào mà tổng bên trái bằng tổng bên phải.

**Ví dụ 3:**
- **Input**: `nums = [2, 1, -1]`
- **Output**: `0`
- **Giải thích**:
    - Tổng các phần tử bên trái của chỉ số `0`: `0` (không có phần tử bên trái).
    - Tổng các phần tử bên phải của chỉ số `0`: `1 + (-1) = 0`.

### Phân tích
Ta có thể tính tổng toàn bộ các phần tử mảng.
Sau đó duyệt từ trái qua phải, tính tổng bên trái. So sánh tổng bên trái với tổng bên phải bằng cách `sumRight = sum - sumLeft - nums[i]`.
