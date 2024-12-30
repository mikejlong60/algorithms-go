The Problem:
Bank account `a` has `n` transactions that we suspect are the result of illegal activity.  But we don't know the exact time of occurrence,
only a time range. Another back account `b` has `n` transactions that seem roughly correlated. Prove that back account `b` is the account that
is producing the transactions on account `a`.  We know the exact time of the transactions for bank account `b`.

Solution:
Algorithm efficiency O(n)

1. Order bank account `a` transactions by start timestamp.
2. Order bank account `b` transactions by timestamp.
3. Iterate over all the transactions in account `b`.
4. Look to see if account `b` transaction n has a timestamp that is ge account `a` transaction n and le account `a` transaction n.
5. If so account `b` is the source of the illegal activity. Arrest the account holder.

