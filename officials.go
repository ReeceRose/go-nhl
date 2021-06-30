package gonhl

import "encoding/json"

func (c *Client) getOfficials(appendedURL string) (OfficialsResponse, int, error) {
	var response OfficialsResponse
	data, statusCode, err := c.get(RECORDS_BASE_URL + "/officials/" + appendedURL)
	if err != nil {
		return response, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response, statusCode, nil
}

func (c *Client) GetOfficials() (OfficialsResponse, int, error) {
	return c.getOfficials("")
}

func (c *Client) GetActiveOfficials() (OfficialsResponse, int, error) {
	return c.getOfficials("?cayenneExp=active=true")
}

// TODO: refactor
type OfficialsResponse struct {
	Officials []struct {
		ID                    int         `json:"id"`
		Active                bool        `json:"active"`
		AssociationURL        interface{} `json:"associationUrl"`
		BirthCity             interface{} `json:"birthCity"`
		BirthDate             interface{} `json:"birthDate"`
		CountryCode           interface{} `json:"countryCode"`
		FirstName             string      `json:"firstName"`
		FirstPlayoffGameID    interface{} `json:"firstPlayoffGameId"`
		FirstRegularGameID    int         `json:"firstRegularGameId"`
		HeadshotURL           interface{} `json:"headshot_url"`
		LastName              string      `json:"lastName"`
		OfficialType          string      `json:"officialType"`
		OfficialsSchemaID     interface{} `json:"officialsSchemaId"`
		ReferreeAssociationID interface{} `json:"referreeAssociationId"`
		StateProvinceCode     interface{} `json:"stateProvinceCode"`
		SweaterNumber         interface{} `json:"sweaterNumber"`
		ThumbURL              interface{} `json:"thumb_url"`
	} `json:"data"`
	Total int `json:"total"`
}
