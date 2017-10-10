package main

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	"encoding/json"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB = "square_holes"
	username = "bubba"
	password = "bumblebeetuna"
)

func writePoints(cInt client.Client) {
	sampleSize := 1000

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "systemstats",
		Precision: "us",
	})
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sampleSize; i++ {
		regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
		tags := map[string]string{
			"cpu": "cpu-total",
			"host": fmt.Sprintf("host%d", rand.Intn(1000)),
			"region": regions[rand.Intn(len(regions))],
		}

		idle := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"idle": idle,
			"busy": 100.0 - idle,
		}

		pt, err := client.NewPoint(
			"cpu_usage",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	if err := cInt.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func queryDB(cInt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command: cmd,
		Database: "systemstats",
	}
	if response, err := cInt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}

	return res, nil
}

func WriteUDP() {
	// Make client
	c, err := client.NewUDPClient("localhost:8089")
	if err != nil {
		panic(err.Error())
	}

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		panic(err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
}

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	/*
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	writePoints(c)
    */


	/*
	q := fmt.Sprintf("SELECT COUNT(%s) FROM %s", "busy", "cpu_usage")
	res, err := queryDB(c, q)
	if err != nil {
		log.Fatal(err)
	}
	count := res[0].Series[0].Values[0][1]
	log.Printf("Found a total of %v records %V\n", count, res)
    */

	q := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	res, err := queryDB(c, q)
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		val := string(row[1].(json.Number))
		//log.Printf("[%2d] %s: %v\n", i, t.Format(time.Stamp), val)
		//log.Printf("[%2d] %s: %T-%v", i, t.Format(time.Stamp), row[1], row[1])

		log.Printf("[%2d] %s %T %v %T %v\n", i, t.Format(time.Stamp), row[1], row[1], val, val)
	}


	WriteUDP()
}
