package services


type Result struct {
	passed bool
}

type Conclusion struct {
	passed bool
}

type Assignment struct {
	name        string
	identifier  string
	description string

	course Course

	uuid string
}

