package main

import "fmt"

var authors = make(map[int][]int)

// longestChain recursively searches for the longest chain of influencers
func longestChain(author int) int {
    length := 1
    for _,i := range authors[author] {
        tmpLength := longestChain(i) + 1
        if tmpLength > length {
            length = tmpLength
        }
    }
    return length
}

func main() {
    var n int 
    fmt.Scanf("%d", &n)
    
    var author int
    var influenced int
    for i := 0; i < n; i++ {
        fmt.Scanf("%d %d",&author,&influenced)
        authors[author] = append(authors[author],influenced)
    }
    longest := 0
    for a := range authors {
        length :=  longestChain(a)
        if length > longest {
            longest = length
        }
    }
    fmt.Println(longest)
}

