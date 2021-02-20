package main

import (
	"observerExample/expert"
	"observerExample/advanced"
	"observerExample/basic"
)

func ObserverDemo (kind string) {
	switch kind {
	case "basic":
		basic.Demo()
	case "advanced":
		advanced.Demo()
	case "expert":
		expert.Demo()
	}	
	

}

func main() {
	ObserverDemo("expert")
}



