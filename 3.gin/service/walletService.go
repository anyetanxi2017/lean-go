package walletService

import (
	"lean-go/3.gin/entity/vo"
)

func GetInfo() vo.WalletInfoVo {
	return vo.WalletInfoVo{Uid: 1, Money: 20}
}
func ListWallet() []vo.WalletVo {
	list := []vo.WalletVo{
		{1, "yy", 1000},
		{1, "sky", 1000},
	}
	return list
}
