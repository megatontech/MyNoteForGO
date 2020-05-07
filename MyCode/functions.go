package main
import "fmt"
func add() int{
    return 1+1
}


func addNumbers(x int,y int) int{
    return x+y
}
func main (){
    fmt.Println(1+1);
    fmt.Println(add());
    fmt.Println(addNumbers(1,2));

}