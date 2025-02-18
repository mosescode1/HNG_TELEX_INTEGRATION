package data

import (
	"github.com/mosescode1/types"
)


func GetIntegrationJson() types.RootJson {


	return types.RootJson{
		Data: types.Data{
			Author: "Eteng Moses Efa",
			Date: types.Date{
				CreatedAt: "2025-02-13",
				UpdatedAt: "2025-02-13",
			},
			Descriptions: types.Descriptions{
				AppDescription:  "Fiber Logs is a Telex integration that formats messages based on predefined templates or and send back the server last 10 logs.",
				AppURL: "https://5gsd8sxs-3000.uks1.devtunnels.ms",
				AppLogo:        "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ09bQSfOMsbzr0ObHC6VtFpBtdvbBEmLY7fQ&s",
				AppName:         "Fiber Logs",
				BackgroundColor: "#ffffff",
			},
			IntegrationCategory: "Monitoring & Logging",
			IntegrationType:     "modifier",
			IsActive:            true,
			KeyFeatures: []string{
				"Receive messages from Telex channels.",
				"Format messages based on predefined templates or logic.",
				"Send Server Logs back to the channel.",
			},
			Settings: []types.Setting{
				{Default: 100, Label: "maxMessageLength", Required: true, Type: "text"},
				{Default: "world,happy", Label: "repeatWords", Required: true, Type: "text"},
				{Default: 2, Label: "noOfRepetitions", Required: true, Type: "text"},
			},
			TargetURL: "https://5gsd8sxs-3000.uks1.devtunnels.ms/webhook",
		},
	}
}
