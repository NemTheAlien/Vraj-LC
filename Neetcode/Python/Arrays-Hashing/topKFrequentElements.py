class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = Counter(nums)
        heap = []
        for i in freq.keys():
            heappush(heap,(freq[i],i))
        freq = nlargest(k,heap)
        ans = []
        for i,j in freq:
            ans.append(j)
        
        return ans