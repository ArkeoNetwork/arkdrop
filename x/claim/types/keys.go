package types

const (
	// ModuleName defines the module name
	ModuleName = "claim"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_claim"

	// ClaimRecordsStorePrefix defines the store prefix for the claim records (by arkeo address)
	ClaimRecordsArkeoStorePrefix = "claimrecordsarkeo"

	// ClaimRecordsStorePrefix defines the store prefix for the claim records (by eth address)
	ClaimRecordsEthStorePrefix = "claimrecordsethereum"

	// ClaimRecordsStorePrefix defines the store prefix for the claim records (by thor address)
	ClaimRecordsThorStorePrefix = "claimrecordsthorchain"

	// ParamsKey defines the store key for claim module parameters
	ParamsKey = "params"

	// ActionKey defines the store key to store user accomplished actions
	ActionKey = "action"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
