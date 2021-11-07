package db

import (
	"log"

	pg "github.com/go-pg/pg/v10"
	orm "github.com/go-pg/pg/v10/orm"
)

type Project struct {
	ProjectID   int32  `sql:"ProjectID, pk"`
	ProjectName string `sql:"ProjectName, unique"`
	Promoter    string `sql:"promoter"`
	Investments int16  `sql:"investments"`
	Description string `sql:"Desription"`
}

type Investment struct {
	InvestmentID int32  `sql:"InvestmentID, pk"`
	Name         string `sql:"Name"`
	ProjectID    int32  `sql:"ProjectID, unique"`
}

type Status struct {
	InvestmentID int32  `sql:"InvestmentID, pk"`
	StatusID     string `sql:"StatusID"`
	ExpectedYear int32  `sql:"ExpectedYear"`
}

type StatusType struct {
	StatusID int32  `sql:"StatusID, pk"`
	Name     string `sql:"Name, unique"`
}

type FromTo struct {
	InvestmentID int32  `sql:"InvestmentID, pk"`
	FromCity     string `sql:"FromCity"`
	ToCity       string `sql:"ToCity"`
	FromTSO      string `sql:"FromTSO"`
	ToTSO        string `sql:"ToTSO"`
}

type Country struct {
	InvestmentID int32  `sql:"InvestmentID, pk"`
	Country1     string `sql:"Country1"`
	Country2     string `sql:"Country2"`
	Country3     string `sql:"Country3"`
}

type Technology struct {
	InvestmentID int32 `sql:"InvestmentID, pk"`
	TypeID       int32 `sql:"TypeID"`
	Voltage      int32 `sql:"Voltage"`
}

type Type struct {
	TypeID      int32  `sql:"TypeID, pk"`
	TypeName    string `sql:"CurrentType"`
	ElementType string `sql:"ElementType"`
}

// Insert new Item to DB
func (pi *Project) Save(db *pg.DB) error {
	_, insertErr := db.Model(pi).Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Project %s has inserted successfully.\n", pi.ProjectName)
	return nil
}

// Insert new Item and show value
func (pi *Project) SaveAndReturn(db *pg.DB) (Project, error) {
	InsertResult, insertErr := db.Model(pi).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", insertErr)
		return *pi, insertErr
	}
	log.Printf("Project %s has inserted successfully.\n", pi.ProjectName)
	log.Printf("Project has received new value: %v", InsertResult.Model())
	return *pi, nil
}

// Used for building the different Tables
func CreateProjectTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&Project{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table Project, Reason: %v\n", createErr)
	}
	log.Printf("Table Project created successfully \n")
	return nil
}

func CreateInvestmentTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&Investment{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table Investment, Reason: %v\n", createErr)
	}
	log.Printf("Table Investment created successfully \n")
	return nil
}

func CreateStatusTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&Status{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table Status, Reason: %v\n", createErr)
	}
	log.Printf("Table Status created successfully \n")
	return nil
}

func CreateStatusTypeTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&StatusType{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table StatusType, Reason: %v\n", createErr)
	}
	log.Printf("Table StatusType created successfully \n")
	return nil
}

func CreateFromToTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&FromTo{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table FromTo, Reason: %v\n", createErr)
	}
	log.Printf("Table FromTo created successfully \n")
	return nil
}

func CreateCountryTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&Country{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table Country, Reason: %v\n", createErr)
	}
	log.Printf("Table Country created successfully \n")
	return nil
}

func CreateTechnologyTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&Technology{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table Technology, Reason: %v\n", createErr)
	}
	log.Printf("Table Technology created successfully \n")
	return nil
}

func CreateTypeTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&Type{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table Type, Reason: %v\n", createErr)
	}
	log.Printf("Table Type created successfully \n")
	return nil
}
