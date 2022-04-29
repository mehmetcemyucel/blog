package request

type ExampleReq struct {
	Hash string `json:"hash" minLength:"4" maxLength:"16"`
	Name string `json:"name" minLength:"4" maxLength:"16"`
}
