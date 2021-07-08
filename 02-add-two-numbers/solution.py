# Definition for singly-linked list.
class ListNode:
  def __init__(self, val=0, next=None):
    self.val = val
    self.next = next


def solution(l1: ListNode, l2: ListNode) -> ListNode:
  carry = 0
  head = ListNode()
  cur = head
  while(l1 or l2):
    x = l1.val if l1 else 0
    y = l2.val if l2 else 0
    sum = carry + x + y
    carry = sum // 10
    cur.next = ListNode(sum % 10)
    cur = cur.next
    l1 = l1.next if l1 else None
    l2 = l2.next if l2 else None
  if carry > 0:
    cur.next = ListNode(carry)
  return head.next
