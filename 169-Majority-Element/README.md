## 169. Majority Element

chi tiết đề bài xem [tại đây](https://leetcode.com/problems/majority-element/description/)

### Vấn đề

Cho một mảng số nguyên `nums` có size là `n`. tìm giá trị `majority`.

`Majority` là số có số lần xuất hiện nhiều hơn n/2 lần trong một mảng.

### ý tưởng.
Do `majority` là số có số lần xuất hiện nhiều hơn n/2 lần trong một mảng, nên ta có thể duyệt qua mảng và count, nếu số nào xuất hiện nhiều hơn n/2 thì số đó là `majority`.

Chi tiết các bước của thuật toán như sau:
 - Chọn `count` = 0, và `ans` = 0
 - Duyệt qua từng phần tử của mảng `nums`.
    - Nếu `count == 0`: gán `ans = nums[i]` và tăng `count` lên 1
    - Nếu `count != 0` thì ta sẽ so sánh `ans` với `nums[i]`.
      - Nếu `ans == nums[i]`: Tăng `count`.
      - Nếu không thì giảm `count`.

Cách xử lý trên hoạt động là do: `majority` có số lần xuất hiện lớn hơn n/2 lần, nên dùng phương pháp loại trừ trên thì phần tử `majority` sẽ luôn còn xót lại sau cùng.
