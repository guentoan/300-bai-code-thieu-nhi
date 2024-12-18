## 11. Container With Most Water

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/container-with-most-water/description)

### Đề bài
Bạn được cung cấp một mảng số nguyên `height` có độ dài `n`. Có `n` đường thẳng dọc được vẽ sao cho hai điểm kết thúc của đường thẳng thứ `i` là `(i, 0)` và `(i, height[i])`.

Hãy tìm hai đường thẳng, cùng với trục x, tạo thành một vùng chứa nước sao cho vùng chứa này có thể chứa được lượng nước lớn nhất.

Trả về lượng nước tối đa mà vùng chứa có thể lưu trữ.

**Lưu ý**: Bạn không được làm nghiêng vùng chứa.


### Phân tích
Dùng hai con trỏ `left` và `right`, sau đó tính diện tích vùng chứa nước: `s = (right-left)*min(height[left],height[right)])`.
Di chuyển con trỏ và so sánh diện tích để được lượng nước tối đa.
Cách di chuyển con trỏ theo logic sau: Nếu `height[left] < height[right]` tăng `left`, ngược lại giảm `right` (mục đích là tìm thành của thùng chứa cao hơn)
