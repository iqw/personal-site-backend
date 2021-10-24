package models

import (
	"gorm.io/gorm"
	"time"
)

type Level int

const (
	Beginner Level = iota
	Medium
	High
	Pro
)

type Skill struct {
	gorm.Model
	ID        uint
	Name      string
	Level     Level
	CreatedAt time.Time
	UpdatedAt int
}
