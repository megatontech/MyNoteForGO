// package main
// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"text/tabwriter"
// )
// func main(){
// 	//rand 包实现了伪随机数生成器.

// /*随机数由一个 Source 生成。像 Float64 和 Int 这样的顶级函数使用默认共享的 Source， 
// 它会在每次程序运行时产生一系列确定的值。若每次运行都需要不同的行为，需使用 Seed 函数来
// 初始化默认的 Source。对于多Go程并发来说，默认的 Source 是安全的。*/
// 	//不同随机数需要不同随机数种子，否则生成的会相同
// 	r:=rand.New(rand.NewSource(256));
// 	w:=tabwriter.NewWriter(os.Stdout,1,1,1,' ',0);
// 	defer w.Flush()
// 	show:=func(name string,v1,v2,v3 interface{}){
// 		fmt.Fprintf(w,"%s\t%v\t%v\t%v\n",name,v1,v2,v3)
// 	}
// 	show("Float32", r.Float32(), r.Float32(), r.Float32())
// 	show("Float64", r.Float64(), r.Float64(), r.Float64())

// 	// ExpFloat64 values have an average of 1 but decay exponentially.
// 	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())

// 	// NormFloat64 values have an average of 0 and a standard deviation of 1.
// 	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())

// 	// Int31, Int63, and Uint32 generate values of the given width.
// 	// The Int method (not shown) is like either Int31 or Int63
// 	// depending on the size of 'int'.
// 	show("Int31", r.Int31(), r.Int31(), r.Int31())
// 	show("Int63", r.Int63(), r.Int63(), r.Int63())
// 	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())

// 	// Intn, Int31n, and Int63n limit their output to be < n.
// 	// They do so more carefully than using r.Int()%n.
// 	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
// 	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
// 	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))

// 	// Perm generates a random permutation of the numbers [0, n).
// 	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
// }