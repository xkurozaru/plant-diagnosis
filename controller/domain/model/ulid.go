package model

import (
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

type ULID string

func NewULID() ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return ULID(strings.ToLower(id.String()))
}

func (u ULID) ToString() string {
	return string(u)
}

func (u ULID) IsEmpty() bool {
	return u == ""
}
