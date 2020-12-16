package main

import "fmt"

func function2() { 
    
         fmt.Println("fun2 line 1")
          }
 
func function1() {  
        fmt.Println("fun1 line 1")    
          function2()  
              fmt.Println("fun1 line 2")
               }
 
func main() {  
        fmt.Println("main line 1")    
          function1()   
             fmt.Println("main line 2") 
             }
 