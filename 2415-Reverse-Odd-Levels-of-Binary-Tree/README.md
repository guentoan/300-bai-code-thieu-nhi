## 2415. Reverse Odd Levels of Binary Tree

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/description)

### Đề bài
Cho gốc của một cây nhị phân hoàn hảo, đảo ngược giá trị của các nút ở mỗi cấp lẻ của cây.

Ví dụ: Giả sử các giá trị nút tại cấp 3 là `[2,1,3,4,7,11,29,18]`, thì sau khi đảo ngược, chúng sẽ trở thành `[18,29,11,7,4,3,1,2]`.

Trả về gốc của cây sau khi các giá trị nút đã được đảo ngược ở các cấp lẻ.

Cây nhị phân được gọi là hoàn hảo nếu tất cả các nút cha đều có hai nút con và tất cả các lá đều nằm ở cùng một cấp.

Cấp của một nút là số cạnh dọc theo đường đi từ nút đó đến nút gốc của cây.

### Phân tích
Để giải quyết bài toán này ta có thể **duyệt cây theo chiều rộng(BFS)**, hoặc **duyệt cây theo chiều sâu (DFS)**.
