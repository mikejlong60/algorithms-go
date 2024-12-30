The Problem:
Given a schedule of n processes with a start and end time, choose the minimum number of processes from that set of n
to serve as supervisors over the processes with which their start and end times overlap.  Every process must have at least 
one overlapping process.

Solution:
supervisors is a stack and that is the returned stack of supervisors

1. Order the set of n processes by start time ascending and store it in a stack _schedule_ with smallest time on top.
2. If _schedule_ is empty return _supervisors_
3. Pop _schedule_ and mark it as a supervisor by placing it in stack _supervisors_
3.   Pop the next process in _schedule_ and call it _i_.
   4. If i's end time is greater than last element in _supervisors_
      5. Pop _supervisors_
      6. Push i onto  _supervisors_
   6. goto step 2 -- use recursion of course


The cost of this algorithm is O(n log n)
func makeSchedule(schedule Stack[Process], supervisors Stack[Processes]) Stack[Process] {
      if len(schedule) == 0 return supervisors
      
      a = Pop(schedule)
      supervisors = Push(supervisors, a)
      i = Pop(schedule)
      if i.endTime > Peek(supervisors).endTime {
            Pop(supervisors)
            supervisors = Push(supervisors, i)
      }
      makeSchedule(schedule, supervisors)
}
