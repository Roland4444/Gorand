package  main
import (
    "fmt"
	"os"
	"math/rand"	
	"time"
)

func gen() (float32){
	return (float32(rand.Float64()))
}

func main(){	
	const delta_mult float32 = 10000
	const delta float32 = 300
    var outs float32 = 75.0
    const step int64 = 150//300;
	const iteration int32 = 16*8640000
	var outs___ [iteration]float32
	start___ := time.Now()

	var i int32
    for i=0; i<iteration; i++{
        if (gen() < 0.5){
            if (gen() < 0.5) {outs = outs - gen() / delta}
            if (gen() > 0.5) {outs = outs + gen() / delta}  //adding
        }

        if (gen() > 0.5){
            if (gen() < 0.5) {outs = outs * (1 + gen() / delta_mult)}
            if (gen() > 0.5) {outs = outs * (1 - gen() / delta_mult)} //multyple
        }
        outs___[i] = outs;
    }
	open_, err := os.Create("open")
	high_, err := os.Create("high")
	low_, err := os.Create("low")
	close_, err := os.Create("close")
	if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    var open, high, low, close, it float32
    var counter int64 = 1
    var start bool = true
    var current float32 = 0.0
    open = 0; high =0; low = 0; close = 0;    
    for i=0; i<iteration; i++{
        it=outs___[i];
        if ((counter % step == 1)   && (start == true)){
            open = it; high = close;  close = low;  low = it;
            start = false;
        };
        if ((counter % step == 1)   && (start == false)){
            high = close; close = low; low = it;
            start = false;
        }
        if (counter % step == 0){
			close = it;			
			fmt.Fprintln(open_, open)
			fmt.Fprintln(close_, close)
			fmt.Fprintln(high_, high)
			fmt.Fprintln(low_, low)	
			open = close
        }
        current = it
        if current > high { high = current}
        if (current < low) {low = current}
        counter += 1
    }
    open_.Close() 
    high_.Close() 
    low_.Close() 
	close_.Close()     
	duration := time.Since(start___)
	fmt.Println("Execution time::", duration)
}


