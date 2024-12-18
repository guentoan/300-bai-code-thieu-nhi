## 1475. Final Prices With a Special Discount in a Shop

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/description)

### Đề bài
Bạn được cho một mảng số nguyên `prices` trong đó `prices[i]` là giá của món hàng thứ `i` trong cửa hàng.

Có một chương trình giảm giá đặc biệt cho các món hàng trong cửa hàng. Nếu bạn mua món hàng thứ i, bạn sẽ nhận được một khoản giảm giá tương đương với `prices[j]`, trong đó `j` là chỉ số nhỏ nhất sao cho `j > i` và `prices[j] <= prices[i]`. Nếu không có chỉ số nào thỏa mãn, bạn sẽ không nhận được bất kỳ giảm giá nào.

Trả về một mảng số nguyên `answer` trong đó `answer[i]` là giá cuối cùng mà bạn phải trả cho món hàng thứ i, xét đến chương trình giảm giá đặc biệt.

### Phân tích
Trước tiên, ta sẽ tạo một mảng `answer` là copy của mảng `prices`.
Tạo một mảng `stack`, dùng để lưu gía trị của các phần tử chưa tìm thấy phần tử nhỏ hơn hoặc bằng giá của chúng.
Duyệt mảng `prices`, nếu `prices[i] <= prices[stack]`: áp dụng mã giảm giá.