package algorand

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/crypto-checkout/core/lib/log"
	"github.com/pkg/errors"
)

const (
	uriScheme = "algorand://%s?label=%s&xnote=%s&amount=%d&asset=%d"
)

// QRCodeGenerator contains dependencies for a algorand QR code generator
type QRCodeGenerator struct {
	WalletAdress string
	Label        string
}

// NewQRCodeGenerator returns a new algorand QR code generator
func NewQRCodeGenerator(walletAddress, label string) *QRCodeGenerator {
	return &QRCodeGenerator{
		WalletAdress: walletAddress,
		Label:        label,
	}
}

type GenerateQRCodeParams struct {
	Amount  int64
	AssetID int64
	Note    string
}

// GenerateQRCode generates a QR code for purchases
func (g *QRCodeGenerator) GenerateQRCode(ctx context.Context, params GenerateQRCodeParams) ([]byte, error) {

	dataString := fmt.Sprintf(uriScheme,
		g.WalletAdress,
		g.Label,
		params.Note,
		params.Amount,
		params.AssetID)

	log.Info(ctx, "QR code data string", log.WithValue("string", dataString))

	qrCode, err := qr.Encode(dataString, qr.L, qr.Auto)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode")
	}

	qrCode, err = barcode.Scale(qrCode, 512, 512)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scale")
	}

	var b bytes.Buffer
	bw := bufio.NewWriter(&b)

	if err := png.Encode(bw, qrCode); err != nil {
		return nil, errors.Wrap(err, "failed to encode")
	}

	if err := bw.Flush(); err != nil {
		return nil, errors.Wrap(err, "failed to flush")
	}

	return b.Bytes(), nil
}
