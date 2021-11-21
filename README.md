# go-state

```bash
(cd ../go-state-sample && go run .)
curl --unix-socket /tmp/sample.state http://localhost/
curl --unix-socket /tmp/sample.state http://localhost/environ/
curl --unix-socket /tmp/sample.state http://localhost/node/state/
curl --unix-socket /tmp/sample.state 'http://localhost/node/state/?values=true'
curl --unix-socket /tmp/sample.state http://localhost/node/root/
curl --unix-socket /tmp/sample.state 'http://localhost/node/root/?values=true'
curl --unix-socket /tmp/sample.state http://localhost/debug/pprof/
curl --unix-socket /tmp/sample.state 'http://localhost/debug/pprof/goroutine?debug=2'
```
