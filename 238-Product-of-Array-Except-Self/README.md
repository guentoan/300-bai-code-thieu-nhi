## 238. Product of Array Except Self

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150)

### Mục Tiêu

Cho một mảng số nguyên `nums`, hãy trả về một mảng `answer` sao cho `answer[i]` bằng tích của tất cả các phần tử trong `nums` ngoại trừ `nums[i]`.

- Tích của bất kỳ dãy con nào (prefix hoặc suffix) trong `nums` được đảm bảo nằm trong phạm vi của số nguyên 32-bit.
- Bạn phải viết một thuật toán có độ phức tạp thời gian là `O(n)` và **không được sử dụng phép chia**.

### Ví dụ 1:
**Input**:  
`nums = [1,2,3,4]`  
**Output**:  
`[24,12,8,6]`  
**Giải thích**:
- `answer[0] = 2 * 3 * 4 = 24`
- `answer[1] = 1 * 3 * 4 = 12`
- `answer[2] = 1 * 2 * 4 = 8`
- `answer[3] = 1 * 2 * 3 = 6`

### Ví dụ 2:
**Input**:  
`nums = [-1,1,0,-3,3]`  
**Output**:  
`[0,0,9,0,0]`  
**Giải thích**:
- `answer[0] = 1 * 0 * -3 * 3 = 0`
- `answer[1] = -1 * 0 * -3 * 3 = 0`
- `answer[2] = -1 * 1 * -3 * 3 = 9`
- `answer[3] = -1 * 1 * 0 * 3 = 0`
- `answer[4] = -1 * 1 * 0 * -3 = 0`

### Ý Tưởng
ta có thể tạo một mảng `answer` có kích thước bằng mảng `nums`.
Sau đó ta duyệt từ `0 -> n-1` và lưu tích của các phần tử từ `0 -> i-1` vào phần tử `i` của mảng `answer`.
Sau đó ta lại duyệt từ `n-1 -> 0` và lưu tích của các phần tử từ `n-1 -> i+i` vào phần tử `i` của mảng `answer`.