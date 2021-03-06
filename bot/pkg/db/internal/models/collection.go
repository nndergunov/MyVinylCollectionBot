// Code generated by SQLBoiler 4.9.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Collection is an object representing the database table.
type Collection struct {
	ID         int `boil:"id" json:"id" toml:"id" yaml:"id"`
	AlbumID    int `boil:"album_id" json:"album_id" toml:"album_id" yaml:"album_id"`
	LocationID int `boil:"location_id" json:"location_id" toml:"location_id" yaml:"location_id"`

	R *collectionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L collectionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CollectionColumns = struct {
	ID         string
	AlbumID    string
	LocationID string
}{
	ID:         "id",
	AlbumID:    "album_id",
	LocationID: "location_id",
}

var CollectionTableColumns = struct {
	ID         string
	AlbumID    string
	LocationID string
}{
	ID:         "collection.id",
	AlbumID:    "collection.album_id",
	LocationID: "collection.location_id",
}

// Generated where

var CollectionWhere = struct {
	ID         whereHelperint
	AlbumID    whereHelperint
	LocationID whereHelperint
}{
	ID:         whereHelperint{field: "\"collection\".\"id\""},
	AlbumID:    whereHelperint{field: "\"collection\".\"album_id\""},
	LocationID: whereHelperint{field: "\"collection\".\"location_id\""},
}

// CollectionRels is where relationship names are stored.
var CollectionRels = struct {
	Album    string
	Location string
}{
	Album:    "Album",
	Location: "Location",
}

// collectionR is where relationships are stored.
type collectionR struct {
	Album    *Album    `boil:"Album" json:"Album" toml:"Album" yaml:"Album"`
	Location *Location `boil:"Location" json:"Location" toml:"Location" yaml:"Location"`
}

// NewStruct creates a new relationship struct
func (*collectionR) NewStruct() *collectionR {
	return &collectionR{}
}

// collectionL is where Load methods for each relationship are stored.
type collectionL struct{}

var (
	collectionAllColumns            = []string{"id", "album_id", "location_id"}
	collectionColumnsWithoutDefault = []string{"album_id", "location_id"}
	collectionColumnsWithDefault    = []string{"id"}
	collectionPrimaryKeyColumns     = []string{"id"}
	collectionGeneratedColumns      = []string{}
)

type (
	// CollectionSlice is an alias for a slice of pointers to Collection.
	// This should almost always be used instead of []Collection.
	CollectionSlice []*Collection
	// CollectionHook is the signature for custom Collection hook methods
	CollectionHook func(boil.Executor, *Collection) error

	collectionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	collectionType                 = reflect.TypeOf(&Collection{})
	collectionMapping              = queries.MakeStructMapping(collectionType)
	collectionPrimaryKeyMapping, _ = queries.BindMapping(collectionType, collectionMapping, collectionPrimaryKeyColumns)
	collectionInsertCacheMut       sync.RWMutex
	collectionInsertCache          = make(map[string]insertCache)
	collectionUpdateCacheMut       sync.RWMutex
	collectionUpdateCache          = make(map[string]updateCache)
	collectionUpsertCacheMut       sync.RWMutex
	collectionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var collectionAfterSelectHooks []CollectionHook

var collectionBeforeInsertHooks []CollectionHook
var collectionAfterInsertHooks []CollectionHook

var collectionBeforeUpdateHooks []CollectionHook
var collectionAfterUpdateHooks []CollectionHook

var collectionBeforeDeleteHooks []CollectionHook
var collectionAfterDeleteHooks []CollectionHook

var collectionBeforeUpsertHooks []CollectionHook
var collectionAfterUpsertHooks []CollectionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Collection) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Collection) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Collection) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Collection) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Collection) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Collection) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Collection) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Collection) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Collection) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range collectionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCollectionHook registers your hook function for all future operations.
func AddCollectionHook(hookPoint boil.HookPoint, collectionHook CollectionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		collectionAfterSelectHooks = append(collectionAfterSelectHooks, collectionHook)
	case boil.BeforeInsertHook:
		collectionBeforeInsertHooks = append(collectionBeforeInsertHooks, collectionHook)
	case boil.AfterInsertHook:
		collectionAfterInsertHooks = append(collectionAfterInsertHooks, collectionHook)
	case boil.BeforeUpdateHook:
		collectionBeforeUpdateHooks = append(collectionBeforeUpdateHooks, collectionHook)
	case boil.AfterUpdateHook:
		collectionAfterUpdateHooks = append(collectionAfterUpdateHooks, collectionHook)
	case boil.BeforeDeleteHook:
		collectionBeforeDeleteHooks = append(collectionBeforeDeleteHooks, collectionHook)
	case boil.AfterDeleteHook:
		collectionAfterDeleteHooks = append(collectionAfterDeleteHooks, collectionHook)
	case boil.BeforeUpsertHook:
		collectionBeforeUpsertHooks = append(collectionBeforeUpsertHooks, collectionHook)
	case boil.AfterUpsertHook:
		collectionAfterUpsertHooks = append(collectionAfterUpsertHooks, collectionHook)
	}
}

// One returns a single collection record from the query.
func (q collectionQuery) One(exec boil.Executor) (*Collection, error) {
	o := &Collection{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for collection")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Collection records from the query.
func (q collectionQuery) All(exec boil.Executor) (CollectionSlice, error) {
	var o []*Collection

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Collection slice")
	}

	if len(collectionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Collection records in the query.
func (q collectionQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count collection rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q collectionQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if collection exists")
	}

	return count > 0, nil
}

// Album pointed to by the foreign key.
func (o *Collection) Album(mods ...qm.QueryMod) albumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AlbumID),
	}

	queryMods = append(queryMods, mods...)

	query := Albums(queryMods...)
	queries.SetFrom(query.Query, "\"albums\"")

	return query
}

// Location pointed to by the foreign key.
func (o *Collection) Location(mods ...qm.QueryMod) locationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LocationID),
	}

	queryMods = append(queryMods, mods...)

	query := Locations(queryMods...)
	queries.SetFrom(query.Query, "\"locations\"")

	return query
}

// LoadAlbum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (collectionL) LoadAlbum(e boil.Executor, singular bool, maybeCollection interface{}, mods queries.Applicator) error {
	var slice []*Collection
	var object *Collection

	if singular {
		object = maybeCollection.(*Collection)
	} else {
		slice = *maybeCollection.(*[]*Collection)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &collectionR{}
		}
		args = append(args, object.AlbumID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &collectionR{}
			}

			for _, a := range args {
				if a == obj.AlbumID {
					continue Outer
				}
			}

			args = append(args, obj.AlbumID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`albums`),
		qm.WhereIn(`albums.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Album")
	}

	var resultSlice []*Album
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Album")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for albums")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for albums")
	}

	if len(collectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Album = foreign
		if foreign.R == nil {
			foreign.R = &albumR{}
		}
		foreign.R.Collections = append(foreign.R.Collections, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AlbumID == foreign.ID {
				local.R.Album = foreign
				if foreign.R == nil {
					foreign.R = &albumR{}
				}
				foreign.R.Collections = append(foreign.R.Collections, local)
				break
			}
		}
	}

	return nil
}

// LoadLocation allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (collectionL) LoadLocation(e boil.Executor, singular bool, maybeCollection interface{}, mods queries.Applicator) error {
	var slice []*Collection
	var object *Collection

	if singular {
		object = maybeCollection.(*Collection)
	} else {
		slice = *maybeCollection.(*[]*Collection)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &collectionR{}
		}
		args = append(args, object.LocationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &collectionR{}
			}

			for _, a := range args {
				if a == obj.LocationID {
					continue Outer
				}
			}

			args = append(args, obj.LocationID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`locations`),
		qm.WhereIn(`locations.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Location")
	}

	var resultSlice []*Location
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Location")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for locations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for locations")
	}

	if len(collectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Location = foreign
		if foreign.R == nil {
			foreign.R = &locationR{}
		}
		foreign.R.Collections = append(foreign.R.Collections, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LocationID == foreign.ID {
				local.R.Location = foreign
				if foreign.R == nil {
					foreign.R = &locationR{}
				}
				foreign.R.Collections = append(foreign.R.Collections, local)
				break
			}
		}
	}

	return nil
}

// SetAlbum of the collection to the related item.
// Sets o.R.Album to related.
// Adds o to related.R.Collections.
func (o *Collection) SetAlbum(exec boil.Executor, insert bool, related *Album) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"collection\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"album_id"}),
		strmangle.WhereClause("\"", "\"", 2, collectionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AlbumID = related.ID
	if o.R == nil {
		o.R = &collectionR{
			Album: related,
		}
	} else {
		o.R.Album = related
	}

	if related.R == nil {
		related.R = &albumR{
			Collections: CollectionSlice{o},
		}
	} else {
		related.R.Collections = append(related.R.Collections, o)
	}

	return nil
}

// SetLocation of the collection to the related item.
// Sets o.R.Location to related.
// Adds o to related.R.Collections.
func (o *Collection) SetLocation(exec boil.Executor, insert bool, related *Location) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"collection\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"location_id"}),
		strmangle.WhereClause("\"", "\"", 2, collectionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LocationID = related.ID
	if o.R == nil {
		o.R = &collectionR{
			Location: related,
		}
	} else {
		o.R.Location = related
	}

	if related.R == nil {
		related.R = &locationR{
			Collections: CollectionSlice{o},
		}
	} else {
		related.R.Collections = append(related.R.Collections, o)
	}

	return nil
}

// Collections retrieves all the records using an executor.
func Collections(mods ...qm.QueryMod) collectionQuery {
	mods = append(mods, qm.From("\"collection\""))
	return collectionQuery{NewQuery(mods...)}
}

// FindCollection retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCollection(exec boil.Executor, iD int, selectCols ...string) (*Collection, error) {
	collectionObj := &Collection{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"collection\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, collectionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from collection")
	}

	if err = collectionObj.doAfterSelectHooks(exec); err != nil {
		return collectionObj, err
	}

	return collectionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Collection) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no collection provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(collectionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	collectionInsertCacheMut.RLock()
	cache, cached := collectionInsertCache[key]
	collectionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			collectionAllColumns,
			collectionColumnsWithDefault,
			collectionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(collectionType, collectionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(collectionType, collectionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"collection\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"collection\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into collection")
	}

	if !cached {
		collectionInsertCacheMut.Lock()
		collectionInsertCache[key] = cache
		collectionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the Collection.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Collection) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	collectionUpdateCacheMut.RLock()
	cache, cached := collectionUpdateCache[key]
	collectionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			collectionAllColumns,
			collectionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update collection, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"collection\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, collectionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(collectionType, collectionMapping, append(wl, collectionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update collection row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for collection")
	}

	if !cached {
		collectionUpdateCacheMut.Lock()
		collectionUpdateCache[key] = cache
		collectionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q collectionQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for collection")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for collection")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CollectionSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), collectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"collection\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, collectionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in collection slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all collection")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Collection) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no collection provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(collectionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	collectionUpsertCacheMut.RLock()
	cache, cached := collectionUpsertCache[key]
	collectionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			collectionAllColumns,
			collectionColumnsWithDefault,
			collectionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			collectionAllColumns,
			collectionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert collection, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(collectionPrimaryKeyColumns))
			copy(conflict, collectionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"collection\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(collectionType, collectionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(collectionType, collectionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert collection")
	}

	if !cached {
		collectionUpsertCacheMut.Lock()
		collectionUpsertCache[key] = cache
		collectionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single Collection record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Collection) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Collection provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), collectionPrimaryKeyMapping)
	sql := "DELETE FROM \"collection\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from collection")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for collection")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q collectionQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no collectionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from collection")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for collection")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CollectionSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(collectionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), collectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"collection\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, collectionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from collection slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for collection")
	}

	if len(collectionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Collection) Reload(exec boil.Executor) error {
	ret, err := FindCollection(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CollectionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CollectionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), collectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"collection\".* FROM \"collection\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, collectionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CollectionSlice")
	}

	*o = slice

	return nil
}

// CollectionExists checks if the Collection row exists.
func CollectionExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"collection\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if collection exists")
	}

	return exists, nil
}
