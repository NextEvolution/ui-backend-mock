package types

type FbLoginReq struct {
	FbToken string `json:"fb_token"`
}

type Config struct {
	Port int `json:"port"`
	StaticFilePath string `json:"static_file_path"`
	Origins []string `json:"origins"`
}