package models

import "time"

//user_id      BIGSERIAL PRIMARY KEY,
//author       varchar(100),
//title        VARCHAR(100),
//text         VARCHAR(2000),
//image        varchar(100),
//date_created timestamp not null default now()

type Posts struct {
	ID            uint64    `json:"id"`
	UserID        uint64    `json:"userID"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Text          string    `json:"text"`
	Image         string    `json:"image"`
	DateCreatedAt time.Time `json:"dateCreatedAt"`
}

type PostDTO struct {
	ID            uint64    `json:"id"`
	UserID        uint64    `json:"userID"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Text          string    `json:"text"`
	Image         string    `json:"image"`
	DateCreatedAt time.Time `json:"dateCreatedAt"`
}
