# Problem - What is value returned by following function? Express answer as a function of n.
Give worst-case running time using Big Oh notation.

Mystery(n)
  r = 0
  for i = 1 to n -1 do
     for j = i + 1 to n do
       for k = 1 to j do
         r = r + 1
  return r


## Answer:
  answer = n * (n -1) * (n -1)
  Worst-case: O(n^3)
