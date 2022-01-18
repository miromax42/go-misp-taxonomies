package taxonomies

import "errors"

var (
	ErrWrongFormat = errors.New(`wrong format of taxonomy, use namespace:predicate:"value" or namespace:predicate`)
)
