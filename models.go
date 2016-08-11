package gografana

import (
	"net/http"
	"net/url"
)

type Client struct {
	ApiKey   string
	BaseURL  *url.URL
	User     string
	Password string
	*http.Client
}

type DashboardUploader struct {
	Dashboard *Dashboard `json:"dashboard"`
	Overwrite bool       `json:"overwrite"`
}

type DashboardRes struct {
	Meta      *Meta      `json:"meta"`
	Dashboard *Dashboard `json:"dashboard"`
}

type Meta struct {
	IsStarred bool   `json:"isStarred,omitempty"`
	Type      string `json:"type,omitempty"`
	CanSave   bool   `json:"canSave,omitempty"`
	CanEdit   bool   `json:"canEdit,omitempty"`
	CanStar   bool   `json:"canStar,omitempty"`
	Slug      string `json:"slug,omitempty"`
	Expires   string `json:"expires,omitempty"`
	Created   string `json:"created,omitempty"`
	Updated   string `json:"updated,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	Version   int64  `json:"version,omitempty"`
}

type Dashboard struct {
	Annotations     Annotations `json:"annotations,omitempty"`
	Editable        bool        `json:"editable,omitempty"`
	GnetId          int64       `json:"gnetId"`
	HideControls    bool        `json:"hideControls"`
	ID              int64       `json:"id,omitempty"`
	Links           []string    `json:"links"`
	OriginalTitle   string      `json:"originalTitle,omitempty"`
	Refresh         string      `json:"refresh,omitempty"`
	Rows            []*Row      `json:"rows,omitempty"`
	SchemaVersion   int64       `json:"schemaVersion,omitempty"`
	SharedCrosshair bool        `json:"sharedCrosshair,omitempty"`
	Style           string      `json:"style,omitempty"`
	Tags            []string    `json:"tags,omitempty"`
	Templating      Templating  `json:"templating,omitempty"`
	Time            Time        `json:"time,omitempty"`
	Timepicker      Timepicker  `json:"timepicker,omitempty"`
	Timezone        string      `json:"timezone,omitempty"`
	Title           string      `json:"title,omitempty"`
	Version         int64       `json:"version,omitempty"`
}

type Annotations struct {
	Enable bool     `json:"enable,omitempty"`
	List   []string `json:"list"`
}

type Row struct {
	Collapse bool     `json:"collapse"`
	Editable bool     `json:"editable,omitempty"`
	Height   string   `json:"height,omitempty"`
	Panels   []*Panel `json:"panels,omitempty"`
	Title    string   `json:"title,omitempty"`
}

type Panel struct {
	AliasColors     AliasColors `json:"aliasColors"`
	Bars            bool        `json:"bars"`
	Datasource      string      `json:"datasource,omitempty"`
	Editable        bool        `json:"editable,omitempty"`
	Error           bool        `json:"error,omitempty"`
	Fill            int64       `json:"fill,omitempty"`
	Grid            Grid        `json:"grid,omitempty"`
	ID              int64       `json:"id,omitempty"`
	IsNew           bool        `json:"isNew,omitempty"`
	Legend          Legend      `json:"legend,omitempty"`
	Lines           bool        `json:"lines,omitempty"`
	Linewidth       int64       `json:"linewidth,omitempty"`
	Links           []string    `json:"links,omitempty"`
	NullPointMode   string      `json:"nullPointMode,omitempty"`
	Percentage      bool        `json:"percentage"`
	Pointradius     int64       `json:"pointradius,omitempty"`
	Points          bool        `json:"points"`
	Renderer        string      `json:"renderer,omitempty"`
	SeriesOverrides []string    `json:"seriesOverrides"`
	Span            int64       `json:"span,omitempty"`
	Stack           bool        `json:"stack"`
	SteppedLine     bool        `json:"steppedLine"`
	Targets         []Target    `json:"targets,omitempty"`
	TimeFrom        string      `json:"timeFrom,omitempty"`
	TimeShift       string      `json:"timeShift,omitempty"`
	Title           string      `json:"title,omitempty"`
	Tooltip         Tooltip     `json:"tooltip,omitempty"`
	Type            string      `json:"type,omitempty"`
	Xaxis           Xaxis       `json:"xaxis,omitempty"`
	Yaxes           []Yax       `json:"yaxes,omitempty"`
}

type AliasColors struct{}

type Grid struct {
	Threshold1      string `json:"threshold1"`
	Threshold1Color string `json:"threshold1Color,omitempty"`
	Threshold2      string `json:"threshold2"`
	Threshold2Color string `json:"threshold2Color,omitempty"`
}

type Legend struct {
	Avg     bool `json:"avg"`
	Current bool `json:"current"`
	Max     bool `json:"max"`
	Min     bool `json:"min"`
	Show    bool `json:"show"`
	Total   bool `json:"total"`
	Values  bool `json:"values"`
}

type Target struct {
	Alias        string    `json:"alias,omitempty"`
	DsType       string    `json:"dsType,omitempty"`
	GroupBy      []GroupBy `json:"groupBy,omitempty"`
	Measurement  string    `json:"measurement,omitempty"`
	Policy       string    `json:"policy,omitempty"`
	Query        string    `json:"query,omitempty"`
	RawQuery     bool      `json:"rawQuery,omitempty"`
	RefId        string    `json:"refId,omitempty"`
	ResultFormat string    `json:"resultFormat,omitempty"`
	Select       []Selects `json:"select,omitempty"`
	Tags         []Tag     `json:"tags,omitempty"`
	Transform    string    `json:"transform,omitempty"`
}

type GroupBy struct {
	Params   []string `json:"params,omitempty"`
	Type     string   `json:"type,omitempty"`
	Interval string   `json:"interval,omitempty"`
}

type Selects []Select

type Select struct {
	Type   string   `json:"type,omitempty"`
	Params []string `json:"params,omitempty"`
}

type Tag struct {
	Condition string `json:"condition,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Tooltip struct {
	MsResolution bool   `json:"msResolution,omitempty"`
	Shared       bool   `json:"shared,omitempty"`
	Value_type   string `json:"value_type,omitempty"`
}

type Xaxis struct {
	Show bool `json:"show,omitempty"`
}

type Yax struct {
	Format  string `json:"format,omitempty"`
	Lable   string `json:"lable,omitempty"`
	LogBase int64  `json:"logBase,omitempty"`
	Max     int64  `json:"max,omitempty"`
	Min     int64  `json:"min,omitempty"`
	Show    bool   `json:"show,omitempty"`
}

type Time struct {
	From string `json:"from,omitempty"`
	Now  bool   `json:"now,omitempty"`
	To   string `json:"to,omitempty"`
}

type Templating struct {
	List []Template `json:"list,omitempty"`
}

type Template struct {
	AllFormat string `json:"allFormat,omitempty"`
	Current   struct {
		Tags  []interface{} `json:"tags,omitempty"`
		Text  string        `json:"text,omitempty"`
		Value interface{}   `json:"value,omitempty"`
	} `json:"current,omitempty"`
	Datasource  string `json:"datasource,omitempty"`
	IncludeAll  bool   `json:"includeAll,omitempty"`
	Multi       bool   `json:"multi,omitempty"`
	MultiFormat string `json:"multiFormat,omitempty"`
	Name        string `json:"name,omitempty"`
	Options     []struct {
		Selected bool   `json:"selected,omitempty"`
		Text     string `json:"text,omitempty"`
		Value    string `json:"value,omitempty"`
	} `json:"options,omitempty"`
	Query         string `json:"query,omitempty"`
	Refresh       int    `json:"refresh,omitempty"`
	RefreshOnLoad bool   `json:"refresh_on_load,omitempty"`
	Regex         string `json:"regex,omitempty"`
	Type          string `json:"type,omitempty"`
}

type Timepicker struct {
	Refresh_intervals []string `json:"refresh_intervals,omitempty"`
	Time_options      []string `json:"time_options,omitempty"`
}
