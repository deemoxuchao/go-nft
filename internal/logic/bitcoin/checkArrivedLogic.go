package bitcoin

import (
	"context"

	"nft/internal/svc"
	"nft/internal/types"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckArrivedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type Address struct {
	Address string `json:"address"`
	Balance int64  `json:"balance"`
}

func NewCheckArrivedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckArrivedLogic {
	return &CheckArrivedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckArrivedLogic) CheckArrived(req *types.CheckArrivedReq) (resp *types.CheckArrivedResp, err error) {
	// todo: add your logic here and delete this line
	address := req.Address
	url := fmt.Sprintf("https://api.blockcypher.com/v1/btc/main/addrs/%s/balance", address)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var addr Address
	err = json.Unmarshal(body, &addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Balance of address %s: %d\n", req.Address, addr.Balance)

	arrived := addr.Balance >= req.Amount

	return &types.CheckArrivedResp{
		Address: req.Address,
		Arrived: arrived,
		Balance: addr.Balance,
	}, nil
}
