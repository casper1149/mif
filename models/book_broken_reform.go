// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type bookBrokenTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *bookBrokenTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("books_broken").
func (v *bookBrokenTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *bookBrokenTableType) Columns() []string {
	return []string{"id", "url", "created_at", "updated_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *bookBrokenTableType) NewStruct() reform.Struct {
	return new(BookBroken)
}

// NewRecord makes a new record for that table.
func (v *bookBrokenTableType) NewRecord() reform.Record {
	return new(BookBroken)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *bookBrokenTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// BookBrokenTable represents books_broken view or table in SQL database.
var BookBrokenTable = &bookBrokenTableType{
	s: parse.StructInfo{Type: "BookBroken", SQLSchema: "", SQLName: "books_broken", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "URL", Type: "string", Column: "url"}, {Name: "CreatedAt", Type: "time.Time", Column: "created_at"}, {Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"}}, PKFieldIndex: 0},
	z: new(BookBroken).Values(),
}

// String returns a string representation of this struct or record.
func (s BookBroken) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "URL: " + reform.Inspect(s.URL, true)
	res[2] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[3] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *BookBroken) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.URL,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *BookBroken) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.URL,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *BookBroken) View() reform.View {
	return BookBrokenTable
}

// Table returns Table object for that record.
func (s *BookBroken) Table() reform.Table {
	return BookBrokenTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *BookBroken) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *BookBroken) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *BookBroken) HasPK() bool {
	return s.ID != BookBrokenTable.z[BookBrokenTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *BookBroken) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = BookBrokenTable
	_ reform.Struct = (*BookBroken)(nil)
	_ reform.Table  = BookBrokenTable
	_ reform.Record = (*BookBroken)(nil)
	_ fmt.Stringer  = (*BookBroken)(nil)
)

func init() {
	parse.AssertUpToDate(&BookBrokenTable.s, new(BookBroken))
}
