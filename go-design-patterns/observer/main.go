package main

func main() {

	listener1 := DataListener{
		Name: "Listener 1",
	}

	listener2 := DataListener{
		Name: "Listener 2",
	}

	subject := &DataSubject{}

	subject.registerObserver(listener1)
	subject.registerObserver(listener2)

	subject.ChangeItem("Monday!")
	subject.ChangeItem("Wednesday!")

	subject.deregisterObserver(listener2)

	subject.ChangeItem("Monday!")
	subject.ChangeItem("Wednesday!")

}
