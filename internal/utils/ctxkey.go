package utils

import "net/http"

// CtxKey is a custom context key type.
//
// Path parameters are handled by adding the matches slice
// to the request context, so the handlers can pick them up from there.
type CtxKey struct{}

// GetField gets a slug or id from route.
//
// Examples:
//
//	slug := utils.GetField(r, 0)
//	id, _ := strconv.Atoi(utils.GetField(r, 1))
func GetField(r *http.Request, index int) string {
	fields := r.Context().Value(CtxKey{}).([]string)
	return fields[index]
}
