func main (){
	replicas := 3
	if replicas > 2 {
		fmt.Println("Multiple replicas running")
	}else{
		fmt.Println("Single replica running")
	}
}