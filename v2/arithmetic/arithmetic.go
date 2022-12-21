package arithmetic
import "errors"
func IsPrime(num int) bool {
    
    for i := 2; i < int(num/2); i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}
func add(num... int)int {
    sum:=0
    for _,v := range num {
        sum+=v
    }
    return sum
}
func Factorial(n int) (int,error) {
    if n<0  {
     return 0,errors.New("Please provide a +ve number greater than zero")
    }
    var f int = 1
    for i := 2; i <= n; i++ {
        f *= i
    }
    return f,nil
}