
def solution(s):
  subs = {}
  l, idx = 0, 0
  for i, c in enumerate(s):
    if c in subs and subs[c] >= idx:
      idx = subs[c] + 1
    subs[c] = i
    l = max(l, i - idx + 1)
  return l