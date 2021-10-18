package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MachineStatus interface {
	getStatus()
	setStatus(state string)
	getTrigger()
	setTriggers()
	getStatusTypes()
}

type statusType struct {
	status  []string
	trigger []string
	action  []string
	counter int
}

func (st *statusType) getStatus() []string {

	return st.status
}

func (st *statusType) getTrigger() []string {

	return st.trigger
}

func (st *statusType) setStatus(state *string) {
	st.status = []string{*state}
}

func (st *statusType) setStatusTriggerTypes(statusCount *int, triggerCount *int) {

	st.status = make([]string, *statusCount)
	st.trigger = make([]string, *triggerCount)

	// needs to make it one instead writing 2 times

	for i := 0; i < *statusCount; i++ {

		fmt.Printf(" %d. Status Type: \n", i+1)
		fmt.Scanln(&st.status[i])

	}

	for i := 0; i < *triggerCount; i++ {

		fmt.Printf("%d. Trigger KeyWord: \n", i+1)
		fmt.Scanln(&st.trigger[i])

	}

}

func contains(keyword string, keywordArrays []string) bool {
	for _, key := range keywordArrays {
		if keyword == key {
			return true
		}
	}
	return false
}

func main() {

	rand.Seed(time.Now().UnixNano())

	var st statusType
	var statusCount int
	var triggerCount int

	var defaultStatus, nextStatus, triggerWord string

	fmt.Println("How many Status Types the Program will be ? ")
	fmt.Scanln(&statusCount)
	fmt.Println("How many Trigger Types the Program will be ?")
	fmt.Scanln(&triggerCount)
	st.setStatusTriggerTypes(&statusCount, &triggerCount)
	fmt.Println("Which Status you want to be default ?")
	fmt.Scanln(&defaultStatus)
	st.setStatus(&defaultStatus)
	fmt.Println("Next Status ?")
	fmt.Scanln(&nextStatus)
	fmt.Println("Initial Counter Value : ", st.counter)

	fmt.Println("..........Program Running....... ")
	for {
		fmt.Println("Trigger Word please ?")
		fmt.Scanln(&triggerWord)

		if contains(triggerWord, st.getTrigger()) {

			st.setStatus(&defaultStatus)
			fmt.Println("Current Status : ", st.getStatus())

			fmt.Println("Time Limit Set at 5 Seconds ")
			st.counter++
			time.Sleep(5 * time.Second)
			st.setStatus(&nextStatus)
			fmt.Println("Current Status : ", st.getStatus())
			// breaking Condition
			if st.counter == 6 {
				break
			}
		} else {
			fmt.Println("Keyword Mis-Match.. Exiting")
			break
		}
	}
	fmt.Println("Final Value of the counter is  :", st.counter)
}
