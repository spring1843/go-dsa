package strings

import (
	"strings"
)

var dbs []map[string]string

// RunDBCommand solves the problem in O(1) time and O(n) space.
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
