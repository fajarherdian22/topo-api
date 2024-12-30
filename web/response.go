package web

import (
	"database/sql"
	"time"

	"github.com/fajarherdian22/topo-api/repository"
)

type ListLevel struct {
	LevelName string `json:"level_name"`
}

func NewListLevels(name []string) []ListLevel {
	var resp []ListLevel
	for _, names := range name {
		resp = append(resp, ListLevel{LevelName: names})
	}
	return resp
}

type RanResponse struct {
	Date              string      `json:"date"`
	Level             string      `json:"level"`
	LevelName         string      `json:"level_name"`
	ReferenceName     string      `json:"reference_name"`
	IohDataTraffic4g  interface{} `json:"ioh_data_traffic_4g"`
	Availability      interface{} `json:"availability"`
	Eut               interface{} `json:"eut"`
	Cqi               interface{} `json:"cqi"`
	Prb               interface{} `json:"prb"`
	Accesibility      interface{} `json:"accesibility"`
	S1Sr              interface{} `json:"s1_sr"`
	ErabDrop          interface{} `json:"erab_drop"`
	RrcSr             interface{} `json:"rrc_sr"`
	ErabSr            interface{} `json:"erab_sr"`
	InterFreq         interface{} `json:"inter_freq"`
	IntraFreq         interface{} `json:"intra_freq"`
	IratHosr          interface{} `json:"irat_hosr"`
	UlRssi            interface{} `json:"ul_rssi"`
	SeDl              interface{} `json:"se_dl"`
	CsfbPrepSr        interface{} `json:"csfb_prep_sr"`
	CsfbSr            interface{} `json:"csfb_sr"`
	IohVolteTraffic4g interface{} `json:"ioh_volte_traffic_4g"`
	Rank2Mimo         interface{} `json:"rank_2_mimo"`
}

// ConvertRanToWebResponse converts a BhDataLevel object into a RanResponse object.
func ConvertRanToWebResponse(data repository.BhDataLevel) RanResponse {
	return RanResponse{
		Date:              formatDate(data.Date),
		Level:             data.Level,
		LevelName:         data.LevelName,
		ReferenceName:     data.ReferenceName,
		IohDataTraffic4g:  nullFloatToInterface(data.IohDataTraffic4g),
		Availability:      nullFloatToInterface(data.Availability),
		Eut:               nullFloatToInterface(data.Eut),
		Cqi:               nullFloatToInterface(data.Cqi),
		Prb:               nullFloatToInterface(data.Prb),
		Accesibility:      nullFloatToInterface(data.Accesibility),
		S1Sr:              nullFloatToInterface(data.S1Sr),
		ErabDrop:          nullFloatToInterface(data.ErabDrop),
		RrcSr:             nullFloatToInterface(data.RrcSr),
		ErabSr:            nullFloatToInterface(data.ErabSr),
		InterFreq:         nullFloatToInterface(data.InterFreq),
		IntraFreq:         nullFloatToInterface(data.IntraFreq),
		IratHosr:          nullFloatToInterface(data.IratHosr),
		UlRssi:            nullFloatToInterface(data.UlRssi),
		SeDl:              nullFloatToInterface(data.SeDl),
		CsfbPrepSr:        nullFloatToInterface(data.CsfbPrepSr),
		CsfbSr:            nullFloatToInterface(data.CsfbSr),
		IohVolteTraffic4g: nullFloatToInterface(data.IohVolteTraffic4g),
		Rank2Mimo:         nullFloatToInterface(data.Rank2Mimo),
	}
}

func NewRanResponses(ranData []repository.BhDataLevel) []RanResponse {
	resp := make([]RanResponse, len(ranData))
	for i, data := range ranData {
		resp[i] = ConvertRanToWebResponse(data)
	}
	return resp
}

func nullFloatToInterface(value sql.NullFloat64) interface{} {
	if value.Valid {
		return value.Float64
	}
	return nil
}

func formatDate(date time.Time) string {
	return date.Format("2006-01-02")
}
