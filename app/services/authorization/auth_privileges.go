package authorization

type privilege struct {
	Method string
	Route  string
	Key    string
}

type privileges []privilege

func (privs privileges) getPrivilegeByPriority(keyPriority ...string) string {
	if len(keyPriority) == 0 {
		return ""
	}

	found := make([]bool, len(keyPriority))
	for i := range privs {
		for j := range keyPriority {
			if privs[i].Key == keyPriority[j] {
				if j == 0 { // highest priority
					return keyPriority[j]
				}
				found[j] = true
			}
		}
	}
	for i := range found {
		if found[i] {
			return keyPriority[i]
		}
	}
	return ""
}
