# Problem - There are 25 horses that can race 5 at a time. Find the minimum number of races it takes to determine the first three fastest horses.
## Assumptions:
  1. The only outcome you know is who won the heat,  not a horse's time.
  2. The races are conducted as heats.  The winner of a heat goes forward to next heat.  And this is not necessarily the order of fastest times
     lower place finishers of a given heat may have faster times than first place finishers of other heats. But all athletic tournaments work this way.

## Solution:
  1. Race them in groups of 5. That takes 5 races.  And from those races you have five winners. 5 races so far.
  2. Race the 5 winners together.  6 races so far.  That horse is fastest
  3. Remove the winner from that race and race the four remaining horses.  7 races so far. That winner is second fastest.
  4. Remove the winner from that race and race the three remaining horses.  8 races so far. That winner is third fastest.