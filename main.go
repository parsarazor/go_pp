package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Product struct {
    Weight int
    Index  int
}

type Shelf struct {
    Capacity int
    Level    int
    Used     int
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    // Read N and L
    firstLine, _ := reader.ReadString('\n')
    firstLine = strings.TrimSpace(firstLine)
    parts := strings.Split(firstLine, " ")
    N, _ := strconv.Atoi(parts[0])
    L, _ := strconv.Atoi(parts[1])

    // Read weights
    weightsLine, _ := reader.ReadString('\n')
    weightsLine = strings.TrimSpace(weightsLine)
    weightsStr := strings.Split(weightsLine, " ")
    weights := make([]int, N)
    for i := 0; i < N; i++ {
        weights[i], _ = strconv.Atoi(weightsStr[i])
    }

    // Read capacities
    capsLine, _ := reader.ReadString('\n')
    capsLine = strings.TrimSpace(capsLine)
    capsStr := strings.Split(capsLine, " ")
    shelves := make([]Shelf, L)
    for i := 0; i < L; i++ {
        cap, _ := strconv.Atoi(capsStr[i])
        shelves[i] = Shelf{Capacity: cap, Level: i + 1, Used: 0}
    }

    // Sort products by weight descending
    products := make([]Product, N)
    for i := 0; i < N; i++ {
        products[i] = Product{Weight: weights[i], Index: i}
    }
    sort.Slice(products, func(i, j int) bool {
        return products[i].Weight > products[j].Weight
    })

    // Sort shelves by initial capacity descending, but level ascending if tie
    sort.Slice(shelves, func(i, j int) bool {
        if shelves[i].Capacity == shelves[j].Capacity {
            return shelves[i].Level < shelves[j].Level
        }
        return shelves[i].Capacity > shelves[j].Capacity
    })

    // Greedy assignment: heavier products on lower-shelf-with-room
    score := 0
    for _, prod := range products {
        for si := 0; si < L; si++ {
            if shelves[si].Used+prod.Weight <= shelves[si].Capacity {
                shelves[si].Used += prod.Weight
                score += prod.Weight * shelves[si].Level
                break
            }
        }
    }

    fmt.Println(score)
}
