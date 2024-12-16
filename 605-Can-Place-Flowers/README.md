## 605. Can Place Flowers

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/can-place-flowers/description/)

### Mục Tiêu
Bạn có một luống hoa dài, trong đó một số ô đã được trồng hoa và một số ô thì chưa. 
Tuy nhiên, hoa không thể được trồng trên các ô liền kề nhau.

Cho một mảng số nguyên flowerbed chứa các số `0` và `1`, trong đó `0` có nghĩa là ô trống và `1` có nghĩa là ô đã được trồng hoa.
Đồng thời cho một số nguyên `n`.
Hãy trả về `true` nếu có thể trồng thêm n bông hoa mới trên luống hoa mà không vi phạm quy tắc **"không có hoa liền kề"**, 
ngược lại trả về `false`.

### Ý Tưởng
Bởi vì, chúng ta có quy tắc **không có hoa liền kề**, nên ta có thể nhận thấy quy tắc như sau:

 - trồng 1 cây hoa cần it nhất 3 ô trống.
 - trồng 2 cây hoa cần it nhất 5 ô trống.
 - trồng 3 cây hoa cần it nhất 7 ô trống.
 - trồng 4 cây hoa cần it nhất 9 ô trống.
 - trồng 5 cây hoa cần it nhất 11 ô trống.
 - ...
 - trồng n cây hoa cần ít nhất `2n+1` ô trống.

Theo quy luật trên, ta có thể đếm số ô trống liền kế nhau, sau đó tính được số lượng cây tối đa có thể trồng.
