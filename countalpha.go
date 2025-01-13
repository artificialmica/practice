package practice

func CountAlpha(s string) int {
    counter:=0
    arr:=[]rune(s)

    for i:=0;i<len(arr);i++{
        if arr[i] >= 'A' && arr[i] <= 'Z' ||  arr[i] >= 'a' && arr[i] <= 'z'{
            counter++
        }
    }
    return counter
}