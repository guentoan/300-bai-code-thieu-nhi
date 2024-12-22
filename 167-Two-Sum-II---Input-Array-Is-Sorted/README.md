## 167. Two Sum II - Input Array Is Sorted

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description)

### Đề bài
Cho một mảng số nguyên đã được sắp xếp theo thứ tự không giảm với chỉ số bắt đầu từ 1, hãy tìm hai số sao cho tổng của chúng bằng một giá trị mục tiêu cụ thể. Hai số này là `numbers[index1]` và `numbers[index2]` trong đó `1 <= index1 < index2 <= numbers.length`.

Hãy trả về chỉ số của hai số, `index1` và `index2`, đã cộng thêm một, dưới dạng mảng số nguyên `[index1, index2]` có độ dài 2.

Các bài kiểm tra được tạo ra sao cho luôn có đúng một lời giải. Bạn không được sử dụng cùng một phần tử hai lần.

Giải pháp của bạn phải sử dụng chỉ không gian bổ sung không đổi.

### Phân tích
Có thể tiếp cận bài toán bằng phương pháp **Two-Pointer**.
Một con trỏ `left` bắt đầu từ đầu mảng (chỉ số 1).
Một con trỏ `right` bắt đầu từ cuối mảng.
Tính tổng của hai số tại các chỉ số `left` và `right`.
Nếu tổng bằng giá trị `target`, ta đã tìm được hai số cần tìm.
Nếu tổng nhỏ hơn `target` thì tăng `left`, ngược lại giảm `right`.