## 38. Count and Say

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/count-and-say/description/)

### Đề bài
Dãy số **count-and-say** là một dãy các chuỗi chữ số được định nghĩa bởi công thức đệ quy:

- `countAndSay(1) = "1"`
- `countAndSay(n)` là run-length encoding của `countAndSay(n - 1)`.

**Run-Length Encoding - RLE** là một phương pháp nén chuỗi, hoạt động bằng cách thay thế các ký tự liên tiếp giống nhau (lặp lại từ 2 lần trở lên) bằng cách nối ký tự với số chỉ định số lần xuất hiện của ký tự đó (chiều dài của chuỗi lặp lại). Ví dụ, để nén chuỗi "3322251", ta thay thế "33" bằng "23", thay thế "222" bằng "32", thay thế "5" bằng "15" và thay thế "1" bằng "11". Do đó, chuỗi nén trở thành "23321511".

**Yêu cầu:** Cho một số nguyên dương `n`, hãy trả về chuỗi `count-and-say` thứ `n`.
