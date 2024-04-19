class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        res = []

        for left in range(0,len(nums)-2):
            i=left+1
            right = len(nums)-1
            while i<right and i>left:
                sums = nums[left] + nums[i] + nums[right]
                if sums == 0:
                    res.append([nums[left],nums[i],nums[right]])
                    i += 1
                elif sums > 0:
                    right -= 1
                else:
                    i += 1
        res = set(map(tuple,res))
        return list(res)