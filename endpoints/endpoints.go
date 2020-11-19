package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/junereycasuga/gokit-grpc-demo/service"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	Add      endpoint.Endpoint
	Subtract endpoint.Endpoint
	Multiply endpoint.Endpoint
	Divide   endpoint.Endpoint
}

// MathReq struct holds the endpoint request definition
type MathReq struct {
	NumA float32
	NumB float32
}

// MathResp struct holds the endpoint response definition
type MathResp struct {
	Result float32
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Add:      makeAddEndpoint(s),
		Subtract: makeSubtractEndpoint(s),
		Multiply: makeMultiplyEndpoint(s),
		Divide:   makeDivideEndpoint(s),
	}
}

func makeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Add(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeSubtractEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Subtract(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeMultiplyEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Multiply(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeDivideEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Divide(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}
