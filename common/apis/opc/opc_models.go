package opc

type Params struct {
	Symbol       string
	CurrentPrice float64
	// todo
}

type OPCResponse struct {
	Status  int     `json:"status"`
	TabDesc TabDesc `json:"tab_desc"`
}

type TabDesc struct {
	Description string `json:"desc"`
	Type        string `json:"type"`
}

type Results struct {
	Data map[string][]Point `json:"data"`
}

type Point struct {
	Gain int `json:"g"`
}
