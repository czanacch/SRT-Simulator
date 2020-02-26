package main

import (
	"fmt"
)


/****************************************** SCHEDULABILITY TEST ********************************************/


func ANALYSIS(core *Core) {

	response_time_t := make(map[*Task]float64)

	countTrue := 0
	countFalse := 0

	for l := 0; l < total_system_criticality; l++ {				
		core.criticity_level = l


		for nei := 0; nei < len(core.neighbors); nei++ {
			MIGRATION_ALL(core.neighbors[nei])
		}

		
		// Calculate the number of tests to do
		tests_number := total_system_criticality - l - 1
		
		for k := 0; k <= tests_number; k++ {

			core.criticity_level = l + k

			for t := 0; t < len(core.tasks); t++ {
				
				if core.tasks[t].criticity_level >= core.criticity_level {
					

					if k == 0 { // Run Steady State Mode
						response_time_t[core.tasks[t]] = SteadyStateMode(core.tasks[t], 0)
					} else { // Run the "recursive case" of the analysis
						response_time_t[core.tasks[t]] = ResponseTimeAnalysis(l+k, total_system_criticality, core.tasks[t], 0, response_time_t[core.tasks[t]])
					}

					if test_analysis(response_time_t[core.tasks[t]], core.tasks[t].deadline) {
						countTrue = countTrue + 1
					} else {
						countFalse = countFalse + 1
					}

				}
			}		
		}

		MIGRATION_ME(core)
		LEAVE_CPU(core)
		
	}

	if countFalse != 0 {
		fmt.Println("Not schedulable")
	} else {
		fmt.Println("Schedulable")
	}

}


func delta(level int, ref int) int {
	res := 0
	if level < ref {
		res = level 
	} else if level == ref {
		res = ref
	} else if level > ref {
		res = 0
	}
	return res
}



func test_analysis(result float64, deadline float64) bool {
	if result < deadline {
		return true
	} else {
		return false
	}
}


func MIGRATION_ALL(core *Core) {

	for l := 0; l < total_system_criticality; l++ {

		migratable_tasks := []*Task{}

		for t := 0; t < len(core.tasks); t++ {
			if core.tasks[t].criticity_level == l && core.tasks[t].migrant {
				migratable_tasks = append(migratable_tasks, core.tasks[t])
			}
		}

		WF_migration(core, migratable_tasks)
	}
	
}

func MIGRATION_ME(core *Core) {
	migration_level := core.criticity_level - 1
	migratable_tasks := []*Task{}

	for t := 0; t < len(core.tasks); t++ {
		if core.tasks[t].criticity_level == migration_level && core.tasks[t].migrant {
			migratable_tasks = append(migratable_tasks, core.tasks[t])
		}
	}

	WF_migration(core, migratable_tasks)

}


func migration(task *Task,core_start *Core, core_end *Core) {

	response_time := 0.0
			
	if task.criticity_level == 0 { // Run Steady State Mode
		response_time = SteadyStateMode(task, 0)
	} else { // Run the "recursive case" of the analysis
		response_time = SteadyStateMode(task, 0)
		for level := 1; level <= task.criticity_level; level++ {
			response_time = ResponseTimeAnalysis(level, total_system_criticality, task, 0, response_time)
		}
	}

	task.release_jitter = response_time + task.WCET[task.criticity_level] // New Release Jitter
	task.deadline = task.deadline - task.release_jitter // New deadline of the task

	// Do the actual migration
	core_start.tasks = removeElementOfSlice(core_start.tasks, task) // Remove the task from the initial core
	core_start.capacity = core_start.capacity - task.utilization
	task.myCore = core_end // Associate the new core with the task
	core_end.tasks = append(core_end.tasks, task) // Add the task to the final core
	core_end.capacity = core_end.capacity + task.utilization
}

func LEAVE_CPU(core *Core) {
	level_threshold := core.criticity_level - 2

	if level_threshold >= 0 {
		task_to_remove := []*Task{}


		for t:= 0; t < len(core.tasks); t++ {
			if core.tasks[t].criticity_level <= level_threshold {
				task_to_remove = append(task_to_remove, core.tasks[t])
			}
		}
	
		// Delete all tasks
		for t := 0; t < len(task_to_remove); t++ {
			core.tasks = removeElementOfSlice(core.tasks, task_to_remove[t])
			core.capacity = core.capacity - task_to_remove[t].utilization
		}
	}

}
