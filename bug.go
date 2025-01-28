```go
func main() {
    slice := []int{1, 2, 3, 4, 5}
    for i := range slice {
        if i == 3 {
            slice = append(slice[:i], slice[i+1:]...)
        }
    }
    fmt.Println(slice)
}
```