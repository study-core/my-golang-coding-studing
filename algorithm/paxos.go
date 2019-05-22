package main

import "sync"

func main() {

}

type Paxos struct {
	mu sync.Mutex
	rounds  map[int]Round
	completes  [] int
}

func (px Paxos) newInstance() Round {
	return Round{}
}

type   PrepareArgs struct {
	Seq, PNum  int
}

type PrepareReply struct {
	Err, AcceptPnum int
	AcceptValue string

}

type Round struct {
	proposeNumber, acceptorNumber, state int
	acceptValue string
}

const (
	OK  int = iota
	Reject
	Decided
)

func (px *Paxos)Prepare(args *PrepareArgs, reply *PrepareReply) error {
	px.mu.Lock()
	defer px.mu.Unlock()
	round, exist := px.rounds[args.Seq]
	if !exist {
		//new seq  of commit,so need new
		px.rounds[args.Seq] = px.newInstance()
		round, _ = px.rounds[args.Seq]
		reply.Err = OK
	}else {
		if args.PNum > round.proposeNumber {
			reply.Err = OK
		}else {
			reply.Err = Reject
		}
	}
	if reply.Err == OK {
		reply.AcceptPnum = round.acceptorNumber
		reply.AcceptValue = round.acceptValue
		round := px.rounds[args.Seq]
		round.proposeNumber = args.PNum
		px.rounds[args.Seq] = round
	}else {
		//reject
	}
	return nil
}

type  AcceptArgs struct {
	Seq, PNum int
	Value string
}

type AcceptReply struct {
	Err int
}
func (px Paxos)Accept(args *AcceptArgs, reply *AcceptReply) error {
	px.mu.Lock()
	defer px.mu.Unlock()
	round, exist := px.rounds[args.Seq]
	if !exist {
		px.rounds[args.Seq] = px.newInstance()
		reply.Err = OK
	}else {
		if args.PNum >= round.proposeNumber {
			reply.Err = OK
		}else {
			reply.Err = Reject
		}
	}
	if reply.Err == OK {
		round := px.rounds[args.Seq]
		round.acceptorNumber = args.PNum
		round.proposeNumber = args.PNum
		round.acceptValue = args.Value
		px.rounds[args.Seq] = round
	}else {
		//reject
	}
	return nil
}

type  DecideArgs struct {
	Seq, PNum, Me,  Done int
	Value string
}

type  DecideReply struct {

}

func (px *Paxos)Decide(args *DecideArgs, reply *DecideReply) error {
	px.mu.Lock()
	defer px.mu.Unlock()
	_, exist := px.rounds[args.Seq]
	if !exist {
		px.rounds[args.Seq] = px.newInstance()
	}
	round := px.rounds[args.Seq]
	round.acceptorNumber = args.PNum
	round.acceptValue = args.Value
	round.proposeNumber = args.PNum
	round.state = Decided
	px.completes[args.Me] = args.Done
	return nil
}