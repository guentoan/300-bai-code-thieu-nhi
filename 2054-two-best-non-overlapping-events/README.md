## 2054. Two Best Non-Overlapping Events

You are given a **0-indexed** 2D integer array of `events` where `events[i] = [startTime, endTime, value]`. The `ith` event starts at `startTime` and ends at `endTime`, and if you attend this event, you will receive a value of `value`. You can choose **at most two non-overlapping** events to attend such that the sum of their values is **maximized**.

Return *this **maximum** sum*.

Note that the start time and end time is **inclusive**: that is, you cannot attend two events where one of them starts and the other ends at the same time. More specifically, if you attend an event with end time `t`, the next event must start at or after `t + 1`.

**Example 1:**
```
Input: events = [[1,3,2],[4,5,2],[2,4,3]]
Output: 4
Explanation: Choose the green events, 0 and 1 for a sum of 2 + 2 = 4.
```

**Example 2:**

```
Input: events = [[1,3,2],[4,5,2],[1,5,5]]
Output: 5
Explanation: Choose event 2 for a sum of 5.
```

**Example 3:**

```
Input: events = [[1,5,3],[1,5,1],[6,6,5]]
Output: 8
Explanation: Choose events 0 and 2 for a sum of 3 + 5 = 8.
```

**Constraints:**

 - 2 <= events.length <= 10^5 
 - events[i].length == 3 
 - 1 <= startTime <= endTime <= 10^9
 - 1 <= value <= 10^6

### mục tiêu
Chúng ta cần tìm giá trị lớn nhất có thể đạt được khi tham gia sự kiện.
Tối đa có thể tham gia 2 sự kiện không chồng chéo nhau về mặt thời gian.

### ý tưởng
Vì chúng ta có thể chọn tối đa hai sự kiện không chồng chéo nhau về mặt thời gian, 
nên chúng ta có thể sắp xếp list các sự kiện theo thời gian với thứ tự tăng dần.
Sau đó, ta sẽ duyệt qua từng sự kiện, trong quá trình duyệt qua các sự kiện đó, ta sẽ sử dụng hai biến.

 - maxValue: là giá trị lớn nhất của các sự kiện đã kết thúc.
 - res: tổng giá trị lớn nhất khi tham gia hai sự kiện.

Tuy nhiên, nếu để nguyên cấu trúc `[startTime, endTime, value]` chúng ta sẽ khó phân biệt khi nào sự kiện bắt đầu, và kết thúc, cũng như sự kiện tiếp theo có chồng chéo lên sự kiện hiện tại hay không.
vì thế, chúng ta cần phải tách cấu trúc này thành một cấu trúc mới là `[time, timeType, value]`. 
với `time` là thời gian bắt đầu, hoặc kết thúc sự kiện, `timeType` là kiểu thời gian, `1` là thời gian bắt đầu sự kiện, `0` là thời gian kết thúc sự kiện.
và thời gian kết thúc chúng ta cần `+1` để có thể tracking các sự kiện không chồng chéo với 1 đơn vị thời gian.

