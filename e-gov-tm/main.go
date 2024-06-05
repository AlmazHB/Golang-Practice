package main

import (
	"fmt"
	// "log"
	"net/http"
	handler "nuraly/Handler"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Server runnig ...")

	mH := handler.ServerProvider()
	mH.MyServer()
	// s := echo.New()

	// s.GET("/soltan", Handler)
	// err := s.Start(":8081")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

type Soltan struct {
	d1 int
	d2 int
	d3 int
}

func (s *Soltan) ulaltmak() {
	s.d1 *= 2
	s.d2 *= 3
	s.d3 *= 4
}
func (s *Soltan) capEtmek() {
	fmt.Println(s.d1)
	fmt.Println(s.d2)
	fmt.Println(s.d3)
}
func Handler(c echo.Context) error {

	s1 := Soltan{
		d1: 2,
		d2: 4,
		d3: 9,
	}
	s := fmt.Sprintf("%d,%d,%d", s1.d1, s1.d2, s1.d3)
	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	s1.capEtmek()
	s1.ulaltmak()
	s1.capEtmek()
	return nil
}
