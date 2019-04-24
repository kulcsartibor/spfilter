package spfilter

import (
	"fmt"
	"io"
	"os"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/metrics"
	clog "github.com/coredns/coredns/plugin/pkg/log"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

// Define log to be a logger with the plugin name in it. This way we can just use log.Info and
// friends to log.
var log = clog.NewWithPlugin("spfilter")

// SpFilter is an spfilter plugin to show how to write a plugin.
type SpFilter struct {
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface. This method gets called when spfilter is used
// in a Server.
func (e SpFilter) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	// This function could be simpler. I.e. just fmt.Println("spfilter") here, but we want to show
	// a slightly more complex spfilter as to make this more interesting.
	// Here we wrap the dns.ResponseWriter in a new ResponseWriter and call the next plugin, when the
	// answer comes back, it will print "spfilter".

	// Debug log that we've have seen the query. This will only be shown when the debug plugin is loaded.
	clog.Info("Received response")
	clog.Info(w.LocalAddr().String() + " -> " + w.RemoteAddr().String())
	clog.Info("Question: " + r.Question[1].Name)

	// Wrap.
	//pw := NewResponsePrinter(w)

	// Export metric with the server label set to the current server handling the request.
	requestCount.WithLabelValues(metrics.WithServer(ctx)).Inc()

	// Call next plugin (if any).
	return dns.RcodeRefused, nil
	//return plugin.NextOrFailure(e.Name(), e.Next, ctx, pw, r)
}

// Name implements the Handler interface.
func (e SpFilter) Name() string { return "spfilter" }

// ResponsePrinter wrap a dns.ResponseWriter and will write spfilter to standard output when WriteMsg is called.
type ResponsePrinter struct {
	dns.ResponseWriter
}

// NewResponsePrinter returns ResponseWriter.
func NewResponsePrinter(w dns.ResponseWriter) *ResponsePrinter {
	return &ResponsePrinter{ResponseWriter: w}
}

// WriteMsg calls the underlying ResponseWriter's WriteMsg method and prints "spfilter" to standard output.
func (r *ResponsePrinter) WriteMsg(res *dns.Msg) error {
	fmt.Fprintln(out, ex)

	res.Rcode = dns.RcodeRefused

	return r.ResponseWriter.WriteMsg(res)
}

// Make out a reference to os.Stdout so we can easily overwrite it for testing.
var out io.Writer = os.Stdout

// What we will print.
const ex = "spfilter"
