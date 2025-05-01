package routes

import "time"

type request struct {
	URL         string        "json:url"
	CustomShort string        "json:custom_short"
	Expiry      time.Duration "json:expiry"
}

type response struct {
	URL             string        "json:url"
	CustomShortURL  string        "json:short_url"
	Expiry          time.Duration "json:expiry"
	XRateRemaining  int           "json:x_rate_limit"
	XRateLimitReset int           "json:x_rate_limit_reset"
}
