package main


type Core struct {
	type_core string 
	criticity_level int 
	tasks []*Task
	boundary_number int
	neighbors []*Core // list of pointers to other adjacent cores: they are the direct neighbors of the core (from 2 to 4 cores)

	capacity float64
}