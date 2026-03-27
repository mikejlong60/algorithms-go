# You are given 12 coins.  One of them is lighter or heavier than the rest.  Find this 
# coin in three weighings of a balance scale.

# Answer
 - Split 12 coins into two groups of 6, A and B.
 - weighing 1 - Pick one group of 6 and split it into two groups and place it on the balance scale
 - If the scale balances, the different weight coin is in the other group of 6.  Pick either A or B and call it C based upon that condition for remaining tests.
 - Split C into two groups of three and remove one item D1 and E1 from each group, D and E. D and E now have two items.   
 - weighing 2 - Place 2-coin-groups D and E on the scale.
 - If scale balances, the different weight is either D1 or E1 - Place those two items with any two items from D or E on the scale to decide which is a different weight. - weighing 3
 - If it does not balance, remove one coin, F1 and G1, from each group F and G which have one coin. 
 - weighing 3 - Place F and E1 on the scale. If the scale balances. G is the differing coin.  Otherwise, F is the differing coin.
 