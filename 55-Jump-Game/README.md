## 55. Jump Game

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/jump-game/description/)

### Mục Tiêu
Bạn được cung cấp một mảng số nguyên `nums`. Bạn bắt đầu ở chỉ số đầu tiên của mảng, và mỗi phần tử trong mảng đại diện cho độ dài nhảy tối đa của bạn tại vị trí đó.

Hãy trả về `true` nếu bạn có thể đến được chỉ số cuối cùng của mảng, ngược lại trả về `false`.

### Ý Tưởng
Tại mỗi vị trí, ta có thể nhảy xa `n` bước, tức là vị trí lớn nhất của thể nhày đến là `i+n`.
Ta có thể sử dụng một biến `maxReach` để tracking vị trí xa nhất có thể nhảy được.
Tại mỗi vị trí ta sẽ cập nhật `maxReach = max(maxReach, n+i)`.
Nếu `maxReach > len(nums)` trả về `true`. Nếu `maxReach < i` thì trả về `false`.
