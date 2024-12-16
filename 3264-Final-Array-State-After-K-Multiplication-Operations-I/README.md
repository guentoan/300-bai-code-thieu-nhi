## 3264. Final Array State After K Multiplication Operations I

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/description/)

### Mục Tiêu
Bạn được cho một mảng số nguyên `nums` và hai số nguyên `k` và `multiplier`.
Bạn cần thực hiện `k` thao tác trên mảng `nums`. Trong mỗi thao tác:

 - Tìm giá trị nhỏ nhất `x` trong mảng nums. Nếu có nhiều giá trị nhỏ nhất giống nhau, chọn giá trị đầu tiên xuất hiện.
 - Thay giá trị nhỏ nhất `x` bằng `x * multiplier`

Sau khi thực hiện tất cả `k` thao tác, trả về một mảng số nguyên đại diện cho trạng thái cuối cùng của `nums`.

### Ý Tưởng
Có thể duyệt qua `nums` để tìm giá trị nhỏ nhất, sau đó cập nhật lại giá trị tại vị trí đó.
