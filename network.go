package infura

type Network uint

const (
	Mainnet Network = iota
	Ropsten
	// INFURAnet
	Kovan
	Rinkeby
	// IPFS
)

func (n Network) URL() string {
	switch n {
	case Mainnet:
		return "https://mainnet.infura.io/"
	case Ropsten:
		return "https://ropsten.infura.io/"
	case Kovan:
		return "https://kovan.infura.io/"
	case Rinkeby:
		return "https://rinkeby.infura.io/"
	default:
		return ""
	}
}

func (n Network) ChainID() uint64 {
	switch n {
	case Mainnet:
		return 1
	case Ropsten:
		return 3
	case Kovan:
		return 42
	case Rinkeby:
		return 4
	default:
		return 0
	}
}
