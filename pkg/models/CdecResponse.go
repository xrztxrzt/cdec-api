package models

type CdecResponse struct {
	TariffCodes []struct {
		TariffCode        int     `json:"tariff_code"`
		TariffName        string  `json:"tariff_name"`
		TariffDescription string  `json:"tariff_description"`
		DeliveryMode      int     `json:"delivery_mode"`
		DeliverySum       float64 `json:"delivery_sum"`
		PeriodMin         int     `json:"period_min"`
		PeriodMax         int     `json:"period_max"`
		CalendarMin       int     `json:"calendar_min"`
		CalendarMax       int     `json:"calendar_max"`
	} `json:"tariff_codes"`
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}
