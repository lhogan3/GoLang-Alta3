/* Alta3 Research | RZFeeser
   Custom Sorting - Sorting structs with custom functions   */

package main

import (
    "fmt"
    "sort"
)

// record a "Person" Name, Age
type Person struct {
    Name   string
    Age    int
    Height int
}


/* In Go, if you define "String()" for a struct
   the result is a custom "ToString". That is to say,
   we can control the "view" that is returned when someone tries
   to print our struct */
func (p Person) String() string {
    return fmt.Sprintf("%s: %d years old with a height of %d", p.Name, p.Age, p.Height)
}

// ByAge implements sort.Interface for []Person based on
// the Age field
type ByAge []Person

func (a ByAge) Len() int {
    return len(a)
}
func (a ByAge) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
    return a[i].Age < a[j].Age
}

type ByHeight []Person

func (h ByHeight) Len() int {
    return len(h)
}

func (h ByHeight) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h ByHeight) Less(i, j int) bool {
    return h[i].Height < h[j].Height
}

func main() {

    // Name, Age, Height
    people := []Person{
        {"Bob", 31, 6},
        {"John", 42, 5},
        {"Michael", 17, 6},
        {"Jenny", 26, 4},
    }

    /* If you commented out the "String()" function,
    the way the people struct is displayed will change.
    The formatting it is following is provided by the "String()"
    function */
    fmt.Println(people)

    // We defined the interface for sort.Sort
    // a set of methods for the slice type, as with ByAge, and
    // call sort.Sort.
    sort.Sort(ByAge(people))
    fmt.Println(people)

    sort.Sort(ByHeight(people))
    fmt.Println(people)

}

