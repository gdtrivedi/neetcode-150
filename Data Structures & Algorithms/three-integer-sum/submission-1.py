class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        result_set = set()
        nums = sorted(nums) # nlogn
        for i, n in enumerate(nums): # n * 
            fix_index = i
            l=0
            r=len(nums)-1
            while l < r:
                if l == fix_index:
                    l+=1
                    continue
                
                if r == fix_index:
                    r-=1
                    continue

                sum = nums[l] + nums[r] + nums[fix_index]
                if sum == 0:
                    sorted_lst = sorted([nums[l], nums[r], nums[fix_index]]) # O(1)
                    result_set.add((sorted_lst[0], sorted_lst[1], sorted_lst[2]))

                if sum < 0:
                    l+=1
                elif sum > 0:
                    r-=1
                else:
                    l+=1
                    r-=1

        return [list(item) for item in result_set]