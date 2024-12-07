## 1760. Minimum Limit of Balls in a Bag

You are given an integer array `nums` where the `i` bag contains `nums[i]` balls. You are also given an integer `maxOperations`.

You can perform the following operation at most `maxOperations` times:

 - Take any bag of balls and divide it into two new bags with a positive number of balls.
    - For example, a bag of 5 balls can become two new bags of 1 and 4 balls, or two new bags of 2 and 3 balls.

Your penalty is the **maximum** number of balls in a bag. You want to **minimize** your penalty after the operations.

Return *the minimum possible penalty after performing the operations*.

**Example 1:**
```
Input: nums = [9], maxOperations = 2
Output: 3
Explanation:
- Divide the bag with 9 balls into two bags of sizes 6 and 3. [9] -> [6,3].
- Divide the bag with 6 balls into two bags of sizes 3 and 3. [6,3] -> [3,3,3].
The bag with the most number of balls has 3 balls, so your penalty is 3 and you should return 3.
```

**Example 2:**
```
Input: nums = [2,4,8,2], maxOperations = 4
Output: 2
Explanation:
- Divide the bag with 8 balls into two bags of sizes 4 and 4. [2,4,8,2] -> [2,4,4,4,2].
- Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,4,4,4,2] -> [2,2,2,4,4,2].
- Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,4,4,2] -> [2,2,2,2,2,4,2].
- Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2].
The bag with the most number of balls has 2 balls, so your penalty is 2, and you should return 2.
```

**Constraints:**
 - 1 <= nums.length <= 10^5
 - 1 <= maxOperations, nums[i] <= 10^9


### mục tiêu
tìm penalty nhỏ nhất(giá trị lớn nhất trong các túi bóng) có thể sau khi thực hiện tối đa `maxOperations` thao tác.

### ý tưởng

 - giả sử `p` là giá trị kỳ vọng, khi đó tất cả các túi bóng lớn hơn `p` cần được chia nhỏ để không có túi nào vượt quá `p`.
 - mỗi lần chia một túi thành `k` phần, cần `k-1` thao tác.

=> mục tiêu sẽ là giảm `p` xuống nhỏ nhất có thể, và đảm bảo số thao tác không vượt quá `maxOperations`.

 - ta có thể thấy `p` nhỏ nhất là 1 (trường hợp tất cả các túi có số lượng bằng 1). gọi là `low`.
 - `p` lớn nhất là `max(nums)`. gọi là `high`

=> mục tiêu là tìm một số nhỏ nhất có thể, nằm giữa `low` và `high`, và đảm bảo điều kiện số thao tác nhỏ hơn hoặc bằng `maxOperations`.

### Cách tiếp cận

1. **Xác định phạm vi penalty:**
   - Giá trị penalty nhỏ nhất là `low = 1`.
   - Giá trị penalty lớn nhất là `high = max(nums)`.

2. **Binary Search để tìm penalty tối ưu:**
   - Duyệt tất cả giá trị penalty `p` trong phạm vi `[low, high]`.
   - Kiểm tra liệu có thể chia các túi với penalty `p` mà không vượt quá `maxOperations` thao tác.
   - Nếu có thể chia các túi với `p` thì ta sẽ gán `high = p`, còn không thì `low = p + 1`
   - Bài toán sẽ duyệt đến khi low >= high => low là giá trị penalty nhỏ nhất có thể đạt được. 

3. **Hàm kiểm tra khả năng chia:**
   - Với penalty `p`, mỗi túi có `n` bóng sẽ cần `n/p - 1` thao tác để đảm bảo tất cả túi `≤ p`.
   - Tổng số thao tác không được vượt quá `maxOperations`.

### Phân tích độ phức tạp
1. Time complexity:
 -  Binary Search thực hiện trong `𝑂(log(max(nums)))
 - Kiểm tra khả năng chia mất O(i), i là số lượng túi bóng.
 - Tổng thời gian: O(n⋅log(max(nums)))
2. Space complexity: O(1)
