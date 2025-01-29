package structures

import "log"

func RunHashTests() {
	hash := map[string]int{
		"hi": 3,
	}

	hash["test"] = 2

	hash["anotherTest"] = 25

	hash["otherOne"] = 13

	delete(hash, "test")
	_, exists := hash["test"]
	if !exists {
		log.Println("deleted value")
	}

}
