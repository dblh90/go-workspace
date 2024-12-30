package main

import (
	"errors"
	"io"
	"log"
	"net"

	"github.com/dbl90/grpc-server/src/server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	creds, err := credentials.NewServerTLSFromFile("localhost.crt", "localhost.key")

	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	pb.RegisterEmployeeServiceServer(s, new(employeeService))
	log.Printf("Started GRPC server!, on port %s", "9000")
	s.Serve(lis)
}

type employeeService struct{}

func (s *employeeService) GetByBadgeNumber(ctx context.Context, req *pb.GetByBadgeNumberRequest) (*pb.EmployeeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Something wrong while receiving info!")
	}

	log.Printf("Metadata received:\n%v\n", md)

	for _, e := range getEmps() {
		if req.BadgeNumber == e.BadgeNumber {
			return &pb.EmployeeResponse{Employee: &e}, nil
		}
	}

	return nil, errors.New("Employee not found!")
}
func (s *employeeService) GetAll(req *pb.GetAllRequest, stream pb.EmployeeService_GetAllServer) error {
	for _, e := range getEmps() {
		stream.Send(&pb.EmployeeResponse{Employee: &e})
	}
	return nil
}

func (s *employeeService) mustEmbedUnimplementedEmployeeServiceServer() {
}

func (s *employeeService) AddPhoto(stream pb.EmployeeService_AddPhotoServer) error {
	_, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Println("Error in getting context!")
	}
	imgData := []byte{}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			log.Printf("File received with length:  %v\n", len(imgData))
			return stream.SendAndClose(&pb.AddPhotoResponse{IsOk: true})
		}
		if err != nil {
			return err
		}
		log.Printf("Received %v bytes\n", len(data.Data))
		imgData = append(imgData, data.Data...)
	}
}

func (s *employeeService) Save(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (s *employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error {
	for {
		emp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		emps := getEmps()
		emps = append(emps, *emp.Employee)
		stream.Send(&pb.EmployeeResponse{Employee: emp.Employee})

		for _, e := range emps {
			log.Println(e)
		}
		return nil
	}
	return nil
}

func getEmps() []pb.Employee {
	return []pb.Employee{
		{
			Id:                  1,
			BadgeNumber:         2080,
			FirstName:           "Grace",
			LastName:            "Decker",
			VacationAccrualRate: 2,
			VacationAccrued:     30,
		},
		{
			Id:                  2,
			BadgeNumber:         7538,
			FirstName:           "Amity",
			LastName:            "Fuller",
			VacationAccrualRate: 2.3,
			VacationAccrued:     23.4,
		},
		{
			Id:                  3,
			BadgeNumber:         5144,
			FirstName:           "Keaton",
			LastName:            "Willis",
			VacationAccrualRate: 3,
			VacationAccrued:     31.7,
		},
	}
}
