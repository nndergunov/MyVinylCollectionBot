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

var Collection = newCollectionTable("public", "collection", "")

type collectionTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnInteger
	OwnerID    postgres.ColumnInteger
	AlbumID    postgres.ColumnInteger
	LocationID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CollectionTable struct {
	collectionTable

	EXCLUDED collectionTable
}

// AS creates new CollectionTable with assigned alias
func (a CollectionTable) AS(alias string) *CollectionTable {
	return newCollectionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CollectionTable with assigned schema name
func (a CollectionTable) FromSchema(schemaName string) *CollectionTable {
	return newCollectionTable(schemaName, a.TableName(), a.Alias())
}

func newCollectionTable(schemaName, tableName, alias string) *CollectionTable {
	return &CollectionTable{
		collectionTable: newCollectionTableImpl(schemaName, tableName, alias),
		EXCLUDED:        newCollectionTableImpl("", "excluded", ""),
	}
}

func newCollectionTableImpl(schemaName, tableName, alias string) collectionTable {
	var (
		IDColumn         = postgres.IntegerColumn("id")
		OwnerIDColumn    = postgres.IntegerColumn("owner_id")
		AlbumIDColumn    = postgres.IntegerColumn("album_id")
		LocationIDColumn = postgres.IntegerColumn("location_id")
		allColumns       = postgres.ColumnList{IDColumn, OwnerIDColumn, AlbumIDColumn, LocationIDColumn}
		mutableColumns   = postgres.ColumnList{IDColumn, OwnerIDColumn, AlbumIDColumn, LocationIDColumn}
	)

	return collectionTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		// Columns
		ID:         IDColumn,
		OwnerID:    OwnerIDColumn,
		AlbumID:    AlbumIDColumn,
		LocationID: LocationIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
