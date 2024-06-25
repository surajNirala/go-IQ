/*
Q-
Make a struct Theater with the following fields: Seats, RWMutex, invoice chan string.
Create three methods over a struct.
The first method returns the number of seats available in Theater.

The second method book a seat in the theater. If the seat is equal to zero, no one can book it. ( In the booking method, put simple print statements that show booking has been made if seats are available)
Once the seat is booked in Theater, add the name of the user in the invoice channel.

Create a third Method printInvoice(), and run it as a goroutine that runs independently
It fetches the name from the invoice channel. After fetching the name, send an email to the user (use the print statement to demonstrate this)

Note:-
You can allow simultaneous reading of seats data but make sure only one go routine can book a seat at a time.
When all bookings are completed, channel should stop receiving data and program should quit gracefully
*/

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
