package practice


func DigitLen(n, base int) int {
        counter :=0


    if base < 2 || base > 36{
        return -1
    }

    if n < 0{
        n = -n
    }

    for n > 0{
        n = n/base
        counter++
    }
    return counter
}