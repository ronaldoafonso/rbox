package cmd

import (
	"context"
	"fmt"
	pb "github.com/ronaldoafonso/rbox/gcommand"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&field, "field", "f", "ssid", "config field")
}

var (
	getCmd = &cobra.Command{
		Use:   "get BOXNAME [BOXNAMES...]",
		Short: "Get configuration of remotebox(es)",
		Args:  cobra.MinimumNArgs(1),
		Run:   get,
	}
)

func get(cmd *cobra.Command, boxnames []string) {
	results := make(chan result)

	for _, boxname := range boxnames {
		go func(boxname string) {
			var err error

			server += ":" + port
			log.Println(server)
			conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("fail to connect to rbox server: %v.", err)
			}
			defer conn.Close()

			c := pb.NewRemoteCommandClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			getRequest := &pb.GetRequest{
				Boxname: boxname,
				Field:   field,
			}
			retMessage, err := c.Get(ctx, getRequest)
			results <- result{boxname, err, retMessage.GetReturnMsg()}
		}(boxname)
	}

	for i := 0; i < len(boxnames); i++ {
		r := <-results
		if r.err != nil {
			fmt.Printf("%v box error: %v.\n", r.boxname, r.err)
		} else {
			fmt.Printf("%v: %v\n", r.boxname, r.info)
		}
	}
}
