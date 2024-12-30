package repository

import (
	"database/sql"
	"time"
)

type BhDataLevel struct {
	Date              time.Time       `json:"date"`
	Level             string          `json:"level"`
	LevelName         string          `json:"level_name"`
	ReferenceName     string          `json:"reference_name"`
	IohDataTraffic4g  sql.NullFloat64 `json:"ioh_data_traffic_4g"`
	Availability      sql.NullFloat64 `json:"availability"`
	Eut               sql.NullFloat64 `json:"eut"`
	Cqi               sql.NullFloat64 `json:"cqi"`
	Prb               sql.NullFloat64 `json:"prb"`
	Accesibility      sql.NullFloat64 `json:"accesibility"`
	S1Sr              sql.NullFloat64 `json:"s1_sr"`
	ErabDrop          sql.NullFloat64 `json:"erab_drop"`
	RrcSr             sql.NullFloat64 `json:"rrc_sr"`
	ErabSr            sql.NullFloat64 `json:"erab_sr"`
	InterFreq         sql.NullFloat64 `json:"inter_freq"`
	IntraFreq         sql.NullFloat64 `json:"intra_freq"`
	IratHosr          sql.NullFloat64 `json:"irat_hosr"`
	UlRssi            sql.NullFloat64 `json:"ul_rssi"`
	SeDl              sql.NullFloat64 `json:"se_dl"`
	CsfbPrepSr        sql.NullFloat64 `json:"csfb_prep_sr"`
	CsfbSr            sql.NullFloat64 `json:"csfb_sr"`
	IohVolteTraffic4g sql.NullFloat64 `json:"ioh_volte_traffic_4g"`
	Rank2Mimo         sql.NullFloat64 `json:"rank_2_mimo"`
}

type KabKota struct {
	GID    int     `json:"gid"`
	City   *string `json:"city"`
	Region *string `json:"region"`
	Circle *string `json:"circle"`
	Geom   *string `json:"geom,omitempty"`
}
