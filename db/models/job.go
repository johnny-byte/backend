package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Job struct {
	UUID        string `pg:"uuid,unique,pk,default:gen_random_uuid()" json:"uuid"`
	Title       string `pg:"title" json:"title"`
	Location    string `pg:"location" json:"location"`
	Tag         string `pg:"tag" json:"tag"`
	Description string `pg:"description" json:"description,omitempty"`
}

func (v *Job) GetAllJobs(conn *pg.DB) (*[]Job, error) {
	items := &[]Job{}
	if err := conn.Model(items).
		Select(); err != nil {
		return nil, err
	}
	fmt.Printf("items: %v\n", items)
	return items, nil
}

func (v *Job) GetAllJobsShort(conn *pg.DB) (*[]Job, error) {
	items := &[]Job{}
	if err := conn.Model(items).Column("uuid", "title", "location", "tag").
		Select(); err != nil {
		return nil, err
	}
	for _, job := range *items {
		job.Description = ""
	}
	return items, nil
}

func (job *Job) Insert(conn *pg.DB) error {
	fmt.Printf("Insert job: %v\n", job)
	if _, err := conn.Model(job).Insert(); err != nil {
		return err
	}
	return nil
}

func (job *Job) FindWithUUID(conn *pg.DB) error {
	fmt.Printf("Select job: %v\n", job)
	if err := conn.Model(job).Where("uuid = ?0", job.UUID).Select(); err != nil {
		return err
	}
	return nil
}
