package server

import "github.com/dbl90/grpc-server/src/server/pb"

var employees = []pb.Employee{
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
