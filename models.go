package gografana

type DashBoard struct {
	Annotations   DashBoardAnnotations `json: "annotations"`
	Editable      bool                 `json: "editable"`
	HideControls  bool                 `json: "hideControls"`
	ID            int64                `json: "id"`
	Links         []string             `json: "links"`
	OriginalTitle string               `json: "originalTitle"`
	Refresh       string               `json: "refresh"`
	Rows          []DashBoardRow       `json: "rows"`
	Title         string               `json: "title"`
	Tags          []string             `json: "tags"`
	Timezone      string               `json: "timezone"`
}

type DashBoardAnnotations struct {
	List []string `json: list`
}

type DashBoardRow struct {
	Collapse bool                 `json: "cololapse"`
	Editable bool                 `json: "editable"`
	Height   string               `json: "height"`
	Pannels  []DashBoardRowPannel `json: "pannels"`
}

type DashBoardRowPannel struct {
	AliasColors     DashBoardRowPannelAliasColors `json: "aliasColors"`
	Bars            bool                          `json: "bars"`
	Datasource      string                        `json: "datasource"`
	Editable        bool                          `json: "editable"`
	Error           bool                          `json: "error"`
	Fill            int64                         `json: "fill"`
	Grid            DashBoardRowPannelGrid        `json: "grid"`
	ID              int64                         `json: "id"`
	IsNew           bool                          `json: "isNew"`
	Legend          DashBoardRowPannelLegend      `json: "legend"`
	Lines           bool                          `json: "lines"`
	Linewidth       int64                         `json: "linewidth"`
	Links           []string                      `json: "links"`
	NullPointMode   string                        `json: "nullPointMode"`
	Percentage      bool                          `json: "percentage"`
	Pointradius     int64                         `json: "pointradius"`
	Points          bool                          `json: "points"`
	Renderer        string                        `json: "renderer"`
	SeriesOverrides []string                      `json: "seriesOverrides"`
	Span            int64                         `json: "span"`
	Stack           bool                          `json: "stack"`
	SteppedLine     bool                          `json: "steppedLine"`
	Targets         []DashBoardRowPannelTarget    `json: "targets"`
}

type DashBoardRowPannelAliasColors struct{}

type DashBoardRowPannelGrid struct {
	Threshold1      string `json: "threshold1"`
	Threshold1Color string `json: "threshold1Color"`
	Threshold2      string `json: "threshold2"`
	Threshold2Color string `json: "threshold2Color"`
}

type DashBoardRowPannelLegend struct {
	Avg     bool `json: "avg"`
	Current bool `json: "current"`
	Max     bool `json: "max"`
	Min     bool `json: "min"`
	Show    bool `json: "show"`
	Total   bool `json: "total"`
	Values  bool `json: "values"`
}

type DashBoardRowPannelTarget struct {
}

/*
{


                        "targets": [
                            {
                                "alias": "Kafka",
                                "dsType": "influxdb",
                                "groupBy": [
                                    {
                                        "params": [
                                            "5m"
                                        ],
                                        "type": "time"
                                    },
                                    {
                                        "params": [
                                            "null"
                                        ],
                                        "type": "fill"
                                    }
                                ],
                                "measurement": "m.request",
                                "policy": "default",
                                "refId": "A",
                                "resultFormat": "time_series",
                                "select": [
                                    [
                                        {
                                            "params": [
                                                "Kafka"
                                            ],
                                            "type": "field"
                                        },
                                        {
                                            "params": [ ],
                                            "type": "mean"
                                        }
                                    ]
                                ],
                                "tags": [ ]
                            }
                        ],
                        "timeFrom": null,
                        "timeShift": null,
                        "title": "m.request",
                        "tooltip": {
                            "msResolution": false,
                            "shared": true,
                            "value_type": "cumulative"
                        },
                        "type": "graph",
                        "xaxis": {
                            "show": true
                        },
                        "yaxes": [
                            {
                                "format": "none",
                                "label": null,
                                "logBase": 1,
                                "max": null,
                                "min": null,
                                "show": true
                            },
                            {
                                "format": "short",
                                "label": null,
                                "logBase": 1,
                                "max": null,
                                "min": null,
                                "show": true
                            }
                        ]
                    }
                ],
                "title": "New row"
            }
        ],
        "schemaVersion": 12,
        "sharedCrosshair": false,
        "style": "dark",
        "tags": [ ],
        "templating": {
            "list": [ ]
        },
        "time": {
            "from": "now-3h",
            "to": "now"
        },
        "timepicker": {
            "refresh_intervals": [
                "5s",
                "10s",
                "30s",
                "1m",
                "5m",
                "15m",
                "30m",
                "1h",
                "2h",
                "1d"
            ],
            "time_options": [
                "5m",
                "15m",
                "1h",
                "6h",
                "12h",
                "24h",
                "2d",
                "7d",
                "30d"
            ]
        },
        "timezone": "utc",
        "title": "Market By Minute",
        "version": 162
    }
}

*/
