/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode_test

import (
	commonledger "github.com/hyperledger/fabric/common/ledger"
	"github.com/hyperledger/fabric/core/chaincode"
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/core/container/ccintf"
	"github.com/hyperledger/fabric/core/ledger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChaincode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chaincode Suite")
}

//go:generate counterfeiter -o mock/tx_simulator.go --fake-name TxSimulator . txSimulator
type txSimulator interface {
	ledger.TxSimulator
}

//go:generate counterfeiter -o mock/history_query_executor.go --fake-name HistoryQueryExecutor . historyQueryExecutor
type historyQueryExecutor interface {
	ledger.HistoryQueryExecutor
}

//go:generate counterfeiter -o mock/results_iterator.go --fake-name ResultsIterator . resultsIterator
type resultsIterator interface {
	commonledger.ResultsIterator
}

//go:generate counterfeiter -o mock/runtime.go --fake-name Runtime . chaincodeRuntime
type chaincodeRuntime interface {
	chaincode.Runtime
}

//go:generate counterfeiter -o mock/cert_generator.go --fake-name CertGenerator . certGenerator
type certGenerator interface {
	chaincode.CertGenerator
}

//go:generate counterfeiter -o mock/processor.go --fake-name Processor . processor
type processor interface {
	chaincode.Processor
}

//go:generate counterfeiter -o mock/executor.go --fake-name Executor . executor
type executor interface {
	chaincode.Executor
}

//go:generate counterfeiter -o mock/invoker.go --fake-name Invoker . invoker
type invoker interface {
	chaincode.Invoker
}

//go:generate counterfeiter -o mock/package_provider.go --fake-name PackageProvider . packageProvider
type packageProvider interface {
	chaincode.PackageProvider
}

//go:generate counterfeiter -o mock/cc_package.go --fake-name CCPackage . ccpackage
type ccpackage interface {
	ccprovider.CCPackage
}

//go:generate counterfeiter -o mock/chaincode_stream.go --fake-name ChaincodeStream . chaincodeStream
type chaincodeStream interface {
	ccintf.ChaincodeStream
}

//go:generate counterfeiter -o mock/transaction_registry.go --fake-name TransactionRegistry . transactionRegistry
type transactionRegistry interface {
	chaincode.TransactionRegistry
}

//go:generate counterfeiter -o mock/system_chaincode_provider.go --fake-name SystemCCProvider . systemCCProvider
type systemCCProvider interface {
	chaincode.SystemCCProvider
}

//go:generate counterfeiter -o mock/acl_provider.go --fake-name ACLProvider . aclProvider
type aclProvider interface {
	chaincode.ACLProvider
}

//go:generate counterfeiter -o mock/chaincode_definition_getter.go --fake-name ChaincodeDefinitionGetter . chaincodeDefinitionGetter
type chaincodeDefinitionGetter interface {
	chaincode.ChaincodeDefinitionGetter
}

//go:generate counterfeiter -o mock/instantiation_policy_checker.go --fake-name InstantiationPolicyChecker . instantiationPolicyChecker
type instantiationPolicyChecker interface {
	chaincode.InstantiationPolicyChecker
}

//go:generate counterfeiter -o mock/ledger_getter.go --fake-name LedgerGetter . ledgerGetter
type ledgerGetter interface {
	chaincode.LedgerGetter
}

//go:generate counterfeiter -o mock/peer_ledger.go --fake-name PeerLedger . peerLedger
type peerLedger interface {
	ledger.PeerLedger
}

// NOTE: These are getting generated into the "fake" package to avoid import cycles. We need to revisit this.

//go:generate counterfeiter -o fake/launch_registry.go --fake-name LaunchRegistry . launchRegistry
type launchRegistry interface {
	chaincode.LaunchRegistry
}

//go:generate counterfeiter -o fake/message_handler.go --fake-name MessageHandler . messageHandler
type messageHandler interface {
	chaincode.MessageHandler
}

//go:generate counterfeiter -o fake/context_registry.go --fake-name ContextRegistry  . contextRegistry
type contextRegistry interface {
	chaincode.ContextRegistry
}

//go:generate counterfeiter -o fake/query_response_builder.go --fake-name QueryResponseBuilder . queryResponseBuilder
type queryResponseBuilder interface {
	chaincode.QueryResponseBuilder
}

//go:generate counterfeiter -o fake/registry.go --fake-name Registry . registry
type registry interface {
	chaincode.Registry
}
