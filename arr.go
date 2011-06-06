package main

func checkArray() {
	var arr [10]int;
	arr[0] = 25
	arr[1] = 26
	
	for i:=0;i<10;i++ {
	    println("value at index ",i," = ",arr[i])
	}
}

func checkMultiArray() {
     arr:= [2][4]string{{"a","b","c","d"},{"e","f","g","h"}}
     for i:=0;i<2;i++ {
        for j:=0;j<4;j++ {
            println("value at index ",i,j,"=",arr[i][j])
        }
     }
}

func basicSlice() {
    arr:= [...]int{1,2,3,4,5}
    s1 := arr[2:4]
    for i:=0;i<len(s1);i++ {
	    println("value at index ",i," = ",s1[i])
	}
	println("")

    s2 := arr[1:5]
    for i:=0;i<len(s2);i++ {
	    println("value at index ",i," = ",s2[i])
	}
	println("")

    s3 := arr[:]
    for i:=0;i<len(s3);i++ {
	    println("value at index ",i," = ",s3[i])
	}
	println("")

	s4 := arr[:4]
    for i:=0;i<len(s4);i++ {
	    println("value at index ",i," = ",s4[i])
	}
	println("")

	s5 := s2[:]
    for i:=0;i<len(s5);i++ {
	    println("value at index ",i," = ",s5[i])
	}
	println("")
}

func sliceAppendCopy() {
    arr:= [5]string {"a","b","c","d","e"}
    sl:= arr[0:5]
    for i:=0;i<len(sl);i++ {
	    println("value at index ",i," = ",sl[i])
	}
	println("Max Elements=",cap(sl))
	println("Add elements by append")
	sl=append(sl,"f","g","h")
    for i:=0;i<len(sl);i++ {
	    println("value at index ",i," = ",sl[i])
	}


    println("")
    var arr2 = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    var s = make([]int, 6)
    n1 := copy(s, arr2[0:])
    println("copied element:",n1)
    for i:=0;i<len(s);i++ {
	    println("value at index ",i," = ",s[i])
	}

    n2 := copy(s, s[2:])
    println("copied element:",n2)
    for i:=0;i<len(s);i++ {
	    println("value at index ",i," = ",s[i])
	}
}

func main() {
	println("Hello Arrays")
	checkArray()

	println("\nMultidimensinal Arrays")
	checkMultiArray()

    println("\nSlice")
	basicSlice()

	println("\nSlice Append and Copy")
	sliceAppendCopy()
}
