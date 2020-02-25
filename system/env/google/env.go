package google

import (
	"log"
	"os"

	"cloud.google.com/go/compute/metadata"
)

// GetProject on Google Cloud
func GetProject() string {
	var (
		project string
		err     error
	)

	project, err = metadata.ProjectID()
	if err != nil {
		if project = os.Getenv("GOOGLE_PROJECT"); project == "" {
			log.Println("project id can't be empty")
			os.Exit(1)
		}
	}

	return project
}
