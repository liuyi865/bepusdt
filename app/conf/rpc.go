package conf

func GetTronGrpcNode() string {
	if cfg.TronGrpcNode != "" {

		return cfg.TronGrpcNode
	}

	return defaultTronGrpcNode
}

func GetXlayerRpcEndpoint() string {
	if cfg.EvmRpc.Xlayer != "" {

		return cfg.EvmRpc.Xlayer
	}

	return defaultXlayerRpcEndpoint
}

func GetBscRpcEndpoint() string {
	if cfg.EvmRpc.Bsc != "" {

		return cfg.EvmRpc.Bsc
	}

	return defaultBscRpcEndpoint
}

func GetPolygonRpcEndpoint() string {
	if cfg.EvmRpc.Polygon != "" {

		return cfg.EvmRpc.Polygon
	}

	return defaultPolygonRpcEndpoint
}

func GetEthereumRpcEndpoint() string {
	if cfg.EvmRpc.Ethereum != "" {

		return cfg.EvmRpc.Ethereum
	}

	return defaultEthereumRpcEndpoint
}
