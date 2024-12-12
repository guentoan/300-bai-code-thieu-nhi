## 80. Remove Duplicates from Sorted Array II

Given an integer array `nums` sorted in **non-decreasing order**, remove some duplicates [in-place](https://en.wikipedia.org/wiki/In-place_algorithm) such that each unique element appears **at most twice**. The **relative order** of the elements should be kept the **same**.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the **first part** of the array `nums`. More formally, if there are `k` elements after removing the duplicates, then the first `k` elements of `nums` should hold the final result. It does not matter what you leave beyond the first `k` elements.

Return `k` after placing the final result in the first `k` slots of `nums`.

Do **not** allocate extra space for another array. You must do this by **modifying the input array** in-place with **O(1)** extra memory.

**Custom Judge:**

The judge will test your solution with the following code:

```
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```
If all assertions pass, then your solution will be **accepted**.


**Example 1:**
```
Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

**Example 2:**
```
Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

**Constraints:**

 - 1 <= nums.length <= 3 * 10^4
 - -10^4 <= nums[i] <= 10^4
 - `nums` is sorted in non-decreasing order.

## mục tiêu
Xóa các phần tử trùng lặp quá 2 lần trong mảng đã sắp xếp, giữ nguyên thứ tự và không sử dụng bộ nhớ phụ.

## ý tưởng
 - có thể nhận thấy nếu độ dài của mảng <= 2 thì kết quả trả về sẽ là độ dài của mảng.
 - Nếu trường hợp độ dài của mảng > 2, ta sẽ thực hiện các bước sau:
   - sử dụng biến `k` để lưu kết quả.
   - duyệt mảng `nums` bắt đầu từ vị trí thứ 3 (index = 2)
   - so sánh phần tử hiện tại với phần tử cách đó 2 vị trí trước đó. Nếu phần tử hiện tại khác với phần tử cách 2 vị trí trước đó thì `nums[k] = nums[i]`, sau đó tăng `k` lên 1 đơn vị. Nếu giống nhau thì ta sẽ bỏ qua không làm gì cả.
 - kết quả trả về chính là số `k`.

Ví dụ với input đầu vào là: [1,1,1,2,2,3]

```
Ban đầu: [1,1,1,2,2,3], k = 2
i = 2: nums[2] = 1, nums[0] = 1 → bỏ qua
i = 3: nums[3] = 2, nums[1] = 1 → nums[2] = 2, k = 3
i = 4: nums[4] = 2, nums[2] = 2 → nums[3] = 2, k = 4
i = 5: nums[5] = 3, nums[3] = 2 → nums[4] = 3, k = 5

Kết quả cuối cùng: [1,1,2,2,3], k = 5
```

