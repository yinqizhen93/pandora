// Code generated by entc, DO NOT EDIT.

package stock

const (
	// Label holds the string label denoting the stock type in the database.
	Label = "stock"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMarket holds the string denoting the market field in the database.
	FieldMarket = "market"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldOpen holds the string denoting the open field in the database.
	FieldOpen = "open"
	// FieldClose holds the string denoting the close field in the database.
	FieldClose = "close"
	// FieldHigh holds the string denoting the high field in the database.
	FieldHigh = "high"
	// FieldLow holds the string denoting the low field in the database.
	FieldLow = "low"
	// FieldVolume holds the string denoting the volume field in the database.
	FieldVolume = "volume"
	// FieldOutstandingShare holds the string denoting the outstandingshare field in the database.
	FieldOutstandingShare = "outstanding_share"
	// FieldTurnover holds the string denoting the turnover field in the database.
	FieldTurnover = "turnover"
	// Table holds the table name of the stock in the database.
	Table = "stocks"
)

// Columns holds all SQL columns for stock fields.
var Columns = []string{
	FieldID,
	FieldMarket,
	FieldCode,
	FieldName,
	FieldDate,
	FieldOpen,
	FieldClose,
	FieldHigh,
	FieldLow,
	FieldVolume,
	FieldOutstandingShare,
	FieldTurnover,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}