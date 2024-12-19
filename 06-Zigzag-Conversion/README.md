## 6. Zigzag Conversion

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/zigzag-conversion/description/)

### Phân tích
Với dạng zigzag, các ký tự được lặp lại theo chu kỳ (gồm đi xuống và đi lên).
Một chu kỳ đầy đủ có độ dài `cycle = 2*(numRows-1)`.
Ví dụ: `numRows = 3` thì chu kỳ sẽ là `4`, `numRows = 4` chu kỳ sẽ là `6`.
Với mỗi hàng `row`, duyệt qua các ký tự thuộc hàng đó trong chuỗi `s` theo khoảng cách là `cycle`.
Trong trường hợp không phải là hàng đầu tiên hoặc cuối cùng, mỗi hàng sẽ có thêm ký tự xen giữa nằm trên đường đi lên của zigzag.
Vị trí của ký tự đó sẽ là: `midIndex := j + cycle - 2*row`.
