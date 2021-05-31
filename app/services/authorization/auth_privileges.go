package authorization

import "github.com/jtorz/phoenix-backend/utils/stringset"

type privilege struct {
	Method string
	Route  string
	Key    string
}

type privileges []privilege

func (privs privileges) getPrivilegeByPriority(privileges ...string) string {
	pos := -1
	for i := range privs {
		if pos2, found := stringset.FindInSlice(privileges, privs[i].Key); found {
			if pos2 == 0 {
				return privileges[pos2]
			}
			if pos2 > pos {
				pos = pos2
			}
		}
	}
	if pos == -1 {
		return ""
	}
	return privileges[pos]
}
