package models

type Trade struct {
	Id       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Symbol   string  `json:"symbol"`
	Time     int64   `json:"time"`
}

type CommonResponse struct {
	RetCode    int    `json:"retCode"`
	RetMsg     string `json:"retMsg"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type WalletBalance struct {
	CommonResponse
	Result struct {
		List []struct {
			TotalEquity            string `json:"totalEquity"`
			AccountIMRate          string `json:"accountIMRate"`
			TotalMarginBalance     string `json:"totalMarginBalance"`
			TotalInitialMargin     string `json:"totalInitialMargin"`
			AccountType            string `json:"accountType"`
			TotalAvailableBalance  string `json:"totalAvailableBalance"`
			AccountMMRate          string `json:"accountMMRate"`
			TotalPerpUPL           string `json:"totalPerpUPL"`
			TotalWalletBalance     string `json:"totalWalletBalance"`
			AccountLTV             string `json:"accountLTV"`
			TotalMaintenanceMargin string `json:"totalMaintenanceMargin"`
			Coin                   []struct {
				AvailableToBorrow   string `json:"availableToBorrow"`
				Bonus               string `json:"bonus"`
				AccruedInterest     string `json:"accruedInterest"`
				AvailableToWithdraw string `json:"availableToWithdraw"`
				TotalOrderIM        string `json:"totalOrderIM"`
				Equity              string `json:"equity"`
				TotalPositionMM     string `json:"totalPositionMM"`
				UsdValue            string `json:"usdValue"`
				UnrealisedPnl       string `json:"unrealisedPnl"`
				CollateralSwitch    bool   `json:"collateralSwitch"`
				SpotHedgingQty      string `json:"spotHedgingQty"`
				BorrowAmount        string `json:"borrowAmount"`
				TotalPositionIM     string `json:"totalPositionIM"`
				WalletBalance       string `json:"walletBalance"`
				CumRealisedPnl      string `json:"cumRealisedPnl"`
				Locked              string `json:"locked"`
				MarginCollateral    bool   `json:"marginCollateral"`
				Coin                string `json:"coin"`
			} `json:"coin"`
		} `json:"list"`
	} `json:"result"`
}

type DepositAddressResponse struct {
	CommonResponse
	Result struct {
		Coin   string `json:"coin"`
		Chains struct {
			ChainType         string `json:"chainType"`
			AddressDeposit    string `json:"addressDeposit"`
			TagDeposit        string `json:"tagDeposit"`
			Chain             string `json:"chain"`
			BatchReleaseLimit string `json:"batchReleaseLimit"`
		} `json:"chains"`
	} `json:"result"`
}

type AccountCoinBalance struct {
	CommonResponse
	Result struct {
		MemberID    string `json:"memberId"`
		AccountType string `json:"accountType"`
		Balance     []struct {
			Coin            string `json:"coin"`
			TransferBalance string `json:"transferBalance"`
			WalletBalance   string `json:"walletBalance"`
			Bonus           string `json:"bonus"`
		} `json:"balance"`
	} `json:"result"`
}

type WithdrawalRecords struct {
	CommonResponse
	Result struct {
		Rows []struct {
			Coin              string `json:"coin"`
			Chain             string `json:"chain"`
			Amount            string `json:"amount"`
			TxID              string `json:"txID"`
			Status            string `json:"status"`
			ToAddress         string `json:"toAddress"`
			Tag               string `json:"tag"`
			WithdrawFee       string `json:"withdrawFee"`
			CreateTime        string `json:"createTime"`
			UpdateTime        string `json:"updateTime"`
			WithdrawID        string `json:"withdrawId"`
			WithdrawType      int    `json:"withdrawType"`
			Confirmations     string `json:"confirmations"`
			TxIndex           string `json:"txIndex"`
			BlockHash         string `json:"blockHash"`
			BatchReleaseLimit string `json:"batchReleaseLimit"`
			DepositType       string `json:"depositType"`
		} `json:"rows"`
		NextPageCursor string `json:"nextPageCursor"`
	} `json:"result"`
}

type DepositRecords struct {
	CommonResponse
	Result struct {
		Rows []struct {
			Coin              string `json:"coin"`
			Chain             string `json:"chain"`
			Amount            string `json:"amount"`
			TxID              string `json:"txID"`
			Status            string `json:"status"`
			ToAddress         string `json:"toAddress"`
			Tag               string `json:"tag"`
			DepositFee        string `json:"withdrawFee"`
			CreateTime        string `json:"createTime"`
			UpdateTime        string `json:"updateTime"`
			WithdrawID        string `json:"withdrawId"`
			WithdrawType      int    `json:"withdrawType"`
			Confirmations     string `json:"confirmations"`
			TxIndex           string `json:"txIndex"`
			BlockHash         string `json:"blockHash"`
			BatchReleaseLimit string `json:"batchReleaseLimit"`
			DepositType       string `json:"depositType"`
		} `json:"rows"`
		NextPageCursor string `json:"nextPageCursor"`
	} `json:"result"`
}

type WithdrawableAmount struct {
	CommonResponse
	Result struct {
		LimitAmountUsd     string `json:"limitAmountUsd"`
		WithdrawableAmount struct {
			FUND struct {
				Coin               string `json:"coin"`
				WithdrawableAmount string `json:"withdrawableAmount"`
				AvailableBalance   string `json:"availableBalance"`
			} `json:"FUND"`
			SPOT struct {
				Coin               string `json:"coin"`
				WithdrawableAmount string `json:"withdrawableAmount"`
				AvailableBalance   string `json:"availableBalance"`
			} `json:"SPOT"`
		} `json:"withdrawableAmount"`
	} `json:"result"`
}
