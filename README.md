# graf

[![Build Status](https://travis-ci.org/qianlnk/graf.svg?branch=master)](https://travis-ci.org/qianlnk/graf)

# Install

```golang
go get -u github.com/qianlnk/graf
go install github.com/qianlnk/graf/cmd/graf
echo > $GOPATH/bin/config.json
```

# Config

get `apiKey` from your grafana window. and `baseUrl` is your.grafana.com
 
```golang
vi $GOPATH/bin/config.json

{
	"apiKey":"",
	"baseUrl":"",
	"user":"",
	"password":""
}

```

# Use

```golang
graf -h
```
# Template

make your Template panel at grafana, then you can export it:

```golang
graf export -t "dashboard title" -p <the Template panel id you have made> -n <Template name>
```
or edit the Template json file:

```golang
{
    "aliasColors": {},
    "bars": false,
    "datasource": "market",
    "editable": true,
    "fill": 1,
    "grid": {
        "threshold1": "",
        "threshold1Color": "rgba(216, 200, 27, 0.27)",
        "threshold2": "",
        "threshold2Color": "rgba(234, 112, 112, 0.22)"
    },
    "id": 0,
    "isNew": true,
    "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
    },
    "lines": true,
    "linewidth": 2,
    "nullPointMode": "connected",
    "percentage": false,
    "pointradius": 5,
    "points": false,
    "renderer": "flot",
    "seriesOverrides": [],
    "span": 12,
    "stack": false,
    "steppedLine": false,
    "title": "default",
    "tooltip": {
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
            "logBase": 1,
            "show": true
        },
        {
            "format": "short",
            "logBase": 1
        }
    ]
}
```
# Create dashboard

a json file need, like
```golang
{
	"title":"default",
	"timezone":"utc",
	"templateType":"lacal",
	"inheritTitle":"",
	"inheritPanelid":0,
	"templateFile":"templates/default.json",
	"rows":[
		{
			"title":"New Row",
			"panels":[
				{
					"title":"default",
					"dsType":"infuxdb",
					"sqls":[
						"",
						""
					]
				}
			]
		}
	]
}
```

then run cmd:
```golang
graf -f <name>.json
```