from typing import List

class Solution:

    def search(self, result, candidates, current, index, target):
        if target == 0:
            result.append(current)
            return
        if target < 0:
            return
        for point in range(index, len(candidates)):
            self.search(result, candidates, current + [candidates[point]], point, target - candidates[point])


    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = list()
        self.search(result=result, candidates=sorted(candidates), current=[], index=0, target=target)
        return result
