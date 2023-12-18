package main

import (
	"fmt"
	"time"
)

func main()  {
	// timenow := time.Now()
	// fmt.Println(timenow.Format(time.RFC850))

	timeDateCustom := "2006-01-02 15:04:05";
	timeParse,err := time.Parse(time.DateTime, timeDateCustom)

	if err != nil {
		fmt.Println("error", err)
	}else{
		fmt.Println(timeParse.Format(time.RFC850))
	}
	

	/**
	time duration
	*/
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()
	fmt.Println(t1)
	fmt.Println(t2)

	duration := t2.Sub(t1)

	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())
}