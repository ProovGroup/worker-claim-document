package veos

import (
	"fmt"

	"github.com/ProovGroup/lib-claim-models/report"
	"github.com/ProovGroup/worker-claim-declaration/internal/pkg/veos/model"
	"github.com/ProovGroup/worker-claim-declaration/internal/provider"
)

var (
	BASE_URL       = provider.Getenv(provider.API_VEOS_URL)
	ROUTE_DOCUMENT = BASE_URL + "/document"
)

func Handle(proovCode string) error {
	report, err := report.GetReport(provider.ClaimDB, proovCode)
	if err != nil {
		fmt.Println("[ERROR] report.GetReport:", err)
		return err
	}

	fmt.Println("[DEBUG] veos.Handle finished")

	return nil
}
