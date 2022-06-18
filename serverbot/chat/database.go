package chat

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB Database

type Database interface {
	GetState() bool
	GetById(interface{}, int) error
	GetByPk(a, b interface{}, c string) error
	GetList(interface{}, *[]interface{}, int) error
	InsertRow(interface{}) error
	UpdateRow(interface{}) error
}

type SQLdatabase struct {
	db        *sql.DB
	connected bool
	name      string
}

func NewDatabase(sqldb *sql.DB, name string) Database {
	return &SQLdatabase{
		db:        sqldb,
		connected: true,
		name:      name,
	}
}

func (d SQLdatabase) GetState() bool {
	return d.connected
}

// Performs a connection to the database passed as "name" parameter and it must
// be referenced on the  .ENV file followed by the password in order to connect
// otherwise a panic occurs. There is no need to defer Disconnect.
func ConnectSQL(name, password, database string) (Database, error) {
	dbase, err := sql.Open("mysql", name+":"+password+"@tcp(127.0.0.1:3306)/"+database)
	if err != nil {
		panic("Unable to connect to db")
	}
	db := NewDatabase(dbase, name)
	DB = db
	return db, nil
}

// Useful to produce a slice of interface{} values, from a slice of string vals
// particularly  to pass as variadic parameters to Query , Exec or similar func
func interfaceSlice(strlst []string) []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(strlst))
	for i, d := range strlst {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}

// Sets struct fields' values, given the mysql type , field name and arg to set
// the field pointer ptr,the ptr value must be the address of the field pointer
// thus deferencing twice.  Some of the struct's fields must be of pointer type
// otherwise it'll panic. More type conversions should be added at will . There
// is no nil pointer for values not attributed,these are initialized and should
// have the corresponding Zero value for each type
func SetElem(in_type, field string, arg, ptr interface{}) error {
	arg1 := reflect.ValueOf(arg).Elem()
	ptr1 := reflect.Indirect(reflect.Indirect(ptr.(reflect.Value)))

	// Return if arg is nil, nothing to set on ptr value
	// All vars / struct fields should be pointers
	if arg1.Interface() == nil {
		return nil
	}

	switch in_type {
	case "VARCHAR":
		ptr1.SetString(string(arg1.Interface().([]byte)))
	case "TEXT":
		ptr1.SetString(string(arg1.Interface().([]byte)))
	case "INT":
		ptr1.SetInt(int64(arg1.Interface().(int64)))
	case "DATETIME":
		ptr1.SetString(string(arg1.Interface().([]byte)))
	default:
		return errors.New(fmt.Sprintf("Database Type unknown:%s\n", in_type))
	}
	return nil
}

// To extract the values of each struct's field , we must provide the following
// function with the struct pointer, and a reflect.StructField parameter as the
// latter contains information regarding the field's type. There's noway around
// to extract the concrete type of each field, you must use reflect.StructField
// Add types as needed
func structToStuctFieldString(structure interface{}, strField reflect.StructField) string {
	ptr1 := reflect.ValueOf(structure)
	ret := new(string)
	switch strField.Type.String() {
	case "int":
		*ret = strconv.Itoa(ptr1.Elem().FieldByName(strField.Name).Interface().(int))
	case "*string":
		ret = ptr1.Elem().FieldByName(strField.Name).Interface().(*string)
	}
	return *ret
}

// From a tagged "db" struct Produces a tuple with the 1st being a string slice
// of the values from key tag "db", the 2nd a string slice of the corresponding
// field names of said struct,  and the 3rd a reflect.StructField slice contain
// each the params pertaining to the field itself in struct type. strucutre can
// and should be either a dereferenced pointer or struct value.  "ignore" param
// contains a list of strings with db column names to ignore. It can be nil
func structToSlices(structure interface{}, ignore []string) ([]string, []string, []reflect.StructField) {
	contains := func(list []string, val string) bool {
		for _, v := range list {
			if v == val {
				return true
			}
		}
		return false
	}
	structPtr := reflect.TypeOf(structure)
	columns := []string{}
	fields := []string{}
	values := make([]reflect.StructField, 0)
	for i := 0; i < structPtr.Elem().NumField(); i++ {
		col := structPtr.Elem().Field(i).Tag.Get("db")
		if col != "" && !contains(ignore, structPtr.Elem().Field(i).Name) {
			fields = append(fields, structPtr.Elem().Field(i).Name)
			values = append(values, structPtr.Elem().Field(i))
			columns = append(columns, "`"+col+"`")
		}
	}
	return columns, fields, values
}

// Given a slice of type reflect.StructField, and tag key, as well as a tag val
// it should return the corresponding  relect.StructField to such tag's key and
// value string pair. The common use would be to pass the params tag_key = "db"
// and tag_val = "id" ,  it should return the corresponding reflect.StructField
// on sFieldSlice
func structFieldFromTag(sFieldSlice []reflect.StructField, tag_key, tag_val string) reflect.StructField {
	var alias reflect.StructField
	for i := 0; i < len(sFieldSlice); i++ {
		field := sFieldSlice[i]
		if field.Tag.Get(tag_key) == tag_val {
			alias = field
		} else {
			continue
		}
	}
	return alias
}

func (d *SQLdatabase) GetById(structure interface{}, id int) error {
	var inderface interface{}
	inderface = id
	d.GetByPk(structure, inderface, "id")
	return nil
}

// structure parameter must be an address pointing to a struct type val and its
// fields should be pointers,otherwise it will throw an error. It applies where
// pointers are needed, excluding for example: int, float,etc. Make sure you'll
// set the right types beforehand . This will later fetch the field by tag, and
// not "Id" named field
func (d *SQLdatabase) GetByPk(structure, pk interface{}, field string) error {
	id := reflect.ValueOf(pk)
	var strid string
	if id.Type() == reflect.TypeOf(0) {
		strid = strconv.Itoa(id.Interface().(int))
	} else if id.Type() == reflect.TypeOf("") {
		strid = id.Interface().(string)

	} else {
		log.Panic("pk value must be string or int")
	}
	structPtr := reflect.ValueOf(structure)
	struct_name := structPtr.Type().Elem().Name()

	if structPtr.Type().Kind() != reflect.Ptr {
		return errors.New("You must Dereference Struct")
	}

	columns, fields, _ := structToSlices(structure, nil)

	row, err := d.db.Query("SELECT "+strings.Join(columns[:], ", ")+" FROM "+struct_name+" WHERE "+field+"= ?", strid)
	defer row.Close()
	if err != nil || err == sql.ErrNoRows {
		panic(err.Error())
	}

	colTypes, err := row.ColumnTypes()
	values := make([]interface{}, len(columns))
	scan_args := make([]interface{}, len(columns))
	for i := range values {
		scan_args[i] = &values[i]
	}
	if row.Next() {
		err = row.Scan(scan_args...)
		if err != nil {
			panic(err.Error())
		}
	} else {
		return fmt.Errorf("%s object with Id %d does not exist", struct_name, id)
	}
	for i, arg := range scan_args {
		err := SetElem(colTypes[i].DatabaseTypeName(), fields[i], arg, structPtr.Elem().FieldByName(fields[i]).Addr())
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func NewReflectPtr(structure interface{}) (error, reflect.Value) {
	structType := reflect.New(reflect.TypeOf(structure).Elem())
	switch structType.Type().String() {
	case "*main.Post":
		structType.Elem().Set(reflect.Indirect(reflect.ValueOf(NewSession(nil, nil))))
	case "*main.Message":
		structType.Elem().Set(reflect.Indirect(reflect.ValueOf(NewMessage(nil, nil))))
		//reflect.ValueOf(&structType).Elem().Set(reflect.ValueOf(new(Message)).Elem())
	default:
		return errors.New("GetList undefined Type"), reflect.ValueOf(nil)
	}
	return nil, structType
}

// structure parameter must be an address pointing to a struct type val and its
// fields should be pointers,otherwise it will throw an error. It applies where
// pointers are needed, excluding for example: int, float,etc. Make sure you'll
// set the right types beforehand . This will later fetch the field by tag, and
// not "Id" named field
func (d *SQLdatabase) GetList(structure interface{}, list *[]interface{}, id int) error {
	structPtr := reflect.ValueOf(structure)
	struct_name := structPtr.Type().Elem().Name()
	if structPtr.Type().Kind() != reflect.Ptr {
		return errors.New("You must Dereference Struct")
	}

	columns, fields, _ := structToSlices(structure, nil)
	row, err := d.db.Query("SELECT "+strings.Join(columns[:], ", ")+" FROM "+struct_name+" WHERE ID > ? ORDER BY ID LIMIT 5", strconv.Itoa(id))
	defer row.Close()
	if err != nil || err == sql.ErrNoRows {
		panic(err.Error())
	}

	colTypes, err := row.ColumnTypes()
	values := make([]interface{}, len(columns))
	scan_args := make([]interface{}, len(columns))
	for i := range values {
		scan_args[i] = &values[i]
	}

	interfaceSlice := make([]reflect.Value, 0)

	if row.Next() {
		err = row.Scan(scan_args...)
		if err != nil {
			panic(err.Error())
		}
		_, structType := NewReflectPtr(structure)
		interfaceSlice = append(interfaceSlice, structType)
		for i, arg := range scan_args {
			err := SetElem(colTypes[i].DatabaseTypeName(), fields[i], arg, structType.Elem().FieldByName(fields[i]))
			if err != nil {
				panic(err)
			}
		}
		for row.Next() {
			err = row.Scan(scan_args...)
			_, structType2 := NewReflectPtr(structure)
			interfaceSlice = append(interfaceSlice, structType2)
			for i, arg := range scan_args {
				err := SetElem(colTypes[i].DatabaseTypeName(), fields[i], arg, structType2.Elem().FieldByName(fields[i]).Addr())
				if err != nil {
					panic(err)
				}
			}
		}
	} else {
		return errors.New(fmt.Sprintf("%s objects with equal or greater than Id %d do not exist", struct_name, id))
	}
	for _, val := range interfaceSlice {
		*list = append(*list, val.Elem())
	}
	return nil
}

// To save changes made to an existing interface/struct on the database, update
// using the following func. It shouldn't be used on objects that have not been
// created yet. Use Create() first. If said struct's method UpdateRow is called
// without any prior changes made to it, an error will result showing zero rows
// affected.
func (db *SQLdatabase) UpdateRow(structure interface{}) error {
	structPtr := reflect.ValueOf(structure)
	struct_name := structPtr.Type().Elem().Name()
	columns, _, vals := structToSlices(structure, nil)
	vals_str := make([]string, 0)
	if structPtr.Type().Kind() != reflect.Ptr {
		return errors.New("You must Dereference Struct")
	}

	params := func(columns []string) string {
		x := make([]string, 0)
		for i := 0; i < len(columns); i++ {
			x = append(x, columns[i]+"=?")
		}
		return strings.Join(x[:], ", ")
	}

	for i := 0; i < len(columns); i++ {
		vals_str = append(vals_str, structToStuctFieldString(structure, vals[i]))
	}

	values := append(interfaceSlice(vals_str), structPtr.Elem().FieldByName(structFieldFromTag(vals, "db", "id").Name).Interface())
	res, err := db.db.Exec("UPDATE "+struct_name+" SET "+params(columns)+" WHERE id = ?", values...)

	if err != nil {
		panic(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprint("error:", err))
	}
	if rows != 1 {
		return errors.New(fmt.Sprint("expected single row affected, got: ", rows))
	}
	return nil
}

//
//
//
//
//
func (db *SQLdatabase) InsertRow(structure interface{}) error {
	structPtr := reflect.ValueOf(structure)
	struct_name := structPtr.Type().Elem().Name()
	columns, _, vals := structToSlices(structure, []string{"Id"})
	vals_str := make([]string, 0)
	if structPtr.Type().Kind() != reflect.Ptr {
		return errors.New("You must Dereference Struct")
	}

	params := func(columns []string) string {
		x := make([]string, 0)
		for i := 0; i < len(columns); i++ {
			x = append(x, columns[i]+"=?")
		}
		return strings.Join(x[:], ", ")
	}

	for i := 0; i < len(columns); i++ {
		vals_str = append(vals_str, structToStuctFieldString(structure, vals[i]))
	}
	values := interfaceSlice(vals_str)
	log.Println(values)
	res, err := db.db.Exec("INSERT INTO "+struct_name+" SET "+params(columns), values...)

	if err != nil {
		panic(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprint("error:", err))
	}
	if rows != 1 {
		return errors.New(fmt.Sprint("expected single row affected, got: ", rows))
	}
	return nil
}
