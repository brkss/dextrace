package domain

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Timestamp int64  `json:"timestamp"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      struct {
		AccessToken string `json:"access_token"`
		ExpiresIn  int    `json:"expires_in"`
	} `json:"data"`
	ErrorData interface{} `json:"errorData"`
	Success   bool        `json:"success"`
}

type GlucoseInfo struct {
	I    int     `json:"i"`
	T    string  `json:"t"` // record time 
	V    float64 `json:"v"` // blood glucose value mmol/l
	S    int     `json:"s"`
	Ast  int     `json:"ast"`
	Bl   float64 `json:"bl"`
	Name *string `json:"name"`
}

type Target struct {
	Upper  int `json:"upper"`
	Lower  int `json:"lower"`
	IsRec  int `json:"isRec"`
	DrType int `json:"drType"`
}

type GlucoseData struct {
	UserId             string        `json:"userId"`
	UserNickName       string        `json:"userNickName"`
	DrType             int           `json:"drType"`
	DeviceId           string        `json:"deviceId"`
	DeviceName         string        `json:"deviceName"`
	BlueToothNum       string        `json:"blueToothNum"`
	DeviceEnableTime   string        `json:"deviceEnableTime"`
	DeviceStatus       int           `json:"deviceStatus"`
	DeviceAlarmStatus  int           `json:"deviceAlarmStatus"`
	LatestIndex        int           `json:"latestIndex"`
	LatestGlucoseValue float64       `json:"latestGlucoseValue"`
	BloodGlucoseTrend  int           `json:"bloodGlucoseTrend"`
	LatestGlucoseTime  string        `json:"latestGlucoseTime"`
	DeviceAbnormalTime *string       `json:"deviceAbnormalTime"`
	DeviceLastTime     string        `json:"deviceLastTime"`
	GlucoseInfos       []GlucoseInfo `json:"glucoseInfos"`
	Target             Target        `json:"target"`
}

type GlucoseDataResponse struct {
	Timestamp int64       `json:"timestamp"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      GlucoseData `json:"data"`
	ErrorData interface{} `json:"errorData"`
	Success   bool        `json:"success"`
}

type GetDataResponse struct {
	Timestamp string       `json:"timestamp"`
	Value     float64     `json:"value"`
}

// nightscout config 
type NightscoutConfig struct {
	Token 			string
	NightscoutURL 	string
}
// nightscout glucose entry 
type NightscoutEntry struct {
	Type      string  `json:"type"`      // always "sgv"
	SGV       float64 `json:"sgv"`       // glucose value
	Date      int64   `json:"date"`      // milliseconds since epoch
	DateString string `json:"dateString"`// ISO8601 string
}
