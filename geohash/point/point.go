package point

type Point struct {
	Lat    float64 `form:"lat" json:"lat" validate:"exists"`
	Lng    float64 `form:"lng" json:"lng" validate:"exists"`
	ErrLat float64 `form:"latErr" json:"latErr"`
	ErrLng float64 `form:"lngErr" json:"lngErr"`
}

const (
	MAX_LATITUDE  = 90
	MIN_LATITUDE  = -MAX_LATITUDE
	MAX_LONGITUDE = 180
	MIN_LONGITUDE = -MAX_LONGITUDE
)

func NewPoint(lat, lng, errLat, errLng float64) Point {
	return Point{Lat: lat, Lng: lng, ErrLat: errLat, ErrLng: errLng}
}

func (p Point) WithinErr(other Point) bool {
	return p.Lng-p.ErrLng <= other.Lng+other.ErrLng &&
		p.Lng+p.ErrLng >= other.Lng-other.ErrLng &&
		p.Lat-p.ErrLat <= other.Lat+other.ErrLat &&
		p.Lat+p.ErrLat >= other.Lat-other.ErrLat
}

func (p Point) IsValid() bool {
	return p.Lat <= MAX_LATITUDE &&
		p.Lat >= MIN_LATITUDE &&
		p.Lng <= MAX_LONGITUDE &&
		p.Lng >= MIN_LONGITUDE
}
