package main

import ( 
  "fmt"
  "os"
  )

var sourceDir, destDir string
func main() {
  if len(os.Args) < 3 {
    fmt.Println("Usage: securecopy [directory] [destination]")
    return 
  } else { 
    sourceDir = os.Args[1]
    destDir = os.Args[2]
  }
}
