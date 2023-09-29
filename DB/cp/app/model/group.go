package model

// Group symbolizes group of fitness club
type Group struct {
	GroupID       int
	ProgramID     int
	Notes         string
	TrainerID     int
	ClientsAmount int
}

type Program struct {
	ProgramID   int
	ProgramName string
}
