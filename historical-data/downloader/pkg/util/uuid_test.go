package util

import "testing"

func TestGenerateUUIDV4(t *testing.T) {
	uid, err := GenerateUUIDV4()
	if err != nil{
		t.Fatal(err)
	}

	t.Logf("GenerateUUIDV4 ... string: %s, length: %d.", uid, len(uid))
}
