package models

/**
 * created on 10/10/18.
 * author: nebula-ai-chengkun
 * Copyright defined in blockchainwebbrowser/LICENSE.txt
 */

type WalletRequest struct {
	TxHash    string `json:"tx_hash"`
	Value     string `json:"value"`
	TFrom     string `json:"t_from"`
	TTo       string `json:"t_to"`
	Timestamp string `json:"timestamp"`
}
