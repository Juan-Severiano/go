package main

import (
	"fmt"
	"log"

	"github.com/fonini/go-pix/pix"
	"github.com/skip2/go-qrcode"
)

func main() {
	options := pix.Options{
		Name: "NAME_OF",
		Key:  "PIX_KEY",
		City: "CITY",
		// Amount:        0.1,         // Valor em reais (opcional)
		// Description:   "Pagamento teste", // Descrição do pagamento (opcional)
		// TransactionID: "123456",      // TXID (opcional)
	}

	copyPaste, err := pix.Pix(options)
	if err != nil {
		log.Fatalf("Erro ao gerar código PIX: %v", err)
	}

	fmt.Println("Código PIX (Copia e Cola):", copyPaste)

	qrCodeFile := "qrcode_pix.png"
	err = qrcode.WriteFile(copyPaste, qrcode.Medium, 256, qrCodeFile)
	if err != nil {
		log.Fatalf("Erro ao gerar QR Code: %v", err)
	}

	fmt.Println("QR Code gerado com sucesso e salvo como", qrCodeFile)
}
