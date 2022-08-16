package etl

import df "github.com/go-gota/gota/dataframe"

type Loader interface {
	load(df.DataFrame) error
}
