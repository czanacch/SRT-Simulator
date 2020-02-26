package main

import (
	"math/rand"
)

func floatInSlice(a float64, list []float64) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func removeElementOfSlice(slice []*Task, task *Task ) []*Task {
	newSlice := []*Task{}
	for _, x := range slice {
		if x != task {
			// copy and increment index
			newSlice = append(newSlice, x)
		}
	}
	return newSlice
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func randFloat(min float64, max float64) float64 {
    return min + rand.Float64() * (max - min)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

