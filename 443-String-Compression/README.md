## 443. String Compression

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/string-compression/description)

### Đề bài
Cho một mảng ký tự `chars`, hãy nén mảng này theo thuật toán sau:

1. Bắt đầu với một chuỗi rỗng `s`.
2. Với mỗi nhóm ký tự giống nhau liên tiếp trong `chars`:
    - Nếu độ dài của nhóm là 1, thêm ký tự đó vào `s`.
    - Ngược lại, thêm ký tự đó vào `s`, sau đó là độ dài của nhóm.

Chuỗi đã nén `s` không được trả về riêng biệt, mà thay vào đó, phải được lưu trong mảng ký tự đầu vào `chars`. Lưu ý rằng độ dài của các nhóm nếu lớn hơn hoặc bằng 10 phải được chia thành nhiều ký tự trong `chars`.

Sau khi hoàn tất việc sửa đổi mảng đầu vào, hãy trả về **độ dài mới của mảng**.

Bạn cần viết một thuật toán sử dụng **bộ nhớ bổ sung cố định** (constant extra space).


### Phân tích
Ta có thể giải quyết bài toán theo các bước sau:
- Xử lý từng nhóm ký tự liên tiếp trong mảng.
- Tính độ dài nhóm và chèn ký tự tương ứng cùng độ dài của nhóm vào mảng đầu vào.
- Sử dụng chỉ số để theo dõi vị trí hiện tại khi ghi vào mảng `chars`.
