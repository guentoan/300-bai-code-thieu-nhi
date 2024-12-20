## 68. Text Justification

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/text-justification/description)

### Đề bài

Cho một mảng các chuỗi `words` và một số nguyên `maxWidth`, hãy định dạng văn bản sao cho mỗi dòng có chính xác `maxWidth` ký tự và được căn chỉnh đầy đủ (cả căn trái và căn phải).

Bạn cần đóng gói các từ vào mỗi dòng theo cách tham lam; tức là, đóng gói càng nhiều từ càng tốt vào mỗi dòng. Thêm khoảng trống `' '` nếu cần thiết để mỗi dòng có chính xác `maxWidth` ký tự.

Khoảng trống giữa các từ cần được phân bổ càng đều càng tốt. Nếu số khoảng trống trên một dòng không chia đều giữa các từ, các vị trí trống bên trái sẽ được cấp thêm khoảng trống so với các vị trí trống bên phải.

Đối với dòng cuối cùng của văn bản, nó sẽ được căn trái, và không có khoảng trống thêm giữa các từ.

Ghi chú:

- Một từ được định nghĩa là một chuỗi ký tự chỉ bao gồm các ký tự không phải là khoảng trống.
- Độ dài của mỗi từ được đảm bảo là lớn hơn 0 và không vượt quá `maxWidth`.
- Mảng đầu vào `words` chứa ít nhất một từ.

### Phân tích
Để giải quyết bài toán này, chúng ta cần phải xử lý mỗi dòng theo các bước sau:
1. **Duyệt qua các từ**: Chúng ta cần duyệt qua các từ trong mảng `words` và tính toán độ dài của mỗi dòng.
2. **Căn chỉnh dòng**:
    - Khi thêm một từ vào dòng mà tổng chiều dài của dòng vượt quá `maxWidth`, ta sẽ cần xử lý dòng trước đó và bắt đầu một dòng mới.
    - Đối với các dòng không phải cuối cùng, chúng ta cần phân phối khoảng trống giữa các từ sao cho các khoảng trống này được phân phối đều. Phần dư từ phép chia số khoảng trống cho số từ sẽ được phân bổ vào các từ bên trái.
3. **Dòng cuối cùng**: Dòng cuối cùng sẽ được căn trái, không cần phân phối khoảng trống giữa các từ.
