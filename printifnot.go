package practice


func PrintIfNot(str string) string {
    result:=""
    if len(str) >=3{
        result = "Invalid Input\n"
    }else if len(str) == 0{
        result=  "G\n"
    }else{
        result ="G\n"
    }
    return result
}