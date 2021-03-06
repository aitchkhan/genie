package manager

import (
	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/httpio"
	"github.com/mohuk/genie/models"
)

type mockmanager struct {
	store dbase.Store
}

func (m *mockmanager) GetColumns(dbname, tableName string) (*models.TableForm, error) {
	return nil, nil
}
func (m *mockmanager) GetDatabases() ([]httpio.Database, error) {
	return []httpio.Database{
		httpio.Database{
			Name: "db",
		},
	}, nil
}
func (m *mockmanager) GetTables(dbname string) ([]httpio.Table, error) {
	return nil, nil
}

func NewMockGenieManager(s dbase.Store) GenieManager {
	return &mockmanager{store: s}
}
