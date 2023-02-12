package main

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIDKey string

const UserIDKey = userIDKey("userID")

var db database = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userid string) {
	// send userID information to checkMemberShip through context for
	// map lookup.
	vctx := context.WithValue(ctx, UserIDKey, "jane")
	ch := checkMemberShip(vctx)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShip - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do map lookup
// sends the result on the returned channel.
func checkMemberShip(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do some database lookup
		userid := ctx.Value(UserIDKey).(string)
		status := db[userid]
		ch <- status
	}()
	return ch
}
