package main

func main() {
    println("Hello Maps")

    monthdays := map[string]int{
                    "Jan": 31, "Feb": 28, "Mar": 31,
                    "Apr": 30, "May": 31, "Jun": 30,
                    "Jul": 31, "Aug": 31, "Sep": 30,
                    "Oct": 31, "Nov": 30, "Dec": 31,
                       }
    year:=0
    for _,days := range monthdays {
        year += days
    }
    println("Total days:",year)

    //overriding for leap year
    monthdays["Feb"] = 29
    year=0
    for _,days := range monthdays {
        year += days
    }
    println("Total days in leap year:",year)
}
