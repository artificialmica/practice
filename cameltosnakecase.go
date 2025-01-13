package practice



func CamelToSnakeCase(s string) string{
    CamelCase := true
    if len(s) == 0{
        return s
    }

    if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z'{
        CamelCase = false

    }
    
    for i:=0; i< len(s);i++{
        if i + 1 < len(s) && (s[i] >='A' && s[i] <='Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z'){
            CamelCase = false
        }else if !((s[i] >='A' && s[i] <='Z') || (s[i] >= 'a' && s[i] <= 'z')){
            CamelCase=false
        }
    }

    if !CamelCase {
        return s
    }

    res:=""

if CamelCase {
    for i:=0; i < len(s);i++{
        if i >0 && s[i] >='A' && s[i] <='Z'{
            res = res + "_"+ string(s[i]) 
        }else{
            res = res + string(s[i])
        }
    }
   
}

 return res




}


