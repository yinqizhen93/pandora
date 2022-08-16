package etl

import df "github.com/go-gota/gota/dataframe"

type Extractor interface {
	extract() (df.DataFrame, error)
}
