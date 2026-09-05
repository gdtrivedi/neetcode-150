class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        cnt = Counter(nums)
        # tpls = cnt.most_common(k)
        temp = list(map(lambda item: item[0], cnt.most_common(k)))
        # print(temp)
        return temp