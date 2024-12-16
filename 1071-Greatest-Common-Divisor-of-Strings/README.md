## 1071. Greatest Common Divisor of Strings

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/greatest-common-divisor-of-strings/description/)

### mục tiêu
Tìm chuỗi `x` là dài nhất sao cho `x` chia hết cho cả `str1` và `str2`.

### ý tưởng

Ta có thể dễ dàng nhận thấy rằng, nếu tồn tại một chuỗi `x` chia hết cho cả `str1` và `str2` thì khi đó nối hai chuỗi lại theo thứ tự bất kỳ thì đều cho ra một chuỗi giống nhau. Hay nói khách khác: `str1 + str2 = str2 + str1`.
Nếu không `str1+str2` không bằng `str2+str1`, thì khi đó không tồn tại một chuỗi `x` có thể chia hết cho cả `str1` và `str2`.

Nếu hai chuỗi `str1` và `str2` thoả mãn điều kiện trên. Ta có thể đưa bài toán trên thành tìm **ước chung lớn nhất** của `len(str1)` và `len(str2)`.
**Tại sao lại vậy?** Đó là tại vì chuỗi `x` là chuỗi được lặp lại (**repeating substring**) của cả `srt1` và `str2`, nên `len(str1)` và `len(str2)` đều chia hết cho `len(x)`.
