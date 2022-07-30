package tcknverify

import (
	"net/http"
	"strconv"
	"time"

	"github.com/tiaguinho/gosoap"
)

type ValidateData struct {
	TCKimlikNo string `xml:"TCKimlikNo"`
	Ad         string `xml:"Ad"`
	Soyad      string `xml:"Soyad"`
	DogumYili  string `xml:"DogumYili"`
}

type TcVerifyResponse struct {
	TcVerifyResult string `xml:"TCKimlikNoDogrulaResult"`
}

var response TcVerifyResponse

func Validate(data *ValidateData) (bool, error) {
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}

	soap, err := gosoap.SoapClient("https://tckimlik.nvi.gov.tr/service/kpspublic.asmx?WSDL", httpClient)
	if err != nil {
		return false, err
	}

	params := gosoap.Params{
		"TCKimlikNo": data.TCKimlikNo,
		"Ad":         data.Ad,
		"Soyad":      data.Soyad,
		"DogumYili":  data.DogumYili,
	}

	res, err := soap.Call("TCKimlikNoDogrula", params)
	if err != nil {
		return false, err
	}

	err = res.Unmarshal(&response)
	if err != nil {
		return false, err
	}

	result, _ := strconv.ParseBool(response.TcVerifyResult)
	return result, nil
}
