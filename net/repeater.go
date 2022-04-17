package net

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcEffector func(context.Context) (interface{}, error)

func GrpcRepeater(effector GrpcEffector, retries int, delay time.Duration) GrpcEffector {
	return func(ctx context.Context) (interface{}, error) {
		startDelay := delay

		for r := 0; ; r++ {
			resp, err := effector(ctx)
			switch err {
			case nil:
				return resp, err
			default:
				s, _ := status.FromError(err)
				if s.Code() != codes.Unavailable || r >= retries {
					return resp, err
				}
			}

			delay += startDelay
			fmt.Printf("Attempt %d failed; retrying in %v\n", r+1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
	}
}
