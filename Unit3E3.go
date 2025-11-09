// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-09
// Fileoverview: This program will ask the user to enter four lines of text.
// The program will then display the lines in reverse order (line 4, line 3, line 2, line 1).

package main

import "fmt"

func main() {
    // variables
    var line1 string
    var line2 string
    var line3 string
    var line4 string

    // input
    fmt.Print("Enter line 1: ")
    fmt.Scanln(&line1)

    fmt.Print("Enter line 2: ")
    fmt.Scanln(&line2)

    fmt.Print("Enter line 3: ")
    fmt.Scanln(&line3)

    fmt.Print("Enter line 4: ")
    fmt.Scanln(&line4)

    // output
    fmt.Println()
    fmt.Println("The lines in reverse order are:")
    fmt.Println(line4)
    fmt.Println(line3)
    fmt.Println(line2)
    fmt.Println(line1)
    fmt.Println("\nDone.")
}

