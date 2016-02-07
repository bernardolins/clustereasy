package types

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	cluster := Cluster{}
	expect := "default.tkn"
	got := cluster.GetToken()

	if expect != got {
		t.Fatalf("it should return default token value (%s), but returned %s", expect, got)
	}

	expect = "test"
	cluster.Token = expect

	got = cluster.GetToken()

	if expect != got {
		t.Fatalf("it should return the cluster token (%s), but returned %s", expect, got)
	}
}

func TestGetState(t *testing.T) {
	cluster := Cluster{}
	expect := "new"
	got := cluster.GetState()

	if expect != got {
		t.Fatalf("it should return default state value (%s), but returned %s", expect, got)
	}

	expect = "existed"
	cluster.State = expect

	got = cluster.GetState()

	if expect != got {
		t.Fatalf("it should return the cluster state (%s), but returned %s", expect, got)
	}
}
