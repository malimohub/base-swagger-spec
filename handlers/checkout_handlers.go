package handlers

import (
	"encoding/base64"
	"fmt"

	"github.com/crypto-checkout/core/lib/algorand"
	"github.com/crypto-checkout/core/lib/log"
	"github.com/crypto-checkout/server/models"
	"github.com/crypto-checkout/server/restapi/operations"
	ops "github.com/crypto-checkout/server/restapi/operations/crypto_checkout"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

type orderHandlers struct{}

func newOrderHandlers() *orderHandlers {
	return &orderHandlers{}
}

func (h *orderHandlers) configureRoutes(api *operations.CryptoCheckoutAPI) {
	// health checks
	api.CryptoCheckoutPostCheckoutHandler = ops.PostCheckoutHandlerFunc(h.PostCheckout)
}

// PostCheckout handles a POST request to the /checkout endpoint and instructs
// the API to start a checkout
func (h *orderHandlers) PostCheckout(params ops.PostCheckoutParams) middleware.Responder {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("PostOrder(Handler)")
	ctx := opentracing.ContextWithSpan(params.HTTPRequest.Context(), span)
	defer span.Finish()

	log.Info(ctx, "new order", log.WithValue("params", params))

	var amount int64 = 199 * 100

	label := "Malcolm+Monroe"

	sessionID := uuid.New().String()
	note := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"session_id": %s}`, sessionID)))
	walletAddress := "EWAMATGAOROGS3HW3NO3KPAR67YKQSMA7IMHKYMNV4P5UBQFPPWVIQOAB4"

	qrGen := algorand.NewQRCodeGenerator(walletAddress, label)
	code, err := qrGen.GenerateQRCode(ctx, algorand.GenerateQRCodeParams{
		AssetID: 446355880,
		Amount:  amount,
		Note:    note,
	})
	if err != nil {
		log.Error(ctx, "failed to generate QR code", log.WithError(err))
	}

	return ops.NewPostCheckoutOK().WithPayload(&models.Checkout{
		RefID: sessionID,
		AlgorandQrcode: &models.CheckoutQRCode{
			Base64EncodedQrCode: base64.StdEncoding.EncodeToString(code),
		},
	})
}
