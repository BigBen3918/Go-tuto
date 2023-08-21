// In a previous example we saw how for and range provide iteration over basic data structures. We can also use this syntax to iterate over values received from a channel.

package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}