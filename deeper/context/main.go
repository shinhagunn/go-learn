package main

import (
	"fmt"
)

// Context dung de giai quyet viec cancel mot request, tao mot deadline hay truyen du lieu
// Co hai cach de tao mot context moi la context.TODO() va context.Background()
// Ca hai cach deu tra ve mot context empty tuy nhien chung khac nhau o y nghia
// Khi tao mot context chua biet su dung de lam gi thi dung context.TODO()

// type Context interface {
//     // Done returns a channel that is closed when this Context is canceled
//     // or times out.
//     Done() <-chan struct{}

//     // Err indicates why this context was canceled, after the Done channel
//     // is closed.
//     Err() error

//     // Deadline returns the time when this Context will be canceled, if any.
//     Deadline() (deadline time.Time, ok bool)

//     // Value returns the value associated with key or nil if none.
//     Value(key interface{}) interface{}
// }

func main() {
	fmt.Println("Hello Context!")
}
