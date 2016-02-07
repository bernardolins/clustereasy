package init

import (
	"testing"
)

const data = `
cluster:
  token: test
  state: existed

  nodes:
  - name: test1
  interface: eth0
  ip: 1.1.1.1
`

func TestApply(t *testing.T) {
	clusterData := Apply([]byte(data))

	expect := "test"
	data := clusterData.GetClusterInfo()
	got := data.GetToken()

	if expect != got {
		t.Fatalf("expect %s, got %s", expect, got)
	}

	expect = "existed"
	got = data.GetState()

	if expect != got {
		t.Fatalf("expect %s, got %s", expect, got)
	}

	expect = "test1"
	nodes := data.GetNodes()

	got = nodes[0].NodeName()

	if expect != got {
		t.Fatalf("expect %s, got %s", expect, got)
	}
}
