package usecase

import (
	"fmt"
	"log"
	"project/pkg/config"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/models"
	"strconv"

	"github.com/razorpay/razorpay-go"
)

type paymentUseCase struct {
	payRepo   interfaces.PaymentRepository
	cfg       config.Config
	orderRepo interfaces.OrderRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository, cfg config.Config, order interfaces.OrderRepository) *paymentUseCase {
	return &paymentUseCase{
		payRepo:   repo,
		cfg:       cfg,
		orderRepo: order,
	}
}

func (p paymentUseCase) MakePaymentFromRazorPay(userID, orderID int) (models.OrderPaymentDetails, error) {

	var order models.OrderPaymentDetails

	address, orderData, err := p.orderRepo.GetOrderAddress(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	client := razorpay.NewClient(p.cfg.RAZORPAY_KEY_ID, p.cfg.RAZORPAY_KEY_SECRET)

	orderidStrg := string(rune(orderID))
	receipt := "Uniclub" + orderidStrg
	fmt.Println(receipt, "RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRPPPPPPPPPPPPPPPPPPPPPPPPPPTTTTTTTTTTTTTTTTTTTTTTTTTTTT")
	data := map[string]interface{}{
		"amount":   orderData.Final_Price * 100,
		"currency": "INR",
		"receipt":  receipt,
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println("error in razorpay create function")
		return models.OrderPaymentDetails{}, err
	}

	// for v, d := range body {
	// 	fmt.Println(v, "=", d)
	// }
	// fmt.Println("bodyy of client donnow till now", body)

	razorpayID := body["id"].(string)
	order.UserID = userID
	order.Username = address.Name
	order.Razor_id = razorpayID
	order.OrderID = orderID
	order.FinalPrice = orderData.Final_Price

	return order, nil
}

func (p paymentUseCase) VerifyPaymentFromRazorPay(paymentID, orderID, razorID string) error {

	order_id, err := strconv.Atoi(orderID)
	if err != nil {
		return err

	}

	err = p.payRepo.InsertPaymentDetails(order_id, razorID, paymentID)
	if err != nil {
		return err
	}

	err = p.payRepo.UpdatePaymentDetails(order_id)
	if err != nil {
		return err
	}

	log.Println("\n updated order paymentstatus as PAID with out any errorr")
	return nil
}
