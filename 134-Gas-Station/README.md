## 134. Gas Station

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/gas-station/description)

### Đề bài
Có **n trạm xăng** nằm dọc theo một tuyến đường tròn, với lượng xăng ở trạm thứ `i` được biểu thị bằng `gas[i]`.

Bạn có một chiếc xe hơi với bình xăng **không giới hạn**. Chi phí xăng để đi từ trạm `i` đến trạm tiếp theo `(i + 1)` được biểu thị bằng `cost[i]`. Hành trình bắt đầu với bình xăng **rỗng** tại một trong các trạm xăng.

Cho hai mảng số nguyên `gas` và `cost`:
- Trả về **chỉ số của trạm xăng bắt đầu** nếu bạn có thể hoàn thành một vòng tròn theo chiều kim đồng hồ.
- Nếu không thể hoàn thành hành trình, trả về `-1`.

Nếu tồn tại một lời giải, đảm bảo lời giải là **duy nhất**.


### Phân tích
Ta có thể nhận thấy, để xe có thể đi qua tất cả trạm xăng thì tổng lượng xăng dư sau khi đi hết tất cả các trạm xăng phải lớn hơn hoặc bằng `0`.
Khi tổng xăng từ một trạm không đủ để đi tiếp, thì bất kỳ trạm nào trước đó cũng không thể là trạm bắt đầu, nên có thể loại bỏ tất cả các trạm trước trạm thứ `i`.

Từ đó ta có thể duyệt mảng từ trái sang phải, và dùng hai biến `current, sum` để theo dõi lượng xăng của hành trình nhỏ, và lượng xăng của cả hành trình lớn.
Tại trạm thứ `i`, nếu `current < 0`, có nghĩa là tất cả các trạm trước `i` đều không hợp lệ, do đó gán `startAt = i+1`.
Sau khi duyệt hết mảng, ta sẽ so sánh giá trị `sum`.
 - Nếu `sum >= 0`: trạm bắt đầu sẽ là `startAt`.
 - Nếu `sum < 0`: không có trạm nào hợp lệ, nên sẽ trả về `-1`.