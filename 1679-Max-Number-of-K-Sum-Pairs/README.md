## 1679. Max Number of K-Sum Pairs

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/max-number-of-k-sum-pairs/description)

### Đề bài
Bạn được cho một mảng số nguyên `nums` và một số nguyên `k`.

Trong một phép toán, bạn có thể chọn hai số từ mảng có tổng bằng `k` và xóa chúng khỏi mảng.

Hãy trả về số lượng phép toán tối đa mà bạn có thể thực hiện trên mảng.

### Ví dụ
**Input**:  
`nums = [1,2,3,4]`, `k = 5`  
**Output**:  
`2`  
**Giải thích**:
- Chọn 1 và 4 (tổng bằng 5) -> Xóa.
- Chọn 2 và 3 (tổng bằng 5) -> Xóa.
- Không còn cặp nào thỏa mãn, tổng số phép toán là 2.

**Input**:  
`nums = [3,1,3,4,3]`, `k = 6`  
**Output**:  
`1`  
**Giải thích**:
- Chọn 3 và 3 (tổng bằng 6) -> Xóa.
- Không còn cặp nào thỏa mãn, tổng số phép toán là 1.


### Phân tích
**Cách 1:** bài toán này có thể giải quyết bằng cách sắp xếp mảng `nums` và sử dụng hai con trỏ `left` và `right` để tìm các cặp số có tổng bằng `k`.

**Cách 2:** sử dụng hash map để lưu các phần tử trong `nums` và tần số xuất hiện của mỗi phần tử. Tại mỗi vị trí, kiểm tra xem trong map có chứa phần tử `k-nums[i]` không?
Nếu có, thì ta có một cặp số thoả mã điều kiện. Nếu không thì thêm `nums[i]` vào map.
