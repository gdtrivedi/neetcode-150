class Solution:
    def trap(self, height: List[int]) -> int:
        l = len(height)

        water_based_on_left: list[int] = [0] * len(height)
        pre_left_height = 0
        for i in range(l):
            n = height[i]
            if n < pre_left_height:
                water_based_on_left[i] = pre_left_height - n

            if n >= pre_left_height:
                pre_left_height = n

        water_based_on_right: list[int] = [0] * len(height)
        pre_right_height = 0
        for i in range(l-1,-1,-1):
            n = height[i]
            if n < pre_right_height:
                water_based_on_right[i] = pre_right_height - n

            if n >= pre_right_height:
                pre_right_height = n

        sum_trap_water = 0
        for i in range(l):
            sum_trap_water += min(water_based_on_left[i], water_based_on_right[i])

        return sum_trap_water
        