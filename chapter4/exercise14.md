The Problem:
You have a set of processes K, each taking X amount of time to complete, the difference between the start and end times for a given process.
All the processes run every 24 hours one at a time on a schedule that you do not know. 

You have another process P that probes the machine running k and reports important facts about a given process. The problem is to determine the minimum 
number of times you have to invoke P to make sure you hit every process at least once.

The Algorithm:
1. Find the minimum X and call it G.  This takes O(n).
2. Run probe P continually at interval G, starting at the beginning of the day.
3. The total number of times you run P is the minimum number of probes that guarantee that you will probe every process over a given 24 hour day.
