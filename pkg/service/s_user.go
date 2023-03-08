package service

import (
	"context"
	"log"

	"github.com/borud/tt/pkg/model"
	"github.com/borud/tt/pkg/password"
	ttv1 "github.com/borud/tt/pkg/tt/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) AddUser(ctx context.Context, req *ttv1.AddUserRequest) (*ttv1.AddUserResponse, error) {
	user := model.UserFromProto(req.User)

	if err := user.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	hash, err := password.Hash(user.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "error hashing password")
	}
	user.Password = string(hash)

	// TODO (borud): should not directly return db error
	return &ttv1.AddUserResponse{}, s.config.DB.AddUser(user)
}

func (s *service) GetUser(ctx context.Context, req *ttv1.GetUserRequest) (*ttv1.GetUserResponse, error) {
	user, err := s.config.DB.GetUser(req.Username)
	if err != nil {
		log.Printf("GetUser failed: %v", err)
		return nil, status.Error(codes.NotFound, "user not found")
	}

	user.RemoveSensitive()

	return &ttv1.GetUserResponse{User: user.Proto()}, nil
}

func (s *service) UpdateUser(ctx context.Context, req *ttv1.UpdateUserRequest) (*ttv1.UpdateUserResponse, error) {
	//TODO(borud): implement
	return nil, nil
}

func (s *service) DeleteUser(ctx context.Context, req *ttv1.DeleteUserRequest) (*ttv1.DeleteUserResponse, error) {
	//TODO(borud): implement
	return nil, nil
}

func (s *service) ListUsers(ctx context.Context, req *ttv1.ListUsersRequest) (*ttv1.ListUsersResponse, error) {
	//TODO(borud): implement
	return nil, nil
}
