package services

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "html/template"
	"net/http"
	"submit/templates"
	"time"
)

type CloudKey struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Id        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
}

type Submission struct {
	CloudKey
	Name        string
	Identifier  string
	Description string
	Files       []File `gorm:"foreignKey:Id;references:Id"`
	Units       []Unit
}

func (s *Submission) Get(writer http.ResponseWriter, reader *http.Request) {

}

type File struct {
	CloudKey
	Name   string
	Anchor string
}

type Unit struct {
	CloudKey
	Name   string
	Passed bool
}

func SubmissionRoute(router chi.Router) {
	// POST /		Create
	// GET /{id}	Read
	// UPDATE /{id} Update
	// DELETE /{id}/delete

	router.Get("/", createSubmission)

	router.Post("/", createSubmission)

	router.Route("/{id}", func(r chi.Router) {
		router.Get("/", viewSubmission)
		router.Delete("/", createSubmission)
	})

}

func getAll(writer http.ResponseWriter, reader *http.Request) {
	var submissions []Submission
	Database().Find(&submissions)

}
func createSubmission(writer http.ResponseWriter, reader *http.Request) {
	submission := Submission{
		Name:        "Pineapple",
		Identifier:  "pine-101",
		Description: "Poppers for pineapples!!",
	}
	Database().Create(&submission)
	// _, _ = writer.Write([]byte("OKAY"))
	http.Redirect(writer, reader, "/submission/aaa", 302)

}

func viewSubmission(writer http.ResponseWriter, reader *http.Request) {

	var submission []Submission
	Database().Find(&submission)

	s := struct {
		Submissions []Submission
	}{
		Submissions: submission,
	}

	templates.Write("root.gohtml", writer, s)
}
