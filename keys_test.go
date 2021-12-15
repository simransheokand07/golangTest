package types_test

import (
	"fmt"
	"testing"

	"github.com/Sifchain/sifnode/x/dispensation/types"
	"github.com/stretchr/testify/assert"
)

func TestGetDistributionRecordKey_statusCompleted(t *testing.T) {
	status := types.DistributionStatus_DISTRIBUTION_STATUS_COMPLETED
	name := types.AttributeKeyDistributionName
	recipient := types.AttributeKeyDistributionRecordAddress
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	key := []byte(fmt.Sprintf("%s_%d_%s", name, distributionType, recipient))
	DistributionRecordPrefixCompleted := []byte{0x011}
	input := types.GetDistributionRecordKey(status, name, recipient, distributionType)
	fmt.Println(input)
	output := append(DistributionRecordPrefixCompleted, key...)
	fmt.Println(output)
	assert.Equal(t, input, output)

}

func TestGetDistributionRecordKey_statusPending(t *testing.T) {
	status := types.DistributionStatus_DISTRIBUTION_STATUS_PENDING
	name := types.AttributeKeyDistributionName
	recipient := types.AttributeKeyDistributionRecordAddress
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	DistributionRecordPrefixPending := []byte{0x000}
	key := []byte(fmt.Sprintf("%s_%d_%s", name, distributionType, recipient))
	input := types.GetDistributionRecordKey(status, name, recipient, distributionType)
	fmt.Println(input)
	output := append(DistributionRecordPrefixPending, key...)
	fmt.Println(output)
	assert.Equal(t, input, output)

}

func TestGetDistributionRecordKey_statusFailed(t *testing.T) {
	status := types.DistributionStatus_DISTRIBUTION_STATUS_FAILED
	name := types.AttributeKeyDistributionName
	recipient := types.AttributeKeyDistributionRecordAddress
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	DistributionRecordPrefixFailed := []byte{0x012}
	key := []byte(fmt.Sprintf("%s_%d_%s", name, distributionType, recipient))
	input := types.GetDistributionRecordKey(status, name, recipient, distributionType)
	fmt.Println(input)
	output := append(DistributionRecordPrefixFailed, key...)
	fmt.Println(output)
	assert.Equal(t, input, output)
}

func TestGetDistributionRecordKey_statusDefault(t *testing.T) {
	status := types.DistributionStatus_DISTRIBUTION_STATUS_UNSPECIFIED
	name := types.AttributeKeyDistributionName
	recipient := types.AttributeKeyDistributionRecordAddress
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	DistributionRecordPrefixCompleted := []byte{0x011}
	key := []byte(fmt.Sprintf("%s_%d_%s", name, distributionType, recipient))
	input := types.GetDistributionRecordKey(status, name, recipient, distributionType)
	fmt.Println(input)
	output := append(DistributionRecordPrefixCompleted, key...)
	fmt.Println(output)
	assert.Equal(t, input, output)
}

func TestGetDistributionsKey(t *testing.T) {
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	authorizedRunner := types.AttributeKeyDistributionRunner
	name := types.AttributeKeyDistributionName
	DistributionsPrefix := []byte{0x01}
	key := []byte(fmt.Sprintf("%s_%d_%s", name, distributionType, authorizedRunner))
	input := types.GetDistributionsKey(name, distributionType, authorizedRunner)
	fmt.Println(input)
	output := append(DistributionsPrefix, key...)
	fmt.Println(output)
	assert.Equal(t, input, output)
}

func TestGetUserClaimKey(t *testing.T) {
	userClaimType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	userAddress := types.AttributeKeyDistributionRecordAddress
	UserClaimPrefix := []byte{0x02}
	key := []byte(fmt.Sprintf("%s_%d", userAddress, userClaimType))
	input := types.GetUserClaimKey(userAddress, userClaimType)
	fmt.Println(input)
	output := append(UserClaimPrefix, key...)
	fmt.Println(output)
	assert.Equal(t, input, output)

}
