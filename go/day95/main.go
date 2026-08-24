 func divide(a int, b int)(int, erro){
	if b == 0{
		return errors.New("cannot divide by zero")
	} else{
		return a/b
	}
 }
 func main(){
	div, err := divide(10, 2)
 }