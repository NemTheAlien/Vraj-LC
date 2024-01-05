import ast as List

# This contains the solution for LeetCode problem 27 - Remove element
def removeElement(self, nums: List[int], val: int) -> int:
        while val in nums:
            nums.remove(val)



