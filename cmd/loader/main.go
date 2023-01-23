package main

import (
	"context"
	"encoding/xml"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/encoding/charmap"
)

const database = "valutes"
const collection = "valute"

type Resp struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valute  []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName  xml.Name `xml:"Valute" bson:"-"`
	Name     string   `xml:"Name" bson:"name"`
	Value    string   `xml:"Value" bson:"value"`
	CharCode string   `xml:"CharCode" bson:"char_code"`
	Nominal  float64  `xml:"Nominal" bson:"nominal"`
	Date     int64    `xml:"-" bson:"date"`
	ID       string   `xml:"-" bson:"_id"`
}

type ValuteToDB struct {
	XMLName  xml.Name `xml:"Valute" bson:"-"`
	Name     string   `xml:"Name" bson:"name"`
	Value    float64  `xml:"Value" bson:"value"`
	CharCode string   `xml:"CharCode" bson:"char_code"`
	Nominal  float64  `xml:"Nominal" bson:"nominal"`
	Date     int64    `xml:"-" bson:"date"`
	ID       string   `xml:"-" bson:"_id"`
}

var url = "http://www.cbr.ru/scripts/XML_daily.asp?date_req="

func main() {
	timeStart := time.Date(2015, time.Month(1), 1, 0, 0, 0, 0, time.Local)
	timeStop := time.Date(2020, time.Month(11), 20, 0, 0, 0, 0, time.Local)

	cwt, _ := context.WithTimeout(context.Background(), time.Second*10)
	clientOptions := options.Client().ApplyURI(os.Getenv("LOADER_DB"))
	client, err := mongo.Connect(cwt, clientOptions)
	if err := client.Ping(cwt, readpref.Primary()); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	log.Println("Mongo connected")
	for !timeStart.Equal(timeStop) {
		urlParse := url + timeStart.Format("01/02/2006")
		fmt.Println("----------")
		fmt.Printf("Parse %v\n", urlParse)
		response, err := http.Get(urlParse)
		if err != nil {
			fmt.Println(err)
		}
		var r Resp
		d := xml.NewDecoder(response.Body)
		d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
			switch charset {
			case "windows-1251":
				return charmap.Windows1251.NewDecoder().Reader(input), nil
			default:
				return nil, fmt.Errorf("unknown charset: %s", charset)
			}
		}
		err = d.Decode(&r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Результат парсинга: %v\n", r)
		for _, val := range r.Valute {
			value, err := strconv.ParseFloat(strings.ReplaceAll(val.Value, ",", "."), 64)
			if err != nil {
				fmt.Printf("Error str conf %v\n", err)
			}
			obj := ValuteToDB{
				ID:       uuid.New().String(),
				Name:     val.Name,
				CharCode: val.CharCode,
				Nominal:  val.Nominal,
				Date:     timeStart.Unix(),
				Value:    value,
			}
			_, err = client.Database(database).Collection(collection).InsertOne(context.Background(), obj)
			if err != nil {
				panic(err)
			}
		}
		timeStart.Add(24 * time.Hour)
		fmt.Println("----------")
	}
}
