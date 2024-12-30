The Problem:
You have a bunch of two-step processes that need to be completed by a supercomputer for the first part and then a PC for the 
second part.  The supercomputer can only accept one of the two-tep processes at a time. After the first step you have 
an unlimited number of PCs that can do the secpomd step in parallel. 

Devise a schedule for a list of n two-step processes that has the minimum completion time for the whole list. 

Solution:
Order the two-step process list by PC time minus Supercomputer time in descending order. The heuristic is to order the list of processes by maximum 
idle time for the supercomputer.  This is the fastest way to order the n processes.