package utils

import (
	"github.com/rs/xid"
)

func CreateContainerID() string {
	guid := xid.New()
	return guid.String()
}
