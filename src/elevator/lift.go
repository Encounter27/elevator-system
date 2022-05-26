package elevator

import "fmt"

const (
	OPENED     = 0
	CLOSED     = 1
	GOING_UP   = 2
	GOING_DOWN = 3
)

type Lift struct {
	Id              int
	Floor           int
	State           int
	UPWARDS_TASKS   []bool
	DOWNWARDS_TASKS []bool
	HighestStopage  int
	LowestStopage   int
}

func NewLift(id int, floors int) Lift {
	l := Lift{}

	l.Id = id
	l.Floor = floors
	l.State = OPENED
	l.UPWARDS_TASKS = make([]bool, floors)
	l.DOWNWARDS_TASKS = make([]bool, floors)

	for i := 0; i < floors; i++ {
		l.UPWARDS_TASKS[i] = false
		l.DOWNWARDS_TASKS[i] = false
	}

	l.HighestStopage = -1
	l.LowestStopage = -1

	return l
}

func (l *Lift) Move() {

}

func (l *Lift) CalculateCostOfNewReq() int {
	return 0
}

func (l *Lift) AcceptNewReq() {

}

func (l *Lift) GetStatus() string {
	state := []string{"OPENED", "CLOSED"}

	return fmt.Sprintf("Lift:%d Floor:%d State:%s", l.Id, l.Floor, state[l.State&1])
}
