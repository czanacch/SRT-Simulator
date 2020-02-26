package main

import (
	"math"
)

func SteadyStateMode(task *Task, result float64) float64 {
	my_core := task.myCore

	sum := 0.0
	response_time := result
	for t := 0; t < len(my_core.tasks); t++ {
		if my_core.tasks[t].priority > task.priority {
			sum = sum + (math.Ceil(response_time/my_core.tasks[t].period) * my_core.tasks[t].WCET[0])
		}
	}

	final_result := sum + task.WCET[0]

	if final_result == result {
		return result
	} else {
		return SteadyStateMode(task, final_result)
	}
}


func ResponseTimeAnalysis(ref int, total_levels int, task *Task, result float64, previous_value float64) float64 {
	
	my_core := task.myCore

	ref_migrant := ref - 1


	sum := 0.0

	// MIGRATION LEVEL NON MIGRANT tasks
	delta_parameter := delta(ref_migrant,ref)
	response_time := result
	for j := 0; j < len(my_core.tasks); j++ {
		if my_core.tasks[j].priority > task.priority && my_core.tasks[j].criticity_level == ref_migrant && !my_core.tasks[j].migrant { // Per ogni task a priorità più alta sullo stesso core, che si trova a "livello di migrazione" ma non è un migrante
			sum = sum + (math.Ceil((response_time + my_core.tasks[j].release_jitter) / my_core.tasks[j].period) * my_core.tasks[j].WCET[delta_parameter])
		}
	}

	// Task with GREATER CRITICALITY LEVEL
	for level := ref; level < total_levels; level++ { // ciclo esterno che scorre tutti i vari livelli che interessano

		delta_parameter = delta(level,ref)
		response_time := result // TODO: va messo un response_time := result per ogni nuovo ciclo
		for m := 0; m < len(my_core.tasks); m++ {
			if my_core.tasks[m].priority > task.priority && my_core.tasks[m].criticity_level == level { // Per ogni task a priorità più alta sullo stesso core, che appartiene al livello corrente (i migranti sono compresi)
				sum = sum + (math.Ceil((response_time + my_core.tasks[m].release_jitter) /my_core.tasks[m].period) * my_core.tasks[m].WCET[delta_parameter])
			}
		}

	}

	// MIGRATION LEVEL MIGRANT task
	delta_parameter = delta(ref_migrant,ref)
	for z := 0; z < len(my_core.tasks); z++ {
		if my_core.tasks[z].priority > task.priority && my_core.tasks[z].criticity_level == ref_migrant && my_core.tasks[z].migrant { // Per ogni task a priorità più alta sullo stesso core, che si trova a "livello di migrazione" ma non è un migrante
			sum = sum + (math.Ceil((previous_value + my_core.tasks[z].release_jitter) / my_core.tasks[z].period) * my_core.tasks[z].WCET[delta_parameter])
		}
	}

	// Task with LESS CRITICALITY LEVEL
	for level := 0; level < ref - 1; level++ {

		delta_parameter = delta(level,ref)
		for h := 0; h < len(my_core.tasks); h++ {
			if my_core.tasks[h].priority > task.priority && my_core.tasks[h].criticity_level == level {
				sum = sum + (math.Ceil((previous_value + my_core.tasks[h].release_jitter) / my_core.tasks[h].period) * my_core.tasks[h].WCET[delta_parameter])
			}
		}

	}

	final_result := sum + task.WCET[delta(task.criticity_level, ref)]
	if final_result == result {
		return result
	} else {
		return ResponseTimeAnalysis(ref, total_levels, task, final_result, previous_value)
	}

}