package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
)

type Str struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

func Moon(res http.ResponseWriter, req *http.Request) {
	mug := Str{}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &mug)
	fmt.Println(mug)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(mug)
}

func generateRandomNumber(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	// Never likely to occur, kust for safe side.
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}
func main() {
	for i := 0; i < 10; i++ {
		re, _ := generateRandomNumber(15)
		fmt.Println(re)
	}
	maxLimit := int64(int(math.Pow10(10)) - 1)
	fmt.Println(maxLimit)
	// router := mux.NewRouter()
	// router.HandleFunc("/str", Moon)
	// http.ListenAndServe(":4000", router)
}
