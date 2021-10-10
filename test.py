from typing import Tuple


class Solution:

  # Given a string s, return the longest palindromic substring in s.
  # Input: s = "babad"
  # Output: "bab"
  # Note: "aba" is also a valid answer.
  def longestPalindrome(self, s: str) -> str:
    if not s:
      return ''

    def extend(s:str, i:int, j:int) -> Tuple[int, int]:
      while i>=0 and j < len(s):
        if s[i] != s[j]:
          break
        i -= 1
        j += 1
      
      return i+1, j-1

    result = 0, 0

    for i in range(len(s)):
      a, b = extend(s, i, i)
      if b-a > result[1] - result[0]:
        result = a, b

      if i+1 < len(s):
        a, b = extend(s, i, i+1)
        if b-a > result[1] - result[0]:
          result = a, b

    return s[result[0]: result[1]+1]

result = Solution().longestPalindrome("cbbd")
# result = Solution().longestPalindrome("babad")
print(result)
