package ngapConvert

import (
	"github.com/enable-intelligent-containerized-5g/openapi/models"
	"github.com/enable-intelligent-containerized-5g/ngap/ngapType"
)

func AllowedNssaiToNgap(allowedNssaiModels []models.AllowedSnssai) (allowedNssaiNgap ngapType.AllowedNSSAI) {
	for _, allowedSnssai := range allowedNssaiModels {
		item := ngapType.AllowedNSSAIItem{
			SNSSAI: SNssaiToNgap(*allowedSnssai.AllowedSnssai),
		}
		allowedNssaiNgap.List = append(allowedNssaiNgap.List, item)
	}
	return
}

func AllowedNssaiToModels(allowedNssaiNgap ngapType.AllowedNSSAI) (allowedNssaiModels []models.AllowedSnssai) {
	for _, item := range allowedNssaiNgap.List {
		snssai := SNssaiToModels(item.SNSSAI)
		allowedSnssai := models.AllowedSnssai{
			AllowedSnssai: &snssai,
		}
		allowedNssaiModels = append(allowedNssaiModels, allowedSnssai)
	}
	return
}
