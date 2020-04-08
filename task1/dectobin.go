package main

import (
 "fmt"
 "strconv"
)
func main() {
 var dec int64
 fmt.Print("-----Welcome to DecToBin")
 fmt.Print("Enter Decimal Number:")
 fmt.Scanln(&dec)
 //Inbuilt function to convert input from string to int and calculate binary 
 bin := strconv.FormatInt(dec, 2)
 fmt.Print("The Binary equivalent is:", bin)
}