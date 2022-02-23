package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Job struct {
	UUID        string `pg:"uuid,unique,pk,default:gen_random_uuid()" json:"uuid"`
	Title       string `pg: "title" json: "title" `
	Location    string `pg: "location" json: "location" `
	Tag         string `pg: "tag" json: "tag"`
	Description string `pg: "description" json: "description"`
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

func (job *Job) Insert(conn *pg.DB) error {
	fmt.Printf("Insert job: %v\n", job)
	if _, err := conn.Model(job).Insert(); err != nil {
		return err
	}
	return nil
}
