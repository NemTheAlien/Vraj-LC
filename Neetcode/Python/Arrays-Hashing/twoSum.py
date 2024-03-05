class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}
        for i,val in enumerate(nums):
            lf = target - val
            if lf in seen:
                return [seen[lf],i]
            else:
                seen[val] = i
        return []