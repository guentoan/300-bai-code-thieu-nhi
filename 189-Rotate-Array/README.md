## 189. Rotate Array

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/rotate-array/description/)

### Mục Tiêu
Cho một mảng số nguyên `nums`, hãy xoay mảng sang phải `k` bước, trong đó `k` là một số không âm.

### Ý Tưởng
với mảng `nums` có kích thước là `n`, thì sau khi xoay `n` lần thì mảng đó sẽ quay lại trạng thái ban đầu.
Vì thế, trước tiên ta sẽ tính `k=k%n`.
vì mảng `nums` bị xoay k lần, nên các phần tử `nums[n-k:n-1]` sẽ dịch lên đầu mảng, còn các phần tử `nums[0:k]` sẽ dịch chuyển xuống cuối mảng.
để giải quyết vấn đề này mà không tạo thêm bộ nhớ (`in-place`), 
ta có thể đảo ngược toàn bộ mảng, sau đó đảo ngược `nums[0:k-1]`, và đảo ngược `nums[k:n-1]`
