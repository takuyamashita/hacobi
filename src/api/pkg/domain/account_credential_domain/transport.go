package account_credential_domain

import "fmt"

type Transport string

const (
	USB      Transport = "usb"
	NFC      Transport = "nfc"
	BLE      Transport = "ble"
	Hybrid   Transport = "hybrid"
	Internal Transport = "internal"
)

func NewTransports(ts []string) ([]Transport, error) {

	var transports []Transport

	for _, t := range ts {

		var transport Transport

		switch t {
		case "usb":
			transport = USB
		case "nfc":
			transport = NFC
		case "ble":
			transport = BLE
		case "hybrid":
			transport = Hybrid
		case "internal":
			transport = Internal
		default:
			return nil, fmt.Errorf("invalid transport: %s", t)
		}

		transports = append(transports, transport)
	}

	return transports, nil

}
