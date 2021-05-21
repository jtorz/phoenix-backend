package stringset_test

import (
	"testing"

	"github.com/jtorz/phoenix-backend/utils/stringset"
)

func TestTitleCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		In  string
		Out string
	}{
		{"hello_there", "HelloThere"},
		{"", ""},
		{"____a____a___", "AA"},
		{"_a_a_", "AA"},
		{"fun_id", "FunID"},
		{"_fun_id", "FunID"},
		{"__fun____id_", "FunID"},
		{"uid", "UID"},
		{"guid", "GUID"},
		{"uid", "UID"},
		{"uuid", "UUID"},
		{"ssn", "Ssn"},
		{"tz", "Tz"},
		{"thing_guid", "ThingGUID"},
		{"guid_thing", "GUIDThing"},
		{"thing_guid_thing", "ThingGUIDThing"},
		{"id", "ID"},
		{"gvzxc", "Gvzxc"},
		{"id_trgb_id", "IDTrgbID"},
		{"vzxx_vxccb_nmx", "VzxxVxccbNmx"},
		{"thing_zxc_stuff_vxz", "ThingZxcStuffVxz"},
		{"zxc_thing_vxz_stuff", "ZxcThingVxzStuff"},
		{"zxc_vdf9c9_hello9", "ZxcVdf9c9Hello9"},
		{"id9_uid911_guid9e9", "ID9UID911GUID9E9"},
		{"zxc_vdf0c0_hello0", "ZxcVdf0c0Hello0"},
		{"id0_uid000_guid0e0", "ID0UID000GUID0E0"},
		{"ab_5zxc5d5", "Ab5zxc5d5"},
		{"Identifier", "Identifier"},
	}

	for i, test := range tests {
		if out := stringset.SnakeToGoCase(test.In); out != test.Out {
			t.Errorf("[%d] (%s) Out was wrong: %q, want: %q", i, test.In, out, test.Out)
		}
	}
}
