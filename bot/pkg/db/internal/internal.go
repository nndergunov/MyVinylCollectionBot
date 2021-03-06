package internal

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nndergunov/tgBot/bot/pkg/db/internal/models"
	"github.com/nndergunov/tgBot/bot/pkg/domain/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DB struct {
	db *sql.DB
}

func NewDB(dbURL string) (*DB, error) {
	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("NewDB: %w", err)
	}

	return &DB{
		db: database,
	}, nil
}

func (d DB) DeleteAllInfos() error {
	_, err := models.Infos().DeleteAll(d.db)
	if err != nil {
		return fmt.Errorf("DeleteAllInfos: %w", err)
	}

	return nil
}

func (d DB) InsertInfo(startTime time.Time, timeLayout string) error {
	var info models.Info

	info.Starttime = startTime.Format(timeLayout)

	err := info.Insert(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertInfo: %w", err)
	}

	return nil
}

func (d DB) GetInfo() (*models.Info, error) {
	infos, err := models.Infos().All(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetInfo: %w", err)
	}

	lastInfo := len(infos) - 1

	return infos[lastInfo], nil
}

func (d DB) AddUserIfNotExists(userID string, userName string) error {
	allUsers, err := d.GetUsers()
	if err != nil {
		return fmt.Errorf("AddUserIfNotExists: %w", err)
	}

	for _, user := range allUsers {
		if user.Telegramid == userID {
			return nil
		}
	}

	err = d.InsertUser(userID, userName)
	if err != nil {
		return fmt.Errorf("AddUserIfNotExists: %w", err)
	}

	return nil
}

func (d DB) GetUsers() (models.UserSlice, error) {
	users, err := models.Users().All(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetUsers: %w", err)
	}

	return users, nil
}

func (d DB) InsertUser(userID, userName string) error {
	var user models.User

	user.Telegramid = userID
	user.Username = userName

	err := user.Insert(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertUser: %w", err)
	}

	return nil
}

func (d DB) GetUserInTable(userID string) (int, error) {
	user, err := models.Users(qm.Where("telegramid=?", userID)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("GetUserInTable: %w", err)
	}

	return user.ID, nil
}

func (d DB) AddLocationIfNotExists(locationName, userID string) error {
	allUserLocations, err := d.GetUserLocations(userID)
	if err != nil {
		return fmt.Errorf("AddLocationIfNotExists: %w", err)
	}

	for _, location := range allUserLocations {
		if location.Name == locationName {
			return nil
		}
	}

	err = d.InsertLocation(locationName, userID)
	if err != nil {
		return fmt.Errorf("AddLocationIfNotExists: %w", err)
	}

	return nil
}

func (d DB) GetUserLocations(userID string) (models.LocationSlice, error) {
	userInTable, err := d.GetUserInTable(userID)
	if err != nil {
		return nil, fmt.Errorf("GetUserLocations: %w", err)
	}

	locations, err := models.Locations(qm.Where("user_id=?", userInTable)).All(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetUserLocations: %w", err)
	}

	return locations, nil
}

func (d DB) GetLocationByName(locationName, userID string) (*models.Location, error) {
	userInTable, err := d.GetUserInTable(userID)
	if err != nil {
		return nil, fmt.Errorf("GetUserLocations: %w", err)
	}

	locations, err := models.Locations(qm.Where("user_id=? and name=?", userInTable, locationName)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetUserLocations: %w", err)
	}

	return locations, nil
}

func (d DB) InsertLocation(locationName, userID string) error {
	userInTable, err := d.GetUserInTable(userID)
	if err != nil {
		return fmt.Errorf("InsertLocation: %w", err)
	}

	var location models.Location

	location.Name = locationName
	location.UserID = userInTable

	err = location.Insert(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertLocation: %w", err)
	}

	return nil
}

func (d DB) GetLocationInTable(locationName, userID string) (int, error) {
	userInTable, err := d.GetUserInTable(userID)
	if err != nil {
		return 0, fmt.Errorf("GetLocationInTable: %w", err)
	}

	location, err := models.Locations(qm.Where("user_id=? and name=?", userInTable, locationName)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("GetLocationInTable: %w", err)
	}

	return location.ID, nil
}

func (d DB) AddAlbumToCollection(album entities.Album, locationName, userID string) error {
	locationInTable, err := d.GetLocationInTable(locationName, userID)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	err = d.AddArtistIfNotExists(album.Artist.Name)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	err = d.InsertAlbum(album)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	albumInTable, err := d.GetAlbumInTable(album)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	err = d.InsertToCollection(albumInTable, locationInTable)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	return nil
}

func (d DB) AddArtistIfNotExists(artistName string) error {
	allArtists, err := d.GetArtists()
	if err != nil {
		return fmt.Errorf("AddArtistIfNotExists: %w", err)
	}

	for _, artist := range allArtists {
		if artist.Name == artistName {
			return nil
		}
	}

	err = d.InsertArtist(artistName)
	if err != nil {
		return fmt.Errorf("AddArtistIfNotExists: %w", err)
	}

	return nil
}

func (d DB) GetArtists() (models.ArtistSlice, error) {
	artists, err := models.Artists().All(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetArtists: %w", err)
	}

	return artists, nil
}

func (d DB) InsertArtist(artistName string) error {
	var artist models.Artist

	artist.Name = artistName

	err := artist.Insert(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertUser: %w", err)
	}

	return nil
}

func (d DB) GetArtistInTable(artistName string) (int, error) {
	artist, err := models.Artists(qm.Where("name=?", artistName)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("GetUserInTable: %w", err)
	}

	return artist.ID, nil
}

func (d DB) GetArtistByID(artistID int) (*models.Artist, error) {
	artist, err := models.Artists(qm.Where("id=?", artistID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetUserInTable: %w", err)
	}

	return artist, nil
}

func (d DB) InsertAlbum(albumData entities.Album) error {
	var album models.Album

	artistInTable, err := d.GetArtistInTable(albumData.Artist.Name)
	if err != nil {
		return fmt.Errorf("InsertAlbum: %w", err)
	}

	album.ArtistID = artistInTable
	album.AlbumName = albumData.Name
	album.Genre = albumData.Genre
	album.ReleaseYear = albumData.ReleaseYear
	album.ReissueYear = albumData.ReissueYear
	album.Label = albumData.Label
	album.Coloured = albumData.Coloured
	album.CoverID = albumData.CoverID

	err = album.Insert(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertAlbum: %w", err)
	}

	return nil
}

func (d DB) GetAlbumInTable(albumData entities.Album) (int, error) {
	album, err := models.Albums(qm.Where(
		"album_name=? and genre=? and release_year=? and reissue_year=? and label=?",
		albumData.Name,
		albumData.Genre,
		albumData.ReleaseYear,
		albumData.ReissueYear,
		albumData.Label,
	)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("GetUserInTable: %w", err)
	}

	return album.ID, nil
}

func (d DB) InsertToCollection(albumInTable, locationInTable int) error {
	var collectionElement models.Collection

	collectionElement.AlbumID = albumInTable
	collectionElement.LocationID = locationInTable

	err := collectionElement.Insert(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertToCollection: %w", err)
	}

	return nil
}

func (d DB) GetAlbumsByLocation(locationInTable int) (models.AlbumSlice, error) {
	var albums models.AlbumSlice

	albumIDList, err := models.Collections(qm.Where("location_id=?", locationInTable)).All(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetAlbumsByLocation: %w", err)
	}

	for _, el := range albumIDList {
		album, err := models.Albums(qm.Where("id=?", el.AlbumID)).One(d.db)
		if err != nil {
			return nil, fmt.Errorf("GetAlbumsByLocation: %w", err)
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (d DB) GetCollection(userID string) (models.AlbumSlice, error) {
	var albums models.AlbumSlice

	locations, err := d.GetUserLocations(userID)
	if err != nil {
		return nil, fmt.Errorf("GetCollection: %w", err)
	}

	for _, location := range locations {
		albumsInLoc, err := d.GetAlbumsByLocation(location.ID)
		if err != nil {
			return nil, fmt.Errorf("GetCollection: %w", err)
		}

		albums = append(albums, albumsInLoc...)
	}

	return albums, nil
}

func (d DB) DeleteAlbum(userID string, albumID int) error {
	var albumIDList models.CollectionSlice

	locations, err := d.GetUserLocations(userID)
	if err != nil {
		return fmt.Errorf("DeleteAlbum: %w", err)
	}

	for _, location := range locations {
		albumInLocList, err := models.Collections(qm.Where("location_id=?", location.ID)).All(d.db)
		if err != nil {
			return fmt.Errorf("DeleteAlbum: %w", err)
		}

		albumIDList = append(albumIDList, albumInLocList...)
	}

	_, err = albumIDList[albumID].Delete(d.db)
	if err != nil {
		return fmt.Errorf("DeleteAlbum: %w", err)
	}

	return nil
}
