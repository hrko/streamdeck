package streamdeck_test

import (
	"testing"

	"github.com/hrko/streamdeck"
)

func TestParseRegistrationParams(t *testing.T) {
	// $(GO) run $(GOFLAGS) main.go -port 12345 -pluginUUID 213 -registerEvent test -info "{\"application\":{\"language\":\"en\",\"platform\":\"mac\",\"version\":\"4.1.0\"},\"plugin\":{\"version\":\"1.1\"},\"devicePixelRatio\":2,\"devices\":[{\"id\":\"55F16B35884A859CCE4FFA1FC8D3DE5B\",\"name\":\"Device Name\",\"size\":{\"columns\":5,\"rows\":3},\"type\":0},{\"id\":\"B8F04425B95855CF417199BCB97CD2BB\",\"name\":\"Another Device\",\"size\":{\"columns\":3,\"rows\":2},\"type\":1}]}"
	p, err := streamdeck.ParseRegistrationParams([]string{
		"",
		"-port", "12345",
		"-pluginUUID", "213",
		"-registerEvent", "test",
		"-info", `{"application":{"language":"en","platform":"mac","version":"4.1.0"},"plugin":{"version":"1.1"},"devicePixelRatio":2,"devices":[{"id":"55F16B35884A859CCE4FFA1FC8D3DE5B","name":"Device Name","size":{"columns":5,"rows":3},"type":0},{"id":"B8F04425B95855CF417199BCB97CD2BB","name":"Another Device","size":{"columns":3,"rows":2},"type":1}]}`,
	})
	if err != nil {
		t.Fatalf("Failed to parse params:%v", err)
	}
	t.Logf("Parsed:%v\n", p)
}
