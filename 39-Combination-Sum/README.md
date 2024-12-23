## 39. Combination Sum

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/combination-sum/description/)

### Đề bài
Cho một mảng các số nguyên khác nhau `candidates` và một số nguyên `target`, hãy trả về danh sách tất cả các tổ hợp duy nhất của các số trong `candidates` sao cho tổng của các số trong mỗi tổ hợp bằng `target`. Bạn có thể trả về các tổ hợp theo bất kỳ thứ tự nào.

Số giống nhau có thể được chọn từ mảng `candidates` một số lần không giới hạn. Hai tổ hợp được xem là duy nhất nếu tần suất xuất hiện của ít nhất một số trong các số được chọn là khác nhau.

Các trường hợp kiểm thử được tạo ra sao cho số lượng các tổ hợp duy nhất có tổng bằng `target` ít hơn 150 tổ hợp đối với đầu vào đã cho.


### Phân tích
Ta có thể sử dụng phương pháp **backtrack** để giải quyết bài toán này.

- Ta sẽ duyệt qua các số trong mảng `candidates`.
- Với mỗi số, ta có thể chọn nếu số đó nhỏ hơn `target`. Với số được chọn, ta tiếp tục tìm kiếm tổ hợp tiếp theo với số đó được thêm vào tổ hợp, và tiếp tục cho đến khi tổng đạt đến `target` hoặc vượt qua nó.
- Vì các số có thể được sử dụng nhiều lần, khi chọn một số, ta sẽ không bỏ qua các số ở phía sau trong mảng.
- Khi tìm ra tổ hợp có tổng bằng `target`, ta sẽ lưu lại tổ hợp đó.
- Cuối cùng, trả về tất cả các tổ hợp duy nhất.