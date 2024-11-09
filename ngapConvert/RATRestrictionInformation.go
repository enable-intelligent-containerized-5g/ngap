package ngapConvert

import (
	"github.com/enable-intelligent-containerized-5g/openapi/models"
	"github.com/free5gc/aper"
	"github.com/enable-intelligent-containerized-5g/ngap/ngapType"
)

// TS 38.413 9.3.1.85
func RATRestrictionInformationToNgap(ratType models.RatType) (ratResInfo ngapType.RATRestrictionInformation) {
	// Only support EUTRA & NR in version15.2.0
	switch ratType {
	case models.RatType_EUTRA:
		ratResInfo.Value = aper.BitString{
			Bytes:     []byte{0x80},
			BitLength: 8,
		}
	case models.RatType_NR:
		ratResInfo.Value = aper.BitString{
			Bytes:     []byte{0x40},
			BitLength: 8,
		}
	}
	return
}
