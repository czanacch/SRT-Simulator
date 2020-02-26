package main


/****************************************** PRIORITY ASSIGNMENT ********************************************/


func RandomPriority(cores [][]*Core) {

	for i := 0; i < len(cores); i++ {
		for j := 0; j < len(cores[i]); j++ {
			for t := 0; t < len(cores[i][j].tasks); t++ {
				number_priorities := len(cores[i][j].tasks)
				cores[i][j].tasks[t].priority = randInt(0, number_priorities)
				cores[i][j].tasks[t].priority_assigned = true
			}
		}
	}
}

func RateMonotonic(cores [][]*Core) {
	for i := 0; i < len(cores); i++ {
		for j := 0; j < len(cores[i]); j++ {

			different_periods := []float64{}

			for t := 0; t < len(cores[i][j].tasks); t++ {
				if !floatInSlice(cores[i][j].tasks[t].period, different_periods) {
					different_periods = append(different_periods, cores[i][j].tasks[t].period)
				}
			}

			for i := 1; i < len(different_periods); i++ {
				z := i
				for z > 0 {
					if different_periods[z-1] < different_periods[z] {
						different_periods[z-1], different_periods[z] = different_periods[z], different_periods[z-1]
					}
					z = z - 1
				}
			}

			map_priorities := make(map[float64]int)
			z := 1
			for i := 0; i < len(different_periods); i++ {
				map_priorities[different_periods[i]] = z
				z = z + 1
			}

			for t := 0; t < len(cores[i][j].tasks); t++ {
				cores[i][j].tasks[t].priority = map_priorities[cores[i][j].tasks[t].period]
				cores[i][j].tasks[t].priority_assigned = true
			}
		}
	}
}

func Audsley(taskpool []*Task) {

	total_priority_levels := n

	for k:=0; k < total_priority_levels; k++ {

		found := false

		for t := 0; t < len(taskpool) && found == false; t++ {

			if taskpool[t].priority_assigned == false && taskpool[t].myCore != nil {
				
				taskpool[t].priority = k
			
				for t_j := 0; t_j < len(taskpool); t_j++ {
					if  taskpool[t_j] != taskpool[t] && taskpool[t_j].priority_assigned == false {
						taskpool[t_j].priority = k+1
					}
				}

				response_time := 0.0
				if taskpool[t].criticity_level == 0 {
					response_time = SteadyStateMode(taskpool[t], 0)
				} else {
					response_time = SteadyStateMode(taskpool[t], 0)
					for level := 1; level <= taskpool[t].criticity_level; level++ {
						response_time = ResponseTimeAnalysis(level, total_system_criticality, taskpool[t], 0, response_time)
					}
				}
				
				if test_analysis(response_time, taskpool[t].deadline) == true {
					taskpool[t].priority_provisional = k
					taskpool[t].priority_assigned = true
					found = true
				} else {
					taskpool[t].priority = 0
					taskpool[t].priority_assigned = false
				}
			}

		}

			
	}

	
	for t := 0; t < len(taskpool) && taskpool[t].myCore != nil; t++ {
		taskpool[t].priority = taskpool[t].priority_provisional
	}
	


}