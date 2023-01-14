package core

import (
	"strings"
	"time"
)

type HolidayType int

const (
	Religious HolidayType = iota
	Civil
	Unknown
)

func (ht HolidayType) String() string {
	switch ht {
	case Religious:
		return "religioso"
	case Civil:
		return "civil"
	}
	return "unknown"
}

func NewHolidayType(hType string) HolidayType {
	switch hType {
	case "religioso":
		return Religious
	case "civil":
		return Civil
	}
	return Unknown
}

type Holiday struct {
	Date        string    `json:"date" xml:"date"`
	Title       string    `json:"title" xml:"title"`
	Type        string    `json:"type" xml:"type"`
	Unalienable bool      `json:"inalienable" xml:"inalienable"`
	Extra       string    `json:"extra" xml:"extra"`
	DateTime    time.Time `json:"-" xml:"-"`
}

func NewHoliday(date, title, hType, extra string, unalienable bool) *Holiday {
	return &Holiday{
		Date:        date,
		Title:       title,
		Type:        hType,
		Unalienable: unalienable,
		Extra:       extra,
	}
}

func (h Holiday) Compare(hh Holiday) bool {
	if h.Extra != "" && strings.EqualFold(h.Extra, hh.Extra) {
		return true
	}

	if h.Title != "" && strings.EqualFold(h.Title, hh.Title) {
		return true
	}

	if h.Type != "" && strings.EqualFold(h.Type, hh.Type) {
		return true
	}

	return false
}
