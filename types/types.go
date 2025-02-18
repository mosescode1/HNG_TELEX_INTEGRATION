package types

type Date struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Descriptions struct {
	AppDescription string `json:"app_description"`
	AppLogo        string `json:"app_logo"`
	AppName        string `json:"app_name"`
	AppURL         string `json:"app_url"`
	BackgroundColor string `json:"background_color"`
}

type Permission struct {
	Events []string `json:"events"`
}

type Setting struct {
	Default interface{} `json:"default"`
	Label   string      `json:"label"`
	Required bool       `json:"required"`
	Type    string      `json:"type"`
}

type Data struct {
	Author             string       `json:"author"`
	Date               Date         `json:"date"`
	Descriptions       Descriptions `json:"descriptions"`
	IntegrationCategory string      `json:"integration_category"`
	IntegrationType    string      `json:"integration_type"`
	IsActive           bool        `json:"is_active"`
	KeyFeatures        []string    `json:"key_features"`
	Permissions        Permission  `json:"permissions"`
	Settings           []Setting   `json:"settings"`
	TargetURL          string      `json:"target_url"`
	TickURL            string      `json:"tick_url"`
	Website            string      `json:"website"`
}

type RootJson struct {
	Data Data `json:"data"`
}
