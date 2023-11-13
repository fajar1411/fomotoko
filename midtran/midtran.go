package midtran

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func TopUpFromBankToWallet(orderid string, amount int64, metode string) (respon *coreapi.ChargeResponse) {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	switch {
	case metode == "transfer_va_bca":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderid,
				GrossAmt: amount,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
			CustomExpiry: &coreapi.CustomExpiry{

				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		respon, _ = coreapi.ChargeTransaction(chargeReq)

		return respon
	case metode == "transfer_va_permata":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderid,
				GrossAmt: amount,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
			CustomExpiry: &coreapi.CustomExpiry{

				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		respon, _ = coreapi.ChargeTransaction(chargeReq)

		return respon
	case metode == "transfer_va_bni":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderid,
				GrossAmt: amount,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
			CustomExpiry: &coreapi.CustomExpiry{

				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		respon, _ = coreapi.ChargeTransaction(chargeReq)

		return respon
	case metode == "transfer_va_bri":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderid,
				GrossAmt: amount,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
			CustomExpiry: &coreapi.CustomExpiry{

				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		respon, _ = coreapi.ChargeTransaction(chargeReq)

		return respon
	}
	return &coreapi.ChargeResponse{}
}
