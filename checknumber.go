package practice

func CheckNumber(arg string) bool {
        check:= false
        arr:= []rune(arg)

        for i:= 0; i < len(arg); i++{
            if arr[i] == '0'{
                check = true
            }else if arr[i] == '1'{
                check = true
            }else if arr[i] == '2'{
                check = true
            }else if arr[i] == '3'{
                check = true
            }else if arr[i] == '4'{
                check = true
            }else if arr[i] == '5'{
                check = true
            }else if arr[i] == '6'{
                check = true
            }else if arr[i] == '7'{
                check = true
            }else if arr[i] == '8'{
                check = true
            }else if arr[i] == '9'{
                check = true
            }
        }
        return check
}
