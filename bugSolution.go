```go
func main() {
    slice := []int{1, 2, 3, 4, 5}
    newSlice := []int{}
    for _, v := range slice {
        if v != 4 {
            newSlice = append(newSlice, v)
        }
    }
    fmt.Println(newSlice)
}
```