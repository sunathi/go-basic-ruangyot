package main

import "fmt"

// ควรทำเป็น interface เสมอ
type (
	Database interface {
		Insert() error
		Update() error
	}
	PostgresDb struct{} //real database
	MockDb     struct{}
)

// check use all method for PostgresDb
func NewPostgresDb() Database {
	return &PostgresDb{}
}

func NewMockDb() Database {
	return &MockDb{}
}

func (p *PostgresDb) Insert() error {
	fmt.Println("real insert")
	return nil
}

func (p *PostgresDb) Update() error {
	fmt.Println("real update")
	return nil
}

func (m *MockDb) Insert() error {
	fmt.Println("mock insert")
	return nil
}

func (m *MockDb) Update() error {
	fmt.Println("mock update")
	return nil
}

func InsertPlayerItem(db Database) error {
	return db.Insert()
}

func UpdatePlayerItem(db Database) error {
	return db.Update()
}

func main() {
	postgres := &PostgresDb{}
	mock := &MockDb{}
	InsertPlayerItem(postgres)
	InsertPlayerItem(mock)
	UpdatePlayerItem(postgres)
	UpdatePlayerItem(mock)
}
