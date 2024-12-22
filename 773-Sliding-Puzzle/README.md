## 773. Sliding Puzzle

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/sliding-puzzle)

### Đề bài
Trên một bảng 2 x 3, có 5 ô vuông được đánh số từ 1 đến 5, và một ô trống được biểu diễn bởi số 0. Mỗi lần di chuyển bao gồm việc chọn số 0 và một số liền kề theo 4 hướng (trên, dưới, trái, phải), sau đó hoán đổi vị trí của chúng.

Trạng thái của bảng được xem là đã giải quyết(solved) nếu và chỉ nếu bảng có dạng `[[1,2,3],[4,5,0]]`.

Cho trạng thái ban đầu của bảng `board`, hãy trả về số bước ít nhất cần thiết để đưa bảng về trạng thái đã giải quyết. Nếu không thể đưa bảng về trạng thái đã giải quyết, hãy trả về `-1`.


### Phân tích
Ý tưởng là duyệt qua tất cả các trạng thái mà `board` có thể tạo ra. Nếu gặp trạng thái solved thì trả về số bước đã duỵệt, không thì sẽ trả về `-1`.
