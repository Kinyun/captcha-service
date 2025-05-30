package request

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
)

var reqid uint64

// RequestID is a middleware that injects a request ID into the context of each
// request. A request ID is a string of the form "host.example.com/random-0001",
// where "random" is a base62 random string that uniquely identifies this go
// process, and where the last number is an atomically incremented request
// counter.
//
// Ported from Goji's middleware, source:
// https://github.com/zenazn/goji/tree/master/web/middlewar
func ID() string {
	var sb strings.Builder

	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		_, err := rand.Read(buf[:])
		if err != nil {
			return ""
		}
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}

	sb.WriteString(hostname)
	sb.WriteString("-")
	sb.WriteString(b64[0:10])

	return fmt.Sprintf("%s-%08d", sb.String(), atomic.AddUint64(&reqid, 1))
}
