package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func (x *HistoricalEventAddHfEntityHonor) Html(c *Context) string {
	e := c.World.Entities[x.EntityId]
	h := e.Honor[x.HonorId]
	return fmt.Sprintf("%s received the title %s of %s%s", c.hf(x.Hfid), h.Name(), c.entity(x.EntityId), h.Requirement())
}

func (x *HistoricalEventAddHfEntityLink) Html(c *Context) string {
	h := c.hf(x.Hfid)
	e := c.entity(x.CivId)

	if c.Story {
		return "the ascension of " + h + " to " + c.position(x.CivId, x.PositionId, x.Hfid) + " of " + e
	}

	if x.AppointerHfid != -1 {
		e += fmt.Sprintf(", appointed by %s", c.hf(x.AppointerHfid))
	}
	switch x.Link {
	case HistoricalEventAddHfEntityLinkLink_Enemy:
		return h + " became an enemy of " + e
	case HistoricalEventAddHfEntityLinkLink_Member:
		return h + " became a member of " + e
	case HistoricalEventAddHfEntityLinkLink_Position:
		return h + " became " + c.position(x.CivId, x.PositionId, x.Hfid) + " of " + e
	case HistoricalEventAddHfEntityLinkLink_Prisoner:
		return h + " was imprisoned by " + e
	case HistoricalEventAddHfEntityLinkLink_Slave:
		return h + " was enslaved by " + e
	case HistoricalEventAddHfEntityLinkLink_Squad:
		return h + " became a hearthperson/solder of  " + e // TODO
	}
	return h + " became SOMETHING of " + e
}

func (x *HistoricalEventAddHfHfLink) Html(c *Context) string {
	h := c.hf(x.Hfid)
	t := c.hf(x.HfidTarget)
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

func (x *HistoricalEventAddHfSiteLink) Html(c *Context) string {
	h := c.hf(x.Histfig)
	e := ""
	if x.Civ != -1 {
		e = " of " + c.entity(x.Civ)
	}
	b := ""
	if x.Structure != -1 {
		b = " " + c.structure(x.SiteId, x.Structure)
	}
	s := c.site(x.SiteId, "in")
	switch x.LinkType {
	case HistoricalEventAddHfSiteLinkLinkType_HomeSiteAbstractBuilding:
		return h + " took up residence in " + b + e + " " + s
	case HistoricalEventAddHfSiteLinkLinkType_Occupation:
		return h + " started working at " + b + e + " " + s
	case HistoricalEventAddHfSiteLinkLinkType_PrisonAbstractBuilding:
		return h + " was imprisoned in " + b + e + " " + s
	case HistoricalEventAddHfSiteLinkLinkType_PrisonSiteBuildingProfile:
		return h + " was imprisoned in " + b + e + " " + s
	case HistoricalEventAddHfSiteLinkLinkType_SeatOfPower:
		return h + " ruled from " + b + e + " " + s
	default:
		return h + " LINKED TO " + s
	}
}

func (x *HistoricalEventAgreementConcluded) Html(c *Context) string { // TODO wording
	r := ""
	switch x.Topic {
	case HistoricalEventAgreementConcludedTopic_Treequota:
		r += "a lumber agreement"
	}
	return r + " proposed by " + c.entity(x.Source) + " was concluded by " + c.entity(x.Destination) + c.site(x.Site, " at")
}

func (x *HistoricalEventAgreementFormed) Html(c *Context) string { // TODO no info
	return "UNKNWON HistoricalEventAgreementFormed"
}

func (x *HistoricalEventAgreementMade) Html(c *Context) string {
	r := ""
	switch x.Topic {
	case HistoricalEventAgreementMadeTopic_Becomelandholder:
		r += "the establishment of landed nobility"
	case HistoricalEventAgreementMadeTopic_Promotelandholder:
		r += "the elevation of landed nobility"
	case HistoricalEventAgreementMadeTopic_Treequota:
		r += "a lumber agreement"
	}
	return r + " proposed by " + c.entity(x.Source) + " was accepted by " + c.entity(x.Destination) + c.site(x.SiteId, " at")
}

func (x *HistoricalEventAgreementRejected) Html(c *Context) string {
	r := ""
	switch x.Topic {
	case HistoricalEventAgreementRejectedTopic_Becomelandholder:
		r += "the establishment of landed nobility"
	case HistoricalEventAgreementRejectedTopic_Treequota:
		r += "a lumber agreement"
	case HistoricalEventAgreementRejectedTopic_Tributeagreement:
		r += "a tribute agreement"
	case HistoricalEventAgreementRejectedTopic_Unknown10:
		r += "a demand of unconditional surrender"
	}
	return r + " proposed by " + c.entity(x.Source) + " was rejected by " + c.entity(x.Destination) + c.site(x.SiteId, " at")
}

func (x *HistoricalEventArtifactClaimFormed) Html(c *Context) string {
	a := c.artifact(x.ArtifactId)
	switch x.Claim {
	case HistoricalEventArtifactClaimFormedClaim_Heirloom:
		return a + " was made a family heirloom by " + c.hf(x.HistFigureId)
	case HistoricalEventArtifactClaimFormedClaim_Symbol:
		p := c.World.Entities[x.EntityId].Position(x.PositionProfileId).Name_
		e := c.entity(x.EntityId)
		return a + " was made a symbol of the " + p + " by " + e
	case HistoricalEventArtifactClaimFormedClaim_Treasure:
		circumstance := ""
		if x.Circumstance != HistoricalEventArtifactClaimFormedCircumstance_Unknown {
			circumstance = " " + x.Circumstance.String()
		}
		if x.HistFigureId != -1 {
			return a + " was claimed by " + c.hf(x.HistFigureId) + circumstance
		} else if x.EntityId != -1 {
			return a + " was claimed by " + c.entity(x.EntityId) + circumstance
		}
	}
	return a + " was claimed"
}

func (x *HistoricalEventArtifactCopied) Html(c *Context) string {
	s := util.If(x.FromOriginal, "made a copy of the original", "aquired a copy of")
	return fmt.Sprintf("%s %s %s %s of %s, keeping it%s",
		c.entity(x.DestEntityId), s, c.artifact(x.ArtifactId), c.siteStructure(x.SourceSiteId, x.SourceStructureId, "from"),
		c.entity(x.SourceEntityId), c.siteStructure(x.DestSiteId, x.DestStructureId, "within"))
}

func (x *HistoricalEventArtifactCreated) Html(c *Context) string {
	a := c.artifact(x.ArtifactId)
	h := c.hf(x.HistFigureId)
	s := ""
	if x.SiteId != -1 {
		s = c.site(x.SiteId, " in ")
	}
	if !x.NameOnly {
		return h + " created " + a + s
	}
	e := ""
	if x.Circumstance != nil {
		switch x.Circumstance.Type {
		case HistoricalEventArtifactCreatedCircumstanceType_Defeated:
			e = " after defeating " + c.hf(x.Circumstance.Defeated)
		case HistoricalEventArtifactCreatedCircumstanceType_Favoritepossession:
			e = " as the item was a favorite possession"
		case HistoricalEventArtifactCreatedCircumstanceType_Preservebody:
			e = " by preserving part of the body"
		}
	}
	switch x.Reason {
	case HistoricalEventArtifactCreatedReason_SanctifyHf:
		return fmt.Sprintf("%s received its name%s from %s in order to sanctify %s%s", a, s, h, c.hf(x.SanctifyHf), e)
	default:
		return fmt.Sprintf("%s received its name%s from %s %s", a, s, h, e)
	}
}

func (x *HistoricalEventArtifactDestroyed) Html(c *Context) string {
	return c.artifact(x.ArtifactId) + " was destroyed" + util.If(x.DestroyerEnid != -1, " by "+c.entity(x.DestroyerEnid), "") + c.site(x.SiteId, " in")
}

func (x *HistoricalEventArtifactFound) Html(c *Context) string {
	w := ""
	if x.SiteId != -1 {
		w = c.site(x.SiteId, "")
		if x.SitePropertyId != -1 {
			w = c.property(x.SiteId, x.SitePropertyId) + " in " + w
		}
	}
	return fmt.Sprintf("%s was found in %s by %s", c.artifact(x.ArtifactId), w, util.If(x.HistFigureId != -1, c.hf(x.HistFigureId), "an unknown creature"))
}

func (x *HistoricalEventArtifactGiven) Html(c *Context) string {
	r := ""
	if x.ReceiverHistFigureId != -1 {
		r = c.hf(x.ReceiverHistFigureId)
		if x.ReceiverEntityId != -1 {
			r += " of " + c.entity(x.ReceiverEntityId)
		}
	} else if x.ReceiverEntityId != -1 {
		r += c.entity(x.ReceiverEntityId)
	}
	g := ""
	if x.GiverHistFigureId != -1 {
		g = c.hf(x.GiverHistFigureId)
		if x.GiverEntityId != -1 {
			g += " of " + c.entity(x.GiverEntityId)
		}
	} else if x.GiverEntityId != -1 {
		g += c.entity(x.GiverEntityId)
	}
	reason := ""
	switch x.Reason {
	case HistoricalEventArtifactGivenReason_PartOfTradeNegotiation:
		reason = " as part of a trade negotiation"
	}
	return fmt.Sprintf("%s was offered to %s by %s%s", c.artifact(x.ArtifactId), r, g, reason)
}
func (x *HistoricalEventArtifactLost) Html(c *Context) string {
	w := ""
	if x.SubregionId != -1 {
		w = c.region(x.SubregionId)
	}
	if x.SiteId != -1 {
		w = c.site(x.SiteId, "")
		if x.SitePropertyId != -1 {
			w = c.property(x.SiteId, x.SitePropertyId) + " in " + w
		}
	}
	return fmt.Sprintf("%s was lost in %s", c.artifact(x.ArtifactId), w)
}

func (x *HistoricalEventArtifactPossessed) Html(c *Context) string {
	a := c.artifact(x.ArtifactId)
	h := c.hf(x.HistFigureId)
	w := ""
	if x.SubregionId != -1 {
		w = c.region(x.SubregionId)
	}
	if x.SiteId != -1 {
		w = c.site(x.SiteId, "")
	}
	circumstance := ""
	switch x.Circumstance {
	case HistoricalEventArtifactPossessedCircumstance_HfIsDead:
		circumstance = " after the death of " + c.hf(x.CircumstanceId)
	}

	switch x.Reason {
	case HistoricalEventArtifactPossessedReason_ArtifactIsHeirloomOfFamilyHfid:
		return fmt.Sprintf("%s was aquired in %s by %s as an heirloom of %s%s", a, w, h, c.hf(x.ReasonId), circumstance)
	case HistoricalEventArtifactPossessedReason_ArtifactIsSymbolOfEntityPosition:
		return fmt.Sprintf("%s was aquired in %s by %s as a symbol of authority within %s%s", a, w, h, c.entity(x.ReasonId), circumstance)
	}
	return fmt.Sprintf("%s was claimed in %s by %s%s", a, w, h, circumstance) // TODO wording
}

func (x *HistoricalEventArtifactRecovered) Html(c *Context) string {
	a := c.artifact(x.ArtifactId)
	h := c.hf(x.HistFigureId)
	w := ""
	if x.SubregionId != -1 {
		w = "in " + c.region(x.SubregionId)
	}
	if x.SiteId != -1 {
		w = c.site(x.SiteId, "in ")
		if x.StructureId != -1 {
			w = c.siteStructure(x.SiteId, x.StructureId, "from")
		}
	}
	return fmt.Sprintf("%s was recovered %s by %s", a, w, h)
}

func (x *HistoricalEventArtifactStored) Html(c *Context) string {
	if x.HistFigureId != -1 {
		return fmt.Sprintf("%s stored %s in %s", c.hf(x.HistFigureId), c.artifact(x.ArtifactId), c.site(x.SiteId, ""))
	} else {
		return fmt.Sprintf("%s was stored in %s", c.artifact(x.ArtifactId), c.site(x.SiteId, ""))
	}
}

func (x *HistoricalEventArtifactTransformed) Html(c *Context) string {
	return fmt.Sprintf("%s was made from %s by %s in %s", c.artifact(x.NewArtifactId), c.artifact(x.OldArtifactId), c.hf(x.HistFigureId), c.site(x.SiteId, "")) // TODO wording
}

func (x *HistoricalEventAssumeIdentity) Html(c *Context) string {
	h := c.hf(x.TricksterHfid)
	i := c.identity(x.IdentityId)
	if x.TargetEnid == -1 {
		return fmt.Sprintf(`%s assumed the identity "%s"`, h, i)
	} else {
		return fmt.Sprintf(`%s fooled %s into believing %s was "%s"`, h, c.entity(x.TargetEnid), c.pronoun(x.TricksterHfid), i)
	}
}

func (x *HistoricalEventAttackedSite) Html(c *Context) string {
	atk := c.entity(x.AttackerCivId)
	def := c.siteCiv(x.SiteCivId, x.DefenderCivId)
	generals := ""
	if x.AttackerGeneralHfid != -1 {
		generals += ". " + util.Capitalize(c.hf(x.AttackerGeneralHfid)) + " led the attack"
		if x.DefenderGeneralHfid != -1 {
			generals += ", and the defenders were led by " + c.hf(x.DefenderGeneralHfid)
		}
	}
	mercs := ""
	if x.AttackerMercEnid != -1 {
		mercs += fmt.Sprintf(". %s were hired by the attackers", c.entity(x.AttackerMercEnid))
	}
	if x.ASupportMercEnid != -1 {
		mercs += fmt.Sprintf(". %s were hired as scouts by the attackers", c.entity(x.ASupportMercEnid))
	}
	if x.DefenderMercEnid != -1 {
		mercs += fmt.Sprintf(". The defenders hired %s", c.entity(x.DefenderMercEnid))
	}
	if x.DSupportMercEnid != -1 {
		mercs += fmt.Sprintf(". The defenders hired %s as scouts", c.entity(x.DSupportMercEnid))
	}
	return fmt.Sprintf("%s attacked %s at %s%s%s", atk, def, c.site(x.SiteId, ""), generals, mercs)
}

func (x *HistoricalEventBodyAbused) Html(c *Context) string {
	s := "the " + util.If(len(x.Bodies) > 1, "bodies", "body") + " of " + c.hfList(x.Bodies) + " " + util.If(len(x.Bodies) > 1, "were", "was")

	switch x.AbuseType {
	case HistoricalEventBodyAbusedAbuseType_Animated:
		s += " animated" + util.If(x.Histfig != -1, " by "+c.hf(x.Histfig), "") + c.site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Flayed:
		s += " flayed and the skin stretched over " + c.structure(x.SiteId, x.Structure) + " by " + c.entity(x.Civ) + c.site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Hung:
		s += " hung from a tree by " + c.entity(x.Civ) + c.site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Impaled:
		s += " impaled on " + articled(x.ItemMat+" "+x.ItemSubtype.String()) + " by " + c.entity(x.Civ) + c.site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Mutilated:
		s += " horribly mutilated by " + c.entity(x.Civ) + c.site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Piled:
		s += " added to a "
		switch x.PileType {
		case HistoricalEventBodyAbusedPileType_Grislymound:
			s += "grisly mound"
		case HistoricalEventBodyAbusedPileType_Grotesquepillar:
			s += "grotesque pillar"
		case HistoricalEventBodyAbusedPileType_Gruesomesculpture:
			s += "gruesome sculpture"
		}
		s += " by " + c.entity(x.Civ) + c.site(x.SiteId, " in ")
	}

	return s
}

func (x *HistoricalEventBuildingProfileAcquired) Html(c *Context) string {
	return util.If(x.AcquirerEnid != -1, c.entity(x.AcquirerEnid), c.hf(x.AcquirerHfid)) +
		util.If(x.PurchasedUnowned, " purchased ", " inherited ") +
		c.property(x.SiteId, x.BuildingProfileId) + c.site(x.SiteId, " in") +
		util.If(x.LastOwnerHfid != -1, " formerly owned by "+c.hfRelated(x.LastOwnerHfid, x.AcquirerHfid), "")
}

func (x *HistoricalEventCeremony) Html(c *Context) string {
	r := c.entity(x.CivId) + " held a ceremony in " + c.site(x.SiteId, "")
	if e, ok := c.World.Entities[x.CivId]; ok {
		o := e.Occasion[x.OccasionId]
		r += " as part of " + o.Name()
		s := o.Schedule[x.ScheduleId]
		if len(s.Feature) > 0 {
			r += ". The event featured " + andList(util.Map(s.Feature, c.feature))
		}
	}
	return r
}

func (x *HistoricalEventChangeHfBodyState) Html(c *Context) string {
	r := c.hf(x.Hfid)
	switch x.BodyState {
	case HistoricalEventChangeHfBodyStateBodyState_EntombedAtSite:
		r += " was entombed"
	}
	if x.StructureId != -1 {
		r += " within " + c.structure(x.SiteId, x.StructureId)
	}
	r += c.site(x.SiteId, " in ")
	return r
}

func (x *HistoricalEventChangeHfJob) Html(c *Context) string {
	w := ""
	if x.SubregionId != -1 {
		w = " in " + c.region(x.SubregionId)
	}
	if x.SiteId != -1 {
		w = " in " + c.site(x.SiteId, "")
	}
	old := articled(strcase.ToDelimited(x.OldJob, ' '))
	new := articled(strcase.ToDelimited(x.NewJob, ' '))
	if x.OldJob == "standard" {
		return c.hf(x.Hfid) + " became " + new + w
	} else if x.NewJob == "standard" {
		return c.hf(x.Hfid) + " stopped being " + old + w
	} else {
		return c.hf(x.Hfid) + " gave up being " + old + " to become a " + new + w
	}
}

func (x *HistoricalEventChangeHfState) Html(c *Context) string {
	r := ""
	switch x.Reason {
	case HistoricalEventChangeHfStateReason_BeWithMaster:
		r = " in order to be with the master"
	case HistoricalEventChangeHfStateReason_ConvictionExile, HistoricalEventChangeHfStateReason_ExiledAfterConviction:
		r = " after being exiled following a criminal conviction"
	case HistoricalEventChangeHfStateReason_FailedMood:
		r = " after failing to create an artifact"
	case HistoricalEventChangeHfStateReason_Flight:
	case HistoricalEventChangeHfStateReason_GatherInformation:
		r = " to gather information"
	case HistoricalEventChangeHfStateReason_GreatDealOfStress:
		r = " after a great deal of stress"
	case HistoricalEventChangeHfStateReason_LackOfSleep:
		r = " due to lack of sleep"
	case HistoricalEventChangeHfStateReason_OnAPilgrimage:
		r = " on a pilgrimage"
	case HistoricalEventChangeHfStateReason_Scholarship:
		r = " in order to pursue scholarship"
	case HistoricalEventChangeHfStateReason_UnableToLeaveLocation:
		r = " after being unable to leave a location"
	}

	switch x.State {
	case HistoricalEventChangeHfStateState_Refugee:
		return c.hf(x.Hfid) + " fled " + c.location(x.SiteId, "to", x.SubregionId, "into")
	case HistoricalEventChangeHfStateState_Settled:
		switch x.Reason {
		case HistoricalEventChangeHfStateReason_BeWithMaster, HistoricalEventChangeHfStateReason_Scholarship:
			return c.hf(x.Hfid) + " moved to study " + c.site(x.SiteId, "in") + r
		case HistoricalEventChangeHfStateReason_Flight:
			return c.hf(x.Hfid) + " fled " + c.site(x.SiteId, "to")
		case HistoricalEventChangeHfStateReason_ConvictionExile, HistoricalEventChangeHfStateReason_ExiledAfterConviction:
			return c.hf(x.Hfid) + " departed " + c.site(x.SiteId, "to") + r
		case HistoricalEventChangeHfStateReason_None:
			return c.hf(x.Hfid) + " settled " + c.location(x.SiteId, "in", x.SubregionId, "in")
		}
	case HistoricalEventChangeHfStateState_Visiting:
		return c.hf(x.Hfid) + " visited " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateState_Wandering:
		if x.SubregionId != -1 {
			return c.hf(x.Hfid) + " began wandering " + c.region(x.SubregionId)
		} else {
			return c.hf(x.Hfid) + " began wandering the wilds"
		}
	}

	switch x.Mood {
	case HistoricalEventChangeHfStateMood_Berserk:
		return c.hf(x.Hfid) + " went berserk " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Catatonic:
		return c.hf(x.Hfid) + " stopped responding to the outside world " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Fell:
		return c.hf(x.Hfid) + " was taken by a fell mood " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Fey:
		return c.hf(x.Hfid) + " was taken by a fey mood " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Insane:
		return c.hf(x.Hfid) + " became crazed " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Macabre:
		return c.hf(x.Hfid) + " began to skulk and brood " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Melancholy:
		return c.hf(x.Hfid) + " was striken by melancholy " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Possessed:
		return c.hf(x.Hfid) + " was posessed " + c.site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Secretive:
		return c.hf(x.Hfid) + " withdrew from society " + c.site(x.SiteId, "in") + r
	}
	return "UNKNWON HistoricalEventChangeHfState"
}

func (x *HistoricalEventChangedCreatureType) Html(c *Context) string {
	return c.hf(x.ChangerHfid) + " changed " + c.hfRelated(x.ChangeeHfid, x.ChangerHfid) + " from " + articled(x.OldRace) + " to " + articled(x.NewRace)
}

func (x *HistoricalEventCompetition) Html(c *Context) string {
	e := c.World.Entities[x.CivId]
	o := e.Occasion[x.OccasionId]
	s := o.Schedule[x.ScheduleId]
	return c.entity(x.CivId) + " held a " + strcase.ToDelimited(s.Type.String(), ' ') + c.site(x.SiteId, " in") + " as part of the " + o.Name() +
		". Competing " + util.If(len(x.CompetitorHfid) > 1, "were ", "was ") + c.hfList(x.CompetitorHfid) + ". " +
		util.Capitalize(c.hf(x.WinnerHfid)) + " was the victor"
}

func (x *HistoricalEventCreateEntityPosition) Html(c *Context) string {
	e := c.entity(x.Civ)
	if x.SiteCiv != x.Civ {
		e = c.entity(x.SiteCiv) + " of " + e
	}
	if x.Histfig != -1 {
		e = c.hf(x.Histfig) + " of " + e
	} else {
		e = "members of " + e
	}
	switch x.Reason {
	case HistoricalEventCreateEntityPositionReason_AsAMatterOfCourse:
		return e + " created the position of " + x.Position + " as a matter of course"
	case HistoricalEventCreateEntityPositionReason_Collaboration:
		return e + " collaborated to create the position of " + x.Position
	case HistoricalEventCreateEntityPositionReason_ForceOfArgument:
		return e + " created the position of " + x.Position + " trough force of argument"
	case HistoricalEventCreateEntityPositionReason_ThreatOfViolence:
		return e + " compelled the creation of the position of " + x.Position + " with threats of violence"
	case HistoricalEventCreateEntityPositionReason_WaveOfPopularSupport:
		return e + " created the position of " + x.Position + ", pushed by a wave of popular support"
	}
	return e + " created the position of " + x.Position
}

func (x *HistoricalEventCreatedSite) Html(c *Context) string {
	f := util.If(x.ResidentCivId != -1, " for "+c.entity(x.ResidentCivId), "")
	if x.BuilderHfid != -1 {
		return c.hf(x.BuilderHfid) + " created " + c.site(x.SiteId, "") + f
	}
	return c.siteCiv(x.SiteCivId, x.CivId) + " founded " + c.site(x.SiteId, "") + f

}

func (x *HistoricalEventCreatedStructure) Html(c *Context) string { // TODO rebuild/rebuilt
	if x.BuilderHfid != -1 {
		return c.hf(x.BuilderHfid) + " thrust a spire of slade up from the underworld, naming it " + c.structure(x.SiteId, x.StructureId) +
			", and established a gateway between worlds in " + c.site(x.SiteId, "")
	}
	return c.siteCiv(x.SiteCivId, x.CivId) + util.If(x.Rebuilt, " rebuild ", " constructed ") + c.siteStructure(x.SiteId, x.StructureId, "")
}

func (x *HistoricalEventCreatedWorldConstruction) Html(c *Context) string {
	return c.siteCiv(x.SiteCivId, x.CivId) + " finished the contruction of " + c.worldConstruction(x.Wcid) +
		" connecting " + c.site(x.SiteId1, "") + " with " + c.site(x.SiteId2, "") +
		util.If(x.MasterWcid != -1, " as part of "+c.worldConstruction(x.MasterWcid), "")
}

func (x *HistoricalEventCreatureDevoured) Html(c *Context) string {
	return c.hf(x.Eater) + " devoured " + util.If(x.Victim != -1, c.hfRelated(x.Victim, x.Eater), articled(x.Race)) +
		util.If(x.Entity != -1, " of "+c.entity(x.Entity), "") +
		c.location(x.SiteId, " in", x.SubregionId, " in")
}

func (x *HistoricalEventDanceFormCreated) Html(c *Context) string {
	reason := ""
	switch x.Reason {
	case HistoricalEventDanceFormCreatedReason_GlorifyHf:
		reason = " in order to glorify " + c.hfRelated(x.ReasonId, x.HistFigureId)
	}
	circumstance := ""
	switch x.Circumstance {
	case HistoricalEventDanceFormCreatedCircumstance_Dream:
		circumstance = " after a dream"
	case HistoricalEventDanceFormCreatedCircumstance_DreamAboutHf:
		circumstance = " after a dreaming about " + util.If(x.ReasonId == x.CircumstanceId, c.hfShort(x.CircumstanceId), c.hfRelated(x.CircumstanceId, x.HistFigureId))
	case HistoricalEventDanceFormCreatedCircumstance_Nightmare:
		circumstance = " after a nightmare"
	case HistoricalEventDanceFormCreatedCircumstance_PrayToHf:
		circumstance = " after praying to " + util.If(x.ReasonId == x.CircumstanceId, c.hfShort(x.CircumstanceId), c.hfRelated(x.CircumstanceId, x.HistFigureId))
	}
	return c.danceForm(x.FormId) + " was created by " + c.hf(x.HistFigureId) + c.location(x.SiteId, " in", x.SubregionId, " in") + reason + circumstance
}

func (x *HistoricalEventDestroyedSite) Html(c *Context) string {
	return c.entity(x.AttackerCivId) + " defeated " + c.siteCiv(x.SiteCivId, x.DefenderCivId) + " and destroyed " + c.site(x.SiteId, "")
}

func (x *HistoricalEventDiplomatLost) Html(c *Context) string {
	return c.entity(x.Entity) + " lost a diplomant in " + c.site(x.SiteId, "") + ". They suspected the involvement of " + c.entity(x.Involved)
}

func (x *HistoricalEventEntityAllianceFormed) Html(c *Context) string {
	return c.entityList(x.JoiningEnid) + " swore to support " + c.entity(x.InitiatingEnid) + " in war if the latter did likewise"
}

func (x *HistoricalEventEntityBreachFeatureLayer) Html(c *Context) string {
	return c.siteCiv(x.SiteEntityId, x.CivEntityId) + " breached the Underworld at " + c.site(x.SiteId, "")
}

func (x *HistoricalEventEntityCreated) Html(c *Context) string {
	if x.CreatorHfid != -1 {
		return c.hf(x.CreatorHfid) + " formed " + c.entity(x.EntityId) + c.siteStructure(x.SiteId, x.StructureId, "in")
	} else {
		return c.entity(x.EntityId) + " formed" + c.siteStructure(x.SiteId, x.StructureId, "in")
	}
}

func (x *HistoricalEventEntityDissolved) Html(c *Context) string {
	return c.entity(x.EntityId) + " dissolved after " + x.Reason.String()
}

func (x *HistoricalEventEntityEquipmentPurchase) Html(c *Context) string { // todo check hfid
	return c.entity(x.EntityId) + " purchased " + equipmentLevel(x.NewEquipmentLevel) + " equipment"
}

func (x *HistoricalEventEntityExpelsHf) Html(c *Context) string {
	return c.entity(x.EntityId) + " expelled " + c.hf(x.Hfid) + c.site(x.SiteId, " from")
}

func (x *HistoricalEventEntityFledSite) Html(c *Context) string {
	return c.entity(x.FledCivId) + " fled " + c.site(x.SiteId, "")
}

func (x *HistoricalEventEntityIncorporated) Html(c *Context) string { // TODO site
	return c.entity(x.JoinerEntityId) + util.If(x.PartialIncorporation, " began operating at the direction of ", " fully incorporated into ") +
		c.entity(x.JoinedEntityId) + " under the leadership of " + c.hf(x.LeaderHfid)
}

func (x *HistoricalEventEntityLaw) Html(c *Context) string {
	switch x.LawAdd {
	case HistoricalEventEntityLawLawAdd_Harsh:
		return c.hf(x.HistFigureId) + " laid a series of oppressive edicts upon " + c.entity(x.EntityId)
	}
	switch x.LawRemove {
	case HistoricalEventEntityLawLawRemove_Harsh:
		return c.hf(x.HistFigureId) + " lifted numerous oppressive laws from " + c.entity(x.EntityId)
	}
	return c.hf(x.HistFigureId) + " UNKNOWN LAW upon " + c.entity(x.EntityId)
}

func (x *HistoricalEventEntityOverthrown) Html(c *Context) string {
	return c.hf(x.InstigatorHfid) + " toppled the government of " + util.If(x.OverthrownHfid != -1, c.hfRelated(x.OverthrownHfid, x.InstigatorHfid)+" of ", "") + c.entity(x.EntityId) + " and " +
		util.If(x.PosTakerHfid == x.InstigatorHfid, "assumed control", "placed "+c.hfRelated(x.PosTakerHfid, x.InstigatorHfid)+" in power") + c.site(x.SiteId, " in") +
		util.If(len(x.ConspiratorHfid) > 0, ". The support of "+c.hfListRelated(x.ConspiratorHfid, x.InstigatorHfid)+" was crucial to the coup", "")
}

func (x *HistoricalEventEntityPersecuted) Html(c *Context) string {
	var l []string
	if len(x.ExpelledHfid) > 0 {
		l = append(l, c.hfListRelated(x.ExpelledHfid, x.PersecutorHfid)+util.If(len(x.ExpelledHfid) > 1, " were", " was")+" expelled")
	}
	if len(x.PropertyConfiscatedFromHfid) > 0 {
		l = append(l, "most property was confiscated")
	}
	if x.DestroyedStructureId != -1 {
		l = append(l, c.structure(x.SiteId, x.DestroyedStructureId)+" was destroyed"+util.If(x.ShrineAmountDestroyed > 0, " along with several smaller sacred sites", ""))
	} else if x.ShrineAmountDestroyed > 0 {
		l = append(l, "some sacred sites were desecrated")
	}
	return c.hf(x.PersecutorHfid) + " of " + c.entity(x.PersecutorEnid) + " persecuted " + c.entity(x.TargetEnid) + " in " + c.site(x.SiteId, "") +
		util.If(len(l) > 0, ". "+util.Capitalize(andList(l)), "")
}

func (x *HistoricalEventEntityPrimaryCriminals) Html(c *Context) string { // TODO structure
	return c.entity(x.EntityId) + " became the primary criminal organization in " + c.site(x.SiteId, "")
}

func (x *HistoricalEventEntityRampagedInSite) Html(c *Context) string {
	return "the forces of " + c.entity(x.RampageCivId) + " rampaged throughout " + c.site(x.SiteId, "")
}

func (x *HistoricalEventEntityRelocate) Html(c *Context) string {
	return c.entity(x.EntityId) + " moved" + c.siteStructure(x.SiteId, x.StructureId, "to")
}

func (x *HistoricalEventEntitySearchedSite) Html(c *Context) string {
	return c.entity(x.SearcherCivId) + " searched " + c.site(x.SiteId, "") +
		util.If(x.Result == HistoricalEventEntitySearchedSiteResult_FoundNothing, " and found nothing", "")
}

func (x *HistoricalEventFailedFrameAttempt) Html(c *Context) string {
	return c.hf(x.FramerHfid) + " attempted to frame " + c.hfRelated(x.TargetHfid, x.FramerHfid) + " for " + x.Crime.String() +
		util.If(x.PlotterHfid != -1, " at the behest of "+c.hfRelated(x.PlotterHfid, x.FramerHfid), "") +
		" by fooling " + c.hfRelated(x.FooledHfid, x.FramerHfid) + " and " + c.entity(x.ConvicterEnid) +
		" with fabricated evidence, but nothing came of it"
}

func (x *HistoricalEventFailedIntrigueCorruption) Html(c *Context) string {
	action := ""
	switch x.Action {
	case HistoricalEventFailedIntrigueCorruptionAction_BribeOfficial:
		action = "have law enforcement look the other way"
	case HistoricalEventFailedIntrigueCorruptionAction_BringIntoNetwork:
		action = "have someone to act on plots and schemes"
	case HistoricalEventFailedIntrigueCorruptionAction_CorruptInPlace:
		action = "have an agent"
	case HistoricalEventFailedIntrigueCorruptionAction_InduceToEmbezzle:
		action = "secure embezzled funds"
	}
	method := ""
	switch x.Method {
	case HistoricalEventFailedIntrigueCorruptionMethod_BlackmailOverEmbezzlement:
		method = "made a blackmail threat, due to embezzlement using the position " + c.position(x.RelevantEntityId, x.RelevantPositionProfileId, x.CorruptorHfid) + " of " + c.entity(x.RelevantEntityId)
	case HistoricalEventFailedIntrigueCorruptionMethod_Bribe:
		method = "offered a bribe"
	case HistoricalEventFailedIntrigueCorruptionMethod_Flatter:
		method = "made flattering remarks"
	case HistoricalEventFailedIntrigueCorruptionMethod_Intimidate:
		method = "made a threat"
	case HistoricalEventFailedIntrigueCorruptionMethod_OfferImmortality:
		method = "offered immortality"
	case HistoricalEventFailedIntrigueCorruptionMethod_Precedence:
		method = "pulled rank as " + c.position(x.RelevantEntityId, x.RelevantPositionProfileId, x.CorruptorHfid) + " of " + c.entity(x.RelevantEntityId)
	case HistoricalEventFailedIntrigueCorruptionMethod_ReligiousSympathy:
		method = "played for sympathy" + util.If(x.RelevantIdForMethod != -1, " by appealing to shared worship of "+c.hfRelated(x.RelevantIdForMethod, x.CorruptorHfid), "")
	case HistoricalEventFailedIntrigueCorruptionMethod_RevengeOnGrudge:
		method = "offered revenge upon the persecutor " + c.hfRelated(x.RelevantIdForMethod, x.CorruptorHfid)
	}
	fail := "The plan failed"
	switch x.TopValue {
	case HistoricalEventFailedIntrigueCorruptionTopValue_Law:
		fail = c.hf(x.TargetHfid) + " valued the law and refused"
	case HistoricalEventFailedIntrigueCorruptionTopValue_Power:
	}
	switch x.TopFacet {
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Ambition:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_AnxietyPropensity:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Confidence:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_EnvyPropensity:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Fearlessness:
		fail += ", despite being afraid"
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Greed:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Hope:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Pride:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_StressVulnerability:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Swayable:
		fail += ", despite being swayed by the emotional appeal" // TODO relationship values
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Vanity:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Vengeful:
	}
	return c.hf(x.CorruptorHfid) + " attempted to corrupt " + c.hfRelated(x.TargetHfid, x.CorruptorHfid) +
		" in order to " + action + c.location(x.SiteId, " in", x.SubregionId, " in") + ". " +
		util.Capitalize(util.If(x.LureHfid != -1,
			c.hfRelated(x.LureHfid, x.CorruptorHfid)+" lured "+c.hfShort(x.TargetHfid)+" to a meeting with "+c.hfShort(x.CorruptorHfid)+", where the latter",
			c.hfShort(x.CorruptorHfid)+" met with "+c.hfShort(x.TargetHfid))) +
		util.If(x.FailedJudgmentTest, ", while completely misreading the situation,", "") + " " + method + ". " + fail
}

func (x *HistoricalEventFieldBattle) Html(c *Context) string {
	atk := c.entity(x.AttackerCivId)
	def := c.entity(x.DefenderCivId)
	generals := ""
	if x.AttackerGeneralHfid != -1 {
		generals += ". " + util.Capitalize(c.hf(x.AttackerGeneralHfid)) + " led the attack"
		if x.DefenderGeneralHfid != -1 {
			generals += ", and the defenders were led by " + c.hf(x.DefenderGeneralHfid)
		}
	}
	mercs := ""
	if x.AttackerMercEnid != -1 {
		mercs += fmt.Sprintf(". %s were hired by the attackers", c.entity(x.AttackerMercEnid))
	}
	if x.ASupportMercEnid != -1 {
		mercs += fmt.Sprintf(". %s were hired as scouts by the attackers", c.entity(x.ASupportMercEnid))
	}
	if x.DefenderMercEnid != -1 {
		mercs += fmt.Sprintf(". The defenders hired %s", c.entity(x.DefenderMercEnid))
	}
	if x.DSupportMercEnid != -1 {
		mercs += fmt.Sprintf(". The defenders hired %s as scouts", c.entity(x.DSupportMercEnid))
	}
	return fmt.Sprintf("%s attacked %s at %s%s%s", atk, def, c.region(x.SubregionId), generals, mercs)
}

func (x *HistoricalEventFirstContact) Html(c *Context) string {
	return c.entity(x.ContactorEnid) + " made contact with " + c.entity(x.ContactedEnid) + c.site(x.SiteId, " at")
}

func (x *HistoricalEventGamble) Html(c *Context) string {
	outcome := ""
	switch d := x.NewAccount - x.OldAccount; {
	case d <= -5000:
		outcome = "lost a fortune"
	case d <= -1000:
		outcome = "did poorly"
	case d <= 1000:
		outcome = "did well"
	case d <= 5000:
		outcome = "made a fortune"
	}
	return c.hf(x.GamblerHfid) + " " + outcome + " gambling" + c.siteStructure(x.SiteId, x.StructureId, " in") +
		util.If(x.OldAccount >= 0 && x.NewAccount < 0, " and went into debt", "")
}

func (x *HistoricalEventHfAbducted) Html(c *Context) string {
	return c.hf(x.TargetHfid) + " was abducted " + c.location(x.SiteId, "from", x.SubregionId, "from") + " by " + c.hfRelated(x.SnatcherHfid, x.TargetHfid)
}

func (x *HistoricalEventHfAttackedSite) Html(c *Context) string {
	return c.hf(x.AttackerHfid) + " attacked " + c.siteCiv(x.SiteCivId, x.DefenderCivId) + c.site(x.SiteId, " in")
}

func (x *HistoricalEventHfConfronted) Html(c *Context) string {
	return c.hf(x.Hfid) + " aroused " + x.Situation.String() + c.location(x.SiteId, " in", x.SubregionId, " in") + " after " +
		andList(util.Map(x.Reason, func(r HistoricalEventHfConfrontedReason) string {
			switch r {
			case HistoricalEventHfConfrontedReason_Ageless:
				return " appearing not to age"
			case HistoricalEventHfConfrontedReason_Murder:
				return "a murder"
			}
			return ""
		}))
}

func (x *HistoricalEventHfConvicted) Html(c *Context) string { // TODO no_prison_available, interrogator_hfid
	r := util.If(x.ConfessedAfterApbArrestEnid != -1, "after being recognized and arrested, ", "")
	switch {
	case x.SurveiledCoconspirator:
		r += "due to ongoing surveillance on a coconspiratior, " + c.hfRelated(x.CoconspiratorHfid, x.ConvictedHfid) + ", as the plot unfolded, "
	case x.SurveiledContact:
		r += "due to ongoing surveillance on the contact " + c.hfRelated(x.ContactHfid, x.ConvictedHfid) + " as the plot unfolded, "
	case x.SurveiledConvicted:
		r += "due to ongoing surveillance as the plot unfolded, "
	case x.SurveiledTarget:
		r += "due to ongoing surveillance on the target " + c.hfRelated(x.TargetHfid, x.ConvictedHfid) + " as the plot unfolded, "
	}
	r += c.hf(x.ConvictedHfid) + util.If(x.ConfessedAfterApbArrestEnid != -1, " confessed and", "") + " was " + util.If(x.WrongfulConviction, "wrongfully ", "") + "convicted " +
		util.If(x.ConvictIsContact, "as a go-between in a conspiracy to commit ", "of ") + x.Crime.String() + " by " + c.entity(x.ConvicterEnid)
	if x.FooledHfid != -1 {
		r += " after " + c.hfRelated(x.FramerHfid, x.ConvictedHfid) + " fooled " + c.hfRelated(x.FooledHfid, x.ConvictedHfid) + " with fabricated evidence" +
			util.If(x.PlotterHfid != -1, " at the behest of "+c.hfRelated(x.PlotterHfid, x.ConvictedHfid), "")
	}
	if x.CorruptConvicterHfid != -1 {
		r += " and the corrupt " + c.hfRelated(x.CorruptConvicterHfid, x.ConvictedHfid) + " through the machinations of " + c.hfRelated(x.PlotterHfid, x.ConvictedHfid)
	}
	var penaltiy []string
	penaltiy = append(penaltiy, r)
	if x.Beating {
		penaltiy = append(penaltiy, "beaten")
	}
	if x.Hammerstrokes > 0 {
		penaltiy = append(penaltiy, fmt.Sprintf("sentenced to %d hammerstrokes", x.Hammerstrokes))
	}
	if x.DeathPenalty {
		penaltiy = append(penaltiy, "sentenced to death")
	}
	if x.Exiled {
		penaltiy = append(penaltiy, "exiled")
	}
	if x.PrisonMonths > 0 {
		m := x.PrisonMonths % 12
		y := x.PrisonMonths / 12
		t := ""
		if m != 0 && y != 0 {
			t = fmt.Sprintf("%d %s and %d %s", y, util.If(y > 1, "years", "year"), m, util.If(m > 1, "months", "month"))
		} else if y != 0 {
			t = fmt.Sprintf("%d %s", y, util.If(y > 1, "years", "year"))
		} else {
			t = fmt.Sprintf("%d %s", m, util.If(m > 1, "months", "month"))
		}
		penaltiy = append(penaltiy, "imprisoned for a term of "+t)
	}
	r = andList(penaltiy)
	if x.HeldFirmInInterrogation {
		r += ". " + c.hfShort(x.ConvictedHfid) + " revealed nothing during interrogation"
	} else if len(x.ImplicatedHfid) > 0 {
		r += ". " + c.hfShort(x.ConvictedHfid) + " implicated " + c.hfList(x.ImplicatedHfid) + " during interrogation" +
			util.If(x.DidNotRevealAllInInterrogation, " but did not reveal eaverything", "")
	}
	return r
}

func (x *HistoricalEventHfDestroyedSite) Html(c *Context) string {
	return c.hf(x.AttackerHfid) + " routed " + c.siteCiv(x.SiteCivId, x.DefenderCivId) + " and destroyed " + c.site(x.SiteId, "")
}

func (x *HistoricalEventHfDied) Html(c *Context) string {
	hf := c.hf(x.Hfid)
	loc := c.location(x.SiteId, " in", x.SubregionId, " in")
	slayer := ""
	if x.SlayerHfid != -1 {
		slayer = " by " + c.hfRelated(x.SlayerHfid, x.Hfid)
	}

	if x.SlayerItemId != -1 {
		slayer += " with " + c.artifact(x.SlayerItemId)
	} else {
	}
	switch x.Cause {
	case HistoricalEventHfDiedCause_Behead, HistoricalEventHfDiedCause_ExecBeheaded:
		return hf + " was beheaded" + slayer + loc
	case HistoricalEventHfDiedCause_Bleed, HistoricalEventHfDiedCause_Blood:
		return hf + " bled to death, slain by " + slayer + loc
	case HistoricalEventHfDiedCause_BloodDrained, HistoricalEventHfDiedCause_DrainBlood:
		return hf + " was drained of blood by " + slayer + loc
	case HistoricalEventHfDiedCause_BurnAlive, HistoricalEventHfDiedCause_ExecBurnedAlive:
		return hf + " was burned alive" + slayer + loc
	case HistoricalEventHfDiedCause_BuryAlive, HistoricalEventHfDiedCause_ExecBuriedAlive:
		return hf + " was buried alive" + slayer + loc
	case HistoricalEventHfDiedCause_Cavein:
	case HistoricalEventHfDiedCause_Chasm:
		return hf + " fell into a deep chasm" + loc
	case HistoricalEventHfDiedCause_Collision, HistoricalEventHfDiedCause_Obstacle:
		return hf + " died after colliding with an obstacle, slain by " + slayer + loc
	case HistoricalEventHfDiedCause_Crucify, HistoricalEventHfDiedCause_ExecCrucified:
		return hf + " was crucified" + slayer + loc
	case HistoricalEventHfDiedCause_Crushed:
	case HistoricalEventHfDiedCause_CrushedBridge, HistoricalEventHfDiedCause_Drawbridge:
		return hf + " was crushed by a drawbridge" + loc
	case HistoricalEventHfDiedCause_Drown:
		return hf + " drowned" + loc
	case HistoricalEventHfDiedCause_DrownAlt, HistoricalEventHfDiedCause_ExecDrowned:
		return hf + " was drowned" + slayer + loc
	case HistoricalEventHfDiedCause_EncaseIce, HistoricalEventHfDiedCause_FreezingWater:
		return hf + " was encased in ice" + loc
	case HistoricalEventHfDiedCause_ExecGeneric, HistoricalEventHfDiedCause_ExecutionGeneric:
		return hf + " was executed" + slayer + loc
	case HistoricalEventHfDiedCause_FallingObject:
	case HistoricalEventHfDiedCause_FeedToBeasts, HistoricalEventHfDiedCause_ExecFedToBeasts:
		return hf + " was fed to beasts" + slayer + loc
	case HistoricalEventHfDiedCause_FlyingObject:
	case HistoricalEventHfDiedCause_HackToPieces, HistoricalEventHfDiedCause_ExecHackedToPieces:
		return hf + " was hacked to pieces" + slayer + loc
	case HistoricalEventHfDiedCause_Heat:
	case HistoricalEventHfDiedCause_Hunger:
		return hf + " starved" + loc
	case HistoricalEventHfDiedCause_Infection:
		return hf + " succumbed to infection" + loc
	case HistoricalEventHfDiedCause_Melt:
		return hf + " melted" + loc
	case HistoricalEventHfDiedCause_Murder, HistoricalEventHfDiedCause_Murdered:
		return hf + " was murdered" + slayer + loc
	case HistoricalEventHfDiedCause_OldAge:
		return hf + " died of old age" + loc
	case HistoricalEventHfDiedCause_PutToRest, HistoricalEventHfDiedCause_Memorialize:
		return hf + " was put to rest" + loc
	case HistoricalEventHfDiedCause_Quit:
	case HistoricalEventHfDiedCause_Quitdead:
	case HistoricalEventHfDiedCause_Scare:
	case HistoricalEventHfDiedCause_ScaredToDeath:
	case HistoricalEventHfDiedCause_Shot:
		return hf + " was shot and killed" + slayer + loc
	case HistoricalEventHfDiedCause_Slaughter, HistoricalEventHfDiedCause_Slaughtered:
		return hf + " was slaughtered by " + slayer + loc
	case HistoricalEventHfDiedCause_Spike:
	case HistoricalEventHfDiedCause_Spikes:
	case HistoricalEventHfDiedCause_Struck, HistoricalEventHfDiedCause_StruckDown:
		return hf + " was struck down" + slayer + loc
	case HistoricalEventHfDiedCause_Suffocate, HistoricalEventHfDiedCause_Air:
		return hf + " suffocated, slain by " + slayer + loc
	case HistoricalEventHfDiedCause_SuicideDrowned, HistoricalEventHfDiedCause_DrownAltTwo:
		return hf + " drowned " + util.If(c.World.HistoricalFigures[x.Hfid].Female(), "herself ", "himself ") + loc
	case HistoricalEventHfDiedCause_SuicideLeaping, HistoricalEventHfDiedCause_LeaptFromHeight:
		return hf + " leapt from a great height" + loc
	case HistoricalEventHfDiedCause_Thirst:
		return hf + " died of thirst" + loc
	case HistoricalEventHfDiedCause_Trap:
		return hf + " was killed by a trap" + loc
	case HistoricalEventHfDiedCause_Vanish:
	}
	return hf + " died: " + x.Cause.String() + slayer + loc
}

func (x *HistoricalEventHfDisturbedStructure) Html(c *Context) string {
	return c.hf(x.HistFigId) + " disturbed " + c.siteStructure(x.SiteId, x.StructureId, "")
}

func (x *HistoricalEventHfDoesInteraction) Html(c *Context) string { // TODO ignore source
	i := strings.Index(x.InteractionAction, " ")
	if i > 0 {
		return c.hf(x.DoerHfid) + " " + x.InteractionAction[:i+1] + c.hfRelated(x.TargetHfid, x.DoerHfid) + x.InteractionAction[i:] + util.If(x.Site != -1, c.site(x.Site, " in"), "")
	} else {
		return c.hf(x.DoerHfid) + " UNKNOWN INTERACTION " + c.hfRelated(x.TargetHfid, x.DoerHfid) + util.If(x.Site != -1, c.site(x.Site, " in"), "")
	}
}

func (x *HistoricalEventHfEnslaved) Html(c *Context) string {
	return c.hf(x.SellerHfid) + " sold " + c.hfRelated(x.EnslavedHfid, x.SellerHfid) + " to " + c.entity(x.PayerEntityId) + c.site(x.MovedToSiteId, " in")
}

func (x *HistoricalEventHfEquipmentPurchase) Html(c *Context) string { // TODO site, structure, region
	return c.hf(x.GroupHfid) + " purchased " + equipmentLevel(x.Quality) + " equipment"
}

func (x *HistoricalEventHfFreed) Html(c *Context) string {
	return util.If(x.FreeingHfid != -1, c.hf(x.FreeingHfid), "the forces") + " of " + c.entity(x.FreeingCivId) + " freed " + c.hfList(x.RescuedHfid) +
		c.site(x.SiteId, " from") + " and " + c.siteCiv(x.SiteCivId, x.HoldingCivId)
}

func (x *HistoricalEventHfGainsSecretGoal) Html(c *Context) string {
	switch x.SecretGoal {
	case HistoricalEventHfGainsSecretGoalSecretGoal_Immortality:
		return c.hf(x.Hfid) + " became obsessed with " + c.posessivePronoun(x.Hfid) + " own mortality and sought to extend " + c.posessivePronoun(x.Hfid) + " life by any means"
	}
	return c.hf(x.Hfid) + " UNKNOWN SECRET GOAL"
}

func (x *HistoricalEventHfInterrogated) Html(c *Context) string { // TODO wanted_and_recognized, held_firm_in_interrogation, implicated_hfid
	return c.hf(x.TargetHfid) + " was recognized and arrested by " + c.entity(x.ArrestingEnid) +
		". Despite the interrogation by " + c.hfRelated(x.InterrogatorHfid, x.TargetHfid) + ", " + c.hfShort(x.TargetHfid) + " refused to reveal anything and was released"
}

func (x *HistoricalEventHfLearnsSecret) Html(c *Context) string {
	if x.ArtifactId != -1 {
		return c.hf(x.StudentHfid) + " learned " + x.SecretText.String() + " from " + c.artifact(x.ArtifactId)
	} else {
		return c.hf(x.TeacherHfid) + " taught " + c.hfRelated(x.StudentHfid, x.TeacherHfid) + " " + x.SecretText.String()
	}
}

func (x *HistoricalEventHfNewPet) Html(c *Context) string {
	return c.hf(x.GroupHfid) + " tamed " + articled(x.Pets) + c.location(x.SiteId, " of", x.SubregionId, " of")
}
func (x *HistoricalEventHfPerformedHorribleExperiments) Html(c *Context) string {
	return c.hf(x.GroupHfid) + " performed horrible experiments " + c.place(x.StructureId, x.SiteId, " in", x.SubregionId, " in")
}
func (x *HistoricalEventHfPrayedInsideStructure) Html(c *Context) string {
	return c.hf(x.HistFigId) + " prayed " + c.siteStructure(x.SiteId, x.StructureId, "inside")
}

func (x *HistoricalEventHfPreach) Html(c *Context) string { // relevant site
	topic := ""
	switch x.Topic {
	case HistoricalEventHfPreachTopic_Entity1ShouldLoveEntityTwo:
		topic = ", urging love to be shown to "
	case HistoricalEventHfPreachTopic_SetEntity1AgainstEntityTwo:
		topic = ", inveighing against "
	}
	return c.hf(x.SpeakerHfid) + " preached to " + c.entity(x.Entity1) + topic + c.entity(x.Entity2) + c.site(x.SiteHfid, " in")
}

func (x *HistoricalEventHfProfanedStructure) Html(c *Context) string {
	return c.hf(x.HistFigId) + " profaned " + c.siteStructure(x.SiteId, x.StructureId, "")
}

func (x *HistoricalEventHfRansomed) Html(c *Context) string {
	return c.hf(x.RansomerHfid) + " ransomed " + c.hfRelated(x.RansomedHfid, x.RansomerHfid) + " to " + util.If(x.PayerHfid != -1, c.hfRelated(x.PayerHfid, x.RansomerHfid), c.entity(x.PayerEntityId)) +
		". " + c.hfShort(x.RansomedHfid) + " was sent " + c.site(x.MovedToSiteId, "to")
}

func (x *HistoricalEventHfReachSummit) Html(c *Context) string {
	id, _, _ := util.FindInMap(c.World.MountainPeaks, func(m *MountainPeak) bool { return m.Coords == x.Coords })
	return c.hfList(x.GroupHfid) + util.If(len(x.GroupHfid) > 1, " were", " was") + " the first to reach the summit of " + c.mountain(id) + " which rises above " + c.region(x.SubregionId)
}

func (x *HistoricalEventHfRecruitedUnitTypeForEntity) Html(c *Context) string {
	return c.hf(x.Hfid) + " recruited " + x.UnitType.String() + "s into " + c.entity(x.EntityId) + c.location(x.SiteId, " in", x.SubregionId, " in")
}

func (x *HistoricalEventHfRelationshipDenied) Html(c *Context) string {
	r := c.hf(x.SeekerHfid)
	switch x.Relationship {
	case HistoricalEventHfRelationshipDeniedRelationship_Apprentice:
		r += " was denied an apprenticeship under "
	default:
		r += " was denied an UNKNOWN RELATIONSHIP with "
	}
	r += c.hf(x.TargetHfid)
	switch x.Reason {
	case HistoricalEventHfRelationshipDeniedReason_Jealousy:
		r += " due to " + util.If(x.ReasonId != x.TargetHfid, c.hfRelated(x.ReasonId, x.SeekerHfid), "the latter") + "'s jealousy"
	case HistoricalEventHfRelationshipDeniedReason_PrefersWorkingAlone:
		r += " as " + util.If(x.ReasonId != x.TargetHfid, c.hfRelated(x.ReasonId, x.SeekerHfid), "the latter") + " prefers to work alone"
	}
	return r
}

func (x *HistoricalEventHfReunion) Html(c *Context) string {
	return c.hf(x.Group1Hfid) + " was reunited with " + c.hfListRelated(x.Group2Hfid, x.Group1Hfid) + c.location(x.SiteId, " in", x.SubregionId, " in")
}

func (x *HistoricalEventHfRevived) Html(c *Context) string {
	r := c.hf(x.Hfid)
	if x.ActorHfid != -1 {
		if x.Disturbance {
			r += " was disturbed from eternal rest by " + c.hfRelated(x.ActorHfid, x.Hfid)
		} else {
			r += " was brought back from the dead by " + c.hfRelated(x.ActorHfid, x.Hfid)
		}
	} else {
		r += " came back from the dead"
	}
	return r + util.If(x.RaisedBefore, " once more, this time", "") + " as " + articled(util.If(x.Ghost == HistoricalEventHfRevivedGhost_Unknown, "undead", x.Ghost.String())) +
		c.location(x.SiteId, " in", x.SubregionId, " in")
}

func (x *HistoricalEventHfSimpleBattleEvent) Html(c *Context) string {
	group1 := c.hf(x.Group1Hfid)
	group2 := c.hfRelated(x.Group2Hfid, x.Group1Hfid)
	loc := c.location(x.SiteId, " in", x.SubregionId, " in")
	switch x.Subtype {
	case HistoricalEventHfSimpleBattleEventSubtype_Ambushed:
		return group1 + " ambushed " + group2 + loc
	case HistoricalEventHfSimpleBattleEventSubtype_Attacked:
		return group1 + " attacked " + group2 + loc
	case HistoricalEventHfSimpleBattleEventSubtype_Confront:
		return group1 + " confronted " + group2 + loc
	case HistoricalEventHfSimpleBattleEventSubtype_Corner:
		return group1 + " cornered " + group2 + loc
	case HistoricalEventHfSimpleBattleEventSubtype_GotIntoABrawl:
		return group1 + " got into a brawl with " + group2 + loc
	case HistoricalEventHfSimpleBattleEventSubtype_HappenUpon:
		return group1 + " happened upon " + group2 + loc
	case HistoricalEventHfSimpleBattleEventSubtype_Scuffle:
		return group1 + " fought with " + group2 + loc + ". While defeated the latter escaped unscathed."
	case HistoricalEventHfSimpleBattleEventSubtype_Subdued:
		return group1 + " fought with and subdued " + group2 + loc
	case HistoricalEventHfSimpleBattleEventSubtype_Surprised:
		return group1 + " surprised " + group2 + loc
	case HistoricalEventHfSimpleBattleEventSubtype_TwoLostAfterGivingWounds:
		return group2 + " was forced to retreat from " + group1 + " despite the latters' wounds " + loc
	case HistoricalEventHfSimpleBattleEventSubtype_TwoLostAfterMutualWounds:
		return group2 + " eventually prevailed and " + group1 + " was forced to make a hasty escape" + loc
	case HistoricalEventHfSimpleBattleEventSubtype_TwoLostAfterReceivingWounds:
		return group2 + " managed to escape from " + group1 + "'s onslaught" + loc
	}
	return group1 + " attacked " + group2 + loc
}

func (x *HistoricalEventHfTravel) Html(c *Context) string {
	return c.hfList(x.GroupHfid) + util.If(x.Return, " returned", " made a journey") + c.location(x.SiteId, " to", x.SubregionId, " to")
}

func (x *HistoricalEventHfViewedArtifact) Html(c *Context) string {
	return c.hf(x.HistFigId) + " viewed " + c.artifact(x.ArtifactId) + c.siteStructure(x.SiteId, x.StructureId, " in")
}

func (x *HistoricalEventHfWounded) Html(c *Context) string {
	r := c.hf(x.WoundeeHfid)
	bp := "UNKNOWN BODYPART" // TODO bodyparts
	switch x.InjuryType {
	case HistoricalEventHfWoundedInjuryType_Rip:
		r += "'s " + bp + util.If(x.PartLost == HistoricalEventHfWoundedPartLost_True, " was torn out ", " was ripped ")
	case HistoricalEventHfWoundedInjuryType_Slash:
		r += "'s " + bp + util.If(x.PartLost == HistoricalEventHfWoundedPartLost_True, " was slashed off ", " was slashed ")
	case HistoricalEventHfWoundedInjuryType_Smash:
		r += "'s " + bp + util.If(x.PartLost == HistoricalEventHfWoundedPartLost_True, " was smashed off ", " was smashed ")
	case HistoricalEventHfWoundedInjuryType_Stab:
		r += "'s " + bp + util.If(x.PartLost == HistoricalEventHfWoundedPartLost_True, " was stabbed off ", " was stabbed ")
	default:
		r += " was wounded by "
	}

	return r + c.hfRelated(x.WounderHfid, x.WoundeeHfid) + c.location(x.SiteId, " in", x.SubregionId, " in") + util.If(x.WasTorture, " as a means of torture", "")
}

func (x *HistoricalEventHfsFormedIntrigueRelationship) Html(c *Context) string {
	if x.Circumstance == HistoricalEventHfsFormedIntrigueRelationshipCircumstance_IsEntitySubordinate {
		return c.hf(x.CorruptorHfid) + " subordinated " + c.hfRelated(x.TargetHfid, x.CorruptorHfid) + " as a member of " + c.entity(x.CircumstanceId) +
			" toward the fullfillment of plots and schemes" + c.location(x.SiteId, " in", x.SubregionId, " in")
	}

	action := ""
	switch x.Action {
	case HistoricalEventHfsFormedIntrigueRelationshipAction_BringIntoNetwork:
		action = "have someone to act on plots and schemes"
	case HistoricalEventHfsFormedIntrigueRelationshipAction_CorruptInPlace:
		action = "have an agent"
	}
	method := ""
	switch x.Method {
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_BlackmailOverEmbezzlement:
		method = "made a blackmail threat, due to embezzlement using the position " + c.position(x.RelevantEntityId, x.RelevantPositionProfileId, x.CorruptorHfid) + " of " + c.entity(x.RelevantEntityId)
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_Bribe:
		method = "offered a bribe"
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_Flatter:
		method = "made flattering remarks"
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_Intimidate:
		method = "made a threat"
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_OfferImmortality:
		method = "offered immortality"
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_Precedence:
		method = "pulled rank as " + c.position(x.RelevantEntityId, x.RelevantPositionProfileId, x.CorruptorHfid) + " of " + c.entity(x.RelevantEntityId)
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_ReligiousSympathy:
		method = "played for sympathy" + util.If(x.RelevantIdForMethod != -1, " by appealing to shared worship of "+c.hfRelated(x.RelevantIdForMethod, x.CorruptorHfid), "")
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_RevengeOnGrudge:
		method = "offered revenge upon the persecutor " + c.hfRelated(x.RelevantIdForMethod, x.CorruptorHfid)
	}
	success := "The plan worked"
	switch x.TopValue {
	case HistoricalEventHfsFormedIntrigueRelationshipTopValue_Law:
		// success = c.hf(x.TargetHfid) + " valued the law and refused"
	case HistoricalEventHfsFormedIntrigueRelationshipTopValue_Power:
	}
	switch x.TopFacet {
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Ambition:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_AnxietyPropensity:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Confidence:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_EnvyPropensity:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Fearlessness:
		// success += ", despite being afraid"
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Greed:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Hope:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Pride:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_StressVulnerability:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Swayable:
		// success += ", despite being swayed by the emotional appeal" // TODO relationship values
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Vanity:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Vengeful:
	}
	return c.hf(x.CorruptorHfid) + " corrupted " + c.hfRelated(x.TargetHfid, x.CorruptorHfid) +
		" in order to " + action + c.location(x.SiteId, " in", x.SubregionId, " in") + ". " +
		util.Capitalize(util.If(x.LureHfid != -1,
			c.hfRelated(x.LureHfid, x.CorruptorHfid)+" lured "+c.hfShort(x.TargetHfid)+" to a meeting with "+c.hfShort(x.CorruptorHfid)+", where the latter",
			c.hfShort(x.CorruptorHfid)+" met with "+c.hfShort(x.TargetHfid))) +
		" and " + method + ". " + success
}

func (x *HistoricalEventHfsFormedReputationRelationship) Html(c *Context) string {
	hf1 := c.hf(x.Hfid1) + util.If(x.IdentityId1 != -1, " as "+c.fullIdentity(x.IdentityId1), "")
	hf2 := c.hfRelated(x.Hfid2, x.Hfid1) + util.If(x.IdentityId2 != -1, " as "+c.fullIdentity(x.IdentityId2), "")
	loc := c.location(x.SiteId, " in", x.SubregionId, " in")
	switch x.HfRep2Of1 {
	case HistoricalEventHfsFormedReputationRelationshipHfRep2Of1_Friendly:
		return hf1 + " and " + hf2 + ", formed a false friendship where each used the other for information" + loc
	case HistoricalEventHfsFormedReputationRelationshipHfRep2Of1_InformationSource:
		return hf1 + ", formed a false friendship with " + hf2 + " in order to extract information" + loc
	}
	return hf1 + " and " + hf2 + ", formed an UNKNOWN RELATIONSHIP" + loc
}

func (x *HistoricalEventHolyCityDeclaration) Html(c *Context) string {
	return c.entity(x.ReligionId) + " declared " + c.site(x.SiteId, "") + " to be a holy site"
}
func (x *HistoricalEventInsurrectionStarted) Html(c *Context) string {
	e := util.If(x.TargetCivId != -1, c.entity(x.TargetCivId), "the local government")
	switch x.Outcome {
	case HistoricalEventInsurrectionStartedOutcome_LeadershipOverthrown:
		return "the insurrection " + c.site(x.SiteId, "in") + " concluded with " + e + " overthrown"
	case HistoricalEventInsurrectionStartedOutcome_PopulationGone:
		return "an insurrection " + c.site(x.SiteId, "in") + " against " + e + " ended with the disappearance of the rebelling population"
	default:
		return "an insurrection against " + e + " began " + c.site(x.SiteId, "in")
	}
}
func (x *HistoricalEventItemStolen) Html(c *Context) string {
	i := util.If(x.Item != -1, c.artifact(x.Item), articled(x.Mat+" "+x.ItemType))
	circumstance := ""
	if x.Circumstance != nil {
		switch x.Circumstance.Type {
		case HistoricalEventItemStolenCircumstanceType_Defeated:
			circumstance = " after defeating " + c.hfRelated(x.Circumstance.Defeated, x.Histfig)
		case HistoricalEventItemStolenCircumstanceType_Histeventcollection: // TODO during ...
		case HistoricalEventItemStolenCircumstanceType_Murdered:
			circumstance = " after murdering " + c.hfRelated(x.Circumstance.Defeated, x.Histfig)
		}
	}

	switch x.TheftMethod {
	case HistoricalEventItemStolenTheftMethod_Confiscated:
		return i + " was confiscated by " + c.hf(x.Histfig) + circumstance + util.If(x.Site != -1, c.site(x.Site, " in"), "")
	case HistoricalEventItemStolenTheftMethod_Looted:
		return i + " was looted " + util.If(x.Site != -1, c.site(x.Site, " from"), "") + " by " + c.hf(x.Histfig) + circumstance
	case HistoricalEventItemStolenTheftMethod_Recovered:
		return i + " was recovered by " + c.hf(x.Histfig) + circumstance + util.If(x.Site != -1, c.site(x.Site, " in"), "")
	}
	return i + " was stolen " + c.siteStructure(x.Site, x.Structure, "from") + " by " + c.hf(x.Histfig) + circumstance +
		util.If(x.StashSite != -1, " and brought "+c.site(x.StashSite, "to"), "")
}

func (x *HistoricalEventKnowledgeDiscovered) Html(c *Context) string {
	return c.hf(x.Hfid) + util.If(x.First, " was the very first to discover ", " independently discovered ") + x.Knowledge
}

func (x *HistoricalEventMasterpieceArchConstructed) Html(c *Context) string {
	return c.hf(x.Hfid) + " constructed a masterful " +
		util.If(x.BuildingSubtype != HistoricalEventMasterpieceArchConstructedBuildingSubtype_Unknown, x.BuildingSubtype.String(), x.BuildingType.String()) +
		" for " + c.entity(x.EntityId) + c.site(x.SiteId, " in")
}
func (x *HistoricalEventMasterpieceDye) Html(c *Context) string {
	return c.hf(x.Hfid) + " masterfully dyed a " + x.Mat.String() + " " + x.ItemType.String() + " with " + x.DyeMat +
		" for " + c.entity(x.EntityId) + c.site(x.SiteId, " in")
}
func (x *HistoricalEventMasterpieceEngraving) Html(c *Context) string {
	return c.hf(x.Hfid) + " created a masterful " +
		"engraving" +
		" for " + c.entity(x.EntityId) + c.site(x.SiteId, " in")
}
func (x *HistoricalEventMasterpieceFood) Html(c *Context) string {
	return c.hf(x.Hfid) + " prepared a masterful " +
		x.ItemSubtype.String() +
		" for " + c.entity(x.EntityId) + c.site(x.SiteId, " in")
}
func (x *HistoricalEventMasterpieceItem) Html(c *Context) string {
	return c.hf(x.Hfid) + " created a masterful " +
		x.Mat + " " + util.If(x.ItemSubtype != "", x.ItemSubtype, x.ItemType) +
		" for " + c.entity(x.EntityId) + c.site(x.SiteId, " in")
}
func (x *HistoricalEventMasterpieceItemImprovement) Html(c *Context) string {
	i := ""
	switch x.ImprovementType {
	case HistoricalEventMasterpieceItemImprovementImprovementType_ArtImage:
		i = "a masterful image in " + x.ImpMat
	case HistoricalEventMasterpieceItemImprovementImprovementType_Bands:
		i = "masterful bands in " + x.ImpMat
	case HistoricalEventMasterpieceItemImprovementImprovementType_Covered:
		i = "a masterful covering in " + x.ImpMat
	case HistoricalEventMasterpieceItemImprovementImprovementType_Itemspecific: // TODO check subtypes
		i = "a masterful handle in " + x.ImpMat
	case HistoricalEventMasterpieceItemImprovementImprovementType_RingsHanging:
		i = "masterful rings in " + x.ImpMat
	case HistoricalEventMasterpieceItemImprovementImprovementType_Spikes:
		i = "masterful spikes in " + x.ImpMat
	}
	return c.hf(x.Hfid) + " added " + i + " to " +
		articled(x.Mat+" "+util.If(x.ItemSubtype != "", x.ItemSubtype, x.ItemType)) +
		" for " + c.entity(x.EntityId) + c.site(x.SiteId, " in")
}
func (x *HistoricalEventMasterpieceLost) Html(c *Context) string {
	if e, ok := c.World.HistoricalEvents[x.CreationEvent]; ok {
		switch y := e.Details.(type) {
		case *HistoricalEventMasterpieceArchConstructed:
			return "the " + util.If(y.BuildingSubtype != HistoricalEventMasterpieceArchConstructedBuildingSubtype_Unknown, y.BuildingSubtype.String(), y.BuildingType.String()) +
				" masterfully constructed by " + c.hf(y.Hfid) + " for " + c.entity(y.EntityId) + c.site(x.Site, " at") + " in " + Time(e.Year, e.Seconds72) +
				" was destroyed" + util.If(x.Histfig != -1, " by "+c.hfRelated(x.Histfig, y.Hfid), "") +
				" by " + x.Method + c.site(x.Site, " at")
		case *HistoricalEventMasterpieceEngraving:
			return "a masterful engraving created by " + c.hf(y.Hfid) + " for " + c.entity(y.EntityId) + c.site(x.Site, " at") + " in " + Time(e.Year, e.Seconds72) +
				" was destroyed" + util.If(x.Histfig != -1, " by "+c.hfRelated(x.Histfig, y.Hfid), "") +
				" by " + x.Method + c.site(x.Site, " at")
		case *HistoricalEventMasterpieceItem:
			return "the masterful " +
				y.Mat + " " + util.If(y.ItemSubtype != "", y.ItemSubtype, y.ItemType) +
				" created by " + c.hf(y.Hfid) + " for " + c.entity(y.EntityId) + c.site(x.Site, " at") + " in " + Time(e.Year, e.Seconds72) +
				" was destroyed" + util.If(x.Histfig != -1, " by "+c.hfRelated(x.Histfig, y.Hfid), "") +
				" by " + x.Method + c.site(x.Site, " at")
		default:
			return c.hf(x.Histfig) + " destroyed a masterful item" + c.site(x.Site, " in") + " -- " + fmt.Sprintf("%T", e.Details)
		}
	}
	return c.hf(x.Histfig) + " destroyed a masterful item" + c.site(x.Site, " in")
}

func (x *HistoricalEventMerchant) Html(c *Context) string {
	return "merchants from " + c.entity(x.TraderEntityId) + " visited " + c.entity(x.DepotEntityId) + c.site(x.SiteId, " at") +
		util.If(x.Hardship, " and suffered great hardship", "") +
		util.If(x.LostValue, ". They reported irregularities with their goods", "")
}

func (x *HistoricalEventModifiedBuilding) Html(c *Context) string {
	return c.hf(x.ModifierHfid) + " had " + articled(x.Modification.String()) + " added " + c.siteStructure(x.SiteId, x.StructureId, "to")
}

func (x *HistoricalEventMusicalFormCreated) Html(c *Context) string {
	reason := ""
	switch x.Reason {
	case HistoricalEventMusicalFormCreatedReason_GlorifyHf:
		reason = " in order to glorify " + c.hfRelated(x.ReasonId, x.HistFigureId)
	}
	circumstance := ""
	switch x.Circumstance {
	case HistoricalEventMusicalFormCreatedCircumstance_Dream:
		circumstance = " after a dream"
	case HistoricalEventMusicalFormCreatedCircumstance_DreamAboutHf:
		circumstance = " after a dreaming about " + util.If(x.ReasonId == x.CircumstanceId, c.hfShort(x.CircumstanceId), c.hfRelated(x.CircumstanceId, x.HistFigureId))
	case HistoricalEventMusicalFormCreatedCircumstance_Nightmare:
		circumstance = " after a nightmare"
	case HistoricalEventMusicalFormCreatedCircumstance_PrayToHf:
		circumstance = " after praying to " + util.If(x.ReasonId == x.CircumstanceId, c.hfShort(x.CircumstanceId), c.hfRelated(x.CircumstanceId, x.HistFigureId))
	}
	return c.musicalForm(x.FormId) + " was created by " + c.hf(x.HistFigureId) + c.site(x.SiteId, " in") + reason + circumstance
}

func (x *HistoricalEventNewSiteLeader) Html(c *Context) string {
	return c.entity(x.AttackerCivId) + " defeated " + c.siteCiv(x.SiteCivId, x.DefenderCivId) + " and placed " + c.hf(x.NewLeaderHfid) + " in charge of" + c.site(x.SiteId, "") +
		". The new government was called " + c.entity(x.NewSiteCivId)
}

func (x *HistoricalEventPeaceAccepted) Html(c *Context) string {
	return c.entity(x.Destination) + " accepted an offer of peace from " + c.entity(x.Source)
}

func (x *HistoricalEventPeaceRejected) Html(c *Context) string {
	return c.entity(x.Destination) + " rejected an offer of peace from " + c.entity(x.Source)
}

func (x *HistoricalEventPerformance) Html(c *Context) string {
	r := c.entity(x.CivId) + " held "
	if e, ok := c.World.Entities[x.CivId]; ok {
		o := e.Occasion[x.OccasionId]
		s := o.Schedule[x.ScheduleId]
		r += c.schedule(s)
		r += " as part of " + o.Name()
		r += c.site(x.SiteId, " in")
		r += string(util.Json(s))
	}
	return r
}

func (x *HistoricalEventPlunderedSite) Html(c *Context) string { // TODO no_defeat_mention, took_items, took_livestock, was_raid
	return c.entity(x.AttackerCivId) + " defeated " + c.siteCiv(x.SiteCivId, x.DefenderCivId) + " and pillaged " + c.site(x.SiteId, "")
}

func (x *HistoricalEventPoeticFormCreated) Html(c *Context) string {
	circumstance := ""
	switch x.Circumstance {
	case HistoricalEventPoeticFormCreatedCircumstance_Dream:
		circumstance = " after a dream"
	case HistoricalEventPoeticFormCreatedCircumstance_Nightmare:
		circumstance = " after a nightmare"
	}
	return c.poeticForm(x.FormId) + " was created by " + c.hf(x.HistFigureId) + c.site(x.SiteId, " in") + circumstance
}

func (x *HistoricalEventProcession) Html(c *Context) string {
	r := c.entity(x.CivId) + " held a procession in " + c.site(x.SiteId, "")
	if e, ok := c.World.Entities[x.CivId]; ok {
		o := e.Occasion[x.OccasionId]
		r += " as part of " + o.Name()
		s := o.Schedule[x.ScheduleId]
		if s.Reference != -1 {
			r += ". It started at " + c.structure(x.SiteId, s.Reference)
			if s.Reference2 != -1 && s.Reference2 != s.Reference {
				r += " and ended at " + c.structure(x.SiteId, s.Reference2)
			} else {
				r += " and returned there after following its route"
			}
		}
		if len(s.Feature) > 0 {
			r += ". The event featured " + andList(util.Map(s.Feature, c.feature))
		}
		r += string(util.Json(s))
	}
	return r
}

func (x *HistoricalEventRazedStructure) Html(c *Context) string {
	return c.entity(x.CivId) + " razed " + c.siteStructure(x.SiteId, x.StructureId, "")
}

func (x *HistoricalEventReclaimSite) Html(c *Context) string {
	if x.Unretire {
		return c.siteCiv(x.SiteCivId, x.CivId) + " were taken by a mood to act against their judgment " + c.site(x.SiteId, "at")
	}
	return c.siteCiv(x.SiteCivId, x.CivId) + " launched an expedition to reclaim " + c.site(x.SiteId, "")
}

func (x *HistoricalEventRegionpopIncorporatedIntoEntity) Html(c *Context) string { // TODO Race
	return strconv.Itoa(x.PopNumberMoved) + " of " + strconv.Itoa(x.PopRace) + " from " + c.region(x.PopSrid) + " joined with " + c.entity(x.JoinEntityId) + c.site(x.SiteId, " at")
}

func (x *HistoricalEventRemoveHfEntityLink) Html(c *Context) string {
	hf := c.hf(x.Hfid)
	civ := c.entity(x.CivId)
	switch x.Link {
	case HistoricalEventRemoveHfEntityLinkLink_Member:
		return hf + " left " + civ
	case HistoricalEventRemoveHfEntityLinkLink_Position:
		return hf + " ceased to be the " + c.position(x.CivId, x.PositionId, x.Hfid) + " of " + civ
	case HistoricalEventRemoveHfEntityLinkLink_Prisoner:
		return hf + " escaped from the prisons of " + civ
	case HistoricalEventRemoveHfEntityLinkLink_Slave:
		return hf + " escaped from the slavery of " + civ
	}
	return hf + " left " + civ
}

func (x *HistoricalEventRemoveHfHfLink) Html(c *Context) string { // divorced
	return c.hf(x.Hfid) + " and " + c.hfRelated(x.HfidTarget, x.Hfid) + " broke up"
}

func (x *HistoricalEventRemoveHfSiteLink) Html(c *Context) string {
	switch x.LinkType {
	case HistoricalEventRemoveHfSiteLinkLinkType_HomeSiteAbstractBuilding:
		return c.hf(x.Histfig) + " moved out " + c.siteStructure(x.SiteId, x.Structure, "of")
	case HistoricalEventRemoveHfSiteLinkLinkType_Occupation:
		return c.hf(x.Histfig) + " stopped working " + c.siteStructure(x.SiteId, x.Structure, "at")
	case HistoricalEventRemoveHfSiteLinkLinkType_SeatOfPower:
		return c.hf(x.Histfig) + " stopped ruling " + c.siteStructure(x.SiteId, x.Structure, "from")
	}
	return c.hf(x.Histfig) + " stopped working " + c.siteStructure(x.SiteId, x.Structure, "at")
}

func (x *HistoricalEventReplacedStructure) Html(c *Context) string {
	return c.siteCiv(x.SiteCivId, x.CivId) + " replaced " + c.siteStructure(x.SiteId, x.OldAbId, "") + " with " + c.structure(x.SiteId, x.NewAbId)
}

func (x *HistoricalEventSiteDied) Html(c *Context) string {
	return c.siteCiv(x.SiteCivId, x.CivId) + " abandonned the settlement of " + c.site(x.SiteId, "")
}

func (x *HistoricalEventSiteDispute) Html(c *Context) string {
	return c.entity(x.EntityId1) + " of " + c.site(x.SiteId1, "") + " and " + c.entity(x.EntityId2) + " of " + c.site(x.SiteId2, "") + " became embroiled in a dispute over " + x.Dispute.String()
}

func (x *HistoricalEventSiteRetired) Html(c *Context) string {
	return c.siteCiv(x.SiteCivId, x.CivId) + " at the settlement " + c.site(x.SiteId, "of") + " regained their senses after " + util.If(x.First, "an initial", "another") + " period of questionable judgment"
}

func (x *HistoricalEventSiteSurrendered) Html(c *Context) string {
	return c.siteCiv(x.SiteCivId, x.DefenderCivId) + " surrendered " + c.site(x.SiteId, "") + " to " + c.entity(x.AttackerCivId)
}

func (x *HistoricalEventSiteTakenOver) Html(c *Context) string {
	return c.entity(x.AttackerCivId) + " defeated " + c.siteCiv(x.SiteCivId, x.DefenderCivId) + " and took over " + c.site(x.SiteId, "") + ". The new government was called " + c.entity(x.NewSiteCivId)
}

func (x *HistoricalEventSiteTributeForced) Html(c *Context) string {
	return c.entity(x.AttackerCivId) + " secured tribute from " + c.siteCiv(x.SiteCivId, x.DefenderCivId) +
		util.If(x.SiteId != -1, ", to be delivered"+c.site(x.SiteId, " from"), "") +
		util.If(x.Season != HistoricalEventSiteTributeForcedSeason_Unknown, " every "+x.Season.String(), "")
}

func (x *HistoricalEventSneakIntoSite) Html(c *Context) string {
	return util.If(x.AttackerCivId != -1, c.entity(x.AttackerCivId), "an unknown civilization") + " slipped " + c.site(x.SiteId, "into") +
		util.If(x.SiteCivId != -1 || x.DefenderCivId != -1, ", undetected by "+c.siteCiv(x.SiteCivId, x.DefenderCivId), "")
}

func (x *HistoricalEventSpottedLeavingSite) Html(c *Context) string {
	return c.hf(x.SpotterHfid) + " of " + c.entity(x.SiteCivId) + " spotted the forces of " + util.If(x.LeaverCivId != -1, c.entity(x.LeaverCivId), "an unknown civilization") + " slipping out of " + c.site(x.SiteId, "")
}

func (x *HistoricalEventSquadVsSquad) Html(c *Context) string { // TODO a_leader_hfid
	return c.hfList(x.AHfid) + " clashed with " +
		util.If(len(x.DHfid) > 0, c.hfList(x.DHfid), fmt.Sprintf("%d race_%d", x.DNumber, x.DRace)) +
		c.site(x.SiteId, " in") +
		util.If(x.DSlain > 0, fmt.Sprintf(", slaying %d", x.DSlain), "")
}

func plan(diff int) string { // TODO not exact
	switch {
	case diff > 100:
		return "unrolled a brilliant tactical plan"
	case diff > 30:
		return "put forth a sound plan"
	case diff > 0:
		return "used good tactics"
	case diff > -20:
		return "made a poor plan"
	case diff > -60:
		return "blundered terribly"
	default:
		return "made an outright foolish plan"
	}
}

func (x *HistoricalEventTacticalSituation) Html(c *Context) string {
	r := ""
	if x.ATacticianHfid == -1 && x.DTacticianHfid == -1 {
		r = "the forces shifted"
	} else if x.ATacticianHfid != -1 && x.DTacticianHfid == -1 {
		r += c.hf(x.ATacticianHfid) + " " + plan(x.ATacticsRoll-x.DTacticsRoll)
	} else if x.ATacticianHfid == -1 && x.DTacticianHfid != -1 {
		r += c.hf(x.DTacticianHfid) + " " + plan(x.DTacticsRoll-x.ATacticsRoll)
	} else {
		if x.ATacticsRoll < x.DTacticsRoll {
			r = c.hf(x.DTacticianHfid) + "'s tactical planning was superior to " + c.hf(x.ATacticianHfid) + "'s"
		} else {
			r = c.hf(x.ATacticianHfid) + " outmatched " + c.hf(x.DTacticianHfid) + " with a cunning plan"
		}
	}
	switch x.Situation {
	case HistoricalEventTacticalSituationSituation_AFavored: // TODO wording
	case HistoricalEventTacticalSituationSituation_ASlightlyFavored:
		r += ", " + util.If(x.DTacticsRoll > x.ATacticsRoll, "but", "and") + " the attackers had a slight positional advantage"
	case HistoricalEventTacticalSituationSituation_AStronglyFavored:
		r += ", " + util.If(x.DTacticsRoll > x.ATacticsRoll, "but", "and") + " the attackers had a strong positional advantage"
	case HistoricalEventTacticalSituationSituation_DSlightlyFavored:
		r += ", " + util.If(x.ATacticsRoll > x.DTacticsRoll, "but", "and") + " the defenders had a slight positional advantage"
	case HistoricalEventTacticalSituationSituation_DStronglyFavored:
		r += ", " + util.If(x.ATacticsRoll > x.DTacticsRoll, "but", "and") + " the defenders had a strong positional advantage"
	case HistoricalEventTacticalSituationSituation_NeitherFavored:
		r += ", but neither side had a positional advantage"
	}
	return r + c.site(x.SiteId, " in")
}

func (x *HistoricalEventTrade) Html(c *Context) string {
	outcome := ""
	switch d := x.AccountShift; {
	case d > 1000:
		outcome = " did well"
	case d < -1000:
		outcome = " did poorly"
	default:
		outcome = " broke even"
	}
	return c.hf(x.TraderHfid) + util.If(x.TraderEntityId != -1, " of "+c.entity(x.TraderEntityId), "") + outcome + " trading" + c.site(x.SourceSiteId, " from") + c.site(x.DestSiteId, " to")
}

func (x *HistoricalEventWrittenContentComposed) Html(c *Context) string {
	reason := ""
	switch x.Reason {
	case HistoricalEventWrittenContentComposedReason_GlorifyHf:
		reason = " in order to glorify " + c.hfRelated(x.ReasonId, x.HistFigureId)
	}
	circumstance := ""
	switch x.Circumstance {
	case HistoricalEventWrittenContentComposedCircumstance_Dream:
		circumstance = " after a dream"
	case HistoricalEventWrittenContentComposedCircumstance_DreamAboutHf:
		circumstance = " after a dreaming about " + util.If(x.ReasonId == x.CircumstanceId, c.hfShort(x.CircumstanceId), c.hfRelated(x.CircumstanceId, x.HistFigureId))
	case HistoricalEventWrittenContentComposedCircumstance_Nightmare:
		circumstance = " after a nightmare"
	case HistoricalEventWrittenContentComposedCircumstance_PrayToHf:
		circumstance = " after praying to " + util.If(x.ReasonId == x.CircumstanceId, c.hfShort(x.CircumstanceId), c.hfRelated(x.CircumstanceId, x.HistFigureId))
	}
	return c.writtenContent(x.WcId) + " was authored by " + c.hf(x.HistFigureId) + c.location(x.SiteId, " in", x.SubregionId, " in") + reason + circumstance
}
