package gol

import (
	"flag"
	"math/rand"
	"net"
	"net/rpc"
	"time"
)

/** create game of life struct**/
type GameOfLife struct{}

func (s *GameOfLife) ProcessAllTurns(req Request, res Response) (err error) {
	turn := 1
	turns := req.Turns
	world := req.World
	var newWorld [][]uint8

	for turn < turns {
		newWorld = calculateNextState(req, world)
		newWorldData := makeImmutableMatrix(newWorld)
		world = newWorldData
	}
	res.newWorld = world
	res.completedTurns = turns
	return
}

func main() {
	pAddr := flag.String("port", "8030", "Port to listen on")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	rpc.Register(&GameOfLife{})
	listener, _ := net.Listen("tcp", ":"+*pAddr)
	defer listener.Close()
	rpc.Accept(listener)
}
