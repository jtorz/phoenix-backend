package authorization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrivilegeByPriority(t *testing.T) {
	var actual string
	privs := privileges{
		{Key: "QUERY_ALL"},
		{Key: "QUERY_ACTIVE"},
		{Key: "QUERY_SPECIFIC"},
		{Key: "OTHER1"},
		{Key: "OTHER2"},
	}

	actual = privs.getPrivilegeByPriority("QUERY_ALL", "QUERY_ACTIVE", "QUERY_SPECIFIC")
	assert.Equal(t, "QUERY_ALL", actual)

	actual = privs.getPrivilegeByPriority("OTHER3", "OTHER2", "OTHER1")
	assert.Equal(t, "OTHER2", actual)

	actual = privs.getPrivilegeByPriority("FOO3", "FOO2", "FOO1")
	assert.Equal(t, "", actual)

	actual = privs.getPrivilegeByPriority("")
	assert.Equal(t, "", actual)

	actual = privs.getPrivilegeByPriority()
	assert.Equal(t, "", actual)
}
