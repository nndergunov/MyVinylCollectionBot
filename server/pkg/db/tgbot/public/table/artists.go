//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Artists = newArtistsTable("public", "artists", "")

type artistsTable struct {
	postgres.Table

	//Columns
	ID   postgres.ColumnInteger
	Name postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ArtistsTable struct {
	artistsTable

	EXCLUDED artistsTable
}

// AS creates new ArtistsTable with assigned alias
func (a ArtistsTable) AS(alias string) *ArtistsTable {
	return newArtistsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArtistsTable with assigned schema name
func (a ArtistsTable) FromSchema(schemaName string) *ArtistsTable {
	return newArtistsTable(schemaName, a.TableName(), a.Alias())
}

func newArtistsTable(schemaName, tableName, alias string) *ArtistsTable {
	return &ArtistsTable{
		artistsTable: newArtistsTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newArtistsTableImpl("", "excluded", ""),
	}
}

func newArtistsTableImpl(schemaName, tableName, alias string) artistsTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		NameColumn     = postgres.StringColumn("name")
		allColumns     = postgres.ColumnList{IDColumn, NameColumn}
		mutableColumns = postgres.ColumnList{IDColumn, NameColumn}
	)

	return artistsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:   IDColumn,
		Name: NameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}