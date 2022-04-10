// Code generated by SQLBoiler 4.9.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Artist is an object representing the database table.
type Artist struct {
	ID   int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`

	R *artistR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L artistL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ArtistColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var ArtistTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "artists.id",
	Name: "artists.name",
}

// Generated where

var ArtistWhere = struct {
	ID   whereHelperint
	Name whereHelpernull_String
}{
	ID:   whereHelperint{field: "\"artists\".\"id\""},
	Name: whereHelpernull_String{field: "\"artists\".\"name\""},
}

// ArtistRels is where relationship names are stored.
var ArtistRels = struct {
	Albums string
}{
	Albums: "Albums",
}

// artistR is where relationships are stored.
type artistR struct {
	Albums AlbumSlice `boil:"Albums" json:"Albums" toml:"Albums" yaml:"Albums"`
}

// NewStruct creates a new relationship struct
func (*artistR) NewStruct() *artistR {
	return &artistR{}
}

// artistL is where Load methods for each relationship are stored.
type artistL struct{}

var (
	artistAllColumns            = []string{"id", "name"}
	artistColumnsWithoutDefault = []string{}
	artistColumnsWithDefault    = []string{"id", "name"}
	artistPrimaryKeyColumns     = []string{"id"}
	artistGeneratedColumns      = []string{}
)

type (
	// ArtistSlice is an alias for a slice of pointers to Artist.
	// This should almost always be used instead of []Artist.
	ArtistSlice []*Artist
	// ArtistHook is the signature for custom Artist hook methods
	ArtistHook func(context.Context, boil.ContextExecutor, *Artist) error

	artistQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	artistType                 = reflect.TypeOf(&Artist{})
	artistMapping              = queries.MakeStructMapping(artistType)
	artistPrimaryKeyMapping, _ = queries.BindMapping(artistType, artistMapping, artistPrimaryKeyColumns)
	artistInsertCacheMut       sync.RWMutex
	artistInsertCache          = make(map[string]insertCache)
	artistUpdateCacheMut       sync.RWMutex
	artistUpdateCache          = make(map[string]updateCache)
	artistUpsertCacheMut       sync.RWMutex
	artistUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var artistAfterSelectHooks []ArtistHook

var artistBeforeInsertHooks []ArtistHook
var artistAfterInsertHooks []ArtistHook

var artistBeforeUpdateHooks []ArtistHook
var artistAfterUpdateHooks []ArtistHook

var artistBeforeDeleteHooks []ArtistHook
var artistAfterDeleteHooks []ArtistHook

var artistBeforeUpsertHooks []ArtistHook
var artistAfterUpsertHooks []ArtistHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Artist) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Artist) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Artist) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Artist) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Artist) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Artist) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Artist) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Artist) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Artist) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range artistAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddArtistHook registers your hook function for all future operations.
func AddArtistHook(hookPoint boil.HookPoint, artistHook ArtistHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		artistAfterSelectHooks = append(artistAfterSelectHooks, artistHook)
	case boil.BeforeInsertHook:
		artistBeforeInsertHooks = append(artistBeforeInsertHooks, artistHook)
	case boil.AfterInsertHook:
		artistAfterInsertHooks = append(artistAfterInsertHooks, artistHook)
	case boil.BeforeUpdateHook:
		artistBeforeUpdateHooks = append(artistBeforeUpdateHooks, artistHook)
	case boil.AfterUpdateHook:
		artistAfterUpdateHooks = append(artistAfterUpdateHooks, artistHook)
	case boil.BeforeDeleteHook:
		artistBeforeDeleteHooks = append(artistBeforeDeleteHooks, artistHook)
	case boil.AfterDeleteHook:
		artistAfterDeleteHooks = append(artistAfterDeleteHooks, artistHook)
	case boil.BeforeUpsertHook:
		artistBeforeUpsertHooks = append(artistBeforeUpsertHooks, artistHook)
	case boil.AfterUpsertHook:
		artistAfterUpsertHooks = append(artistAfterUpsertHooks, artistHook)
	}
}

// One returns a single artist record from the query.
func (q artistQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Artist, error) {
	o := &Artist{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for artists")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Artist records from the query.
func (q artistQuery) All(ctx context.Context, exec boil.ContextExecutor) (ArtistSlice, error) {
	var o []*Artist

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Artist slice")
	}

	if len(artistAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Artist records in the query.
func (q artistQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count artists rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q artistQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if artists exists")
	}

	return count > 0, nil
}

// Albums retrieves all the album's Albums with an executor.
func (o *Artist) Albums(mods ...qm.QueryMod) albumQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"albums\".\"artist_id\"=?", o.ID),
	)

	query := Albums(queryMods...)
	queries.SetFrom(query.Query, "\"albums\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"albums\".*"})
	}

	return query
}

// LoadAlbums allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (artistL) LoadAlbums(ctx context.Context, e boil.ContextExecutor, singular bool, maybeArtist interface{}, mods queries.Applicator) error {
	var slice []*Artist
	var object *Artist

	if singular {
		object = maybeArtist.(*Artist)
	} else {
		slice = *maybeArtist.(*[]*Artist)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &artistR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &artistR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`albums`),
		qm.WhereIn(`albums.artist_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load albums")
	}

	var resultSlice []*Album
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice albums")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on albums")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for albums")
	}

	if len(albumAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Albums = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &albumR{}
			}
			foreign.R.Artist = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.ArtistID) {
				local.R.Albums = append(local.R.Albums, foreign)
				if foreign.R == nil {
					foreign.R = &albumR{}
				}
				foreign.R.Artist = local
				break
			}
		}
	}

	return nil
}

// AddAlbums adds the given related objects to the existing relationships
// of the artist, optionally inserting them as new records.
// Appends related to o.R.Albums.
// Sets related.R.Artist appropriately.
func (o *Artist) AddAlbums(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Album) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ArtistID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"albums\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"artist_id"}),
				strmangle.WhereClause("\"", "\"", 2, albumPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ArtistID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &artistR{
			Albums: related,
		}
	} else {
		o.R.Albums = append(o.R.Albums, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &albumR{
				Artist: o,
			}
		} else {
			rel.R.Artist = o
		}
	}
	return nil
}

// SetAlbums removes all previously related items of the
// artist replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Artist's Albums accordingly.
// Replaces o.R.Albums with related.
// Sets related.R.Artist's Albums accordingly.
func (o *Artist) SetAlbums(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Album) error {
	query := "update \"albums\" set \"artist_id\" = null where \"artist_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Albums {
			queries.SetScanner(&rel.ArtistID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Artist = nil
		}
		o.R.Albums = nil
	}

	return o.AddAlbums(ctx, exec, insert, related...)
}

// RemoveAlbums relationships from objects passed in.
// Removes related items from R.Albums (uses pointer comparison, removal does not keep order)
// Sets related.R.Artist.
func (o *Artist) RemoveAlbums(ctx context.Context, exec boil.ContextExecutor, related ...*Album) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ArtistID, nil)
		if rel.R != nil {
			rel.R.Artist = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("artist_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Albums {
			if rel != ri {
				continue
			}

			ln := len(o.R.Albums)
			if ln > 1 && i < ln-1 {
				o.R.Albums[i] = o.R.Albums[ln-1]
			}
			o.R.Albums = o.R.Albums[:ln-1]
			break
		}
	}

	return nil
}

// Artists retrieves all the records using an executor.
func Artists(mods ...qm.QueryMod) artistQuery {
	mods = append(mods, qm.From("\"artists\""))
	return artistQuery{NewQuery(mods...)}
}

// FindArtist retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindArtist(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Artist, error) {
	artistObj := &Artist{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"artists\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, artistObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from artists")
	}

	if err = artistObj.doAfterSelectHooks(ctx, exec); err != nil {
		return artistObj, err
	}

	return artistObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Artist) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no artists provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(artistColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	artistInsertCacheMut.RLock()
	cache, cached := artistInsertCache[key]
	artistInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			artistAllColumns,
			artistColumnsWithDefault,
			artistColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(artistType, artistMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(artistType, artistMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"artists\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"artists\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into artists")
	}

	if !cached {
		artistInsertCacheMut.Lock()
		artistInsertCache[key] = cache
		artistInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Artist.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Artist) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	artistUpdateCacheMut.RLock()
	cache, cached := artistUpdateCache[key]
	artistUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			artistAllColumns,
			artistPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update artists, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"artists\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, artistPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(artistType, artistMapping, append(wl, artistPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update artists row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for artists")
	}

	if !cached {
		artistUpdateCacheMut.Lock()
		artistUpdateCache[key] = cache
		artistUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q artistQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for artists")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for artists")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ArtistSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), artistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"artists\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, artistPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in artist slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all artist")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Artist) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no artists provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(artistColumnsWithDefault, o)

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

	artistUpsertCacheMut.RLock()
	cache, cached := artistUpsertCache[key]
	artistUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			artistAllColumns,
			artistColumnsWithDefault,
			artistColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			artistAllColumns,
			artistPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert artists, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(artistPrimaryKeyColumns))
			copy(conflict, artistPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"artists\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(artistType, artistMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(artistType, artistMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert artists")
	}

	if !cached {
		artistUpsertCacheMut.Lock()
		artistUpsertCache[key] = cache
		artistUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Artist record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Artist) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Artist provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), artistPrimaryKeyMapping)
	sql := "DELETE FROM \"artists\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from artists")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for artists")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q artistQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no artistQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from artists")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for artists")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ArtistSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(artistBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), artistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"artists\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, artistPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from artist slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for artists")
	}

	if len(artistAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Artist) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindArtist(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArtistSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ArtistSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), artistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"artists\".* FROM \"artists\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, artistPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ArtistSlice")
	}

	*o = slice

	return nil
}

// ArtistExists checks if the Artist row exists.
func ArtistExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"artists\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if artists exists")
	}

	return exists, nil
}
