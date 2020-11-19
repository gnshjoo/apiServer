package main

import (
	"context"
	"fmt"
	pb "github.com/gnshjoo/apiServer/gen/proto"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedAuthServer
}

const (
	nextPageKey     = "next_page" // 세션에 저장되는 next page의 키
	authSecurityKey = "auth_security_key"
)

func init() {
	gomniauth.SetSecurityKey(authSecurityKey)
	gomniauth.WithProviders(
		google.New("418841239962-b5i0cikimq8032c3tqe08j773p592pok.apps.googleusercontent.com", "L0jfaclsMlgjvob5p5HqB3Xu",
			"http://127.0.0.1:8081/auth/callback/google"),
	)
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	const provider = "google"
	log.Println("Received profile :", in.Action, in.Provider)

	switch in.Action {
	case "login":
		p, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln(err)
		}
		loginUrl, err := p.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(loginUrl)

	case "callback":
		p, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Print(p)
	}




	return &pb.LoginResponse{Token: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	if err := s.Serve(lis; err != nil  {
		log.Fatalf("failed to server: %v", err)
	}

}