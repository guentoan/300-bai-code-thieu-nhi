## 1431. Kids With the Greatest Number of Candies

Chi tiết bài toán xem [tại đây](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/)

### mục tiêu
Cho một mảng `candies` có size là `n`, mỗi phần tử là số kẹo đang có của đứa trẻ, và một số nguyên `extraCandies`, là một số nguyên biểu thị số kẹo bổ sung. Mục tiêu của bài toán là:
 - Xác định xem một đứa trẻ có thể có số lượng kẹo lớn nhất hay không nếu bạn cho nó toàn bộ số kẹo bổ sung `extraCandies`.
 - Trả về một **mảng boolean** để biểu thị kết quả cho từng đứa trẻ:
   - **True**: Nếu sau khi được cộng thêm `extraCandies`, số kẹo của đứa trẻ đó lớn hơn hoặc bằng số kẹo của bất kỳ đứa trẻ nào khác.
   - **False**: Nếu không thể đạt được số kẹo lớn nhất ngay cả khi cộng thêm `extraCandies`.

### ý tưởng
Bài toán quy về tìm số lớn nhất trong mảng `candies`. Sau đó, duyệt qua từng phần tử của mảng `candies`, so sánh với giá trị lớn nhất sau khi đã cộng `extraCandies`.
