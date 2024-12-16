## 2762. Continuous Subarrays

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/continuous-subarrays/description/)

### Mục Tiêu
Bạn được cung cấp một mảng số nguyên `nums` được đánh chỉ số từ 0. Một mảng con của `nums` được gọi là mảng con liên tục (continuous) nếu:

`i, i + 1, ..., j` là các chỉ số trong mảng con. khi đó với mỗi cặp chỉ số `i <= i1, i2 <= j` thoả mãn điều kiện: `0 <= abs(nums[i1]-nums[i2]) <= 2`

Hãy trả về tổng số lượng mảng con liên tục(continuous).

### Ý Tưởng
Sử dụng hai con trỏ `start` và `end` để xác định khoảng mảng con liên tục(continuous).
Với mỗi phần tử `nums[end]` trong mảng:
 - Xác định giá trị nhỏ nhất và lớn nhất trong mảng `[start, end]`
 - Nếu chênh lệch giữa giá trị lớn nhất và nhỏ nhất trong cửa sổ vượt quá 2 `(max − min) > 2` di chuyển con trỏ `start` sang phải
 - Số lượng mảng con kết thúc tại `end` sẽ là độ dài của mảng con liên tục hiện tại. (end−start+1)