package main

import (
	"math/rand"
)


type Task struct {
	criticity_level int // criticity level of task
	deadline float64
	period float64
	WCET map[int]float64
	utilization float64
	priority int
	myCore *Core // CPU associated with the task
	migrant bool // 'true' if the task can be a migrant, false otherwise
	
	priority_assigned bool
	priority_provisional int // temporary priority for Audsley algorithm

	release_jitter float64
}


func ConstructorTask(utilization float64, criticity_level int) *Task {
	p := new(Task)
	p.priority_assigned = false
	p.criticity_level = criticity_level
	p.utilization = utilization
	p.WCET = make(map[int]float64)
	p.release_jitter = 0.0
	p.migrant = rand.Int() % 2 == 0

	p.period = randFloat(10, 1000)
	p.deadline = p.period

	p.WCET[criticity_level] = p.utilization * p.period
	p.WCET[0] = p.WCET[criticity_level] / 2
	
    return p
}