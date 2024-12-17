## 2182. Construct String With Repeat Limit

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/construct-string-with-repeat-limit/description/?envType=daily-question&envId=2024-12-17)

### Mô Tả
Bạn được cho một chuỗi ký tự `s` và một số nguyên `repeatLimit`. Hãy xây dựng một chuỗi mới `repeatLimitedString` từ các ký tự trong `s`, sao cho không có ký tự nào xuất hiện liên tiếp nhiều hơn `repeatLimit` lần. Bạn không bắt buộc phải sử dụng tất cả các ký tự từ `s`.

Trả về chuỗi `repeatLimitedString` lớn nhất về mặt từ điển có thể.

### Định Nghĩa
- Một chuỗi `a` được xem là lớn hơn về mặt từ điển so với chuỗi `b` nếu:
    1. Tại vị trí đầu tiên mà `a` và `b` khác nhau, ký tự của `a` xuất hiện **sau** ký tự của `b` trong bảng chữ cái.
    2. Nếu tất cả các ký tự đầu tiên của `a` và `b` giống nhau cho đến `min(a.length, b.length)`, thì chuỗi dài hơn sẽ được xem là lớn hơn về mặt từ điển.

### Ví Dụ 1
**Input**:  
`s = "cczazcc"`  
`repeatLimit = 3`

**Output**:  
`"zzcccac"`

**Giải Thích**:
- Chuỗi lớn nhất về mặt từ điển là `"zzcccac"`.
- Ký tự `'z'` chỉ xuất hiện 2 lần liên tiếp (nhỏ hơn `repeatLimit = 3`).
- Ký tự `'c'` xuất hiện 3 lần liên tiếp, nhưng không vượt quá giới hạn `repeatLimit`.

### Ví Dụ 2
**Input**:  
`s = "aababab"`  
`repeatLimit = 2`

**Output**:  
`"bbabaa"`

**Giải Thích**:
- Chuỗi lớn nhất về mặt từ điển là `"bbabaa"`.
- Ký tự `'b'` xuất hiện tối đa 2 lần liên tiếp.
- Ký tự `'a'` cũng xuất hiện tối đa 2 lần liên tiếp.

## Yêu Cầu
1. Xây dựng chuỗi lớn nhất về mặt từ điển từ các ký tự trong `s`.
2. Không để bất kỳ ký tự nào lặp lại nhiều hơn `repeatLimit` lần liên tiếp.

### Ý Tưởng
sử dụng MaxHeap