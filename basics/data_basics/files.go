package data 

import (
  "fmt"
  "io/util"
  "os"
)

func main() {
  data := []byte("Hello World!\n")
  // filename, data, and permissions
  // send byte slics to WriteFile
  err := ioutil.WriteFile("data1", data, 0644)
  if err != nil {
    panic(err)
  }

  // Read files by simply passing the name
  // data is received as byte slices
  // path?
  read1, _ := ioutil.ReadFile("data1")
  fmt.Println(string(read1))
  
  // os.Create creates a File struct 
  file1, _ := os.Create("data2")
  defer file1.Close()
  
  // You can write to which also 
  bytes, _ = file1.Write(data)
  fmt.Printf("Wrote %d bytes to file", bytes)

  // os.Open read a file struct 
  // various methods to determine where in the 
  // file you want to read
  file2, _ := os.Open("data2")
  defer file2.Close()
  
  read2 := make([]byte, len(data))
  bytes, _ = file2.Read(read2)
  fmt.Printf("Read %d bytes from file \n", bytes)
  fmt.Println(string(read2))