package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	// ctx := context.Background()
	ctx := context.WithValue(context.Background(), "foo", "bar")
	userId := 10
	val, err := fetchUserData(ctx, userId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
	fmt.Println("Took: ", time.Since(start))

}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {

	val := ctx.Value("foo")
	fmt.Println(val.(string))
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200) // this will timeout after certain duration
	defer cancel()

	respch := make(chan Response)
	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took to long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 180)
	return 666, nil
}
