package practice

func RetainFirstHalf(str string) string {
    if len(str) == 1 || len(str) == 0{
        return str
    }

   return str[:len(str)/2]

}