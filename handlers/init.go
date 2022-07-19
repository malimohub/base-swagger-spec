package handlers

import (
	"context"

	"github.com/crypto-checkout/core/lib/log"
	"github.com/crypto-checkout/server/restapi/operations"
	"github.com/opentracing/opentracing-go"
)

func ImplementAPI(api *operations.CryptoCheckoutAPI) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("ImplementAPI(Boot)")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	log.Info(ctx, "implementing api....")
	oh := newOrderHandlers()
	oh.configureRoutes(api)

}
