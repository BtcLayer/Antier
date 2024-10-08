package constant

const MEMPOOL_HEIGHT = 0x3fffff // 3fffff, 2^22-1

// brc20 protocal
const (
	BRC20_P        = "brc-20"
	BRC20_P_MODULE = "brc20-module"
	BRC20_P_SWAP   = "brc20-swap"
)

// brc20 op
const (
	BRC20_OP_DEPLOY   = "deploy"
	BRC20_OP_MINT     = "mint"
	BRC20_OP_TRANSFER = "transfer"
)

// brc20 history
const (
	BRC20_HISTORY_TYPE_INSCRIBE_DEPLOY   = "inscribe-deploy"
	BRC20_HISTORY_TYPE_INSCRIBE_MINT     = "inscribe-mint"
	BRC20_HISTORY_TYPE_INSCRIBE_TRANSFER = "inscribe-transfer"
	BRC20_HISTORY_TYPE_TRANSFER          = "transfer"
	BRC20_HISTORY_TYPE_SEND              = "send"
	BRC20_HISTORY_TYPE_RECEIVE           = "receive"
)

const (
	// brc20 history N
	BRC20_HISTORY_TYPE_N_INSCRIBE_DEPLOY   uint8 = 0
	BRC20_HISTORY_TYPE_N_INSCRIBE_MINT     uint8 = 1
	BRC20_HISTORY_TYPE_N_INSCRIBE_TRANSFER uint8 = 2
	BRC20_HISTORY_TYPE_N_TRANSFER          uint8 = 3
	BRC20_HISTORY_TYPE_N_SEND              uint8 = 4
	BRC20_HISTORY_TYPE_N_RECEIVE           uint8 = 5

	// swap history N
	BRC20_HISTORY_MODULE_TYPE_N_INSCRIBE_MODULE uint8 = 6

	BRC20_HISTORY_MODULE_TYPE_N_INSCRIBE_WITHDRAW uint8 = 7
	BRC20_HISTORY_MODULE_TYPE_N_WITHDRAW          uint8 = 8
	BRC20_HISTORY_MODULE_TYPE_N_WITHDRAW_FROM     uint8 = 9
	BRC20_HISTORY_MODULE_TYPE_N_WITHDRAW_TO       uint8 = 10

	BRC20_HISTORY_SWAP_TYPE_N_INSCRIBE_APPROVE uint8 = 11
	BRC20_HISTORY_SWAP_TYPE_N_APPROVE          uint8 = 12

	BRC20_HISTORY_SWAP_TYPE_N_INSCRIBE_CONDITIONAL_APPROVE uint8 = 13
	BRC20_HISTORY_SWAP_TYPE_N_CONDITIONAL_APPROVE          uint8 = 14

	BRC20_HISTORY_SWAP_TYPE_N_APPROVE_FROM uint8 = 15
	BRC20_HISTORY_SWAP_TYPE_N_APPROVE_TO   uint8 = 16

	BRC20_HISTORY_SWAP_TYPE_N_INSCRIBE_COMMIT uint8 = 17
	BRC20_HISTORY_SWAP_TYPE_N_COMMIT          uint8 = 18
)

// module op
const (
	BRC20_OP_MODULE_DEPLOY            = "deploy"
	BRC20_OP_MODULE_WITHDRAW          = "withdraw"
	BRC20_OP_SWAP_APPROVE             = "approve"
	BRC20_OP_SWAP_CONDITIONAL_APPROVE = "conditional-approve"
	BRC20_OP_SWAP_COMMIT              = "commit"
)

// swap history
const (
	BRC20_HISTORY_MODULE_TYPE_INSCRIBE_MODULE   = "inscribe-module"
	BRC20_HISTORY_MODULE_TYPE_INSCRIBE_WITHDRAW = "inscribe-withdraw"
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW          = "withdraw"
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW_FROM     = "withdraw-from"
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW_TO       = "withdraw-to"

	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_APPROVE = "inscribe-approve"
	BRC20_HISTORY_SWAP_TYPE_APPROVE          = "approve"

	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_CONDITIONAL_APPROVE = "inscribe-conditional-approve"
	BRC20_HISTORY_SWAP_TYPE_CONDITIONAL_APPROVE          = "conditional-approve"

	BRC20_HISTORY_SWAP_TYPE_APPROVE_FROM = "approve-from"
	BRC20_HISTORY_SWAP_TYPE_APPROVE_TO   = "approve-to"

	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_COMMIT = "inscribe-commit"
	BRC20_HISTORY_SWAP_TYPE_COMMIT          = "commit"
)

var BRC20_HISTORY_TYPE_NAMES []string = []string{
	BRC20_HISTORY_TYPE_INSCRIBE_DEPLOY,
	BRC20_HISTORY_TYPE_INSCRIBE_MINT,
	BRC20_HISTORY_TYPE_INSCRIBE_TRANSFER,
	BRC20_HISTORY_TYPE_TRANSFER,
	BRC20_HISTORY_TYPE_SEND,
	BRC20_HISTORY_TYPE_RECEIVE,

	// module
	BRC20_HISTORY_MODULE_TYPE_INSCRIBE_MODULE,
	BRC20_HISTORY_MODULE_TYPE_INSCRIBE_WITHDRAW,
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW,
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW_FROM,
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW_TO,

	// swap
	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_APPROVE,
	BRC20_HISTORY_SWAP_TYPE_APPROVE,

	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_CONDITIONAL_APPROVE,
	BRC20_HISTORY_SWAP_TYPE_CONDITIONAL_APPROVE,

	BRC20_HISTORY_SWAP_TYPE_APPROVE_FROM,
	BRC20_HISTORY_SWAP_TYPE_APPROVE_TO,

	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_COMMIT,
	BRC20_HISTORY_SWAP_TYPE_COMMIT,
}

var BRC20_HISTORY_TYPES_TO_N map[string]uint8 = map[string]uint8{
	BRC20_HISTORY_TYPE_INSCRIBE_DEPLOY:   BRC20_HISTORY_TYPE_N_INSCRIBE_DEPLOY,
	BRC20_HISTORY_TYPE_INSCRIBE_MINT:     BRC20_HISTORY_TYPE_N_INSCRIBE_MINT,
	BRC20_HISTORY_TYPE_INSCRIBE_TRANSFER: BRC20_HISTORY_TYPE_N_INSCRIBE_TRANSFER,
	BRC20_HISTORY_TYPE_TRANSFER:          BRC20_HISTORY_TYPE_N_TRANSFER,
	BRC20_HISTORY_TYPE_SEND:              BRC20_HISTORY_TYPE_N_SEND,
	BRC20_HISTORY_TYPE_RECEIVE:           BRC20_HISTORY_TYPE_N_RECEIVE,

	// module
	BRC20_HISTORY_MODULE_TYPE_INSCRIBE_MODULE:   BRC20_HISTORY_MODULE_TYPE_N_INSCRIBE_MODULE,
	BRC20_HISTORY_MODULE_TYPE_INSCRIBE_WITHDRAW: BRC20_HISTORY_MODULE_TYPE_N_INSCRIBE_WITHDRAW,
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW:          BRC20_HISTORY_MODULE_TYPE_N_WITHDRAW,
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW_FROM:     BRC20_HISTORY_MODULE_TYPE_N_WITHDRAW_FROM,
	BRC20_HISTORY_MODULE_TYPE_WITHDRAW_TO:       BRC20_HISTORY_MODULE_TYPE_N_WITHDRAW_TO,

	// swap                                                 	// swap
	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_APPROVE: BRC20_HISTORY_SWAP_TYPE_N_INSCRIBE_APPROVE,
	BRC20_HISTORY_SWAP_TYPE_APPROVE:          BRC20_HISTORY_SWAP_TYPE_N_APPROVE,

	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_CONDITIONAL_APPROVE: BRC20_HISTORY_SWAP_TYPE_N_INSCRIBE_CONDITIONAL_APPROVE,
	BRC20_HISTORY_SWAP_TYPE_CONDITIONAL_APPROVE:          BRC20_HISTORY_SWAP_TYPE_N_CONDITIONAL_APPROVE,

	BRC20_HISTORY_SWAP_TYPE_APPROVE_FROM: BRC20_HISTORY_SWAP_TYPE_N_APPROVE_FROM,
	BRC20_HISTORY_SWAP_TYPE_APPROVE_TO:   BRC20_HISTORY_SWAP_TYPE_N_APPROVE_TO,

	BRC20_HISTORY_SWAP_TYPE_INSCRIBE_COMMIT: BRC20_HISTORY_SWAP_TYPE_N_INSCRIBE_COMMIT,
	BRC20_HISTORY_SWAP_TYPE_COMMIT:          BRC20_HISTORY_SWAP_TYPE_N_COMMIT,
}

// swap function
const (
	BRC20_SWAP_FUNCTION_DEPLOY_POOL       = "deployPool"
	BRC20_SWAP_FUNCTION_ADD_LIQ           = "addLiq"
	BRC20_SWAP_FUNCTION_REMOVE_LIQ        = "removeLiq"
	BRC20_SWAP_FUNCTION_SWAP              = "swap"
	BRC20_SWAP_FUNCTION_SEND              = "send"
	BRC20_SWAP_FUNCTION_SENDLP            = "sendLp"
	BRC20_SWAP_FUNCTION_DECREASE_APPROVAL = "decreaseApproval"
)

const ZERO_ADDRESS_PKSCRIPT = "\x6a\x20\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
