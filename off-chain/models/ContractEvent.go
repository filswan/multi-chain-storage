package models

type ContractEventArray struct {
	MinerCreditArray []string `json:"miner_credit_array"`
	TaskManagerArray []string `json:"task_manager_array"`
	DepositArray     []string `json:"deposit_array"`
}
