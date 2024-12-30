Algorithm for detecting whether or not an an array of A contains a sub-sequence array A' in the same order.

- Iterate over A and add each entry to a map. The key of the map is the A value and the value is its position in A. Only record the position of the first occurrence of A in the A array.
- while iterating over A`.  For each A' look it up in the A map. If it is there continue. Otherwise break, A' is not a sub-sequence of A.
  - If that A' is in the map record it an another map along with its position in A'.
    - Is the previous A' value(not a duplicate one) have a position that is less than the current A' position.
      -   yes - continue A' iteration
      -   no - break. A' is not as sub-sequence of A
    - end if
  - end if
- end while