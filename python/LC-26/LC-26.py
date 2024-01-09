from ast import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        count = 0 

        for i in nums:
            if i not in nums[:count]:
                nums[count] = i
                count+=1
        return count