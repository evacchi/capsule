module github.com/bots-garden/capsule/capsulemodule

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../capsulemodule

replace (
	github.com/bots-garden/capsule/commons => ../commons
	github.com/bots-garden/capsule/commons v0.2.9 => ../commons
)

require github.com/bots-garden/capsule/commons v0.0.0-00010101000000-000000000000

replace github.com/bots-garden/capsule v0.2.9 => ../
