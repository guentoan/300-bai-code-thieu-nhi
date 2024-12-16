## 121. Best Time to Buy and Sell Stock

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/)

### Mục Tiêu
Bạn được cung cấp một mảng `prices`, trong đó `prices[i]` là giá của một cổ phiếu vào ngày thứ `i`.
Bạn muốn tối đa hóa lợi nhuận của mình bằng cách chọn một ngày để mua một cổ phiếu và chọn một ngày khác trong tương lai để bán cổ phiếu đó.
Hãy trả về lợi nhuận tối đa mà bạn có thể đạt được từ giao dịch này. Nếu không thể đạt được lợi nhuận nào, trả về 0.

### Ý Tưởng
Ta có thể nhận thấy khi mua cổ phiếu với giá thấp và bán với giá cao thì ta sẽ có lợi nhuận.
Vậy để đạt được lợi nhuận tối đa, ta cần phải mua ở ngày cổ phiếu có giá thấp nhất, và bán ra ở ngày cổ phiếu có giá cao nhất.
Để làm được điều đó, ta cần có 2 biến để tracking gía thấp nhất của cổ phiếu, và lợi nhuận tối đa.
Sau đó ta sẽ duyệt qua mảng `prices`. Tại mỗi bước:
 - Cập nhật giá thấp nhất (min price) nếu giá hiện tại nhỏ hơn giá thấp nhất trước đó.
 - Tính lợi nhuận tại ngày hiện tại: `profit = prices[i] - min price`
 - Cập nhật lợi nhuận lớn nhất nếu lợi nhuận hiện tại lớn hơn lợi nhuận lớn nhất trước đó.

Sau khi duyệt hết mảng, trả về lợi nhuận lớn nhất.
