## 643. Maximum Average Subarray I

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/maximum-average-subarray-i/description)

### Đề bài

Bạn được cho một mảng số nguyên `nums` bao gồm `n` phần tử và một số nguyên `k`.

Hãy tìm một **mảng con liên tiếp** có độ dài bằng `k` sao cho giá trị trung bình của mảng con này là lớn nhất và trả về giá trị trung bình đó.

Bất kỳ câu trả lời nào có sai số tính toán nhỏ hơn `10⁻⁵` đều được chấp nhận.

### Ví dụ
- **Input**:  
  `nums = [1,12,-5,-6,50,3], k = 4`
- **Output**:  
  `12.75000`
- **Giải thích**:  
  Mảng con liên tiếp có độ dài 4 là `[12,-5,-6,50]`, giá trị trung bình của nó là `(12 + (-5) + (-6) + 50) / 4 = 51 / 4 = 12.75`.

### Phân tích
Sử dụng phương pháp `slice window` để giải bài toán.
Bắt đầu bằng việc tính tổng của mảng con đầu tiên có độ dài `k`.
Sau đó trượt cửa sổ sang bên phải của dãy `nums` để tìm `maxSum`. 
