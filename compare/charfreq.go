package main

import (
    "fmt"
    "os"
    "bufio"
    "sort"
)

func main() {
    filename := "page1.txt"
    if len(os.Args) > 1 {
        filename = os.Args[1]
    }
    
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println("Cannot open file")
        return
    }
    defer f.Close()
    
    cf := map[rune]int{}
    
    reader := bufio.NewReader(f)
    for {
        if line, err := reader.ReadString('\n'); err != nil {
            break
        } else {
            for _, c := range []rune(line) {
                if n, ok := cf[c]; ok {
                    cf[c] = n + 1
                } else {
                    cf[c] = 1
                }
            }
        }
    }
    
    // sort map in python: sorted(m, key=lambda k:m[k])
    cfv := sortMapByVal(cf)
    for _, v := range cfv {
        fmt.Printf("%c: %d\n", v.Key, v.Value)
    }
}

// Borrow from Andrew Gerrand's solution
type Pair struct {
    Key rune
    Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p PairList) Len() int {
    return len(p)
}
func (p PairList) Less(i, j int) bool {
    return p[j].Value < p[i].Value
}
func sortMapByVal(cm map[rune]int) PairList {
    p := make(PairList, len(cm))
    i := 0
    for k, v := range cm {
        p[i] = Pair{k, v}
        i += 1
    }
    sort.Sort(p)
    return p
}
