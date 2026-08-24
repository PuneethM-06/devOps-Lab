 func divide(a int, b int)(int, error){
	if b == 0{
		return 0, errors.New("cannot divide by zero")
	} else{
		return a/b, nil 
	}
 }
 func main(){
	div, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(div)
	}
 }