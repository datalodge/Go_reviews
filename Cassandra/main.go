package cassandra

import (
	"fmt"
	"os"

	"github.com/gocql/gocql"
)

// Session for cassandra connection
var Session *gocql.Session

// creates a cassandra connection
func init() {
	var err error

	//gets the ec2 from env and creates a session to cassandra to get data from the cluster
	ec2 := os.Getenv("DATA_LODGE")
	cluster := gocql.NewCluster(ec2)
	cluster.Keyspace = "data_lodge"

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")

}
