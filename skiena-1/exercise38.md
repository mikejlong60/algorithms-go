# Problem - On average how many times would you have to flip open the Manhattan phone book to find a specific person?
## Approach:
    1. Since the phone book is in alphabetical order the maximum number of times you have to open the book is the base-2-logarithm-of-the-number-of-pages.
    Here is the algorithm:
      1. Crack open the phone book in the middle.
      2. Is the person < or > or == the current 2 pages?
        1. > - crack open the phone book in the middle again starting at the next page and until the end.
        2. < - crack open the phone book in the middle again starting at the beginning and ending at previous page.
        3. == - you are done. exit the algorithm

        Call function 2 recursively and the exit condition is #3.


 Binary search - Every elementary school student knows this algorithm.
