package grpc

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	orderv1 "github.com/vexner67/zento/order/api/order/v1"
	"github.com/vexner67/zento/order/internal/order"
	"github.com/vexner67/zento/order/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderService interface {
	Create(context.Context, service.CreateOrderInput) (*order.Order, error)
}

type Handler struct {
	orderv1.UnimplementedOrderServiceServer
	svc    OrderService
	logger *slog.Logger
}

func NewHandler(svc OrderService, logger *slog.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}

func (h *Handler) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (*orderv1.Order, error) {
	customerID := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
	items := make([]order.Item, 0, len(req.GetItems()))

	for i, item := range req.GetItems() {
		productID, err := uuid.Parse(item.ProductId)
		if err != nil || productID == uuid.Nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"items[%d].product_id must be a nonzero UUID", i,
			)
		}

		quantity, err := order.NewQuantity(item.GetQuantity())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument, "items[%d].quantity: %s", i, err,
			)
		}

		price := item.GetUnitPrice()

		unitPrice, err := order.NewMoney(
			price.GetCurrencyCode(),
			price.GetUnits(),
			price.GetNanos(),
		)
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument, "items[%d].unit_price: %s", i, err,
			)
		}

		items = append(items, order.NewItem(productID, quantity, unitPrice))
	}

	o, err := h.svc.Create(ctx, service.CreateOrderInput{
		CustomerID: customerID,
		Items:      items,
	})
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, "request canceled")

		case errors.Is(err, context.DeadlineExceeded):
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")

		case errors.Is(err, order.ErrOrderHasNoItems):
			return nil, status.Error(codes.InvalidArgument, "order must contain at least one item")

		case errors.Is(err, order.ErrMixedCurrencies):
			return nil, status.Error(codes.InvalidArgument, "order items must use the same currency")

		default:
			h.logger.ErrorContext(ctx, "create order failed", "error", err)
			return nil, status.Error(codes.Internal, "failed to create order")
		}
	}

	return orderToProto(o), nil
}

func orderToProto(o *order.Order) *orderv1.Order {
	return &orderv1.Order{
		Id:         o.ID().String(),
		CustomerId: o.CustomerID().String(),
		Items:      itemsToProto(o.Items()),
		Status:     statusToProto(o.Status()),
		Total:      moneyToProto(o.Total()),
		CreatedAt:  timestamppb.New(o.CreatedAt()),
		UpdatedAt:  timestamppb.New(o.UpdatedAt()),
	}
}

func itemsToProto(domainItems []order.Item) []*orderv1.OrderItem {
	items := make([]*orderv1.OrderItem, 0, len(domainItems))

	for _, item := range domainItems {
		items = append(items, &orderv1.OrderItem{
			ProductId: item.ProductID().String(),
			Quantity:  item.Quantity().Int32(),
			UnitPrice: moneyToProto(item.UnitPrice()),
		})
	}

	return items
}

func moneyToProto(m order.Money) *orderv1.Money {
	return &orderv1.Money{
		CurrencyCode: m.CurrencyCode(),
		Units:        m.Units(),
		Nanos:        m.Nanos(),
	}
}

func statusToProto(s order.Status) orderv1.OrderStatus {
	switch s {
	case order.Pending:
		return orderv1.OrderStatus_ORDER_STATUS_PENDING
	case order.Confirmed:
		return orderv1.OrderStatus_ORDER_STATUS_CONFIRMED
	case order.Cancelled:
		return orderv1.OrderStatus_ORDER_STATUS_CANCELLED
	case order.Completed:
		return orderv1.OrderStatus_ORDER_STATUS_COMPLETED
	default:
		return orderv1.OrderStatus_ORDER_STATUS_UNSPECIFIED
	}
}
