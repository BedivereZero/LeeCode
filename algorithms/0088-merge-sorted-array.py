class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        nums3 = list()
        a, b = 0, 0
        while a < m and b < n:
            if nums1[a] < nums2[b]:
                nums3.append(nums1[a])
                a += 1
            else:
                nums3.append(nums2[b])
                b += 1
        while a < m:
            nums3.append(nums1[a])
            a += 1
        while b < n:
            nums3.append(nums2[b])
            b += 1
        for i in range(m + n):
            nums1[i] = nums3[i]
