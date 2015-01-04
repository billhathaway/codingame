package main

import "fmt"
import "sort"

func main() {
    var N int
    fmt.Scan(&N)
    
    var cost int
    fmt.Scan(&cost)
    budgets := make(sort.IntSlice,N)
    totalBudgets := 0
    for i := 0; i < N; i++ {
        fmt.Scan(&budgets[i])
        totalBudgets += budgets[i]
    }
    if totalBudgets < cost {
         fmt.Println("IMPOSSIBLE")
         return
    }
    sort.Sort(budgets)
    for i,budget := range budgets {
        averageRemaining := cost / (N - i)
        if budget < averageRemaining {
             fmt.Printf("%d\n",budget)
             cost -= budget
             continue
        }
        cost -= averageRemaining
        fmt.Printf("%d\n",averageRemaining)
    }
}
