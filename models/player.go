package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	Id         int       `json:"id"`
	ActivityId int       `json:"activity_id"`
	PlayerId   int       `json:"player_id"`
	PlayerName string    `json:"player_name"`
	Avatar     string    `json:"avatar"`
	Score      int       `json:"score"`
	Desc       string    `json:"desc"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Player) TableName() string {
	return "player"
}

func GetPlayerList(aid int) ([]Player, error) {

	var players []Player

	err := DB.Where("activity_id = ?", aid).Find(&players).Error
	if err != nil {
		return nil, err
	}
	return players, nil
}

func GetPlayerByIDActivityID(player_id int, activity_id int) (Player, error) {

	var player Player

	err := DB.Where("player_id = ? AND activity_id = ?", player_id, activity_id).First(&player).Error
	if err != nil {
		return player, err
	}
	return player, nil
}

func UpdatePlayerScore(player_id, activity_id int) (Player, error) {
	var player Player
	err := DB.Model(&player).Where("player_id = ?1 AND activity_id = ?2", player_id, activity_id).Update("score", gorm.Expr("score + ?", 1)).Error
	return player, err
}
