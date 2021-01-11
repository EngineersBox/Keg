package keg

type RuleType rune

const (
	// BLOCK ... prevent traffic from continuing
	BLOCK RuleType = 'B'
	// ALLOW ... allow traffic to continue
	ALLOW RuleType = 'A'
)
