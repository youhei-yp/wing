// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package invar

// Status status type
type Status int

const (
	// StatePanic     [-5], panic state
	StatePanic Status = iota - 5

	// StateException [-4], exception state
	StateException

	// StateTimeout   [-3], failed state
	StateTimeout

	// StateFailed    [-2], failed state
	StateFailed

	// StateError     [-1], error state
	StateError

	// StateSuccess   [ 0], success state
	StateSuccess

	// StateRecover   [ 1], recover state
	StateRecover
)

// The StaActivate value must be 0 that set in database initialize script,
// if you want change it, plase modify script together
const (
	// StateActive   [0], activate state
	StateActive Status = iota

	// StateFrozen   [1], frozen state
	StateFrozen

	// StateDisabled [2], disable state
	StateDisabled
)

// The StateUnpaid value must be 0 that set in database initialize script,
// if you want change it, plase modify script together
const (
	// StateUnpaied [0], initialization state when trade just created
	StateUnpaid Status = iota

	// StatePaid    [1], the trading have completed
	StatePaid

	// StateExpired [2], the trading has expired
	StateExpired

	// StateInvalid [3], the trading becomes invalid
	StateInvalid
)

// Box box type
type Box int

// The BoxDraft must be 0 that set in database initialize script,
// if you want change it, plase modify script together
const (
	// BoxDraft     [0], draft box
	BoxDraft Box = iota

	// BoxActive    [1], active box
	BoxActive

	// BoxOffShelve [2], offshelve box
	BoxOffShelve

	// BoxSend      [3], send box
	BoxSend

	// BoxReceive   [4], receive box
	BoxReceive

	// BoxSending   [5], sending box
	BoxSending

	// BoxDustbin   [6], dustbin box
	BoxDustbin
)

// Role role type
type Role int

const (
	// RoleUser      [ 0], user role
	RoleUser Role = iota

	// RoleAdmin     [ 1], admin role
	RoleAdmin

	// RoleManager   [ 2], manager role
	RoleManager

	// RoleSuper     [ 3], super role
	RoleSuper

	// RoleConsumer  [ 4], consumer role
	RoleConsumer

	// RoleSeller    [ 5], seller role
	RoleSeller

	// RoleAgent     [ 6], agent role
	RoleAgent

	// RoleVendor    [ 7], vendor role
	RoleVendor

	// RoleOwner     [ 8], owner role
	RoleOwner

	// RoleTarget    [ 9], target role
	RoleTarget

	// RoleGuest     [10], user role
	RoleGuest

	// RoleMaster    [11], user role
	RoleMaster

	// RoleCaller    [12], caller role
	RoleCaller

	// RoleCallee    [13], callee role
	RoleCallee

	// RoleReception [14], company reception role
	RoleReception

	// RoleControl   [15], control room role
	RoleControl

	// RoleRoom      [16], gurst room role
	RoleRoom
)

// Limit limit permission type
type Limit int

const (
	// LimitAddible   [0], addible permission
	LimitAddible Limit = iota

	// LimitMutable   [1], mutable permission
	LimitMutable

	// LimitPayable   [2], payable permission
	LimitPayable

	// LimitSettable  [3], settable permission
	LimitSettable

	// LimitDeletable [4], deletable permission
	LimitDeletable
)