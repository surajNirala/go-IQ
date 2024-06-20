package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Message struct {
	Name  string
	Phone int
}

func main() {

	// reverse1() //Custom
	// reverse2("Hello")               //Infoys
	// swapTwoNumber(10, 20)           //Infoys
	// sortByInts()                    //Using sort package
	// sortByStrings()                 //Using sort package
	// sortByFloat64s()                //Using sort package
	// sortByKeySlice1()               //Using sort package
	// sortByKeySlice2()               // Infoys
	// sortByKeyTwoSlice()             // Infoys
	// getCountDuplicateValueInSlice() // Persistent
	// getUniqueSlice()                // create Custom
	// minimumValueFromSlice()
	// maximumValueFromSlice()
	// sortByValue() // NTT DATA
	// goroutineLogical() // HCL
	// CLICobra() //Persistent (pending)
	// CRUD_operation_No_DB() // Persistent (pending)
	//TheaterBookingProgram_theater() //Bristlecone go run main.go theaterbooking.go
}

func sortByFloat64s() {
	floatSlice := []float64{3.14, 1.59, 2.71, 4.66}
	// Sort the slice of floats
	sort.Float64s(floatSlice)
	fmt.Println("Sorted float slice:", floatSlice)
}

func sortByStrings() {
	// Create a slice of strings
	strSlice := []string{"banana", "apple", "cherry", "date"}
	// Sort the slice of strings
	sort.Strings(strSlice)
	// Print the sorted slice
	fmt.Println("Sorted string slice:", strSlice)
}

func sortByInts() {
	data := []int{1, 4, 6, 2, 1, 8, 7}
	fmt.Println("Before Sorting : ", data)
	sort.Ints(data)
	fmt.Println("After Sorting : ", data)
}

func swapTwoNumber(a int, b int) {
	a = a + b // 30
	b = a - b // 30-20
	a = a - b // 30-10
	fmt.Println(" a : ", a)
	fmt.Println("b :", b)
}

func reverse1() {
	reverseString := "suraj"
	storeReverseString := ""
	for i := len(reverseString) - 1; i >= 0; i-- {
		storeReverseString += string(reverseString[i])
	}
	fmt.Println("storeReverseString :==", storeReverseString)
}

func reverse2(s string) {
	var r []string
	var p []string
	for _, v := range s {
		fmt.Println("v", string(v))
		r = append(r, string(v))
	}
	fmt.Println("r", r)
	for i := len(r) - 1; i >= 0; i-- {
		p = append(p, r[i])
	}
	fmt.Println("reverse value is:", strings.Join(p, ""))
}

func sortByKeySlice2() {
	unsortingSlice := make(map[string]string)
	unsortingSlice["suraj"] = "NIrala"
	unsortingSlice["prakash"] = "Sharma"
	unsortingSlice["avinash"] = "Verma"

	newArrslice := []string{}
	for index, _ := range unsortingSlice {
		newArrslice = append(newArrslice, index)
	}
	sort.Strings(newArrslice)
	fmt.Println(newArrslice) // suraj,prakash,avinash
	for _, v := range newArrslice {
		fmt.Println(v, unsortingSlice[v])
	}
}

func sortByKeySlice1() {
	message := []Message{
		{
			"Suraj", 12345678,
		},
		{
			"Nirala", 23456543,
		},
		{
			"Boby", 987654,
		},
	}
	sort.Slice(message, func(i, j int) bool {
		return message[i].Name < message[j].Name
	})
	fmt.Println("message : ", message)
}

func sortByKeyTwoSlice() {
	strings := []string{"Apple", "banana", "Apple", "cherry", "banana", "date"}
	strings2 := []string{"Apple", "Bana", "orange", "Bana", "Bana", "orange"}

	customSlice1 := make(map[string]int)
	for _, v1 := range strings {
		customSlice1[v1]++
	}
	// fmt.Println("customSlice1 := ", customSlice1)
	sliceData := []string{}
	for index, v := range customSlice1 {
		if v >= 1 {
			sliceData = append(sliceData, index)
		}
	}
	fmt.Println("customUnique1 := ", sliceData)
	customSlice2 := make(map[string]int)
	for _, s1 := range strings2 {
		customSlice2[s1]++
	}
	// fmt.Println("customSlice2 = ", customSlice2)
	customUnique := []string{}
	for val, s2 := range customSlice2 {
		if s2 >= 1 {
			customUnique = append(customUnique, val)
		}
	}
	fmt.Println("customUnique2 := ", customUnique)

	mergeSlices := append(sliceData, customUnique...)

	fmt.Println("mergeSlices : ", mergeSlices)

	getMergeUniqueSlice := make(map[string]int)
	for _, mergeSlice := range mergeSlices {
		getMergeUniqueSlice[mergeSlice]++
		// if mergeSlice >= 1 {
		// 	getMergeUniqueSlice = append(getMergeUniqueSlice, index)
		// }
	}
	getMergeUniqueSlice1 := []string{}
	for index, v := range getMergeUniqueSlice {
		if v >= 1 {
			getMergeUniqueSlice1 = append(getMergeUniqueSlice1, index)
		}
	}

	fmt.Println("getMergeUniqueSlice1 := ", getMergeUniqueSlice1)
}

func getCountDuplicateValueInSlice() {
	data := []string{"Apple", "Banana", "Cherry", "Orange", "Cherry", "Cherry"}
	temp := make(map[string]int)
	for _, value := range data {
		temp[value]++
	}
	fmt.Println("OutPut := ", temp)
}

func getUniqueSlice() {
	strings := []string{"Apple", "Banana", "Apple", "cherry", "date", "cherry", "cherry"}
	temp := make(map[string]int)
	for _, v := range strings {
		temp[v]++
	}
	data := []string{}
	for index, v := range temp {
		if v >= 1 {
			data = append(data, index)
		}
	}
	fmt.Println("Data is : ", data)
}

func minimumValueFromSlice() {
	data := []int{3, 4, 2, 5, 6, 8, 6, 1, 2}
	min := data[0]
	for _, v := range data {
		if min > v {
			min = v
		}
	}
	fmt.Println("Minimum Value is : ", min)
}

func maximumValueFromSlice() {
	data := []int{212, 33, 439, 2, 3, 33}
	max := data[0]
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	fmt.Println("Maximum value is : ", max)
}

func sortByValue() {
	//ip: [1,9,0,3,0,0,1,0]
	//op: [0,0,0,0,1,9,3,1]
	data := []int{1, 9, 0, 3, 0, 0, 1, 0}
	minZero := 0
	zerodataArr := []int{}
	dataArr := []int{}
	for _, v := range data {
		if minZero == v {
			zerodataArr = append(zerodataArr, v)
		} else {
			dataArr = append(dataArr, v)
		}
	}
	finalSlice := []int{}
	finalSlice = append(zerodataArr, dataArr...)
	fmt.Println("finalSlice := ", finalSlice)
}

func goroutineLogical() {
	ch_all_message := make(chan int)
	var wg sync.WaitGroup
	// Launching receiver goroutine
	go func() {
		for msg := range ch_all_message {
			fmt.Println("Received message:", msg)
		}
	}()
	// var input int
	// fmt.Scan(&input)
	// ch_all_message <- 10
	// fmt.Printf("Name is %s\n", input)
	for i := 1; i <= 100; i++ {
		// var data int
		// fmt.Scan(&data)
		wg.Add(1)
		input := i
		go func(input int) {
			defer wg.Done()
			ch_all_message <- input
		}(input)
	}
	wg.Wait()
	close(ch_all_message)
}
