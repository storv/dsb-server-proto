package dsb_server

import (
	pb "gitlab.duoshengbu.com/be/dsb-server-proto/work/sg/server/v1"
	"google.golang.org/grpc"
)

func NewClient(address string) (client pb.DsbApiClient) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	return pb.NewDsbApiClient(conn)
}
