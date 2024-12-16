## 122. Best Time to Buy and Sell Stock II

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/)

### Mục Tiêu
Bạn được cung cấp một mảng số nguyên `prices`, trong đó `prices[i]` là giá của một cổ phiếu vào ngày thứ `i`.
Vào mỗi ngày, bạn có thể quyết định mua và/hoặc bán cổ phiếu. Tuy nhiên, bạn chỉ được giữ tối đa một cổ phiếu tại bất kỳ thời điểm nào. Bạn cũng có thể mua và bán ngay trong cùng một ngày.
Hãy tìm và trả về lợi nhuận tối đa mà bạn có thể đạt được.

### Ý Tưởng
Để tạo được lợi nhuận thì ta cần mua cổ phiếu khi giá thấp và bán khi giá cao. Vậy để đạt được lợi nhuận tối đa, ta cần tính lợi nhuận giữa các ngày, sau đó cộng tất cả lợi nhuận đó lại.