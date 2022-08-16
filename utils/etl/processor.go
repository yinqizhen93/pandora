package etl

import df "github.com/go-gota/gota/dataframe"

type Processor interface {
	process(df.DataFrame) df.DataFrame
}
