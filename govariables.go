package main

import "fmt"

func main() {
    var a int = 25
    var b float32 = 32.25
    fmt.Printf("value of a=%d b=%f \n",a,b)

    var c bool = false
    fmt.Printf("value of c=%t \n",c)

    var d,e = 12,25.36
    fmt.Printf("value of d=%d, e=%f \n",d,e)

    var (
        f int = 25
        g bool = true
    )
    fmt.Printf("value of f=%d, g=%t \n",f,g)

	var s string = "hello go language"
	s2:=[]byte(s)
	s2[6] = 'G'
	s = string(s2)
	fmt.Printf("value of s=%s \n",s)
	
	var cx complex64 = 5 + 5i
	fmt.Printf("value of complex number=%v \n",cx)
	
}
