package main
func Print1to20() int{
	res:=0
	for i:=1; i<=20;i++{
		res+=i
	}
	return res
}

func main(){
	Print1to20()
}