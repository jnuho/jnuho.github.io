class Solution:

  # Given a string s, return the longest palindromic substring in s.
  # Input: s = "babad"
  # Output: "bab"
  # Note: "aba" is also a valid answer.
  def longestPalindrome(self, s: str) -> str:
    longest = ""
    for i,c in enumerate(s):
      s_len = min(i, len(s)-i-1)
      for j in range(s_len+1):
        target = s[i-j:i+j+1]
        if self.isPalindrome(target):
          if len(target) > len(longest):
            longest = target
        target = s[i-j:i+j]
        if self.isPalindrome(target):
          if len(target) > len(longest):
            longest = target
    return longest


    longest = ""
    return longest

  def isPalindrome(self, s:str) -> str:
    for i in range(len(s)//2):
      if s[i] != s[len(s)-i-1]:
        return False
    return True

result = Solution().longestPalindrome("cbbd")
print(result)
