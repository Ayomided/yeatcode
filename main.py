str1: str = 'LEET'
str2: str = 'CODE'

len1 = len(str1)
len2 = len(str2)

def isDivisor(l: int) -> bool:
    if len1 % l or len2 % l:
      return False
    f1, f2 = len1 // l, len2 // l
    return str1[:l] * f1 == str1 and str1[:l] * f2  == str2

x: range = range(min(len1, len2), 0, -1)

for l in x:
  if isDivisor(l):
  	print(str1[:l])
