package logic

type ResponseJSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

const Secretkey = "bookMyShowJWT"
