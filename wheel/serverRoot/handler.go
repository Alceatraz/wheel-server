package serverRoot

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var count int64 = 0

func handler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String())

	hostname, _ := os.Hostname()

	now := time.Now()
	datetime := now.Format(time.DateTime)
	unixtime := now.UnixMilli()

	count++

	var builder strings.Builder

	builder.WriteString("Accessor: " + strconv.FormatInt(count, 10) + "\n")
	builder.WriteString("Datetime: " + datetime + "/" + strconv.FormatInt(unixtime, 10) + "\n")

	builder.WriteString("  Server: " + hostname + "\n")
	builder.WriteString("  Client: " + r.RemoteAddr + "\n")

	builder.WriteString(" Request: " + r.Method + " " + r.URL.String() + "\n")

	for k, v := range r.Header {
		for _, t := range v {
			builder.WriteString(" Headers: " + k + "=" + t + "\n")
		}
	}

	body := []byte(builder.String())

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text; charset=utf-8")
	_, _ = w.Write(body)

}
