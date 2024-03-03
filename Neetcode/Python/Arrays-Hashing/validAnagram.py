class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        seen = collections.defaultdict(int)
        for ch in s:
            seen[ch] += 1
        for ch in t:
            seen[ch] -= 1
            if seen[ch] == 0:
                seen.pop(ch)
        return len(seen) == 0