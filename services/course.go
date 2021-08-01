package services

type Course struct {
	Name        string `json:"name"`
	Identifier  string `json:"identifier"`
	Description string `json:"description"`
	Instructor  User   `json:"instructor"`
	Uid         string `json:"uuid"`
}


