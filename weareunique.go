package practice


import "strings"



func WeAreUnique(str1, str2 string) int {
	s := ""
if str1 == "" && str2 == "" {
	return -1
}

if str1 == "" {
	for _,ch := range str2 {
		if !strings.Contains(s,string(ch)) {
			s+= string(ch)
		}
	}
	return len(s)
}

if str2 == "" {
	for _, ch := range str1 {
		if  !strings.Contains(s,string(ch))  {
			s+= string(ch)
		}
	}
	return len(s)
}
	for _,ch := range str1 {
		if !strings.Contains(str2,string(ch)) &&  !strings.Contains(s,string(ch)) {
s+=string(ch)
		}
	for _,c := range str2 {
if !strings.Contains(str1,string(c)) && !strings.Contains(s,string(c)) {
	s+=string(c)
}
	}
}

return len(s)
}