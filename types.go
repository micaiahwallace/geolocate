package geolocate

// Location defines geo-positional coordinates
type Location struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Accuracy int     `json:"accuracy"`
}

// MacAddr contains a json format for a mac address object
type MacAddr struct {
	Mac string `json:"mac"`
}

// LocateRequest holds data required for the json body of a locate request
type LocateRequest struct {
	Wlan []MacAddr `json:"wlan"`
}

// LocateRespErr contains error data for locate response
type LocateRespErr struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

// LocateResp holds json structure for locate requests response
type LocateResp struct {
	Error    LocateRespErr `json:"error"`
	Location Location      `json:"location"`
}
