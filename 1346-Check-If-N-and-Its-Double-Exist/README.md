## 1346. Check If N and Its Double Exist

Given an array `arr` of integers, check if there exist two indices `i` and `j` such that :

 - i != j
 - 0 <= i, j < arr.length
 - arr[i] == 2 * arr[j]

**Example 1:**
```
Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]
```

**Example 2:**
```
Input: arr = [3,1,7,11]
Output: false
Explanation: There is no i and j that satisfy the conditions.
```

**Constraints:**

 - 2 <= arr.length <= 500
 - -10^3 <= arr[i] <= 10^3

### mục tiêu
Kiểm tra xem mảng `arr` có chứa 2 phần tử ở vị trí `i` `j`, mà `arr[i] = 2*arr[j]` không?
Nếu có thì trả về `true`, còn không thì trả về `false`.

### ý tưởng
Ta có thể sử dụng một map để lưu giá trị của các phần tử của array.
Sau đó duyệt qua từng phần tử của array, đồng thời kiểm tra xem trong map có tồn tại giá trị mà `arr[i] = 2*arr[j]` không?
Nếu có thì trả về `true`.
Khi duyệt qua tất cả các phần tử của mảng `arr` mà chưa tìm được phần tử thoả mãn điều kiện trên thì trả về `false`.
