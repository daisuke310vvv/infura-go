package infura

type Network uint

const (
	Mainnet Network = iota
	Ropsten
	INFURAnet
	Kovan
	Rinkeby
	IPFS
)

func (n Network) URL() string {
	switch n {
	case Mainnet:
		return "https://mainnet.infura.io/"
	case Ropsten:
		return "https://ropsten.infura.io/"
	case INFURAnet:
		return "https://infuranet.infura.io/"
	case Kovan:
		return "https://kovan.infura.io/"
	case Rinkeby:
		return "https://rinkeby.infura.io/"
	case IPFS:
		return "https://ipfs.infura.io/"
	default:
		return ""
	}
}
