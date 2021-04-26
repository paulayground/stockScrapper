package main

import (
	common "github.com/mznx0192/stockScrapper/common"
	"github.com/mznx0192/stockScrapper/request"
)

func main() {
	// get Env
	url := common.GetEnv("URL")

	// Get Request
	request.GetStockList(url)
}
