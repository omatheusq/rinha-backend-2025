package payments

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type PaymentService struct {
	redisClient *redis.Client
}

func NewPaymentService(client *redis.Client) *PaymentService {
	return &PaymentService{
		redisClient: client,
	}
}

func (s *PaymentService) QueuePayment(ctx context.Context, p Payment) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return s.redisClient.LPush(ctx, "payments", data).Err()
}
