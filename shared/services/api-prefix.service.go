package shared_services

import (
	"strings"

	"api.virak.me/shared/constants"
)

func ApiPrefix(route string) string {
	return strings.Join([]string{constants.ApiPrefix, constants.ApiVersion, route}, "/")
}
