package infura

type BlockParameter struct {
	Value string
}

func NewBlockParameter(status string) BlockParameter {
	if status != "latest" && status != "earliest" && status != "pending" {
		return BlockParameter{Value: "latest"}
	}
	return BlockParameter{Value: status}
}

func NewBlockParameterWithNumber(n int) BlockParameter {
	return BlockParameter{Value: "0x" + ToHexString(n)}
}
