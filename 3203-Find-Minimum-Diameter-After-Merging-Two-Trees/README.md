## 3203. Find Minimum Diameter After Merging Two Trees

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/find-minimum-diameter-after-merging-two-trees/description/)

### Đề bài
Có hai cây vô hướng có `n` nodes và `m` nodes.
Đầu vào là 2 mảng số nguyên `edges1` và `edges2` có độ dài lần lượt là `n - 1` và `m - 1`.
Trong đó, `edges1[i] = [ai, bi]` chỉ ra rằng có một cạnh nối giữa hai nút `ai` và `bi` trong cây đầu tiên, 
và `edges2[i] = [ui, vi]` chỉ ra rằng có một cạnh nối giữa hai nút `ui` và `vi` trong cây thứ hai.

Bạn phải nối một nút từ cây đầu tiên với một nút từ cây thứ hai bằng một cạnh.

Hãy trả về **đường kính nhỏ nhất** có thể của cây sau khi đã nối.

**Đường kính** của cây là độ dài của đường đi dài nhất giữa hai nút bất kỳ trong cây.

### Phân tích

Ta có thể tìm đường kính của mỗi cây bằng cách duyệt DFS.
Tìm các nút xa nhất trong mỗi cây. Đây là các nút có thể được sử dụng để nối giữa hai cây, giúp giảm thiểu đường kính cây mới.
Nối các nút sao cho đường kính cây mới là nhỏ nhất.
Tìm đường kính cây mới sau khi nối hai cây lại.

