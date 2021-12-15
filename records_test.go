package types_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/dispensation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestNewDistributionRecord(t *testing.T) {
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	status := types.DistributionStatus_DISTRIBUTION_STATUS_PENDING
	distributionName := types.AttributeKeyDistributionName
	runner := types.AttributeKeyDistributionRunner
	recipientAddress := types.AttributeKeyDistributionRecordAddress
	Coins := sdk.Coins{sdk.Coin{
		Denom:  "rowan",
		Amount: sdk.NewInt(20),
	}}
	distributionStartHeight := int64(1)
	distributionCompletedHeight := int64(10)
	input := types.NewDistributionRecord(status, distributionType, distributionName, recipientAddress, Coins, distributionStartHeight, distributionCompletedHeight, runner)
	assert.Equal(t, distributionName, input.DistributionName)
	assert.Equal(t, distributionType, input.DistributionType)
	assert.Equal(t, status, input.DistributionStatus)
	assert.Equal(t, Coins, input.Coins)
	assert.Equal(t, runner, input.AuthorizedRunner)
	assert.Equal(t, recipientAddress, input.RecipientAddress)
	assert.Equal(t, distributionStartHeight, input.DistributionStartHeight)
	assert.Equal(t, distributionCompletedHeight, input.DistributionCompletedHeight)

}

func TestNewDistributionRecord_validateTrue(t *testing.T) {
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	status := types.DistributionStatus_DISTRIBUTION_STATUS_PENDING
	distributionName := types.AttributeKeyDistributionName
	runner := types.AttributeKeyDistributionRunner
	recipientAddress := types.AttributeKeyDistributionRecordAddress
	Coins := sdk.Coins{sdk.Coin{
		Denom:  "rowan",
		Amount: sdk.NewInt(20)}}
	distributionStartHeight := int64(1)
	distributionCompletedHeight := int64(10)
	input := types.NewDistributionRecord(status, distributionType, distributionName, recipientAddress, Coins, distributionStartHeight, distributionCompletedHeight, runner)
	bool := input.Validate()
	assert.True(t, bool)
}

func TestNewDistributionRecord_validateEmptyAddress(t *testing.T) {
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	status := types.DistributionStatus_DISTRIBUTION_STATUS_PENDING
	distributionName := types.AttributeKeyDistributionName
	runner := types.AttributeKeyDistributionRunner
	recipientAddress := ""
	Coins := sdk.Coins{sdk.Coin{
		Denom:  "denom",
		Amount: sdk.NewInt(20)}}
	distributionStartHeight := int64(1)
	distributionCompletedHeight := int64(10)
	input := types.NewDistributionRecord(status, distributionType, distributionName, recipientAddress, Coins, distributionStartHeight, distributionCompletedHeight, runner)
	bool := input.Validate()
	assert.False(t, bool)
}

func TestNewDistributionRecord_validateCoinIsInvalid(t *testing.T) {
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	status := types.DistributionStatus_DISTRIBUTION_STATUS_PENDING
	distributionName := types.AttributeKeyDistributionName
	runner := types.AttributeKeyDistributionRunner
	recipientAddress := types.AttributeKeyDistributionRecordAddress
	Coins := sdk.Coins{sdk.Coin{
		Denom:  "",
		Amount: sdk.NewInt(20)}}
	distributionStartHeight := int64(1)
	distributionCompletedHeight := int64(10)
	input := types.NewDistributionRecord(status, distributionType, distributionName, recipientAddress, Coins, distributionStartHeight, distributionCompletedHeight, runner)
	bool := input.Validate()
	assert.False(t, bool)
}

func TestNewDistributionRecord_DoesTypeSupportClaim(t *testing.T) {
	distributionType := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	status := types.DistributionStatus_DISTRIBUTION_STATUS_PENDING
	distributionName := types.AttributeKeyDistributionName
	runner := types.AttributeKeyDistributionRunner
	recipientAddress := types.AttributeKeyDistributionRecordAddress
	Coins := sdk.Coins{sdk.Coin{
		Denom:  "rowan",
		Amount: sdk.NewInt(20)}}
	distributionStartHeight := int64(1)
	distributionCompletedHeight := int64(10)
	input := types.NewDistributionRecord(status, distributionType, distributionName, recipientAddress, Coins, distributionStartHeight, distributionCompletedHeight, runner)
	bool := input.DoesTypeSupportClaim()
	assert.True(t, bool)
}

func TestNewDistribution(t *testing.T) {
	distributiontype := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	name := types.AttributeKeyDistributionName
	authorizedRunner := types.AttributeKeyDistributionRunner
	input := types.NewDistribution(distributiontype, name, authorizedRunner)
	assert.Equal(t, name, input.DistributionName)
	assert.Equal(t, distributiontype, input.DistributionType)
	assert.Equal(t, authorizedRunner, input.Runner)
}

func TestNewDistribution_validateTrue(t *testing.T) {
	distributiontype := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	name := types.AttributeKeyDistributionName
	authorizedRunner := types.AttributeKeyDistributionRunner
	input := types.NewDistribution(distributiontype, name, authorizedRunner)
	bool := input.Validate()
	assert.True(t, bool)

}

func TestNewDistribution_validateFalse(t *testing.T) {
	distributiontype := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	name := ""
	authorizedRunner := types.AttributeKeyDistributionRunner
	input := types.NewDistribution(distributiontype, name, authorizedRunner)
	bool := input.Validate()
	assert.False(t, bool)

}

func TestGetDistributionStatus_Completed(t *testing.T) {
	status := "Completed"
	result, output := types.GetDistributionStatus(status)
	Distributionstatus := int32(result)
	assert.NotEmpty(t, Distributionstatus)
	assert.True(t, output)

}

func TestGetDistributionStatus_Pending(t *testing.T) {
	status := "Pending"
	result, output := types.GetDistributionStatus(status)
	Distributionstatus := int32(result)
	assert.NotEmpty(t, Distributionstatus)
	assert.True(t, output)
}

func TestGetDistributionStatus_Failed(t *testing.T) {
	status := "Failed"
	result, output := types.GetDistributionStatus(status)
	Distributionstatus := int32(result)
	assert.NotEmpty(t, Distributionstatus)
	assert.True(t, output)
}

func TestGetDistributionStatus_Default(t *testing.T) {
	status := ""
	result, output := types.GetDistributionStatus(status)
	Distributionstatus := int32(result)
	assert.Empty(t, Distributionstatus)
	assert.False(t, output)
}

func TestGetClaimType_ValidatorSubsidy(t *testing.T) {
	claimType := "ValidatorSubsidy"
	result, output := types.GetClaimType(claimType)
	DistributionType := int32(result)
	assert.NotEmpty(t, DistributionType)
	assert.True(t, output)
}

func TestGetClaimType_LiquidityMining(t *testing.T) {
	claimType := "LiquidityMining"
	result, output := types.GetClaimType(claimType)
	DistributionType := int32(result)
	assert.NotEmpty(t, DistributionType)
	assert.True(t, output)
}

func TestGetClaimType_Default(t *testing.T) {
	claimType := ""
	result, output := types.GetClaimType(claimType)
	DistributionType := int32(result)
	assert.Empty(t, DistributionType)
	assert.False(t, output)
}

func TestGetDistributionTypeFromShortString_LiquidityMining(t *testing.T) {
	distributionType := "LiquidityMining"
	result, output := types.GetDistributionTypeFromShortString(distributionType)
	DistributionType := int32(result)
	assert.NotEmpty(t, DistributionType)
	assert.True(t, output)
}

func TestGetDistributionTypeFromShortString_Airdrop(t *testing.T) {
	distributionType := "Airdrop"
	result, output := types.GetDistributionTypeFromShortString(distributionType)
	DistributionType := int32(result)
	assert.NotEmpty(t, DistributionType)
	assert.True(t, output)
}

func TestGetDistributionTypeFromShortString_ValidatorSubsidy(t *testing.T) {
	distributionType := "ValidatorSubsidy"
	result, output := types.GetDistributionTypeFromShortString(distributionType)
	DistributionType := int32(result)
	assert.NotEmpty(t, DistributionType)
	assert.True(t, output)
}

func TestGetDistributionTypeFromShortString_default(t *testing.T) {
	distributionType := ""
	result, output := types.GetDistributionTypeFromShortString(distributionType)
	DistributionType := int32(result)
	assert.Empty(t, DistributionType)
	assert.False(t, output)
}

func TestIsValidDistributionType_LIQUIDITY_MINING(t *testing.T) {
	distributionType := "DISTRIBUTION_TYPE_LIQUIDITY_MINING"
	result, output := types.IsValidDistributionType(distributionType)
	DistributionType := int32(result)
	assert.NotEmpty(t, DistributionType)
	assert.True(t, output)
}

func TestIsValidDistributionType_VALIDATOR_SUBSIDY(t *testing.T) {
	distributionType := "DISTRIBUTION_TYPE_VALIDATOR_SUBSIDY"
	result, output := types.IsValidDistributionType(distributionType)
	DistributionType := int32(result)
	assert.NotEmpty(t, DistributionType)
	assert.True(t, output)
}

func TestIsValidDistributionType_Default(t *testing.T) {
	distributionType := ""
	result, output := types.IsValidDistributionType(distributionType)
	DistributionType := int32(result)
	assert.Empty(t, DistributionType)
	assert.False(t, output)
}

func TestIsValidClaimType(t *testing.T) {
	ClaimType := "DISTRIBUTION_TYPE_LIQUIDITY_MINING"
	result, output := types.IsValidClaimType(ClaimType)
	DistributionType := int32(result)
	assert.NotEmpty(t, DistributionType)
	assert.True(t, output)
}
