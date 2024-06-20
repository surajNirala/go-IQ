package main

import (
	"fmt"
	"sync"
)

type Theater struct {
	Seat    int
	RWMutex sync.RWMutex
	invoice chan string
}

func TheaterBookingProgram_theater() {
	theater := NumberofSeat_theater(5)
	names := []string{"Suraj", "Vikas", "Chandan", "Avinash", "Dinesh", "Jitendra"}
	for _, name := range names {

		theater.BookingSeat_theater(name)
	}
	var wg sync.WaitGroup
	if theater.Seat != 0 {
		go theater.PrintInvoice_theater(&wg)
	}
	wg.Wait()
}

func NumberofSeat_theater(total_seat int) Theater {
	return Theater{
		Seat:    total_seat,
		invoice: make(chan string, total_seat),
	}
}

func (t *Theater) BookingSeat_theater(name string) {
	if t.Seat > 0 {
		t.Seat--
		fmt.Printf("Booking confirmed for %s. Seats remaing %d\n", name, t.Seat)
		t.invoice <- name
	} else {
		fmt.Println("No Seat available for booking.")
	}
}

func (t *Theater) PrintInvoice_theater(wg *sync.WaitGroup) {
	wg.Add(1)
	invoice := <-t.invoice
	fmt.Println("invoice ", invoice)
	wg.Done()
}
