# cach 1: su dung 2 vong lap de check xem 
# tong cua 2 item nao bang target thi tra ve index cua no
# O(n^2)
def solution1(nums, target):
  for i in range(len(nums) - 1):
    for j in range(i + 1, len(nums)):
      if nums[i] + nums[j] == target:
        return [i, j]


# cach 2: su dung 1 map de luu tru value va index cua cac item da duyet qua
# trong qua trinh duyet se kiem tra xem `target - nums[i]` co nam trong map hay khong
# neu co thi tra ve ket qua
# O(n)
def solution2(nums, target):
  keys={}
  for i in range(len(nums)):
    num = target - nums[i]
    if num in keys:
      return [keys[num], i]
    keys[nums[i]] = i
