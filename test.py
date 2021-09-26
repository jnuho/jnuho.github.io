from typing import List

class Solution:
  def twoSum(self, nums: List[int], target: int) -> List[int]:
    numToIdx = {}

    for i,n in enumerate(nums):
      numToIdx[n] = i