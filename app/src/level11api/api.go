package level11api

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

// ต้อง Refactor
type CardStatusOut struct {
	XMLName      xml.Name `xml:"CardStatusOut"`
	IsError      string   `xml:"IsError"`
	ErrorMessage string   `xml:"ErrorMessage"`
	Code         string   `xml:"Code"`
	Desc         string   `xml:"Desc"`
}

// type CardStatusOutResponse struct {
// 	IsError      string `json:"is_error"`
// 	ErrorMessage string `json:"error_message"`
// 	Code         string `json:"code"`
// 	Desc         string `json:"desc"`
// }

type CardStatusOutResponse struct {
	IsError      string
	ErrorMessage string
	Code         string
	Desc         string
}

//----------
// Handlers
//----------

func TestFetch(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, "TestFetch test"+os.Getenv("POSTGRES_PORT")+os.Getenv("POSTGRES_USER")+os.Getenv("POSTGRES_PASSWORD")+os.Getenv("POSTGRES_DB"))
}

func RequestCheckId(c echo.Context) (err error) {

	url := "https://idcard.bora.dopa.go.th/CheckCardStatus/CheckCardService.asmx/CheckCardByLaser?PID=1101400515901&FirstName=ชานน&LastName=ยาคล้าย&BirthDay=25280713&Laser=ME0120407681"

	//url := "http://www.mocky.io/v2/5d32dc283400006200749ef2"

	req, err := http.NewRequest("GET", url, nil)

	//req.Header.Add("User-Agent", "PostmanRuntime/7.15.2")
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("Accept-Encoding", "utf8")

	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}

	// Timeout = 100 ms
	// SetTimeOut ด้วย http://www.somkiat.cc/dont-forgot-to-set-timeout/
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*1000)
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("http.DefaultClient.Do() failed with:\n'%s'\n", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var cardStatusOut CardStatusOut
	xml.Unmarshal(body, &cardStatusOut)

	// fmt.Println("1 :", cardStatusOut.XMLName)
	// fmt.Println("2 :", cardStatusOut.IsError)
	// fmt.Println("3 :", cardStatusOut.ErrorMessage)
	// fmt.Println("4 :", cardStatusOut.Code)
	// fmt.Println("5 :", cardStatusOut.Desc)

	//cardStatusOutResponse := CardStatusOutResponse{IsError: "TEST", ErrorMessage: "TEST", Code: "TEST", Desc: "TEST"}
	cardStatusOutResponse := CardStatusOutResponse{IsError: cardStatusOut.IsError, ErrorMessage: cardStatusOut.ErrorMessage, Code: cardStatusOut.Code, Desc: cardStatusOut.Desc}

	return c.JSON(http.StatusOK, cardStatusOutResponse)
}
