package api

import (
	"fmt"
	"log"
	"strings"

	"github.com/dizzyfool/genna/lib"
	"github.com/semrush/zenrpc/v2"
	"github.com/vmkteam/mfd-generator/mfd"
)

type ProjectService struct {
	*Store

	zenrpc.Service
}

func NewProjectService(store *Store) *ProjectService {
	return &ProjectService{
		Store: store,
	}
}

// Loads project from file
//zenrpc:filePath		the path to mfd file
//zenrpc:connection 	connection string
//zenrpc:return 		Project
func (s *ProjectService) Open(filePath, connection string) (*mfd.Project, error) {
	project, err := mfd.LoadProject(filePath, false, DefaultGoPGVer)
	if err != nil {
		return nil, err
	}

	var logger *log.Logger

	genna := genna.New(connection, logger)
	if err := genna.Connect(); err != nil {
		return nil, err
	}

	s.CurrentProject = project
	s.CurrentFile = filePath
	s.Genna = &genna

	return project, nil
}

// Updates project in memory
//zenrpc:project	project information
func (s *ProjectService) Update(project mfd.Project) error {
	s.CurrentProject = &project

	return nil
}

// Saves project from memory to disk
//zenrpc:project	project information
func (s *ProjectService) Save() error {
	if err := mfd.SaveMFD(s.CurrentFile, s.CurrentProject); err != nil {
		return err
	}

	if err := mfd.SaveProjectXML(s.CurrentFile, s.CurrentProject); err != nil {
		return err
	}

	if err := mfd.SaveProjectVT(s.CurrentFile, s.CurrentProject); err != nil {
		return err
	}

	//TODO Save translation

	return nil
}

// Gets all tables from database
//zenrpc:url	the connection string to pg database
//zenrpc:return	list of tables
func (s *ProjectService) Tables() ([]string, error) {
	schemas, err := s.Genna.Store.Schemas()
	if err != nil {
		return nil, err
	}

	var filter []string
	for _, schema := range schemas {
		if strings.HasPrefix(schema, "pg_") || schema == "information_schema" {
			continue
		}
		filter = append(filter, fmt.Sprintf("%s.*", schema))
	}

	entities, err := s.Genna.Read(filter, false, false, DefaultGoPGVer, nil)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(entities))
	for i, entity := range entities {
		result[i] = entity.PGFullName
	}

	return result, nil
}
