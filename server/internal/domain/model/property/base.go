package property

import "github.com/google/uuid"

type Base struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Location    string `json:"location"`
	Price       int    `json:"price"`
	Photo       string `json:"photo"`
}

type Property struct {
	ID    uuid.UUID `json:"id"`
	Owner uuid.UUID `json:"owner"`
	Base
}

type Update struct {
	Owner       *uuid.UUID `json:"owner"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Type        *string    `json:"type"`
	Location    *string    `json:"location"`
	Price       *int       `json:"price"`
	Photo       *string    `json:"photo"`
}

type Create struct {
	Base
	Email string `json:"email"`
}

type ListParams struct {
	Skip    *int    `query:"skip"`
	Limit   *int    `query:"limit"`
	Order   *int    `query:"order"`
	Keyword *string `query:"keyword"`
	SortKey *string `query:"sort_keyword"`
	Type    *string `query:"type"`
}
