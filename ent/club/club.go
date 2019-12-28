// Code generated by entc, DO NOT EDIT.

package club

const (
	// Label holds the string label denoting the club type in the database.
	Label = "club"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldRandomID holds the string denoting the random_id vertex property in the database.
	FieldRandomID = "random_id"

	// Table holds the table name of the club in the database.
	Table = "clubs"
	// MusicTable is the table the holds the music relation/edge.
	MusicTable = "musics"
	// MusicInverseTable is the table name for the Music entity.
	// It exists in this package in order to avoid circular dependency with the "music" package.
	MusicInverseTable = "musics"
	// MusicColumn is the table column denoting the music relation/edge.
	MusicColumn = "club_id"
	// DeviceTable is the table the holds the device relation/edge.
	DeviceTable = "devices"
	// DeviceInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DeviceInverseTable = "devices"
	// DeviceColumn is the table column denoting the device relation/edge.
	DeviceColumn = "club_id"
)

// Columns holds all SQL columns are club fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldRandomID,
}