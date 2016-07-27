package main

import(
    "container/list"
    "fmt"
)

type Pair struct{
    j int
    k int
}

//E3. Find roots
func findClass(i int, parent []int){
    for iterate := 1; iterate < len(parent); iterate++{
        if (parent[iterate] == i) && (iterate != i){
            fmt.Printf("%d",iterate)
            findClass(iterate,parent)
          }
    }
}
func generateClasses(parent []int){
    for i := 1; i < len(parent); i++{
        if parent[i] == 0{
            fmt.Printf("%d",i)
            findClass(i,parent)
            fmt.Println()
        }
    }
}

func findRoots(eqPair Pair, parent []int) (int, int){
    j := eqPair.j
    for parent[j] > 0 {
        j = parent[j]
    }

    k := eqPair.k
    for parent[k] > 0 {
        k = parent[k]
    }

    return j,k
}

func main(){
    ls := list.New()
    //E1. Initialize
    parent := []int{0,0,0,0,0,0,0,0,0,0}

    //Todo: E2. Input new pair
    //These are the same values Knuth uses
    ls.PushBack(Pair{1,5})
    ls.PushBack(Pair{6,8})
    ls.PushBack(Pair{7,2})
    ls.PushBack(Pair{9,8})
    ls.PushBack(Pair{3,7})
    ls.PushBack(Pair{4,2})
    ls.PushBack(Pair{9,3})

    for e := ls.Front(); e != nil; e = e.Next() {
        j,k := findRoots(e.Value.(Pair), parent)
        // E4. Merge trees
        if j != k {
            parent[j] = k
        }
    }
    generateClasses(parent)
}
