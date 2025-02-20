package main

type observable interface {
	registerObserver(obs observer)
	deregisterObserver(obs observer)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener
	field     string
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataSubject) deregisterObserver(o DataListener) {
	var newObserverListeners []DataListener
	for _, obs := range ds.observers {
		if o.Name != obs.Name {
			newObserverListeners = append(newObserverListeners, obs)

		}
	}
	ds.observers = newObserverListeners
}

func (ds *DataSubject) notifyAll() {
	for _, observer := range ds.observers {
		observer.onUpdate(ds.field)
	}
}
