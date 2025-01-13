package practice


func PrintIf(str string) string {
    result:=""
    if len(str) >=3{
        result = "G\n"
    }else if len(str) == 0{
        result= "G\n"
    }else{
        result = "Invalid Input\n"
    }
    return result
}