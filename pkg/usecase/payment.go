package usecase

import (
	"errors"
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
	userRepo  interfaces.UserRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository, cfg config.Config, order interfaces.OrderRepository, uRepo interfaces.UserRepository) *paymentUseCase {
	return &paymentUseCase{
		payRepo:   repo,
		cfg:       cfg,
		orderRepo: order,
		userRepo:  uRepo,
	}
}

func (p paymentUseCase) MakePaymentFromRazorPay(userID, orderID int) (models.OrderPaymentDetails, error) {

	paymentStatus, err := p.orderRepo.GetPaymentStatusByID(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	if paymentStatus == "PAID" {
		return models.OrderPaymentDetails{}, errors.New("this order is already paid")
	}

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

func (p paymentUseCase) PaymentFromWallet(orderID, userID int) (models.OrderPaymentDetails, error) {

	paymentStatus, err := p.orderRepo.GetPaymentStatusByID(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	if paymentStatus == "PAID" {
		return models.OrderPaymentDetails{}, errors.New("this order is already paid")
	}

	var order models.OrderPaymentDetails

	address, orderData, err := p.orderRepo.GetOrderAddress(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	wallet, err := p.userRepo.GetWallet(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	/*check whether user has suffiecient balance in his wallet if does
	pay full orderAmount from his wallet else reduce the finalamount and proceed
	and proceed with online payment */

	if wallet.TotalAmount >= orderData.Final_Price {

		var model models.AddMoneytoWallet
		model.UserID = userID
		model.Amount = (-orderData.Final_Price)
		model.TranscationType = "DEBIT"
		err := p.userRepo.AddMoneytoWallet(model)
		if err != nil {
			return models.OrderPaymentDetails{}, err
		}
		//make status as Paid
		err = p.payRepo.UpdatePaymentDetails(orderID)
		if err != nil {
			return models.OrderPaymentDetails{}, err
		}

		return models.OrderPaymentDetails{}, nil

	}
	//since user dont have sufficient balance in wallet reduce the amount and pay rest through online payment
	newFinalAmount := orderData.Final_Price - wallet.TotalAmount

	//update the wallet amount to zero by reducing the Total Amount to zero
	var model models.AddMoneytoWallet
	model.UserID = userID
	model.Amount = (-wallet.TotalAmount)
	model.TranscationType = "DEBIT"
	err = p.userRepo.AddMoneytoWallet(model)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	//pay the rest amount(newFinalAmount) through online

	client := razorpay.NewClient(p.cfg.RAZORPAY_KEY_ID, p.cfg.RAZORPAY_KEY_SECRET)

	orderidStrg := string(rune(orderID))
	receipt := "Uniclub" + orderidStrg
	fmt.Println(receipt, "RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRPPPPPPPPPPPPPPPPPPPPPPPPPPTTTTTTTTTTTTTTTTTTTTTTTTTTTT")
	data := map[string]interface{}{
		"amount":   newFinalAmount * 100,
		"currency": "INR",
		"receipt":  receipt,
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println("error in razorpay create function")
		return models.OrderPaymentDetails{}, err
	}

	razorpayID := body["id"].(string)
	order.UserID = userID
	order.Username = address.Name
	order.Razor_id = razorpayID
	order.OrderID = orderID
	order.FinalPrice = newFinalAmount

	return order, nil
}
