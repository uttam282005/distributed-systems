package types

type GetArgs struct {
	Key string
}

type PutArgs struct {
	Key   string
	Value string
}

type PutReply struct{}

type GetReply struct {
	Value string
}
