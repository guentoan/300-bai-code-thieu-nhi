## 1004. Max Consecutive Ones III

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/max-consecutive-ones-iii)

### Đề bài

Cho một mảng nhị phân `nums` và một số nguyên `k`, hãy trả về số lượng lớn nhất các số 1 liên tiếp trong mảng nếu bạn có thể lật tối đa `k` số 0 thành số 1.

### Phân tích

Bài toán có thể hiểu là tìm một khoảng con lớn nhất của mảng `nums` sao cho số lượng `0` trong mảng không vượt quá `k`.
Vì thế ta cần hai biến `left`, `right` để duy trì `slide window`, một biến `zeroCount` để đếm số phần tử bằng 0, và một biến `maxLength` để lưu độ dài lớn nhất của mảng con.
