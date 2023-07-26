package sharedServices

import (
	"strings"

	"api.virak.me/src/shared/constants"
)

func ApiPrefix(route string) string {
	return strings.Join([]string{constants.ApiPrefix, constants.ApiVersion, route}, "/")
}
