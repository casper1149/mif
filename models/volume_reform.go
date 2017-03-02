package models

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type volumeTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *volumeTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("volumes").
func (v *volumeTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *volumeTableType) Columns() []string {
	return []string{"id", "book_id", "type", "created_at", "updated_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *volumeTableType) NewStruct() reform.Struct {
	return new(Volume)
}

// NewRecord makes a new record for that table.
func (v *volumeTableType) NewRecord() reform.Record {
	return new(Volume)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *volumeTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// VolumeTable represents volumes view or table in SQL database.
var VolumeTable = &volumeTableType{
	s: parse.StructInfo{Type: "Volume", SQLSchema: "", SQLName: "volumes", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "BookID", PKType: "", Column: "book_id"}, {Name: "Type", PKType: "", Column: "type"}, {Name: "CreatedAt", PKType: "", Column: "created_at"}, {Name: "UpdatedAt", PKType: "", Column: "updated_at"}}, PKFieldIndex: 0},
	z: new(Volume).Values(),
}

// String returns a string representation of this struct or record.
func (s Volume) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "BookID: " + reform.Inspect(s.BookID, true)
	res[2] = "Type: " + reform.Inspect(s.Type, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Volume) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.BookID,
		s.Type,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Volume) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.BookID,
		&s.Type,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Volume) View() reform.View {
	return VolumeTable
}

// Table returns Table object for that record.
func (s *Volume) Table() reform.Table {
	return VolumeTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Volume) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Volume) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Volume) HasPK() bool {
	return s.ID != VolumeTable.z[VolumeTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Volume) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = VolumeTable
	_ reform.Struct = new(Volume)
	_ reform.Table  = VolumeTable
	_ reform.Record = new(Volume)
	_ fmt.Stringer  = new(Volume)
)

func init() {
	parse.AssertUpToDate(&VolumeTable.s, new(Volume))
}
