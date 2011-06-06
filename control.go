package main

import (
	"fmt"
)

func checkGoto() {
	i:=0
HERE:
	fmt.Printf("%d \n",i)
	i++
	if i<10 {
			goto HERE
	}
}

func checkFor() {
	sum:=0
	for i:=0;i<11;i++ {
		sum += i
	}
	fmt.Printf("sum=%d \n",sum)

	x:=0
	for x<11 {
		fmt.Printf("x=%d \n",x)
		x++
	}
	
	for a,b := 0,10 ; a < b ; a,b =a+2,b+1 {
			fmt.Printf("a=%d b=%d \n",a,b)
		}
}

func checkBreak() {
	for i:=0;i<11;i++ {
			if i==5 {
				break
			}
			fmt.Printf("i=%d \n",i)
		}
}

func checkContinue() {
	for i:=0;i<11;i++ {
			if i==5 {
				continue
			}
			fmt.Printf("i=%d \n",i)
		}
}

func checkRange() {
	list:= []string{"Ram","Rahim","Rohan","Suresh","Munda"}
	for k,v := range list {
		fmt.Printf("value at index %d: %s\n",k,v)
	}
}

func checkSwitch()  {
	day:=20
	switch(day) {
		case 0:
			fmt.Printf("Sunday \n")
		case 1:
			fmt.Printf("Monday \n")
		case 2:
			fmt.Printf("Tuesday \n")
		case 3:
			fmt.Printf("Wednesday \n")
		case 4:
			fmt.Printf("Thursday \n")
		case 5:
			fmt.Printf("Friday \n")
		case 6:
			fmt.Printf("Saturday \n")
		default:
			fmt.Printf("Wrong Value \n")
	}
	
	month:=5
	switch(month) {
		case 1,3,5,7,8,10,12:
			fmt.Printf("No of days 31")
		case 4,6,9,11:
			fmt.Printf("No of days 30")
		case 2:
			fmt.Printf("No of days 28")
		default:
			fmt.Printf("Wrong Value")	
	}
}

func main() {
	x:= 2
	y:= 5
	if x>y {
			fmt.Printf("x is greater \n")
		}else{
			fmt.Printf("y is greater \n")
		}
		
	if x=20/2; x>y {
		fmt.Printf("x is greater \n")
	} else {
		fmt.Printf("y is greater \n")
	}
	
	checkGoto()	
	checkFor()
	checkBreak()
	checkContinue()
	checkRange()
	checkSwitch()
}
