package infura

type Network uint64

const (
	Mainnet Network = 1
	Ropsten         = 3
	Kovan           = 42
	Rinkeby         = 4
)

func (n Network) URL() string {
	switch n {
	case Mainnet:
		return "https://mainnet.infura.io/v3/"
	case Ropsten:
		return "https://ropsten.infura.io/v3/"
	case Kovan:
		return "https://kovan.infura.io/v3/"
	case Rinkeby:
		return "https://rinkeby.infura.io/v3/"
	default:
		return ""
	}
}
