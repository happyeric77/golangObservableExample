package basic


import (
	"fmt"
	"time"
)

type Observer interface {
	Update()
	GetName() string
}

type Subscriber struct {
	CreatedTime time.Time
	Name string
}

func (s Subscriber) Update() {
	fmt.Printf("%s (Created at %v) got a notification.\n", s.Name, s.CreatedTime)
}
func (s Subscriber) GetName() string{
	return s.Name
}

type Observable interface {
	Register()
	Unregister()
	Dispatch()
}

type Notifier struct {
	Observers []Observer
}

func (n *Notifier) Register(obs Observer) {
	n.Observers = append(n.Observers, obs)
}

func (n *Notifier) Unregister(obs Observer) {
	var newObservers []Observer
	for i, observer := range n.Observers{
		if obs.GetName() == observer.GetName() {
			if i+1 == len(n.Observers){
				newObservers = n.Observers[:len(n.Observers)-1]
			} else if i == 0 {
				newObservers = n.Observers[1:]
			} else{
				preCut := n.Observers[:i]
				postCut := n.Observers[i+i:]
				newObservers := append(preCut, postCut...)
				n.Observers = newObservers
			}
			n.Observers = newObservers		
		}
	}
}

func (n *Notifier) Dispatch() {
	for _, observer := range n.Observers {
		observer.Update()		
	}
}

func Demo() {
	obs := Notifier{Observers: []Observer{}}
	eric := Subscriber{Name: "Eric", CreatedTime: time.Now()}
	yuko := Subscriber{Name: "Yuko", CreatedTime: time.Now()}
	pipu := Subscriber{Name: "Pipu", CreatedTime: time.Now()}
	obs.Register(eric)
	obs.Register(yuko)
	obs.Register(pipu)
	obs.Dispatch()
	obs.Unregister(eric)
	obs.Dispatch()
}


