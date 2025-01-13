package practice

func StrRev(s string)string{
	arr:=[]rune(s)
	rev:=""
	for i := len(arr) - 1 ; i >=0; i--{
		rev = rev + string(arr[i])
		
	}
	return rev
	
}