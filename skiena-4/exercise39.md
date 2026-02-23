Describe a sorting algorithm with a lower bound of O(n (log n)).  This seems
weird given the existence of a lower of O(n log n).

I say iterate once over the array and add every value to a hash map with the key
being the value and the value being a count of that value.  Then take the keys of the
map and sort them.  Then create a new array of size the sum of the  key.values and 
iterate once more over that map, adding the value until you have exausted that number
to and move on to the next key in the map.