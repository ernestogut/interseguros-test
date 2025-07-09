package qr

type (
	MatrixRequest struct {
		Data [][]float64 `json:"data"`
	}

	QRResponse struct {
		Q [][]float64 `json:"q"`
		R [][]float64 `json:"r"`
	}
)
