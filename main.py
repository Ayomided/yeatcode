str1: str = 'LEET'
str2: str = 'CODE'

len1 = len(str1)
len2 = len(str2)

def isDivisor(num_l: int) -> bool:
    if len1 % num_l or len2 % num_l:
      return False
    f1, f2 = len1 // num_l, len2 // num_l
    return str1[:num_l] * f1 == str1 and str1[:num_l] * f2  == str2

x: range = range(min(len1, len2), 0, -1)

for num_l in x:
  if isDivisor(num_l):
     	print(str1[:num_l])
