## 135. Candy

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/candy/description)

### Đề bài
Có **n** đứa trẻ đứng thành một hàng. Mỗi đứa trẻ được gán một giá trị đánh giá, được biểu diễn dưới dạng mảng số nguyên `ratings`, trong đó `ratings[i]` là giá trị đánh giá của đứa trẻ thứ `i`.

Bạn cần phát kẹo cho các đứa trẻ sao cho thỏa mãn các điều kiện sau:

1. Mỗi đứa trẻ phải nhận ít nhất **một viên kẹo**.
2. Những đứa trẻ có đánh giá cao hơn những đứa bên cạnh của mình phải nhận được **nhiều kẹo hơn**.

Hãy trả về số lượng kẹo tối thiểu cần thiết để phân phát cho tất cả các đứa trẻ.

### Ví dụ

**Ví dụ 1:**

**Input**:  
`ratings = [1, 0, 2]`  
**Output**:  
`5`

**Giải thích**:
- Đứa trẻ thứ 2 có đánh giá thấp nhất nên nhận 1 viên kẹo.
- Đứa trẻ thứ 1 có đánh giá cao hơn đứa trẻ thứ 2 nên nhận nhiều hơn (2 viên).
- Đứa trẻ thứ 3 cũng nhận 2 viên kẹo vì đánh giá cao hơn đứa trẻ thứ 2.
- Tổng số kẹo là: `1 + 2 + 2 = 5`.

**Ví dụ 2:**

**Input**:  
`ratings = [1, 2, 2]`  
**Output**:  
`4`

**Giải thích**:
- Đứa trẻ đầu tiên nhận 1 viên kẹo.
- Đứa trẻ thứ hai nhận 2 viên kẹo vì đánh giá cao hơn đứa trẻ đầu tiên.
- Đứa trẻ thứ ba nhận 1 viên kẹo vì đánh giá bằng đứa trẻ thứ hai.
- Tổng số kẹo là: `1 + 2 + 1 = 4`.

### Phân tích
