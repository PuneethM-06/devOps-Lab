func main (){
	cpuUsage := 75.0
	if cpuUsage > 80 {
		fmt.Println("High CPU Usage")
	} else if cpuUsage >50 {
		fmt.Println("Moderate CPU usage")
	}else{
		fmt.Println("low CPU usage")
	}
}