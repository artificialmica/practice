package practice



func FromTo(from int, to int) string {
if from > 99 || from < 0 || to > 99 || to < 0 {
	return "Invalid" + "\n"
} 
var s []string
if from == to {
	return Itoa(from)+"\n"
}
if from < to {
	for i := from ; i <= to; i++ {
		if i < 10 {
			s = append(s, "0" +Itoa(i)+", ")
			}  else {
				if i == to {
					s = append(s, Itoa(i))
				} else {
				s = append(s, Itoa(i)+", ")
			}
		}
	}
}  
if from > to {
	for i := from ; i >= to; i-- {
		if i < 10 {
			if i == to {
				s = append(s, "0"+Itoa(i))
			} else { s = append(s, "0" +Itoa(i)+", ")
			}
			}  else {
				if i == to {
					s = append(s, Itoa(i))
				} else {
				s = append(s, Itoa(i)+", ")
			}
		}
	}
}

str := ""
for _,ch:= range s {
	str += ch
}
return str + "\n"
}

func Itoa(n int) string {
	var r []rune 
	s := ""
var i int
	if n == 0 {
		return "0"
	}
for n > 0 {
	i = n % 10

r = append(r,rune(i +'0'))
n=n/10
}

for i := 0 ; i < len(r);i++ {
	for j := len(r)-1; j > 0;j-- {
		r[i],r[j] = r[j],r[i]
	}
}

for _,ch := range r {
	s += string(ch)
}
return s
}