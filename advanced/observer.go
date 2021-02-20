package advanced


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

type Subscriber1 struct {
	Name string
	Age int
}

func (s Subscriber) Update() {
	fmt.Printf("%s (Created at %v), Time to have Lunch.\n", s.Name, s.CreatedTime)
}

func (s Subscriber1) Update() {
	fmt.Printf("%s (age: %v), Time to have dinner\n", s.Name, s.Age)
}


func (s Subscriber) GetName() string{
	return s.Name
}

func (s Subscriber1) GetName() string{
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
	yuko := Subscriber1{Name: "Yuko", Age: 18}
	pipu := Subscriber{Name: "Pipu", CreatedTime: time.Now()}
	obs.Register(eric)
	obs.Register(pipu)
	obs.Dispatch()
	obs.Unregister(eric)
	obs.Unregister(pipu)
	obs.Register(yuko)
	obs.Dispatch()
}

