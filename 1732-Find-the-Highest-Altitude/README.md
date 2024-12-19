## 1732. Find the Highest Altitude

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/find-the-highest-altitude/description)

### Đề bài

Một người đi xe đạp đang thực hiện một chuyến đi đường trường. Chuyến đi bao gồm **n + 1 điểm** ở các độ cao khác nhau. Người đi xe đạp bắt đầu hành trình từ **điểm 0** với độ cao ban đầu là **0**.

Bạn được cung cấp một mảng số nguyên `gain` có độ dài **n**, trong đó `gain[i]` là độ thay đổi **độ cao ròng** (**net gain in altitude**) giữa các điểm **i** và **i + 1** với mọi **0 <= i < n**.

Nhiệm vụ của bạn là **trả về độ cao lớn nhất** tại một điểm bất kỳ trong chuyến đi.


### Ví dụ
**Ví dụ 1:**
- **Input**: `gain = [-5, 1, 5, 0, -7]`
- **Output**: `1`
- **Giải thích**:
    - Độ cao tại các điểm lần lượt là:
        - Điểm 0: 0
        - Điểm 1: 0 + (-5) = -5
        - Điểm 2: -5 + 1 = -4
        - Điểm 3: -4 + 5 = 1
        - Điểm 4: 1 + 0 = 1
        - Điểm 5: 1 + (-7) = -6
    - Độ cao lớn nhất là `1`.

**Ví dụ 2:**
- **Input**: `gain = [-4, -3, -2, -1, 4, 3, 2]`
- **Output**: `0`
- **Giải thích**:
    - Độ cao tại các điểm lần lượt là:
        - Điểm 0: 0
        - Điểm 1: 0 + (-4) = -4
        - Điểm 2: -4 + (-3) = -7
        - Điểm 3: -7 + (-2) = -9
        - Điểm 4: -9 + (-1) = -10
        - Điểm 5: -10 + 4 = -6
        - Điểm 6: -6 + 3 = -3
        - Điểm 7: -3 + 2 = -1
    - Độ cao lớn nhất là `0`.

### Phân tích
Bài này có thể dùng biến `high` để tính độ cao tại mỗi điểm, và `maxHigh` để lưu trữ điểm cao nhất trong hành trình. 