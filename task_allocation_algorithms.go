package main

import (
	"math"
)


/****************************************** TASK ALLOCATION ********************************************/

func FirstFit(tasks []*Task, cores [][]*Core) {
	
	for t := 0; t < len(tasks); t++ {

		found := false
		
		for i := 0; i < len(cores) && found == false; i++ { 
			for j := 0; j < len(cores[i]) && found == false; j++ { 
				if tasks[t].utilization <= cores[i][j].capacity {

					cores[i][j].capacity = cores[i][j].capacity - tasks[t].utilization
					cores[i][j].tasks = append(cores[i][j].tasks, tasks[t])
					tasks[t].myCore = cores[i][j]
					
					found = true
				}
			}
		}

	}
}


func BestFit(tasks []*Task, cores [][]*Core) {

	for t := 0; t < len(tasks); t++ {

		candidates_core := make(map[*Core]float64)

		for i := 0; i < len(cores); i++ { 
			for j := 0; j < len(cores[i]); j++ {

				if tasks[t].utilization < cores[i][j].capacity {

					candidates_core[cores[i][j]] = cores[i][j].capacity - tasks[t].utilization
				}

			}
		}

		min_residual := math.MaxFloat64
		for _, residual := range candidates_core {
			if residual < min_residual {
				min_residual = residual
			}
		}


		found := false
		for i := 0; i < len(cores) && found == false; i++ { 
			for j := 0; j < len(cores[i]) && found == false; j++ {

				if cores[i][j].capacity - tasks[t].utilization == min_residual {
					cores[i][j].capacity = cores[i][j].capacity - tasks[t].utilization
					cores[i][j].tasks = append(cores[i][j].tasks, tasks[t])
					tasks[t].myCore = cores[i][j]
					
					found = true
				}

			}
		}		

	}


}

func WorstFit(tasks []*Task, cores [][]*Core) {

	for t := 0; t < len(tasks); t++ {

		candidates_core := make(map[*Core]float64)

		for i := 0; i < len(cores); i++ { 
			for j := 0; j < len(cores[i]); j++ {

				if tasks[t].utilization < cores[i][j].capacity {

					candidates_core[cores[i][j]] = cores[i][j].capacity - tasks[t].utilization
				}

			}
		}

		max_residual := 0.0
		for _, residual := range candidates_core {
			if residual > max_residual {
				max_residual = residual
			}
		}


		found := false
		for i := 0; i < len(cores) && found == false; i++ { 
			for j := 0; j < len(cores[i]) && found == false; j++ {

				if cores[i][j].capacity - tasks[t].utilization == max_residual {
					cores[i][j].capacity = cores[i][j].capacity - tasks[t].utilization
					cores[i][j].tasks = append(cores[i][j].tasks, tasks[t])
					tasks[t].myCore = cores[i][j]
					
					found = true
				}

			}
		}		

	}

}

func FF_migration(core *Core, migratable_tasks []*Task) {
	for t := 0; t < len(migratable_tasks); t++ {
		found := false
		for CPU := 0; CPU < len(core.neighbors) && !found; CPU++ {
			if migratable_tasks[t].utilization <= core.neighbors[CPU].capacity {
				found = true
				migration(migratable_tasks[t], core, core.neighbors[CPU])
			}
		}
	}
}

func BF_migration(core *Core, tasks []*Task) {
	
	for t := 0; t < len(tasks); t++ {
		candidates_core := make(map[*Core]float64)

		for i := 0; i < len(core.neighbors); i++ { 
			if tasks[t].utilization < core.neighbors[i].capacity {
				candidates_core[core.neighbors[i]] = core.neighbors[i].capacity - tasks[t].utilization
			}
		}
		min_residual := math.MaxFloat64
		for _, residual := range candidates_core {
			if residual < min_residual {
				min_residual = residual
			}
		}

		found := false
		for i := 0; i < len(core.neighbors) && found == false; i++ { 

			if core.neighbors[i].capacity - tasks[t].utilization == min_residual {
								
				found = true
				migration(tasks[t], core, core.neighbors[i])
			}

		}	
		
	}
	
}

func WF_migration(core *Core, tasks []*Task) {
	
	for t := 0; t < len(tasks); t++ {
		candidates_core := make(map[*Core]float64)

		for i := 0; i < len(core.neighbors); i++ { 
			if tasks[t].utilization < core.neighbors[i].capacity {
				candidates_core[core.neighbors[i]] = core.neighbors[i].capacity - tasks[t].utilization
			}
		}
		max_residual := 0.0
		for _, residual := range candidates_core {
			if residual > max_residual {
				max_residual = residual
			}
		}

		found := false
		for i := 0; i < len(core.neighbors) && found == false; i++ { 

			if core.neighbors[i].capacity - tasks[t].utilization == max_residual {
								
				found = true
				migration(tasks[t], core, core.neighbors[i])
			}

		}	
		
	}
	
}