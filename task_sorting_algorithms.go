package main

/****************************************** TASK SORTING ********************************************/

// Decreasing Utilization
func DecreasingUtilizationOrder(tasks []*Task) {
	// Insertion-sort used for the order
	var n = len(tasks)
    for i := 1; i < n; i++ {
        j := i
        for j > 0 {
            if tasks[j-1].utilization < tasks[j].utilization {
                tasks[j-1], tasks[j] = tasks[j], tasks[j-1]
            }
            j = j - 1
        }
    }
}

// Decreasing Criticality
func DecreasingCriticalityOrder(tasks []*Task) {
	// Insertion-sort used for the order
	var n = len(tasks)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if tasks[j-1].criticity_level < tasks[j].criticity_level {
				tasks[j-1], tasks[j] = tasks[j], tasks[j-1]
			}
			j = j - 1
		}
	}
}