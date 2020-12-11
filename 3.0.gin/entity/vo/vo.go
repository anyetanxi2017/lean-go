package vo

type WalletInfoVo struct {
	Uid   int     `json:"uid"`
	Money float32 `json:"money"`
}
type WalletVo struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Money    float32 `json:"money"`
}
