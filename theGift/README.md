Solution
--
Read in cost, number of participants.  

Create a sortable slice to hold participant budgets.  

Keep track of total available budget, if it is less than cost, print IMPOSSIBLE and exit.  

Sort slice of budgets.  

Go through sorted slice of budgets, calculating during each loop interval the average remaining contribution.  

If the current budget is less than the average contribution, just donate the current budget, otherwise donate the
average remaining contribution.  


