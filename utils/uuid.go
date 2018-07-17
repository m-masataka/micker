package utils

import (
	"github.com/rs/xid"
)

func CreateUUID() string {
	guid := xid.New()
	return guid.String()
}
