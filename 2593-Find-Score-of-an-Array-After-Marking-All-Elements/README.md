## 2593. Find Score of an Array After Marking All Elements

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/find-score-of-an-array-after-marking-all-elements/description/)

### Mục Tiêu
Bạn được cho một mảng `nums` bao gồm các số nguyên dương.

Bắt đầu với `score = 0`, áp dụng thuật toán sau:

 - Chọn số nguyên nhỏ nhất trong mảng mà chưa được đánh dấu. Nếu có nhiều số cùng giá trị nhỏ nhất, chọn số có chỉ số nhỏ nhất.
 - Thêm giá trị của số được chọn vào `score`.
 - Đánh dấu số đã chọn và hai phần tử liền kề của nó nếu có.
 - Lặp lại cho đến khi tất cả các phần tử trong mảng đều được đánh dấu.

Trả về điểm số bạn nhận được sau khi áp dụng thuật toán trên.

### Ý Tưởng
Với bài toán này, chúng ta có thể sử dụng `MinHeap` hoặc `sortedList`.
Với cách nào thì chúng ta đều cần phải có một mảng để đánh dấu các phần tử đã được chọn.
