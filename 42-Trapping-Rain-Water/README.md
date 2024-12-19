## 42. Trapping Rain Water

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/trapping-rain-water/description)

### Đề bài
Cho `n` số nguyên không âm đại diện cho một bản đồ độ cao, trong đó chiều rộng của mỗi thanh là 1, hãy tính xem lượng nước mà nó có thể giữ lại được sau khi mưa là bao nhiêu.

### Phân tích
Ta có thể nhận ra rằng, lượng nước tại một điểm **không thể vượt qua mức thấp hơn của hai bức tường bên trái và bên phải**.  
Điều này xảy ra vì nếu nước vượt qua ngưỡng thấp hơn, nó sẽ bị tràn ra ngoài.

Do đó, để giải quyết bài toán, ta có thể sử dụng phương pháp **two-pointers**.
  - `left`: bắt đầu từ đầu mảng (vị trí 0).
  - `right`: bắt đầu từ cuối mảng (vị trí `n - 1`).
  - `maxLeft`: lưu chiều cao lớn nhất của tường bên trái tại mỗi bước.
  - `maxRight`: lưu chiều cao lớn nhất của tường bên phải tại mỗi bước.

