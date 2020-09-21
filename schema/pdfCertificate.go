package schema

//Information is certificate information struct
type Information struct {
	ID                string `json:"no" description:""`
	Color             string `json:"color" description:""`
	Transparency      string `json:"transparency" description:""`
	Shape             string `json:"shape" description:""`
	Measurement       string `json:"measurement" description:""`
	Weight            string `json:"weight" description:""`
	ViewCommentFirst  string `json:"comment" description:""`
	ViewCommentSecond string `json:"comment1" description:""`
	ViewCommentThird  string `json:"comment2" description:""`
	Hardness          string `json:"hardness" description:""`
	SpecificGravity   string `json:"specific_gravity" description:""`
	RefractiveIndex   string `json:"refractive_index" description:""`
	PolarisCope       string `json:"polariscope" description:""`
	Fluorescence      string `json:"fluorescence" description:""`
	Magnification     string `json:"magnification" description:""`
	AnalysisComment   string `json:"remark" description:""`
	ChineseConclusion string `json:"conclusion" description:""`
	EnglishConclusion string `json:"conclusion1" description:""`
	DataImage         string `json:"image" description:""`
	Image             string `json:"image2" description:""`
	QRCodeURL         string `json:"qrcode" description:"QRCod"`
}

// ReplyInformation is return certificate information
type ReplyInformation struct {
	Result []Information `json:"result" description:"回傳一筆"`
}

//InformationX is certificate information struct
type InformationX struct {
	ColorX             float64
	TransparencyX      float64
	ShapeX             float64
	MeasurementX       float64
	WeightX            float64
	HardnessX          float64
	SpecificGravityX   float64
	RefractiveIndexX   float64
	PolarisCopeX       float64
	FluorescenceX      float64
	MagnificationX     float64
	AnalysisCommentX   float64
	ChineseConclusionX float64
	EnglishConclusionX float64
	ViewCommentFirstX  float64
	Point              float64
}
