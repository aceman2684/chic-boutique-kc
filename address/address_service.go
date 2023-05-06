package address

import (
	"database/sql"
	"errors"
	"strings"
)

// loadSQLFile is a type of function for loading a SQL file. This should take in a fil ename and return the contents of
// the SQL file containing the given file name. If no file with the given file name is found, or if there were any
// issues loading the SQL file, an error should be returned.
type loadSQLFile func(fileName string) (string, error)

// AddressService is a struct representing a service used for handling addresses
type AddressService struct {
	db *sql.DB
	loadSQLFile
}

// NewAddressService creates a new *AddressService
func NewAddressService(db *sql.DB, loadloadSqlFile loadSQLFile) *AddressService {
	return &AddressService{
		db,
		loadloadSqlFile,
	}
}

// CreateAddress creates a new address. This returns a *Address containing the information of the created address. This
// returns an error if there were any issues creating the address or if any of the given criteria is considered invalid.
//
// The criteria is considered invalid if any of the following are true:
//   - lineOne is empty or blank
//   - lineTwo is not nil and is empty or blank
//   - city is empty or blank
//   - state is empty or blank
//   - zipCode is empty or blank
func (as AddressService) CreateAddress(
	lineOne string,
	lineTwo *string,
	city string,
	state string,
	zipCode string,
) (*Address, error) {
	if err := validateAddress(lineOne, lineTwo, city, state, zipCode); err != nil {
		return nil, err
	}

	sqlStatement, err := as.loadSQLFile("create_address")
	if err != nil {
		return nil, err
	}

	tx, err := as.db.Begin()
	if err != nil {
		return nil, err
	}

	var addressId int
	if err := tx.QueryRow(sqlStatement, lineOne, lineTwo, city, state, zipCode).Scan(&addressId); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	address := &Address{
		Id:      addressId,
		LineOne: lineOne,
		LineTwo: lineTwo,
		City:    city,
		State:   state,
		ZipCode: zipCode,
	}

	return address, nil
}

func validateAddress(
	lineOne string,
	lineTwo *string,
	city string,
	state string,
	zipCode string,
) error {
	if lineOne = strings.TrimSpace(lineOne); lineOne == "" {
		return errors.New("lineOne cannot be empty or blank")
	}

	if lineTwo != nil {
		lineTwoValue := *lineTwo

		if lineTwoValue = strings.TrimSpace(lineTwoValue); lineTwoValue == "" {
			return errors.New("lineTwo cannot be empty or blank")
		}
	}

	if city = strings.TrimSpace(city); city == "" {
		return errors.New("city cannot be empty or blank")
	}

	if state = strings.TrimSpace(state); state == "" {
		return errors.New("state cannot be empty or blank")
	}

	if zipCode = strings.TrimSpace(zipCode); zipCode == "" {
		return errors.New("zipCode cannot be empty or blank")
	}

	return nil
}
