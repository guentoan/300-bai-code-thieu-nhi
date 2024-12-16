## 283. Move Zeroes

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/move-zeroes/description/)

### Mục Tiêu
Cho một mảng `nums`, dịch chuyển tất cả các phần tử có giá trị `0` về cuối mảng, và giữ nguyên thứ tự của các phần tử còn lại.
Bài toán yêu cầu giải pháp ` in-place`.

### Ý Tưởng
để giải quyết bài toán này, ta có thể sử dụng hai vòng for lồng nhau. Tại mỗi vị trí có phần tử bằng `0`, ta sẽ tiến hành duyệt một vòng for bên trong, nhằm mục đích tìm phần tử khác 0 gần nhất. sau đó, tiến hành swap hai phần tử đó với nhau.
Nếu `i=j` thì vòng lặp sẽ kết thúc.
