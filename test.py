from collections import Counter
from typing import Tuple


class Solution:

  # Given a string s, find the length of the longest substring without repeating characters.

  def lengthOfLongestSubstring(self, s: str) -> int:
    ans = 0
    count = Counter()

    l = 0
    for r,c in enumerate(s):
      count[c] += 1


# result = Solution().lengthOfLongestSubstring("abcabcbb");
result = Solution().lengthOfLongestSubstring("abc");
print(result)
