# bAuthBrute
basic auth endpoint brute forcer

# usage
go run main.go -users userListLocation -pass passwordListLocation -u https://site.com/thing -m POST

# TODO
- [ ] add feedback, filter results based on 200/301 or maybe custom responses
- [X] add ability to GET vs. POST
- [ ] add some ability to slow this down, currently it's gonna get you ratelimited quickly
  * probably should throw all potential permutations into a list and create requests in batches
- [ ] add ability to add cookies in case these are needed to hit endpoint
- [ ] add ability to "go get" package - fix import names for something non-local
