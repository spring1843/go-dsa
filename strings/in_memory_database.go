package strings

import (
	"maps"
	"strings"
)

var dbs = []map[string]string{}

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
	dbs[len(dbs)-1][key] = value
}

func get(key string) string {
	v, ok := dbs[len(dbs)-1][key]
	if !ok {
		return "nil"
	}
	return v
}

func exists(key string) string {
	if _, ok := dbs[len(dbs)-1][key]; !ok {
		return "false"
	}
	return "true"
}

func unset(key string) {
	delete(dbs[len(dbs)-1], key)
}

func begin() {
	dbs = append(dbs, maps.Clone(dbs[len(dbs)-1]))
}

func rollback() {
	dbs = dbs[:len(dbs)-1]
}

func commit() {
	dbs[len(dbs)-2] = dbs[len(dbs)-1]
	rollback()
}
