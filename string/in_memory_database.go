package string

import (
	"strings"
)

var dbs []map[string]string

// RunDBCommand executes commands on an in memory database that stores string key-value pairs.
// Supported commands are SET, GET, EXISTS, and UNSET
// It also allows transactions with BEGIN, ROLLBACK and COMMIT commands.
func RunDBCommand(cmd string) string {
	splitCmd := strings.Split(cmd, " ")
	switch splitCmd[0] {
	case `SET`:
		set(splitCmd[1], splitCmd[2])
	case `GET`:
		return get(splitCmd[1])
	case `EXISTS`:
		return exists(splitCmd[1])
	case `UNSET`:
		unset(splitCmd[1])
	case `BEGIN`:
		begin()
	case `ROLLBACK`:
		rollback()
	case `COMMIT`:
		commit()
	}
	return ""
}

func set(key, value string) {
	db := dbs[len(dbs)-1]
	db[key] = value
}

func get(key string) string {
	db := dbs[len(dbs)-1]
	v, ok := db[key]
	if !ok {
		return "nil"
	}
	return v
}

func exists(key string) string {
	db := dbs[len(dbs)-1]
	if _, ok := db[key]; !ok {
		return "false"
	}
	return "true"
}

func unset(key string) {
	db := dbs[len(dbs)-1]
	delete(db, key)
}

func begin() {
	clonedDB := make(map[string]string)
	dbs = append(dbs, clonedDB)
}

func rollback() {
	dbs = dbs[:len(dbs)-1]
}

func commit() {
	dbs[len(dbs)-2] = dbs[len(dbs)-1]
	rollback()
}
