# str1: str = 'LEET'
# str2: str = 'CODE'

# len1 = len(str1)
# len2 = len(str2)

# def isDivisor(num_l: int) -> bool:
#     if len1 % num_l or len2 % num_l:
#       return False
#     f1, f2 = len1 // num_l, len2 // num_l
#     return str1[:num_l] * f1 == str1 and str1[:num_l] * f2  == str2

# x: range = range(min(len1, len2), 0, -1)

# for num_l in x:
#   if isDivisor(num_l):
#      	print(str1[:num_l])

# def reverseVowels(s:str) -> str:
#     s: list[str] = list(s)
#     vowels = set('aeiouAEIOU')
#     i, j = 0, len(s)-1

#     while (i<j):
#         if(s[i] not in vowels):
#             i += 1
#         elif(s[j] not in vowels):
#             j -= 1
#         else:
#             s[i], s[j] = s[j], s[i]
#             i += 1
#             j -= 1

#     return "".join(s)

# out = reverseVowels("IceCreAm")
# print(out)



class Solution:
    def reverseWords(self, s: str) -> str:
        m: list[str] = []
        p = s.split()
        for i in range(len(p), 0, -1):
            m.append(p[i-1])
        return " ".join(m)

    def rotate(self, nums: list[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        k = k % len(nums)
        nums[:] = nums[-k:] + nums[:-k]

        # if k%2 == 0:
        #     ans = nums[k:]+nums[:k]
        # else:
        #     ans = nums[k+1:]+nums[:k+1]

        # nums = ans

        print(nums)

soln =  Solution()
# print(soln.reverseWords("a good   example"))
soln.rotate([1,2,3,4,5,6,7], 3)
soln.rotate([-1,-100,3,99], 2)
