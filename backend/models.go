package main

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey" json:"id"`
	GroupCode string  `json:"group_code"`
	GroupName string  `json:"group_name"`
	Albums    []Album `gorm:"foreignKey:GroupID" json:"albums"`
}

type Album struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	AlbumName string `json:"album_name"`
	Songs     []Song `gorm:"foreignKey:AlbumID" json:"songs"`
	GroupID   uint   `json:"groupID"`
}

type Song struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey" json:"id"`
	Artist       string `json:"artist"`
	SoundCloudID string `json:"sound_cloud_id"`
	SongName     string `json:"song_name"`
	Duration     string `json:"duration"`
	AlbumID      uint   `json:"albumID"`
}
