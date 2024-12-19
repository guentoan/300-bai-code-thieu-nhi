## 769. Max Chunks To Make Sorted

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/max-chunks-to-make-sorted/description)

### Đề bài
Cho một mảng số nguyên `arr` có độ dài `n`, đại diện cho một hoán vị của các số trong khoảng `[0, n - 1]`.

Ta chia mảng `arr` thành một số đoạn nhất định (chunks), và sắp xếp từng đoạn một cách độc lập. Sau khi ghép nối tất cả các đoạn đã sắp xếp, kết quả thu được phải bằng mảng đã được sắp xếp theo thứ tự tăng dần.

Hãy trả về số lượng đoạn lớn nhất mà ta có thể chia để sắp xếp mảng.

### Phân tích

Mảng arr là một **hoán vị** của các số từ `0` đến `n-1`.
Điều này có nghĩa là mảng sẽ chứa đủ tất cả các số trong khoảng `[0, n-1]`, nhưng có thể theo bất kỳ thứ tự nào.
Ví dụ: Nếu `n = 5`, thì arr có thể là `[4, 3, 1, 0, 2]`, hoặc `[0, 1, 2, 3, 4]`, hoặc bất kỳ sắp xếp nào khác của các số `[0, 1, 2, 3, 4]`.

Mục tiêu là chia mảng đầu vào thành nhiều mảng con, để sau khi sắp xếp các mảng con đó và ghép lại thì được một mảng được sắp xếp theo thứ tự tăng dần.

Ta có thể nhận thấy, để có được một đoạn con thoả mãn điều kiện bài toán thì giá trị lớn nhất của mảng con phải bằng chỉ số thứ i của mảng đó.
Tức là `maxLeft = i`.