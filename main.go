package main

import (
	"math"
	"math/rand"
)

var total_system_criticality int = 3
var n int = 15
var U float64 = 4
var r int = 2
var c int = 2


func UUnifastDiscard(n int, u float64) []float64 {

	sets := []float64{}
	utilizations := []float64{} 
	sumU := u 
	for i := 0; i < n; i++ {
		nextSumU := sumU * math.Pow(rand.Float64(),(1.0 / float64(n - i)))

		utilizations = append(utilizations, sumU - nextSumU)

		sumU = nextSumU
	}
	utilizations = append(utilizations, sumU)

	for i := 0; i < len(utilizations); i++ {
		if utilizations[i] < 1 {
			sets = append(sets, utilizations[i])
		}
	}
	return sets
}

func CreateCPUMatrix(righe int, colonne int) [][]*Core {

	var matrixCore = make([][]*Core, righe)
	for i := range matrixCore {
		matrixCore[i] = make([]*Core, colonne)
	}
	
	for i := 0; i < righe; i++ {
		for j := 0; j < colonne; j++ {

			var new_core *Core = new(Core) // Create the new core
			new_core.capacity = 1

			if (i == 0 || i == righe-1) && (j == 0 || j == colonne-1) {
				new_core.type_core = "Corner"
				new_core.boundary_number = 0
			} else if (i == 0 || i == righe - 1 || j == 0 || j == colonne-1) {
				new_core.type_core = "Edge"
				new_core.boundary_number = 1
			} else {
				new_core.type_core = "Normal"
				new_core.boundary_number = 2
			}

			matrixCore[i][j] = new_core
		}
	}

	for i := 0; i < righe; i++ {
		for j := 0; j < colonne; j++ {
			if matrixCore[i][j].type_core == "Corner" {
				if i == 0 && j == 0 {
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i+1][j])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j+1])
				} else if i == 0 && j == colonne - 1 {
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j-1])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i+1][j])
				} else if i == righe-1 && j == 0 {
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i-1][j])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j+1])
				} else if i == righe-1 && j == colonne - 1 {
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i-1][j])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j-1])
				}
			} else if matrixCore[i][j].type_core == "Edge" {
				if i == 0 {
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i+1][j])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j+1])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j-1])
				} else if i == righe - 1 {
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i-1][j])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j+1])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j-1])
				} else if j == 0 {
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j+1])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i+1][j])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i-1][j])
				} else if j == colonne - 1 {
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j-1])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i+1][j])
					matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i-1][j])
				}
			} else if matrixCore[i][j].type_core == "Normal" {
				matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j-1])
				matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i][j+1])
				matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i-1][j])
				matrixCore[i][j].neighbors = append(matrixCore[i][j].neighbors, matrixCore[i+1][j])
			}
		}
	}

	return matrixCore
}

// Function that creates a balanced task set
func createTaskset(utilizzazioni []float64) []*Task {
	var taskset []*Task
	criticity_levels := makeRange(0,total_system_criticality-1)
	c := 0
	for i := 1; i < len(utilizzazioni); i++ {

		new_task := ConstructorTask(utilizzazioni[i], criticity_levels[c])	
		taskset = append(taskset, new_task)
		c++
		if c == total_system_criticality {
			c = 0
		}
	}

	return taskset
}


func main() {


	matrixCore := CreateCPUMatrix(r,c)

	utilizzazioni := UUnifastDiscard(n, U)

	taskset := createTaskset(utilizzazioni)

	// STEP 1: Task sorting
	DecreasingUtilizationOrder(taskset)
	//DecreasingCriticalityOrder(taskset)

	// STEP 2: Task allocation
	//FirstFit(taskset, matrixCore)
	//BestFit(taskset, matrixCore)
	//WorstFit(taskset, matrixCore)

	// How many tasks have been associated?
	/*
	tasks_allocated := 0
	for t := 0; t < len(taskset); t++ { 
		if taskset[t].myCore != nil {
			tasks_allocated = tasks_allocated + 1
		}

	}
	fmt.Println("Task allocati: ", tasks_allocated, " su: ", n)
	*/


	// STEP 3: Priority assignment
	//RandomPriority(matrixCore)
	RateMonotonic(matrixCore)
	//Audsley(taskset)


	// STEP 4: SCHEDULABILITY TEST
	for i := 0; i < len(matrixCore); i++ { 
		for j := 0; j < len(matrixCore[i]); j++ {
			ANALYSIS(matrixCore[i][j])
		}
	}	
}