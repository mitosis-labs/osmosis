package chain

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/osmosis-labs/osmosis/v7/tests/e2e/util"
)

func Init(id, dataDir string, nodeConfigs []*NodeConfig, votingPeriod time.Duration) (*Chain, error) {
	chain, err := new(id, dataDir)
	if err != nil {
		return nil, err
	}

	for _, nodeConfig := range nodeConfigs {
		newNode, err := newNode(chain, nodeConfig)
		if err != nil {
			return nil, err
		}
		chain.nodes = append(chain.nodes, newNode)
	}

	if err := initGenesis(chain, votingPeriod); err != nil {
		return nil, err
	}

	var peers []string
	for i, peer := range chain.nodes {
		peerID := fmt.Sprintf("%s@%s%d:26656", peer.getNodeKey().ID(), peer.getMoniker(), i)
		peer.setPeerId(peerID)
		peers = append(peers, peerID)
	}

	for _, node := range chain.nodes {
		if node.isValidator {
			if err := node.initValidatorConfigs(chain, peers); err != nil {
				return nil, err
			}
		}
	}
	return chain.export(), nil
}

func InitSingleNode(chainId, dataDir string, existingGenesisDir string, nodeConfig *NodeConfig, votingPeriod time.Duration, trustHeight int64, trustHash string, stateSyncRPCServers []string) (*Node, error) {
	if nodeConfig.IsValidator {
		return nil, errors.New("creating individual validator nodes after starting up chain is not currently supported")
	}

	chain, err := new(chainId, dataDir)
	if err != nil {
		return nil, err
	}

	newNode, err := newNode(chain, nodeConfig)

	_, err = util.CopyFile(
		existingGenesisDir,
		filepath.Join(newNode.configDir(), "config", "genesis.json"),
	)
	if err != nil {
		return nil, err
	}

	if err := newNode.initStateSyncConfig(trustHeight, trustHash, stateSyncRPCServers); err != nil {
		return nil, err
	}

	return newNode.export(), nil
}
