class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        result: List[List[int]] = []
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
                    if sorted_lst not in result:
                        result.append(sorted_lst)

                if sum < 0:
                    l+=1
                elif sum > 0:
                    r-=1
                else:
                    l+=1
                    r-=1

        return result