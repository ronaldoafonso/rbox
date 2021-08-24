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
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&field, "field", "f", "ssid", "config field")
	setCmd.Flags().StringVarP(&value, "value", "v", "", "config field value")
}

var (
	setCmd = &cobra.Command{
		Use:   "set BOXNAME [BOXNAMES...]",
		Short: "Set configuration of remotebox(es)",
		Args:  cobra.MinimumNArgs(1),
		Run:   set,
	}
)

func set(cmd *cobra.Command, boxnames []string) {
	results := make(chan result)

	for _, boxname := range boxnames {
		go func(boxname string) {
			var err error

			server += ":50051"
			conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("fail to connect to rbox server: %v.", err)
			}
			defer conn.Close()

			c := pb.NewRemoteCommandClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			setRequest := &pb.SetRequest{
				Boxname: boxname,
				Field:   field,
				Value:   value,
			}
			retMessage, err := c.Set(ctx, setRequest)
			results <- result{boxname, err, retMessage.GetReturnMsg()}
		}(boxname)
	}

	for i := 0; i < len(boxnames); i++ {
		r := <-results
		if r.err != nil {
			fmt.Printf("%v box error: %v.\n", r.boxname, r.err)
		} else {
			fmt.Printf("%v: %v.\n", r.boxname, r.info)
		}
	}
}
