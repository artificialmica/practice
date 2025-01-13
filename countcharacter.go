package practice

func CountChar(str string, c rune) int {
    counter :=0
    arr:= []rune(str)

    for i := 0; i < len(str) ; i++{
        if arr[i] == c {
            counter++
        }
    }
    return counter
}
