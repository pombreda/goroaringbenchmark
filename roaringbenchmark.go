
// prior to running this code you should do go get 
// github.com/fzandona/goroar and go get github.com/tgruben/roaring
package main

import (
    "fmt"
    "github.com/fzandona/goroar"
    "github.com/tgruben/roaring"
)

func goroartest() {
    fmt.Println("==goroar==")
    rb1 := goroar.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
    fmt.Println(rb1.String())

    rb2 := goroar.BitmapOf(3, 4, 1000)
    fmt.Println(rb2.String())

    rb3 := goroar.New()
    fmt.Println(rb3.String())

    fmt.Println("Cardinality: ", rb1.Cardinality())

    fmt.Println("Contains 3? ", rb1.Contains(3))

    rb1.And(rb2)

    rb3.Add(1)
    rb3.Add(5)

    rb3.Or(rb1)

    //prints 1, 3, 4, 5, 1000
    for value := range rb3.Iterator() {
        fmt.Println(value)
    }
    fmt.Println()
}

func roaringtest() {
    fmt.Println("==roaring==")
    rb1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
    fmt.Println(rb1.String())

    rb2 := roaring.BitmapOf(3, 4, 1000)    
    fmt.Println(rb2.String())
    
    rb3 := roaring.NewRoaringBitmap()
    fmt.Println(rb3.String())

    fmt.Println("Cardinality: ", rb1.GetCardinality())

    fmt.Println("Contains 3? ", rb1.Contains(3))

    rb1.And(rb2)

    rb3.Add(1)
    rb3.Add(5)

    rb3.Or(rb1)

    // prints 1, 3, 4, 5, 1000
    i := rb3.Iterator()
    for i.HasNext() {
        fmt.Println(i.Next())
    }
    fmt.Println()
}

func main() {
    goroartest()
    roaringtest()
}