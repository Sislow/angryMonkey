# angryMonkey
## Password hash via golang

#### Quick Notes:
- Using global structs might not be allowed (or safe in golang). Not sure. Didn't see a lot of negative notes around the web
- Mutex likely should have been used rather than bools. The bools worked fine, but I think a mutex is built for the purpose of restricting a shutdown while processing
- Race condition exists in this code (curl post to /hash with '&&' and then shutdown). Doesn't always show up, but I have seen it preemptively shutdown. Seems to work fine with form via index or multiple calls in different multi terminals
- Little weak on the test cases. Don't have a full understanding of how the injection is being utilized. Watching a tutorial :)

#### Example Commands:
```
// collect html frontend
curl http://localhost:8080/

// examples of hash command
curl http://localhost:8080/hash --data password="angryMonkey"
curl http://localhost:8080/hash -d "angryMonkey"

// closes running operation
curl http://localhost:8080/shutdown

// returns json of average hash run time and count of calls to hash
curl http://localhost:8080/stats
```
