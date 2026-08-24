func calculate (a int, b int)(int, int){
	return a + b, a - b
}
func main(){
	sum, diff := calculate(10,20)
	fmt.Println("sum:", sum, "Diff:", diff)
}