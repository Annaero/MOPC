package helpers

import (
	"errors"
	"net/http"
)

func GetRequestParams(r *http.Request, paramsList map[string]bool) (map[string]string, error) {

	missedParams := ""
	retrivedParams := make(map[string]string)
	for paramName, required := range paramsList {
		param := r.URL.Query().Get(paramName)
		if required && param == "" {
			missedParams += "start_date, "
		} else {
			retrivedParams[paramName] = param
		}
	}

	if missedParams != "" {
		return nil, errors.New("Following params are missed " + missedParams)
	}

	return retrivedParams, nil
}
