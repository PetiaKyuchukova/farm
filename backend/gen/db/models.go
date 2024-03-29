// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"database/sql"
	"time"
)

type Cow struct {
	ID          string         `json:"id"`
	Birthdate   sql.NullTime   `json:"birthdate"`
	Gender      sql.NullString `json:"gender"`
	Breed       sql.NullString `json:"breed"`
	Colour      sql.NullString `json:"colour"`
	Motherid    sql.NullString `json:"motherid"`
	Motherbreed sql.NullString `json:"motherbreed"`
	Fatherid    sql.NullString `json:"fatherid"`
	Fatherbreed sql.NullString `json:"fatherbreed"`
	Ispregnant  sql.NullBool   `json:"ispregnant"`
	Ovulation   sql.NullTime   `json:"ovulation"`
}

type Insemination struct {
	Cowid        string         `json:"cowid"`
	Date         time.Time      `json:"date"`
	Breed        sql.NullString `json:"breed"`
	Isartificial sql.NullBool   `json:"isartificial"`
}

type Milk struct {
	Date   time.Time       `json:"date"`
	Liters sql.NullFloat64 `json:"liters"`
	Price  sql.NullFloat64 `json:"price"`
}

type Pregnancy struct {
	Cowid      string       `json:"cowid"`
	Detectedat time.Time    `json:"detectedat"`
	Firstday   sql.NullTime `json:"firstday"`
	Lastday    sql.NullTime `json:"lastday"`
}

type Task struct {
	Cowid string       `json:"cowid"`
	Date  time.Time    `json:"date"`
	Type  string       `json:"type"`
	Text  string       `json:"text"`
	Done  sql.NullBool `json:"done"`
}
