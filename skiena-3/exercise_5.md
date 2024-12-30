# Problem - Design a resizeable array.
# Answer - It's a slice in Golang.
-  Array is a type primitive type in Go. int[10] is a different type than int[20].  Slice is not that way.
-  Slice has a special constructor(make, [:] notation)
-  Slice has functions for appending and copying. Array does not.
-  Underlying array doubles in size once slice len goes past 50% without panicing. Array panics in this case.
      It is copied to new underlying array with 2x size.
-  Both slice and array have len and capacity functions sensitive to their underlying type.
-  Capacity always equals Len for an array.  For slice len is # of elements in slice. Capacity is size of underlying array.
-  Once slice len < 50% capacity, underlying array copied to new array 50% smaller and slice points to that new underlying array.