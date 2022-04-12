package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/SantiagoZuluaga/drawflowapi/app/config"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

var (
	once   sync.Once
	client *dgo.Dgraph
	err    error
)

func init() {
	once.Do(func() {
		conn, errorConnection := grpc.Dial(
			config.DATABASE_URI,
			grpc.WithInsecure(),
		)
		if err != nil {
			fmt.Println("DATABASE CONNECTION ERROR")
			err = errorConnection
			return
		}

		fmt.Println("DATABASE CONNECTION SUCCESSFUL")
		client = dgo.NewDgraphClient(api.NewDgraphClient(conn))

		schema := &api.Operation{
			RunInBackground: true,
			Schema: `				
				fullname: string .
				email: string @index(exact) .
				password: string .
				blocked: bool .
				createdAt: datetime .
				updatedAt: datetime .

				type User {
					fullname
					email
					password
					blocked
					createdAt
					updatedAt
				}
			`,
		}

		errorSchema := client.Alter(context.Background(), schema)
		if err != nil {
			fmt.Println(err)
			err = errorSchema
			return
		}
	})
}

func GetDatabase() (*dgo.Dgraph, error) {
	return client, err
}
