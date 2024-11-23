package event

import (
	"github.com/hectorgimenez/d2go/pkg/data"
)

const (
	FinishedOK          FinishReason = "ok"
	FinishedDied        FinishReason = "death"
	FinishedChicken     FinishReason = "chicken"
	FinishedMercChicken FinishReason = "merc chicken"
	FinishedError       FinishReason = "error"

	InteractionTypeEntrance InteractionType = "entrance"
	InteractionTypeNPC      InteractionType = "npc"
	InteractionTypeObject   InteractionType = "object"
)

type UsedPotionEvent struct {
	BaseEvent
	PotionType data.PotionType
	OnMerc     bool
}

func UsedPotion(be BaseEvent, pt data.PotionType, onMerc bool) UsedPotionEvent {
	return UsedPotionEvent{
		BaseEvent:  be,
		PotionType: pt,
		OnMerc:     onMerc,
	}
}

type GameCreatedEvent struct {
	BaseEvent
	Name     string
	Password string
}

func GameCreated(be BaseEvent, name string, password string) GameCreatedEvent {
	return GameCreatedEvent{
		BaseEvent: be,
		Name:      name,
		Password:  password,
	}
}

type GameFinishedEvent struct {
	BaseEvent
	Reason FinishReason
}

func GameFinished(be BaseEvent, reason FinishReason) GameFinishedEvent {
	return GameFinishedEvent{
		BaseEvent: be,
		Reason:    reason,
	}
}

type RunFinishedEvent struct {
	BaseEvent
	RunName string
	Reason  FinishReason
}

func RunFinished(be BaseEvent, runName string, reason FinishReason) RunFinishedEvent {
	return RunFinishedEvent{
		BaseEvent: be,
		RunName:   runName,
		Reason:    reason,
	}
}

type ItemStashedEvent struct {
	BaseEvent
	Item data.Drop
}

func ItemStashed(be BaseEvent, drop data.Drop) ItemStashedEvent {
	return ItemStashedEvent{
		BaseEvent: be,
		Item:      drop,
	}
}

type RunStartedEvent struct {
	BaseEvent
	RunName string
}

func RunStarted(be BaseEvent, runName string) RunStartedEvent {
	return RunStartedEvent{
		BaseEvent: be,
		RunName:   runName,
	}
}

type CompanionLeaderAttackEvent struct {
	BaseEvent
	TargetUnitID data.UnitID
}

func CompanionLeaderAttack(be BaseEvent, targetUnitID data.UnitID) CompanionLeaderAttackEvent {
	return CompanionLeaderAttackEvent{
		BaseEvent:    be,
		TargetUnitID: targetUnitID,
	}
}

type CompanionRequestedTPEvent struct {
	BaseEvent
}

func CompanionRequestedTP(be BaseEvent) CompanionRequestedTPEvent {
	return CompanionRequestedTPEvent{BaseEvent: be}
}

type InteractedToEvent struct {
	BaseEvent
	ID              int
	InteractionType InteractionType
}

func InteractedTo(be BaseEvent, id int, it InteractionType) InteractedToEvent {
	return InteractedToEvent{
		BaseEvent:       be,
		ID:              id,
		InteractionType: it,
	}
}

type GamePausedEvent struct {
	BaseEvent
	Paused bool
}

func GamePaused(be BaseEvent, paused bool) GamePausedEvent {
	return GamePausedEvent{
		BaseEvent: be,
		Paused:    paused,
	}
}