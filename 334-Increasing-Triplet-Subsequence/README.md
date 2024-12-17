## 334. Increasing Triplet Subsequence

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/increasing-triplet-subsequence/)

### Bài Toán
Cho một mảng số nguyên `nums`, hãy trả về `true` nếu tồn tại bộ ba chỉ số `(i, j, k)` thỏa mãn điều kiện:

 - i < j < k 
 - nums[i] < nums[j] < nums[k]

Nếu không tồn tại bộ ba nào như vậy, trả về `false`.

### Mục tiêu

Viết một thuật toán với độ phức tạp thời gian **O(n)** để giải quyết bài toán.

### Phân Tích
có thể sử dụng hai biến `first` và `second` để tracking gía trị nhỏ nhất và giá trị nhỏ thứ hai.
 - `first`: giữ giá trị nhỏ nhất tìm được trong mảng.
 - `second`: giữ giá trị nhỏ thứ hai lớn hơn `first`.
Ta sẽ duyệt qua mảng và thực hiện logic sau:
 - nếu phần tử hiện tại nhỏ hơn `first`, cập nhật `first`.
 - nếu phần tử hiện tại nằm giữa `first` và `second` cập nhật `second`.
 - nếu phần tử hiện tại lớn hơn second thì trả về `true`.
 - duyệt hết mảng mà chưa tìm được phần tử nào lớn hơn `second` thì trả về `false`.