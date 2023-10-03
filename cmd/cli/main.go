package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/Liar233/cronos-shovel/internal/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	if len(os.Args) != 3 {
		_, _ = fmt.Fprintln(os.Stderr, "invalid arguments")
		os.Exit(1)
	}

	host := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "invalid port")
		os.Exit(1)
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewTikTackClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.TikTack(ctx, &pb.TikRequest{Input: "Tik"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Resp: %s", r.GetOutput())
}
