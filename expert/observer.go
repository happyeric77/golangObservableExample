package expert

import (
	"errors"
	"fmt"
	"time"
)

func NewObserver(kind, name string) Observer {
	switch kind {
	case "westenFood":
		return WestenFoodObserver{
			Name: name,
			CreatedTime: time.Now(),
		}
	case "japanFood":
		return JapanFoodObserver{
			Name: name,
			CreatedTime: time.Now(),
		}
	}
	return nil
}

type Observer interface {
	Update()
	GetName() string
}

type WestenFoodObserver struct {
	CreatedTime time.Time
	Name string
}

func (w WestenFoodObserver) Update() {
	fmt.Printf("%s (Created at %v) Time for meal --> order Italia dishes.\n", w.Name, w.CreatedTime)
}

func (w WestenFoodObserver) GetName() string{
	return w.Name
}

type JapanFoodObserver struct {
	CreatedTime time.Time
	Name string
}

func (j JapanFoodObserver) Update() {
	fmt.Printf("%s (Created at %v)! Time for meal --> order Japanese dishes.\n", j.Name, j.CreatedTime)
}

func (j JapanFoodObserver) GetName() string{
	return j.Name
}

type Observable struct {}

func (o *Observable) Register(obs Observer, evt Eventer) {
	evt.AddEventListener(obs)
}

func (o *Observable) Unregister(obs Observer, evt Eventer) error {
	var newObservers []Observer
	for i, observer := range evt.GetObservers() {
		if obs.GetName() == observer.GetName() {
			if i+1 == len(evt.GetObservers()){
				newObservers = evt.GetObservers()[:len(evt.GetObservers())-1]
			} else if i == 0 {
				newObservers = evt.GetObservers()[1:]
			} else{
				preCut := evt.GetObservers()[:i]
				postCut := evt.GetObservers()[i+i:]
				newObservers = append(preCut, postCut...)
			}
			evt.RemoveEventListener(newObservers)
			return nil
		}
	}
	return errors.New("expertUnregister")
}

func (o *Observable) Dispatch(evt Eventer) {
	for _, observer := range evt.GetObservers() {
		observer.Update()
		fmt.Println(observer)		
	}
}

type Eventer interface {
	GetObservers() []Observer
	AddEventListener(Observer)
	RemoveEventListener([]Observer)
}

type MorningEvent struct {
	Observers []Observer
}

func (m *MorningEvent) GetObservers() []Observer {
	return m.Observers
}

func (m *MorningEvent) AddEventListener(obs Observer) {
	m.Observers = append(m.Observers, obs)
}

func (m *MorningEvent) RemoveEventListener(new []Observer) {
	m.Observers = new
}

type DinnerEvent struct {
	Observers []Observer
}

func (d *DinnerEvent) GetObservers() []Observer {
	return d.Observers
}

func (d *DinnerEvent) AddEventListener(obs Observer) {
	d.Observers = append(d.Observers, obs)
}

func (d *DinnerEvent) RemoveEventListener(new []Observer) {
	d.Observers = new
}

func Demo() {
	morningEvt := &MorningEvent{[]Observer{}}
	dinnerEvt := &DinnerEvent{[]Observer{}}
	obs := Observable{}
	eric := NewObserver("japanFood", "Eric")
	yuko := NewObserver("westenFood", "Yuko")
	pipu := NewObserver("westenFood", "Pipu")
	obs.Register(eric, morningEvt)
	obs.Register(yuko, morningEvt)
	obs.Register(pipu, dinnerEvt)
	obs.Dispatch(morningEvt)
	obs.Dispatch(dinnerEvt)
	obs.Unregister(eric, morningEvt)
	obs.Dispatch(morningEvt)
}
