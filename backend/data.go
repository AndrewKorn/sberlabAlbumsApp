package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBEngine struct {
	DB *gorm.DB
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
	Tz       string
}

func (dbe *DBEngine) initTables() error {
	err := dbe.DB.AutoMigrate(&Group{}, &Album{}, &Song{})
	if err != nil {
		return err
	}
	return nil
}

func NewDBEngine(dbc DBConfig) (*DBEngine, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbc.Host,
		dbc.User,
		dbc.Password,
		dbc.Name,
		dbc.Port,
		dbc.SSLMode,
		dbc.Tz)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbe := &DBEngine{}
	dbe.DB = db

	if err := dbe.initTables(); err != nil {
		return nil, err
	}

	return dbe, nil
}

func (dbe *DBEngine) createGroup(group *Group) *Group {
	dbe.DB.Create(group)
	return group
}

func (dbe *DBEngine) getGroups() []Group {
	var groups []Group
	dbe.DB.Find(&groups)
	return groups
}

func (dbe *DBEngine) getGroup(groupName string) *Group {
	group := &Group{}
	dbe.DB.Where(Group{GroupName: groupName}).First(&group)
	return group
}

func (dbe *DBEngine) getGroupsAlbums(groupName string) []Album {
	group := dbe.getGroup(groupName)

	var albums []Album
	err := dbe.DB.Model(&group).Association("Albums").Find(&albums)
	if err != nil {
		return nil
	}
	return albums
}

func (dbe *DBEngine) createAlbum(groupName, albumName string) *Album {
	group := dbe.getGroup(groupName)
	album := &Album{AlbumName: albumName, Songs: []Song{}, GroupID: group.ID}
	err := dbe.DB.Model(&group).Association("Albums").Append(album)
	if err != nil {
		return nil
	}
	return album
}

func (dbe *DBEngine) deleteAlbum(groupName, albumName string) *Album {
	group := dbe.getGroup(groupName)

	album := &Album{}
	dbe.DB.Where("group_id = ? AND album_name = ?", group.ID, albumName).Delete(&album)
	return album
}

func (dbe *DBEngine) getAlbum(groupName, albumName string) *Album {
	group := dbe.getGroup(groupName)

	var albums []Album
	err := dbe.DB.Model(&group).Where("album_name = ?", albumName).Association("Albums").Find(&albums)
	if err != nil {
		return nil
	}
	return &albums[0]
}

func (dbe *DBEngine) getAlbumSongs(groupName, albumName string) []Song {
	album := dbe.getAlbum(groupName, albumName)

	var songs []Song
	err := dbe.DB.Model(&album).Association("Songs").Find(&songs)
	if err != nil {
		return nil
	}
	return songs
}

func (dbe *DBEngine) createSong(groupName, albumName string, song *Song) *Song {
	album := dbe.getAlbum(groupName, albumName)
	err := dbe.DB.Model(&album).Association("Songs").Append(song)
	if err != nil {
		return nil
	}
	return song
}

func (dbe *DBEngine) deleteSong(groupName, albumName string, id uint) *Song {
	album := dbe.getAlbum(groupName, albumName)

	song := &Song{}
	dbe.DB.Where("id = ? AND album_id = ?", id, album.ID).Delete(&song)
	return song
}
