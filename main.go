package main

// func main() {
//     message := NewMessage()
//     greeter := NewGreeter(message)
//     event := NewEvent(greeter)

//     event.Start()
// }

func main() {
    e := InitializeEvent()

    e.Start()
}