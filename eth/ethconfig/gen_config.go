// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ethconfig

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool/blobpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/miner"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		DisablePeerTxBroadcast  bool
		EthDiscoveryURLs        []string
		SnapDiscoveryURLs       []string
		TrustDiscoveryURLs      []string
		BscDiscoveryURLs        []string
		NoPruning               bool
		NoPrefetch              bool
		DirectBroadcast         bool
		DisableSnapProtocol     bool
		EnableTrustProtocol     bool
		PipeCommit              bool
		RangeLimit              bool
		TxLookupLimit           uint64 `toml:",omitempty"`
		TransactionHistory      uint64 `toml:",omitempty"`
		StateHistory            uint64 `toml:",omitempty"`
		StateScheme             string `toml:",omitempty"`
		PathSyncFlush           bool   `toml:",omitempty"`
		JournalFileEnabled      bool
		RequiredBlocks          map[uint64]common.Hash `toml:"-"`
		LightServ               int                    `toml:",omitempty"`
		LightIngress            int                    `toml:",omitempty"`
		LightEgress             int                    `toml:",omitempty"`
		LightPeers              int                    `toml:",omitempty"`
		LightNoPrune            bool                   `toml:",omitempty"`
		LightNoSyncServe        bool                   `toml:",omitempty"`
		SkipBcVersionCheck      bool                   `toml:"-"`
		DatabaseHandles         int                    `toml:"-"`
		DatabaseCache           int
		DatabaseFreezer         string
		DatabaseDiff            string
		PersistDiff             bool
		DiffBlock               uint64
		PruneAncientData        bool
		TrieCleanCache          int
		TrieDirtyCache          int
		TrieTimeout             time.Duration
		SnapshotCache           int
		TriesInMemory           uint64
		TriesVerifyMode         core.VerifyMode
		Preimages               bool
		FilterLogCacheSize      int
		Miner                   miner.Config
		TxPool                  legacypool.Config
		BlobPool                blobpool.Config
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		DocRoot                 string `toml:"-"`
		RPCGasCap               uint64
		RPCEVMTimeout           time.Duration
		RPCTxFeeCap             float64
		OverrideCancun          *uint64 `toml:",omitempty"`
		OverrideVerkle          *uint64 `toml:",omitempty"`
		BlobExtraReserve        uint64
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.DisablePeerTxBroadcast = c.DisablePeerTxBroadcast
	enc.EthDiscoveryURLs = c.EthDiscoveryURLs
	enc.SnapDiscoveryURLs = c.SnapDiscoveryURLs
	enc.TrustDiscoveryURLs = c.TrustDiscoveryURLs
	enc.BscDiscoveryURLs = c.BscDiscoveryURLs
	enc.NoPruning = c.NoPruning
	enc.NoPrefetch = c.NoPrefetch
	enc.DirectBroadcast = c.DirectBroadcast
	enc.DisableSnapProtocol = c.DisableSnapProtocol
	enc.EnableTrustProtocol = c.EnableTrustProtocol
	enc.PipeCommit = c.PipeCommit
	enc.RangeLimit = c.RangeLimit
	enc.TxLookupLimit = c.TxLookupLimit
	enc.TransactionHistory = c.TransactionHistory
	enc.StateHistory = c.StateHistory
	enc.StateScheme = c.StateScheme
	enc.PathSyncFlush = c.PathSyncFlush
	enc.JournalFileEnabled = c.JournalFileEnabled
	enc.RequiredBlocks = c.RequiredBlocks
	enc.LightServ = c.LightServ
	enc.LightIngress = c.LightIngress
	enc.LightEgress = c.LightEgress
	enc.LightPeers = c.LightPeers
	enc.LightNoPrune = c.LightNoPrune
	enc.LightNoSyncServe = c.LightNoSyncServe
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.DatabaseFreezer = c.DatabaseFreezer
	enc.DatabaseDiff = c.DatabaseDiff
	enc.PersistDiff = c.PersistDiff
	enc.DiffBlock = c.DiffBlock
	enc.PruneAncientData = c.PruneAncientData
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.SnapshotCache = c.SnapshotCache
	enc.TriesInMemory = c.TriesInMemory
	enc.TriesVerifyMode = c.TriesVerifyMode
	enc.Preimages = c.Preimages
	enc.FilterLogCacheSize = c.FilterLogCacheSize
	enc.Miner = c.Miner
	enc.TxPool = c.TxPool
	enc.BlobPool = c.BlobPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.DocRoot = c.DocRoot
	enc.RPCGasCap = c.RPCGasCap
	enc.RPCEVMTimeout = c.RPCEVMTimeout
	enc.RPCTxFeeCap = c.RPCTxFeeCap
	enc.OverrideCancun = c.OverrideCancun
	enc.OverrideVerkle = c.OverrideVerkle
	enc.BlobExtraReserve = c.BlobExtraReserve
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		DisablePeerTxBroadcast  *bool
		EthDiscoveryURLs        []string
		SnapDiscoveryURLs       []string
		TrustDiscoveryURLs      []string
		BscDiscoveryURLs        []string
		NoPruning               *bool
		NoPrefetch              *bool
		DirectBroadcast         *bool
		DisableSnapProtocol     *bool
		EnableTrustProtocol     *bool
		PipeCommit              *bool
		RangeLimit              *bool
		TxLookupLimit           *uint64 `toml:",omitempty"`
		TransactionHistory      *uint64 `toml:",omitempty"`
		StateHistory            *uint64 `toml:",omitempty"`
		StateScheme             *string `toml:",omitempty"`
		PathSyncFlush           *bool   `toml:",omitempty"`
		JournalFileEnabled      *bool
		RequiredBlocks          map[uint64]common.Hash `toml:"-"`
		LightServ               *int                   `toml:",omitempty"`
		LightIngress            *int                   `toml:",omitempty"`
		LightEgress             *int                   `toml:",omitempty"`
		LightPeers              *int                   `toml:",omitempty"`
		LightNoPrune            *bool                  `toml:",omitempty"`
		LightNoSyncServe        *bool                  `toml:",omitempty"`
		SkipBcVersionCheck      *bool                  `toml:"-"`
		DatabaseHandles         *int                   `toml:"-"`
		DatabaseCache           *int
		DatabaseFreezer         *string
		DatabaseDiff            *string
		PersistDiff             *bool
		DiffBlock               *uint64
		PruneAncientData        *bool
		TrieCleanCache          *int
		TrieDirtyCache          *int
		TrieTimeout             *time.Duration
		SnapshotCache           *int
		TriesInMemory           *uint64
		TriesVerifyMode         *core.VerifyMode
		Preimages               *bool
		FilterLogCacheSize      *int
		Miner                   *miner.Config
		TxPool                  *legacypool.Config
		BlobPool                *blobpool.Config
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		DocRoot                 *string `toml:"-"`
		RPCGasCap               *uint64
		RPCEVMTimeout           *time.Duration
		RPCTxFeeCap             *float64
		OverrideCancun          *uint64 `toml:",omitempty"`
		OverrideVerkle          *uint64 `toml:",omitempty"`
		BlobExtraReserve        *uint64
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.DisablePeerTxBroadcast != nil {
		c.DisablePeerTxBroadcast = *dec.DisablePeerTxBroadcast
	}
	if dec.EthDiscoveryURLs != nil {
		c.EthDiscoveryURLs = dec.EthDiscoveryURLs
	}
	if dec.SnapDiscoveryURLs != nil {
		c.SnapDiscoveryURLs = dec.SnapDiscoveryURLs
	}
	if dec.TrustDiscoveryURLs != nil {
		c.TrustDiscoveryURLs = dec.TrustDiscoveryURLs
	}
	if dec.BscDiscoveryURLs != nil {
		c.BscDiscoveryURLs = dec.BscDiscoveryURLs
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.NoPrefetch != nil {
		c.NoPrefetch = *dec.NoPrefetch
	}
	if dec.DirectBroadcast != nil {
		c.DirectBroadcast = *dec.DirectBroadcast
	}
	if dec.DisableSnapProtocol != nil {
		c.DisableSnapProtocol = *dec.DisableSnapProtocol
	}
	if dec.EnableTrustProtocol != nil {
		c.EnableTrustProtocol = *dec.EnableTrustProtocol
	}
	if dec.PipeCommit != nil {
		c.PipeCommit = *dec.PipeCommit
	}
	if dec.RangeLimit != nil {
		c.RangeLimit = *dec.RangeLimit
	}
	if dec.TxLookupLimit != nil {
		c.TxLookupLimit = *dec.TxLookupLimit
	}
	if dec.TransactionHistory != nil {
		c.TransactionHistory = *dec.TransactionHistory
	}
	if dec.StateHistory != nil {
		c.StateHistory = *dec.StateHistory
	}
	if dec.StateScheme != nil {
		c.StateScheme = *dec.StateScheme
	}
	if dec.PathSyncFlush != nil {
		c.PathSyncFlush = *dec.PathSyncFlush
	}
	if dec.JournalFileEnabled != nil {
		c.JournalFileEnabled = *dec.JournalFileEnabled
	}
	if dec.RequiredBlocks != nil {
		c.RequiredBlocks = dec.RequiredBlocks
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightIngress != nil {
		c.LightIngress = *dec.LightIngress
	}
	if dec.LightEgress != nil {
		c.LightEgress = *dec.LightEgress
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.LightNoPrune != nil {
		c.LightNoPrune = *dec.LightNoPrune
	}
	if dec.LightNoSyncServe != nil {
		c.LightNoSyncServe = *dec.LightNoSyncServe
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.DatabaseFreezer != nil {
		c.DatabaseFreezer = *dec.DatabaseFreezer
	}
	if dec.DatabaseDiff != nil {
		c.DatabaseDiff = *dec.DatabaseDiff
	}
	if dec.PersistDiff != nil {
		c.PersistDiff = *dec.PersistDiff
	}
	if dec.DiffBlock != nil {
		c.DiffBlock = *dec.DiffBlock
	}
	if dec.PruneAncientData != nil {
		c.PruneAncientData = *dec.PruneAncientData
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.SnapshotCache != nil {
		c.SnapshotCache = *dec.SnapshotCache
	}
	if dec.TriesInMemory != nil {
		c.TriesInMemory = *dec.TriesInMemory
	}
	if dec.TriesVerifyMode != nil {
		c.TriesVerifyMode = *dec.TriesVerifyMode
	}
	if dec.Preimages != nil {
		c.Preimages = *dec.Preimages
	}
	if dec.FilterLogCacheSize != nil {
		c.FilterLogCacheSize = *dec.FilterLogCacheSize
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.BlobPool != nil {
		c.BlobPool = *dec.BlobPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = *dec.RPCGasCap
	}
	if dec.RPCEVMTimeout != nil {
		c.RPCEVMTimeout = *dec.RPCEVMTimeout
	}
	if dec.RPCTxFeeCap != nil {
		c.RPCTxFeeCap = *dec.RPCTxFeeCap
	}
	if dec.OverrideCancun != nil {
		c.OverrideCancun = dec.OverrideCancun
	}
	if dec.OverrideVerkle != nil {
		c.OverrideVerkle = dec.OverrideVerkle
	}
	if dec.BlobExtraReserve != nil {
		c.BlobExtraReserve = *dec.BlobExtraReserve
	}
	return nil
}
