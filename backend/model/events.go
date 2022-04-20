package model

import (
	"fmt"
	"strings"
)

func andList(list []string) string {
	if len(list) > 1 {
		return strings.Join(list[:len(list)-1], ", ") + " and " + list[len(list)-1]
	}
	return strings.Join(list, ", ")
}

func (x *Honor) Requirement() string {
	var list []string
	if x.RequiresAnyMeleeOrRangedSkill {
		list = append(list, "attaining sufficent skill with a weapon or technique")
	}
	if x.RequiredSkill != HonorRequiredSkill_Unknown {
		list = append(list, "attaining enough skill with the "+x.RequiredSkill.String())
	}
	if x.RequiredBattles == 1 {
		list = append(list, "serving in combat")
	}
	if x.RequiredBattles > 1 {
		list = append(list, fmt.Sprintf("participating in %d battles", x.RequiredBattles))
	}
	if x.RequiredYears >= 1 {
		list = append(list, fmt.Sprintf("%d years of membership", x.RequiredYears))
	}
	if x.RequiredKills >= 1 {
		list = append(list, fmt.Sprintf("slaying %d enemies", x.RequiredKills))
	}

	return " after " + andList(list)
}

func (x *HistoricalEventAddHfEntityHonor) Html() string {
	e := world.Entities[x.EntityId]
	h := e.Honor[x.HonorId]
	return fmt.Sprintf("%s received the title %s of %s%s", hf(x.Hfid), h.Name(), entity(x.EntityId), h.Requirement())
}

func (x *HistoricalEventAddHfEntityLink) Html() string {
	h := hf(x.Hfid)
	c := entity(x.CivId)
	if x.AppointerHfid != -1 {
		c += fmt.Sprintf(", appointed by %s", hf(x.AppointerHfid))
	}
	switch x.Link {
	case HistoricalEventAddHfEntityLinkLink_Enemy:
		return h + " became an enemy of " + c
	case HistoricalEventAddHfEntityLinkLink_Member:
		return h + " became a member of " + c
	case HistoricalEventAddHfEntityLinkLink_Position:
		return h + " became " + world.Entities[x.CivId].Position(x.PositionId).GenderName(world.HistoricalFigures[x.Hfid]) + " of " + c
	case HistoricalEventAddHfEntityLinkLink_Prisoner:
		return h + " was imprisoned by " + c
	case HistoricalEventAddHfEntityLinkLink_Slave:
		return h + " was enslaved by " + c
	case HistoricalEventAddHfEntityLinkLink_Squad:
		return h + " became a hearthperson/solder of  " + c // TODO
	}
	return h + " became SOMETHING of " + c
}

func (x *HistoricalEventAddHfHfLink) Html() string {
	h := hf(x.Hfid)
	t := hf(x.HfidTarget)
	switch x.LinkType {
	case HistoricalEventAddHfHfLinkLinkType_Apprentice:
		return h + " became the master of " + t
	case HistoricalEventAddHfHfLinkLinkType_Deity:
		return h + " began worshipping " + t
	case HistoricalEventAddHfHfLinkLinkType_FormerMaster:
		return h + " ceased being the apprentice of " + t
	case HistoricalEventAddHfHfLinkLinkType_Lover:
		return h + " became romantically involved with " + t
	case HistoricalEventAddHfHfLinkLinkType_Master:
		return h + " began an apprenticeship under " + t
	case HistoricalEventAddHfHfLinkLinkType_PetOwner:
		return h + " became the owner of " + t
	case HistoricalEventAddHfHfLinkLinkType_Prisoner:
		return h + " imprisoned " + t
	case HistoricalEventAddHfHfLinkLinkType_Spouse:
		return h + " married " + t
	default:
		return h + " LINKED TO " + t
	}
}

func (x *HistoricalEventAddHfSiteLink) Html() string {
	h := hf(x.Histfig)
	c := ""
	if x.Civ != -1 {
		c = " of " + entity(x.Civ)
	}
	b := ""
	if x.Structure != -1 {
		b = " " + structure(x.SiteId, x.Structure)
	}
	s := site(x.SiteId, "in")
	switch x.LinkType {
	case HistoricalEventAddHfSiteLinkLinkType_HomeSiteAbstractBuilding:
		return h + " took up residence in " + b + c + " " + s
	case HistoricalEventAddHfSiteLinkLinkType_Occupation:
		return h + " started working at " + b + c + " " + s
	case HistoricalEventAddHfSiteLinkLinkType_SeatOfPower:
		return h + " ruled from " + b + c + " " + s
	default:
		return h + " LINKED TO " + s
	}
}

func (x *HistoricalEventAgreementFormed) Html() string { // TODO
	return "UNKNWON HistoricalEventAgreementFormed"
}

func (x *HistoricalEventAgreementMade) Html() string { return "UNKNWON HistoricalEventAgreementMade" } // TODO

func (x *HistoricalEventAgreementRejected) Html() string { // TODO
	return "UNKNWON HistoricalEventAgreementRejected"
}

func (x *HistoricalEventArtifactClaimFormed) Html() string {
	a := artifact(x.ArtifactId)
	switch x.Claim {
	case HistoricalEventArtifactClaimFormedClaim_Heirloom:
		return a + " was made a family heirloom by " + hf(x.HistFigureId)
	case HistoricalEventArtifactClaimFormedClaim_Symbol:
		p := world.Entities[x.EntityId].Position(x.PositionProfileId).Name_
		e := entity(x.EntityId)
		return a + " was made a symbol of the " + p + " by " + e
	case HistoricalEventArtifactClaimFormedClaim_Treasure:
		c := ""
		if x.Circumstance != HistoricalEventArtifactClaimFormedCircumstance_Unknown {
			c = x.Circumstance.String()
		}
		if x.HistFigureId != -1 {
			return a + " was claimed by " + hf(x.HistFigureId) + c
		} else if x.EntityId != -1 {
			return a + " was claimed by " + entity(x.EntityId) + c
		}
	}
	return a + " was claimed"
}

func (x *HistoricalEventArtifactCopied) Html() string {
	s := "aquired a copy of"
	if x.FromOriginal {
		s = "made a copy of the original"
	}
	return fmt.Sprintf("%s %s %s from %s%s of %s, keeping it within %s%s",
		entity(x.DestEntityId), s, artifact(x.ArtifactId), structure(x.SourceSiteId, x.SourceStructureId), site(x.SourceSiteId, " in "),
		entity(x.SourceEntityId), structure(x.DestSiteId, x.DestStructureId), site(x.DestSiteId, " in "))
}

func (x *HistoricalEventArtifactCreated) Html() string {
	a := artifact(x.ArtifactId)
	h := hf(x.HistFigureId)
	s := ""
	if x.SiteId != -1 {
		s = site(x.SiteId, " in ")
	}
	if !x.NameOnly {
		return h + " created " + a + s
	}
	c := ""
	if x.Circumstance != nil {
		switch x.Circumstance.Type {
		case HistoricalEventArtifactCreatedCircumstanceType_Defeated:
			c = " after defeating " + hf(x.Circumstance.Defeated)
		case HistoricalEventArtifactCreatedCircumstanceType_Favoritepossession:
			c = " as the item was a favorite possession"
		case HistoricalEventArtifactCreatedCircumstanceType_Preservebody:
			c = " by preserving part of the body"
		}
	}
	switch x.Reason {
	case HistoricalEventArtifactCreatedReason_SanctifyHf:
		return fmt.Sprintf("%s received its name%s from %s in order to sanctify %s%s", a, s, h, hf(x.SanctifyHf), c)
	default:
		return fmt.Sprintf("%s received its name%s from %s %s", a, s, h, c)
	}
}
func (x *HistoricalEventArtifactDestroyed) Html() string {
	return fmt.Sprintf("%s was destroyed by %s in %s", artifact(x.ArtifactId), entity(x.DestroyerEnid), site(x.SiteId, ""))
}

func (x *HistoricalEventArtifactFound) Html() string {
	return fmt.Sprintf("%s was found in %s by %s", artifact(x.ArtifactId), site(x.SiteId, ""), hf(x.HistFigureId))
}
func (x *HistoricalEventArtifactGiven) Html() string {
	r := ""
	if x.ReceiverHistFigureId != -1 {
		r = hf(x.ReceiverHistFigureId)
		if x.ReceiverEntityId != -1 {
			r += " of " + entity(x.ReceiverEntityId)
		}
	} else if x.ReceiverEntityId != -1 {
		r += entity(x.ReceiverEntityId)
	}
	g := ""
	if x.GiverHistFigureId != -1 {
		g = hf(x.GiverHistFigureId)
		if x.GiverEntityId != -1 {
			g += " of " + entity(x.GiverEntityId)
		}
	} else if x.GiverEntityId != -1 {
		g += entity(x.GiverEntityId)
	}
	reason := ""
	switch x.Reason {
	case HistoricalEventArtifactGivenReason_PartOfTradeNegotiation:
		reason = " as part of a trade negotiation"
	}
	return fmt.Sprintf("%s was offered to %s by %s%s", artifact(x.ArtifactId), r, g, reason)
}
func (x *HistoricalEventArtifactLost) Html() string {
	w := ""
	if x.SiteId != -1 {
		w = site(x.SiteId, "")
	}
	if x.SubregionId != -1 {
		w = region(x.SubregionId) // TODO optional
	}

	return fmt.Sprintf("%s was lost in %s", artifact(x.ArtifactId), w)
}
func (x *HistoricalEventArtifactPossessed) Html() string {
	return "UNKNWON HistoricalEventArtifactPossessed"
}
func (x *HistoricalEventArtifactRecovered) Html() string {
	return "UNKNWON HistoricalEventArtifactRecovered"
}
func (x *HistoricalEventArtifactStored) Html() string { return "UNKNWON HistoricalEventArtifactStored" }
func (x *HistoricalEventArtifactTransformed) Html() string {
	return "UNKNWON HistoricalEventArtifactTransformed"
}
func (x *HistoricalEventAssumeIdentity) Html() string { return "UNKNWON HistoricalEventAssumeIdentity" }
func (x *HistoricalEventAttackedSite) Html() string   { return "UNKNWON HistoricalEventAttackedSite" }
func (x *HistoricalEventBodyAbused) Html() string     { return "UNKNWON HistoricalEventBodyAbused" }
func (x *HistoricalEventBuildingProfileAcquired) Html() string {
	return "UNKNWON HistoricalEventBuildingProfileAcquired"
}
func (x *HistoricalEventCeremony) Html() string { return "UNKNWON HistoricalEventCeremony" }
func (x *HistoricalEventChangeHfBodyState) Html() string {
	return "UNKNWON HistoricalEventChangeHfBodyState"
}
func (x *HistoricalEventChangeHfJob) Html() string { return "UNKNWON HistoricalEventChangeHfJob" }
func (x *HistoricalEventChangeHfState) Html() string {
	switch x.State {
	case HistoricalEventChangeHfStateState_Settled:
		switch x.Reason {
		case HistoricalEventChangeHfStateReason_BeWithMaster:
			return hf(x.Hfid) + " moved to study " + site(x.SiteId, "in") + " in order to be with the master"
		default:
			return hf(x.Hfid) + " settled " + site(x.SiteId, "in")
		}
	}
	return "UNKNWON HistoricalEventChangeHfState"
}

func (x *HistoricalEventChangedCreatureType) Html() string {
	return "UNKNWON HistoricalEventChangedCreatureType"
}
func (x *HistoricalEventCompetition) Html() string { return "UNKNWON HistoricalEventCompetition" }
func (x *HistoricalEventCreateEntityPosition) Html() string {
	return "UNKNWON HistoricalEventCreateEntityPosition"
}
func (x *HistoricalEventCreatedSite) Html() string { return "UNKNWON HistoricalEventCreatedSite" }
func (x *HistoricalEventCreatedStructure) Html() string {
	return "UNKNWON HistoricalEventCreatedStructure"
}
func (x *HistoricalEventCreatedWorldConstruction) Html() string {
	return "UNKNWON HistoricalEventCreatedWorldConstruction"
}
func (x *HistoricalEventCreatureDevoured) Html() string {
	return "UNKNWON HistoricalEventCreatureDevoured"
}
func (x *HistoricalEventDanceFormCreated) Html() string {
	return "UNKNWON HistoricalEventDanceFormCreated"
}
func (x *HistoricalEventDestroyedSite) Html() string { return "UNKNWON HistoricalEventDestroyedSite" }
func (x *HistoricalEventDiplomatLost) Html() string  { return "UNKNWON HistoricalEventDiplomatLost" }
func (x *HistoricalEventEntityAllianceFormed) Html() string {
	return "UNKNWON HistoricalEventEntityAllianceFormed"
}
func (x *HistoricalEventEntityBreachFeatureLayer) Html() string {
	return "UNKNWON HistoricalEventEntityBreachFeatureLayer"
}
func (x *HistoricalEventEntityCreated) Html() string { return "UNKNWON HistoricalEventEntityCreated" }
func (x *HistoricalEventEntityDissolved) Html() string {
	return "UNKNWON HistoricalEventEntityDissolved"
}
func (x *HistoricalEventEntityEquipmentPurchase) Html() string {
	return "UNKNWON HistoricalEventEntityEquipmentPurchase"
}
func (x *HistoricalEventEntityExpelsHf) Html() string { return "UNKNWON HistoricalEventEntityExpelsHf" }
func (x *HistoricalEventEntityFledSite) Html() string { return "UNKNWON HistoricalEventEntityFledSite" }
func (x *HistoricalEventEntityIncorporated) Html() string {
	return "UNKNWON HistoricalEventEntityIncorporated"
}
func (x *HistoricalEventEntityLaw) Html() string { return "UNKNWON HistoricalEventEntityLaw" }
func (x *HistoricalEventEntityOverthrown) Html() string {
	return "UNKNWON HistoricalEventEntityOverthrown"
}
func (x *HistoricalEventEntityPersecuted) Html() string {
	return "UNKNWON HistoricalEventEntityPersecuted"
}
func (x *HistoricalEventEntityPrimaryCriminals) Html() string {
	return "UNKNWON HistoricalEventEntityPrimaryCriminals"
}
func (x *HistoricalEventEntityRampagedInSite) Html() string {
	return "UNKNWON HistoricalEventEntityRampagedInSite"
}
func (x *HistoricalEventEntityRelocate) Html() string { return "UNKNWON HistoricalEventEntityRelocate" }
func (x *HistoricalEventEntitySearchedSite) Html() string {
	return "UNKNWON HistoricalEventEntitySearchedSite"
}
func (x *HistoricalEventFailedFrameAttempt) Html() string {
	return "UNKNWON HistoricalEventFailedFrameAttempt"
}
func (x *HistoricalEventFailedIntrigueCorruption) Html() string {
	return "UNKNWON HistoricalEventFailedIntrigueCorruption"
}
func (x *HistoricalEventFieldBattle) Html() string    { return "UNKNWON HistoricalEventFieldBattle" }
func (x *HistoricalEventFirstContact) Html() string   { return "UNKNWON HistoricalEventFirstContact" }
func (x *HistoricalEventGamble) Html() string         { return "UNKNWON HistoricalEventGamble" }
func (x *HistoricalEventHfAbducted) Html() string     { return "UNKNWON HistoricalEventHfAbducted" }
func (x *HistoricalEventHfAttackedSite) Html() string { return "UNKNWON HistoricalEventHfAttackedSite" }
func (x *HistoricalEventHfConfronted) Html() string   { return "UNKNWON HistoricalEventHfConfronted" }
func (x *HistoricalEventHfConvicted) Html() string    { return "UNKNWON HistoricalEventHfConvicted" }
func (x *HistoricalEventHfDestroyedSite) Html() string {
	return "UNKNWON HistoricalEventHfDestroyedSite"
}
func (x *HistoricalEventHfDied) Html() string { return "UNKNWON HistoricalEventHfDied" }
func (x *HistoricalEventHfDisturbedStructure) Html() string {
	return "UNKNWON HistoricalEventHfDisturbedStructure"
}
func (x *HistoricalEventHfDoesInteraction) Html() string {
	return "UNKNWON HistoricalEventHfDoesInteraction"
}
func (x *HistoricalEventHfEnslaved) Html() string { return "UNKNWON HistoricalEventHfEnslaved" }
func (x *HistoricalEventHfEquipmentPurchase) Html() string {
	return "UNKNWON HistoricalEventHfEquipmentPurchase"
}
func (x *HistoricalEventHfFreed) Html() string { return "UNKNWON HistoricalEventHfFreed" }
func (x *HistoricalEventHfGainsSecretGoal) Html() string {
	return "UNKNWON HistoricalEventHfGainsSecretGoal"
}
func (x *HistoricalEventHfInterrogated) Html() string { return "UNKNWON HistoricalEventHfInterrogated" }
func (x *HistoricalEventHfLearnsSecret) Html() string { return "UNKNWON HistoricalEventHfLearnsSecret" }
func (x *HistoricalEventHfNewPet) Html() string       { return "UNKNWON HistoricalEventHfNewPet" }
func (x *HistoricalEventHfPerformedHorribleExperiments) Html() string {
	return "UNKNWON HistoricalEventHfPerformedHorribleExperiments"
}
func (x *HistoricalEventHfPrayedInsideStructure) Html() string {
	return "UNKNWON HistoricalEventHfPrayedInsideStructure"
}
func (x *HistoricalEventHfPreach) Html() string { return "UNKNWON HistoricalEventHfPreach" }
func (x *HistoricalEventHfProfanedStructure) Html() string {
	return "UNKNWON HistoricalEventHfProfanedStructure"
}
func (x *HistoricalEventHfRansomed) Html() string    { return "UNKNWON HistoricalEventHfRansomed" }
func (x *HistoricalEventHfReachSummit) Html() string { return "UNKNWON HistoricalEventHfReachSummit" }
func (x *HistoricalEventHfRecruitedUnitTypeForEntity) Html() string {
	return "UNKNWON HistoricalEventHfRecruitedUnitTypeForEntity"
}
func (x *HistoricalEventHfRelationshipDenied) Html() string {
	return "UNKNWON HistoricalEventHfRelationshipDenied"
}
func (x *HistoricalEventHfReunion) Html() string { return "UNKNWON HistoricalEventHfReunion" }
func (x *HistoricalEventHfRevived) Html() string { return "UNKNWON HistoricalEventHfRevived" }
func (x *HistoricalEventHfSimpleBattleEvent) Html() string {
	return "UNKNWON HistoricalEventHfSimpleBattleEvent"
}
func (x *HistoricalEventHfTravel) Html() string { return "UNKNWON HistoricalEventHfTravel" }
func (x *HistoricalEventHfViewedArtifact) Html() string {
	return "UNKNWON HistoricalEventHfViewedArtifact"
}
func (x *HistoricalEventHfWounded) Html() string { return "UNKNWON HistoricalEventHfWounded" }
func (x *HistoricalEventHfsFormedIntrigueRelationship) Html() string {
	return "UNKNWON HistoricalEventHfsFormedIntrigueRelationship"
}
func (x *HistoricalEventHfsFormedReputationRelationship) Html() string {
	return "UNKNWON HistoricalEventHfsFormedReputationRelationship"
}
func (x *HistoricalEventHolyCityDeclaration) Html() string {
	return "UNKNWON HistoricalEventHolyCityDeclaration"
}
func (x *HistoricalEventInsurrectionStarted) Html() string {
	return "UNKNWON HistoricalEventInsurrectionStarted"
}
func (x *HistoricalEventItemStolen) Html() string { return "UNKNWON HistoricalEventItemStolen" }
func (x *HistoricalEventKnowledgeDiscovered) Html() string {
	return "UNKNWON HistoricalEventKnowledgeDiscovered"
}
func (x *HistoricalEventMasterpieceArchConstructed) Html() string {
	return "UNKNWON HistoricalEventMasterpieceArchConstructed"
}
func (x *HistoricalEventMasterpieceEngraving) Html() string {
	return "UNKNWON HistoricalEventMasterpieceEngraving"
}
func (x *HistoricalEventMasterpieceFood) Html() string {
	return "UNKNWON HistoricalEventMasterpieceFood"
}
func (x *HistoricalEventMasterpieceItem) Html() string {
	return "UNKNWON HistoricalEventMasterpieceItem"
}
func (x *HistoricalEventMasterpieceItemImprovement) Html() string {
	return "UNKNWON HistoricalEventMasterpieceItemImprovement"
}
func (x *HistoricalEventMasterpieceLost) Html() string {
	return "UNKNWON HistoricalEventMasterpieceLost"
}
func (x *HistoricalEventMerchant) Html() string { return "UNKNWON HistoricalEventMerchant" }
func (x *HistoricalEventModifiedBuilding) Html() string {
	return "UNKNWON HistoricalEventModifiedBuilding"
}
func (x *HistoricalEventMusicalFormCreated) Html() string {
	return "UNKNWON HistoricalEventMusicalFormCreated"
}
func (x *HistoricalEventNewSiteLeader) Html() string { return "UNKNWON HistoricalEventNewSiteLeader" }
func (x *HistoricalEventPeaceAccepted) Html() string { return "UNKNWON HistoricalEventPeaceAccepted" }
func (x *HistoricalEventPeaceRejected) Html() string { return "UNKNWON HistoricalEventPeaceRejected" }
func (x *HistoricalEventPerformance) Html() string   { return "UNKNWON HistoricalEventPerformance" }
func (x *HistoricalEventPlunderedSite) Html() string { return "UNKNWON HistoricalEventPlunderedSite" }
func (x *HistoricalEventPoeticFormCreated) Html() string {
	return "UNKNWON HistoricalEventPoeticFormCreated"
}
func (x *HistoricalEventProcession) Html() string     { return "UNKNWON HistoricalEventProcession" }
func (x *HistoricalEventRazedStructure) Html() string { return "UNKNWON HistoricalEventRazedStructure" }
func (x *HistoricalEventReclaimSite) Html() string    { return "UNKNWON HistoricalEventReclaimSite" }
func (x *HistoricalEventRegionpopIncorporatedIntoEntity) Html() string {
	return "UNKNWON HistoricalEventRegionpopIncorporatedIntoEntity"
}
func (x *HistoricalEventRemoveHfEntityLink) Html() string {
	return "UNKNWON HistoricalEventRemoveHfEntityLink"
}
func (x *HistoricalEventRemoveHfHfLink) Html() string { return "UNKNWON HistoricalEventRemoveHfHfLink" }
func (x *HistoricalEventRemoveHfSiteLink) Html() string {
	return "UNKNWON HistoricalEventRemoveHfSiteLink"
}
func (x *HistoricalEventReplacedStructure) Html() string {
	return "UNKNWON HistoricalEventReplacedStructure"
}
func (x *HistoricalEventSiteDied) Html() string    { return "UNKNWON HistoricalEventSiteDied" }
func (x *HistoricalEventSiteDispute) Html() string { return "UNKNWON HistoricalEventSiteDispute" }
func (x *HistoricalEventSiteRetired) Html() string { return "UNKNWON HistoricalEventSiteRetired" }
func (x *HistoricalEventSiteSurrendered) Html() string {
	return "UNKNWON HistoricalEventSiteSurrendered"
}
func (x *HistoricalEventSiteTakenOver) Html() string { return "UNKNWON HistoricalEventSiteTakenOver" }
func (x *HistoricalEventSiteTributeForced) Html() string {
	return "UNKNWON HistoricalEventSiteTributeForced"
}
func (x *HistoricalEventSneakIntoSite) Html() string { return "UNKNWON HistoricalEventSneakIntoSite" }
func (x *HistoricalEventSpottedLeavingSite) Html() string {
	return "UNKNWON HistoricalEventSpottedLeavingSite"
}
func (x *HistoricalEventSquadVsSquad) Html() string { return "UNKNWON HistoricalEventSquadVsSquad" }
func (x *HistoricalEventTacticalSituation) Html() string {
	return "UNKNWON HistoricalEventTacticalSituation"
}
func (x *HistoricalEventTrade) Html() string { return "UNKNWON HistoricalEventTrade" }
func (x *HistoricalEventWrittenContentComposed) Html() string {
	return "UNKNWON HistoricalEventWrittenContentComposed"
}
