package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type rowTestDbGetNames struct {
	names       []string
	errExpected error
}

type rowTestDbSelectUniqueValues struct {
	uniquValues []string
	errExpected error
}

type inputDataSelectUniqueValues struct {
	columnName string
	tableName  string
}

var testTableGetNames = []rowTestDbGetNames{
	{
		names: []string{"Ivan, Gena228"},
	},
	{
		names:       nil,
		errExpected: errors.New("ExpectedError"),
	},
}

var testTableSelectUniqueValues = []rowTestDbSelectUniqueValues{
	{
		uniquValues: []string{"1, 2"},
	},
	{
		uniquValues: nil,
		errExpected: errors.New("ExpectedError"),
	},
}

var tesInputDataSelectUniqueValues = inputDataSelectUniqueValues{
	columnName: "COLUMNNAME",
	tableName:  "TABLENAME",
}

func TestGetNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal()
	}
	dbService := Service{DB: mockDB}
	for i, row := range testTableGetNames {
		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDbRowsName(row.names)).
			WillReturnError(row.errExpected)
		names, err := dbService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}

}

func mockDbRowsName(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal()
	}
	dbService := Service{DB: mockDB}
	for i, row := range testTableSelectUniqueValues {
		mock.ExpectQuery("SELECT DISTINCT " + tesInputDataSelectUniqueValues.columnName + " FROM " + tesInputDataSelectUniqueValues.tableName).
			WillReturnRows(mockDbRowsUniqueValues(row.uniquValues)).
			WillReturnError(row.errExpected)
		names, err := dbService.SelectUniqueValues(tesInputDataSelectUniqueValues.columnName, tesInputDataSelectUniqueValues.tableName)

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.uniquValues, names, "row: %d, expected names: %s, actual names: %s", i, row.uniquValues, names)
	}

}

func mockDbRowsUniqueValues(uniquValues []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"uniquValue"})
	for _, name := range uniquValues {
		rows = rows.AddRow(name)
	}
	return rows
}
