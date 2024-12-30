package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/dbl90/grpc-client/src/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const port = ":9000"

func main() {
	log.Println("Starting client!")
	option := flag.Int("o", 1, "Option to choose what to run")
	flag.Parse()
	creds, err := credentials.NewClientTLSFromFile("localhost.crt", "")
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial("localhost"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)
	switch *option {
	case 1:
		log.Println("Calling SendMetadata")
		SendMetadata(client)
	case 2:
		log.Println("Calling GetByBadgeNumber")
		GetByBadgeNumber(client)
	case 3:
		GetAll(client)
	case 4:
		AddPhoto(client)
	case 5:
		SaveAll(client)
	}
	log.Println("Stopping client!")
}

func SendMetadata(client pb.EmployeeServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"Hamza"}
	md["password"] = []string{"100§$%"}
	client.GetByBadgeNumber(metadata.NewOutgoingContext(context.Background(), md), &pb.GetByBadgeNumberRequest{})
}

func GetByBadgeNumber(client pb.EmployeeServiceClient) {
	client.GetByBadgeNumber(context.Background(), &pb.GetByBadgeNumberRequest{BadgeNumber: 2000})
}

func GetAll(client pb.EmployeeServiceClient) {
	stream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.Employee)
	}

}

func AddPhoto(client pb.EmployeeServiceClient) {
	f, err := os.Open("Penguins.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	md := metadata.New(map[string]string{"badgenumber": "2080"})
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := client.AddPhoto(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for {
		chunk := make([]byte, 64*1024)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n < len(chunk) {
			chunk = chunk[:n]
		}
		stream.Send(&pb.AddPhotoRequest{Data: chunk})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.IsOk)
}

func SaveAll(client pb.EmployeeServiceClient) {
	stream, err := client.SaveAll(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	doneChannel := make(chan struct{})

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				doneChannel <- struct{}{}
			}

			if err != nil {
				log.Fatal(err)
			}
			log.Println(res.Employee)
		}
	}()

	for _, emp := range getEmps() {
		log.Println("---")
		log.Println(emp)
		log.Println("---")
		err := stream.Send(&pb.EmployeeRequest{Employee: &emp})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println()
	}
	stream.CloseSend()
	<-doneChannel
}

func getEmps() []pb.Employee {
	return []pb.Employee{
		{
			Id:                  4,
			BadgeNumber:         5144,
			FirstName:           "Hamza",
			LastName:            "Hasan",
			VacationAccrualRate: 3,
			VacationAccrued:     31.7,
		},
		{
			Id:                  5,
			BadgeNumber:         5145,
			FirstName:           "Mais",
			LastName:            "Willis",
			VacationAccrualRate: 3,
			VacationAccrued:     31.7,
		},
	}
}
