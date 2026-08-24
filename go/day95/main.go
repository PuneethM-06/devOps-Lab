func calculate (a int, b int)(int, int){
	return (a + b, a - b)
}
func main(){
	sum, diff := calculate(10,20)
}