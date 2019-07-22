package level11api

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"level11infrastructure"
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

// func CreateUser(c echo.Context) (err error) {
// 	u := new(User)
// 	if err = c.Bind(u); err != nil {
// 		return
// 	}

// 	params := make(map[interface{}]interface{})
// 	params["sql"] = fmt.Sprintf("INSERT INTO users (name,email) VALUES ('%v','%v')", u.Name, u.Email)
// 	level11infrastructure.Execute(params)
// 	return c.JSON(http.StatusOK, u)
// }

func TestFetch(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, "TestFetch test"+os.Getenv("POSTGRES_PORT")+os.Getenv("POSTGRES_USER")+os.Getenv("POSTGRES_PASSWORD")+os.Getenv("POSTGRES_DB"))
}

func RequestCheckId(c echo.Context) (err error) {

	url := "https://idcard.bora.dopa.go.th/CheckCardStatus/CheckCardService.asmx/CheckCardByLaser?PID=1101400515901&FirstName=ชานน&LastName=ยาคล้าย&BirthDay=25280713&Laser=ME0120407681"

	// connection reset by peer
	// เหมือนต้อง set ด้วยไม่งั้นมัน reset ไม่รู็เพราะอะไร ??
	// https://stackoverflow.com/questions/37774624/go-http-get-concurrency-and-connection-reset-by-peer
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
		temp := fmt.Sprintf("http.DefaultClient.Do() failed with:\n'%s'\n", err)
		fmt.Println(temp)
		//log.Fatalf("http.DefaultClient.Do() failed with:\n'%s'\n", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var cardStatusOut CardStatusOut
	resUnmarshal := xml.Unmarshal(body, &cardStatusOut)
	fmt.Println(resUnmarshal)  // => เอาไปเก็บไว้ใน DB
	fmt.Println(cardStatusOut) // => เอาไปเก็บไว้ใน DB

	// fmt.Println("1 :", cardStatusOut.XMLName)
	// fmt.Println("2 :", cardStatusOut.IsError)
	// fmt.Println("3 :", cardStatusOut.ErrorMessage)
	// fmt.Println("4 :", cardStatusOut.Code)
	// fmt.Println("5 :", cardStatusOut.Desc)

	//cardStatusOutResponse := CardStatusOutResponse{IsError: "TEST", ErrorMessage: "TEST", Code: "TEST", Desc: "TEST"}
	cardStatusOutResponse := CardStatusOutResponse{IsError: cardStatusOut.IsError, ErrorMessage: cardStatusOut.ErrorMessage, Code: cardStatusOut.Code, Desc: cardStatusOut.Desc}

	_resUnmarshal, _ := json.Marshal(resUnmarshal)
	_cardStatusOutResponse, _ := json.Marshal(cardStatusOutResponse)

	params := make(map[interface{}]interface{})
	params["sql"] = fmt.Sprintf("INSERT INTO log (card_details,xml_details) VALUES ('%s','%s')", _cardStatusOutResponse, _resUnmarshal) //  ใช้ '%v' ไม่ได้สำหรับ jsonb
	level11infrastructure.Execute(params)

	return c.JSON(http.StatusOK, cardStatusOutResponse)
}
