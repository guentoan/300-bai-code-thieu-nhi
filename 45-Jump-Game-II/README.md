## 45. Jump Game II

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/jump-game-ii/description/)

### Mục Tiêu
Hãy trả về **số lần nhảy tối thiểu** để bạn có thể đến được vị trí `nums[n - 1]`.

### Ý Tưởng
Ý tưởng là sử dụng biến `maxReach` để lưu lại vị trí nhảy xa nhất có thể tại vị trí `i`.
`current` để lưu lại vị trí hiện tại đang đứng. `count` để lưu lại số lần đã nhảy.

Duyệt qua mảng `nums`. Cập nhật `maxReach` tại mỗi bước. Nếu `curent = i`, ta sẽ tăng `count` và cập nhật `current = maxReach`.
Khi duyệt hết mảng thì `count` là số lần nhảy tối thiểu cần thiết để đến cuối mảng.
