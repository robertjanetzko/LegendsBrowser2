package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

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

func (x *HistoricalEventAddHfEntityHonor) Html(c *context) string {
	e := world.Entities[x.EntityId]
	h := e.Honor[x.HonorId]
	return fmt.Sprintf("%s received the title %s of %s%s", c.hf(x.Hfid), h.Name(), entity(x.EntityId), h.Requirement())
}

func (x *HistoricalEventAddHfEntityLink) Html(c *context) string {
	h := c.hf(x.Hfid)
	e := entity(x.CivId)

	if c.Story {
		return "the ascension of " + h + " to " + position(x.CivId, x.PositionId, x.Hfid) + " of " + e
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
		return h + " became " + position(x.CivId, x.PositionId, x.Hfid) + " of " + e
	case HistoricalEventAddHfEntityLinkLink_Prisoner:
		return h + " was imprisoned by " + e
	case HistoricalEventAddHfEntityLinkLink_Slave:
		return h + " was enslaved by " + e
	case HistoricalEventAddHfEntityLinkLink_Squad:
		return h + " became a hearthperson/solder of  " + e // TODO
	}
	return h + " became SOMETHING of " + e
}

func (x *HistoricalEventAddHfHfLink) Html(c *context) string {
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

func (x *HistoricalEventAddHfSiteLink) Html(c *context) string {
	h := c.hf(x.Histfig)
	e := ""
	if x.Civ != -1 {
		e = " of " + entity(x.Civ)
	}
	b := ""
	if x.Structure != -1 {
		b = " " + structure(x.SiteId, x.Structure)
	}
	s := site(x.SiteId, "in")
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

func (x *HistoricalEventAgreementFormed) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventAgreementFormed"
}

func (x *HistoricalEventAgreementMade) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventAgreementMade"
}

func (x *HistoricalEventAgreementRejected) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventAgreementRejected"
}

func (x *HistoricalEventArtifactClaimFormed) Html(c *context) string {
	a := artifact(x.ArtifactId)
	switch x.Claim {
	case HistoricalEventArtifactClaimFormedClaim_Heirloom:
		return a + " was made a family heirloom by " + c.hf(x.HistFigureId)
	case HistoricalEventArtifactClaimFormedClaim_Symbol:
		p := world.Entities[x.EntityId].Position(x.PositionProfileId).Name_
		e := entity(x.EntityId)
		return a + " was made a symbol of the " + p + " by " + e
	case HistoricalEventArtifactClaimFormedClaim_Treasure:
		circumstance := ""
		if x.Circumstance != HistoricalEventArtifactClaimFormedCircumstance_Unknown {
			circumstance = " " + x.Circumstance.String()
		}
		if x.HistFigureId != -1 {
			return a + " was claimed by " + c.hf(x.HistFigureId) + circumstance
		} else if x.EntityId != -1 {
			return a + " was claimed by " + entity(x.EntityId) + circumstance
		}
	}
	return a + " was claimed"
}

func (x *HistoricalEventArtifactCopied) Html(c *context) string {
	s := util.If(x.FromOriginal, "made a copy of the original", "aquired a copy of")
	return fmt.Sprintf("%s %s %s %s of %s, keeping it%s",
		entity(x.DestEntityId), s, artifact(x.ArtifactId), siteStructure(x.SourceSiteId, x.SourceStructureId, "from"),
		entity(x.SourceEntityId), siteStructure(x.DestSiteId, x.DestStructureId, "within"))
}

func (x *HistoricalEventArtifactCreated) Html(c *context) string {
	a := artifact(x.ArtifactId)
	h := c.hf(x.HistFigureId)
	s := ""
	if x.SiteId != -1 {
		s = site(x.SiteId, " in ")
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

func (x *HistoricalEventArtifactDestroyed) Html(c *context) string {
	return fmt.Sprintf("%s was destroyed by %s in %s", artifact(x.ArtifactId), entity(x.DestroyerEnid), site(x.SiteId, ""))
}

func (x *HistoricalEventArtifactFound) Html(c *context) string {
	w := ""
	if x.SiteId != -1 {
		w = site(x.SiteId, "")
		if x.SitePropertyId != -1 {
			w = property(x.SiteId, x.SitePropertyId) + " in " + w
		}
	}
	return fmt.Sprintf("%s was found in %s by %s", artifact(x.ArtifactId), w, util.If(x.HistFigureId != -1, c.hf(x.HistFigureId), "an unknown creature"))
}

func (x *HistoricalEventArtifactGiven) Html(c *context) string {
	r := ""
	if x.ReceiverHistFigureId != -1 {
		r = c.hf(x.ReceiverHistFigureId)
		if x.ReceiverEntityId != -1 {
			r += " of " + entity(x.ReceiverEntityId)
		}
	} else if x.ReceiverEntityId != -1 {
		r += entity(x.ReceiverEntityId)
	}
	g := ""
	if x.GiverHistFigureId != -1 {
		g = c.hf(x.GiverHistFigureId)
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
func (x *HistoricalEventArtifactLost) Html(c *context) string {
	w := ""
	if x.SubregionId != -1 {
		w = region(x.SubregionId)
	}
	if x.SiteId != -1 {
		w = site(x.SiteId, "")
		if x.SitePropertyId != -1 {
			w = property(x.SiteId, x.SitePropertyId) + " in " + w
		}
	}
	return fmt.Sprintf("%s was lost in %s", artifact(x.ArtifactId), w)
}

func (x *HistoricalEventArtifactPossessed) Html(c *context) string {
	a := artifact(x.ArtifactId)
	h := c.hf(x.HistFigureId)
	w := ""
	if x.SubregionId != -1 {
		w = region(x.SubregionId)
	}
	if x.SiteId != -1 {
		w = site(x.SiteId, "")
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
		return fmt.Sprintf("%s was aquired in %s by %s as a symbol of authority within %s%s", a, w, h, entity(x.ReasonId), circumstance)
	}
	return fmt.Sprintf("%s was claimed in %s by %s%s", a, w, h, circumstance) // TODO wording
}

func (x *HistoricalEventArtifactRecovered) Html(c *context) string {
	a := artifact(x.ArtifactId)
	h := c.hf(x.HistFigureId)
	w := ""
	if x.SubregionId != -1 {
		w = "in " + region(x.SubregionId)
	}
	if x.SiteId != -1 {
		w = site(x.SiteId, "in ")
		if x.StructureId != -1 {
			w = siteStructure(x.SiteId, x.StructureId, "from")
		}
	}
	return fmt.Sprintf("%s was recovered %s by %s", a, w, h)
}

func (x *HistoricalEventArtifactStored) Html(c *context) string { // TODO export siteProperty
	if x.HistFigureId != -1 {
		return fmt.Sprintf("%s stored %s in %s", c.hf(x.HistFigureId), artifact(x.ArtifactId), site(x.SiteId, ""))
	} else {
		return fmt.Sprintf("%s was stored in %s", artifact(x.ArtifactId), site(x.SiteId, ""))
	}
}

func (x *HistoricalEventArtifactTransformed) Html(c *context) string {
	return fmt.Sprintf("%s was made from %s by %s in %s", artifact(x.NewArtifactId), artifact(x.OldArtifactId), c.hf(x.HistFigureId), site(x.SiteId, "")) // TODO wording
}

func (x *HistoricalEventAssumeIdentity) Html(c *context) string {
	h := c.hf(x.TricksterHfid)
	i := identity(x.IdentityId)
	if x.TargetEnid == -1 {
		return fmt.Sprintf(`%s assumed the identity "%s"`, h, i)
	} else {
		return fmt.Sprintf(`%s fooled %s into believing %s was "%s"`, h, entity(x.TargetEnid), pronoun(x.TricksterHfid), i)
	}
}

func (x *HistoricalEventAttackedSite) Html(c *context) string {
	atk := entity(x.AttackerCivId)
	def := siteCiv(x.SiteCivId, x.DefenderCivId)
	generals := ""
	if x.AttackerGeneralHfid != -1 {
		generals += ". " + util.Capitalize(c.hf(x.AttackerGeneralHfid)) + " led the attack"
		if x.DefenderGeneralHfid != -1 {
			generals += ", and the defenders were led by " + c.hf(x.DefenderGeneralHfid)
		}
	}
	mercs := ""
	if x.AttackerMercEnid != -1 {
		mercs += fmt.Sprintf(". %s were hired by the attackers", entity(x.AttackerMercEnid))
	}
	if x.ASupportMercEnid != -1 {
		mercs += fmt.Sprintf(". %s were hired as scouts by the attackers", entity(x.ASupportMercEnid))
	}
	if x.DefenderMercEnid != -1 {
		mercs += fmt.Sprintf(". The defenders hired %s", entity(x.DefenderMercEnid))
	}
	if x.DSupportMercEnid != -1 {
		mercs += fmt.Sprintf(". The defenders hired %s as scouts", entity(x.DSupportMercEnid))
	}
	return fmt.Sprintf("%s attacked %s at %s%s%s", atk, def, site(x.SiteId, ""), generals, mercs)
}

func (x *HistoricalEventBodyAbused) Html(c *context) string {
	s := "the " + util.If(len(x.Bodies) > 1, "bodies", "body") + " of " + c.hfList(x.Bodies) + " " + util.If(len(x.Bodies) > 1, "were", "was")

	switch x.AbuseType {
	case HistoricalEventBodyAbusedAbuseType_Animated:
		s += " animated" + util.If(x.Histfig != -1, " by "+c.hf(x.Histfig), "") + site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Flayed:
		s += " flayed and the skin stretched over " + structure(x.SiteId, x.Structure) + " by " + entity(x.Civ) + site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Hung:
		s += " hung from a tree by " + entity(x.Civ) + site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Impaled:
		s += " impaled on " + articled(x.ItemMat+" "+x.ItemSubtype.String()) + " by " + entity(x.Civ) + site(x.SiteId, " in ")
	case HistoricalEventBodyAbusedAbuseType_Mutilated:
		s += " horribly mutilated by " + entity(x.Civ) + site(x.SiteId, " in ")
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
		s += " by " + entity(x.Civ) + site(x.SiteId, " in ")
	}

	return s
}

func (x *HistoricalEventBuildingProfileAcquired) Html(c *context) string {
	return util.If(x.AcquirerEnid != -1, entity(x.AcquirerEnid), c.hf(x.AcquirerHfid)) +
		util.If(x.PurchasedUnowned, " purchased ", " inherited ") +
		property(x.SiteId, x.BuildingProfileId) + site(x.SiteId, " in") +
		util.If(x.LastOwnerHfid != -1, " formerly owned by "+c.hfRelated(x.LastOwnerHfid, x.AcquirerHfid), "")
}

func (x *HistoricalEventCeremony) Html(c *context) string {
	r := entity(x.CivId) + " held a ceremony in " + site(x.SiteId, "")
	if e, ok := world.Entities[x.CivId]; ok {
		o := e.Occasion[x.OccasionId]
		r += " as part of " + o.Name()
		s := o.Schedule[x.ScheduleId]
		if len(s.Feature) > 0 {
			r += ". The event featured " + andList(util.Map(s.Feature, feature))
		}
	}
	return r
}

func (x *HistoricalEventChangeHfBodyState) Html(c *context) string {
	r := c.hf(x.Hfid)
	switch x.BodyState {
	case HistoricalEventChangeHfBodyStateBodyState_EntombedAtSite:
		r += " was entombed"
	}
	if x.StructureId != -1 {
		r += " within " + structure(x.SiteId, x.StructureId)
	}
	r += site(x.SiteId, " in ")
	return r
}

func (x *HistoricalEventChangeHfJob) Html(c *context) string {
	w := ""
	if x.SubregionId != -1 {
		w = " in " + region(x.SubregionId)
	}
	if x.SiteId != -1 {
		w = " in " + site(x.SiteId, "")
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

func (x *HistoricalEventChangeHfState) Html(c *context) string {
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
		r = " after a great deal of stress" // TODO check
	case HistoricalEventChangeHfStateReason_LackOfSleep:
		r = " after a lack of sleep" // TODO check
	case HistoricalEventChangeHfStateReason_OnAPilgrimage:
		r = " on a pilgrimage"
	case HistoricalEventChangeHfStateReason_Scholarship:
		r = " in order to pursue scholarship"
	case HistoricalEventChangeHfStateReason_UnableToLeaveLocation:
		r = " after being unable to leave the location" // TODO check
	}

	switch x.State {
	case HistoricalEventChangeHfStateState_Refugee:
		return c.hf(x.Hfid) + " fled " + location(x.SiteId, "to", x.SubregionId, "into")
	case HistoricalEventChangeHfStateState_Settled:
		switch x.Reason {
		case HistoricalEventChangeHfStateReason_BeWithMaster, HistoricalEventChangeHfStateReason_Scholarship:
			return c.hf(x.Hfid) + " moved to study " + site(x.SiteId, "in") + r
		case HistoricalEventChangeHfStateReason_Flight:
			return c.hf(x.Hfid) + " fled " + site(x.SiteId, "to")
		case HistoricalEventChangeHfStateReason_ConvictionExile, HistoricalEventChangeHfStateReason_ExiledAfterConviction:
			return c.hf(x.Hfid) + " departed " + site(x.SiteId, "to") + r
		case HistoricalEventChangeHfStateReason_None:
			return c.hf(x.Hfid) + " settled " + location(x.SiteId, "in", x.SubregionId, "in")
		}
	case HistoricalEventChangeHfStateState_Visiting:
		return c.hf(x.Hfid) + " visited " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateState_Wandering:
		if x.SubregionId != -1 {
			return c.hf(x.Hfid) + " began wandering " + region(x.SubregionId)
		} else {
			return c.hf(x.Hfid) + " began wandering the wilds"
		}
	}

	switch x.Mood {
	case HistoricalEventChangeHfStateMood_Berserk:
		return c.hf(x.Hfid) + " went berserk " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Catatonic:
		return c.hf(x.Hfid) + " stopped responding to the outside world " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Fell:
		return c.hf(x.Hfid) + " was taken by a fell mood " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Fey:
		return c.hf(x.Hfid) + " was taken by a fey mood " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Insane:
		return c.hf(x.Hfid) + " became crazed " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Macabre:
		return c.hf(x.Hfid) + " began to skulk and brood " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Melancholy:
		return c.hf(x.Hfid) + " was striken by melancholy " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Possessed:
		return c.hf(x.Hfid) + " was posessed " + site(x.SiteId, "in") + r
	case HistoricalEventChangeHfStateMood_Secretive:
		return c.hf(x.Hfid) + " withdrew from society " + site(x.SiteId, "in") + r
	}
	return "UNKNWON HistoricalEventChangeHfState"
}

func (x *HistoricalEventChangedCreatureType) Html(c *context) string {
	return c.hf(x.ChangerHfid) + " changed " + c.hfRelated(x.ChangeeHfid, x.ChangerHfid) + " from " + articled(x.OldRace) + " to " + articled(x.NewRace)
}

func (x *HistoricalEventCompetition) Html(c *context) string {
	e := world.Entities[x.CivId]
	o := e.Occasion[x.OccasionId]
	s := o.Schedule[x.ScheduleId]
	return entity(x.CivId) + " held a " + strcase.ToDelimited(s.Type.String(), ' ') + site(x.SiteId, " in") + " as part of the " + o.Name() +
		". Competing " + util.If(len(x.CompetitorHfid) > 1, "were ", "was ") + c.hfList(x.CompetitorHfid) + ". " +
		util.Capitalize(c.hf(x.WinnerHfid)) + " was the victor"
}

func (x *HistoricalEventCreateEntityPosition) Html(c *context) string {
	e := entity(x.Civ)
	if x.SiteCiv != x.Civ {
		e = entity(x.SiteCiv) + " of " + e
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

func (x *HistoricalEventCreatedSite) Html(c *context) string {
	f := util.If(x.ResidentCivId != -1, " for "+entity(x.ResidentCivId), "")
	if x.BuilderHfid != -1 {
		return c.hf(x.BuilderHfid) + " created " + site(x.SiteId, "") + f
	}
	return siteCiv(x.SiteCivId, x.CivId) + " founded " + site(x.SiteId, "") + f

}

func (x *HistoricalEventCreatedStructure) Html(c *context) string { // TODO rebuild/rebuilt
	if x.BuilderHfid != -1 {
		return c.hf(x.BuilderHfid) + " thrust a spire of slade up from the underworld, naming it " + structure(x.SiteId, x.StructureId) +
			", and established a gateway between worlds in " + site(x.SiteId, "")
	}
	return siteCiv(x.SiteCivId, x.CivId) + util.If(x.Rebuilt, " rebuild ", " constructed ") + siteStructure(x.SiteId, x.StructureId, "")
}

func (x *HistoricalEventCreatedWorldConstruction) Html(c *context) string {
	return siteCiv(x.SiteCivId, x.CivId) + " finished the contruction of " + worldConstruction(x.Wcid) +
		" connecting " + site(x.SiteId1, "") + " with " + site(x.SiteId2, "") +
		util.If(x.MasterWcid != -1, " as part of "+worldConstruction(x.MasterWcid), "")
}

func (x *HistoricalEventCreatureDevoured) Html(c *context) string {
	return c.hf(x.Eater) + " devoured " + util.If(x.Victim != -1, c.hfRelated(x.Victim, x.Eater), articled(x.Race)) +
		util.If(x.Entity != -1, " of "+entity(x.Entity), "") +
		location(x.SiteId, " in", x.SubregionId, " in")
}

func (x *HistoricalEventDanceFormCreated) Html(c *context) string {
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
	return danceForm(x.FormId) + " was created by " + c.hf(x.HistFigureId) + location(x.SiteId, " in", x.SubregionId, " in") + reason + circumstance
}

func (x *HistoricalEventDestroyedSite) Html(c *context) string { // TODO NoDefeatMention
	return entity(x.AttackerCivId) + " defeated " + siteCiv(x.SiteCivId, x.DefenderCivId) + " and destroyed " + site(x.SiteId, "")
}

func (x *HistoricalEventDiplomatLost) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventDiplomatLost"
}

func (x *HistoricalEventEntityAllianceFormed) Html(c *context) string {
	return entityList(x.JoiningEnid) + " swore to support " + entity(x.InitiatingEnid) + " in war if the latter did likewise"
}

func (x *HistoricalEventEntityBreachFeatureLayer) Html(c *context) string {
	return siteCiv(x.SiteEntityId, x.CivEntityId) + " breached the Underworld at " + site(x.SiteId, "")
}

func (x *HistoricalEventEntityCreated) Html(c *context) string {
	if x.CreatorHfid != -1 {
		return c.hf(x.CreatorHfid) + " formed " + entity(x.EntityId) + siteStructure(x.SiteId, x.StructureId, "in")
	} else {
		return entity(x.EntityId) + " formed" + siteStructure(x.SiteId, x.StructureId, "in")
	}
}

func (x *HistoricalEventEntityDissolved) Html(c *context) string {
	return entity(x.EntityId) + " dissolved after " + x.Reason.String()
}

func (x *HistoricalEventEntityEquipmentPurchase) Html(c *context) string { // todo check hfid
	return entity(x.EntityId) + " purchased " + equipmentLevel(x.NewEquipmentLevel) + " equipment"
}

func (x *HistoricalEventEntityExpelsHf) Html(c *context) string {
	return "UNKNWON HistoricalEventEntityExpelsHf"
}

func (x *HistoricalEventEntityFledSite) Html(c *context) string {
	return "UNKNWON HistoricalEventEntityFledSite"
}

func (x *HistoricalEventEntityIncorporated) Html(c *context) string { // TODO site
	return entity(x.JoinerEntityId) + util.If(x.PartialIncorporation, " began operating at the direction of ", " fully incorporated into ") +
		entity(x.JoinedEntityId) + " under the leadership of " + c.hf(x.LeaderHfid)
}

func (x *HistoricalEventEntityLaw) Html(c *context) string {
	switch x.LawAdd {
	case HistoricalEventEntityLawLawAdd_Harsh:
		return c.hf(x.HistFigureId) + " laid a series of oppressive edicts upon " + entity(x.EntityId)
	}
	switch x.LawRemove {
	case HistoricalEventEntityLawLawRemove_Harsh:
		return c.hf(x.HistFigureId) + " lifted numerous oppressive laws from " + entity(x.EntityId)
	}
	return c.hf(x.HistFigureId) + " UNKNOWN LAW upon " + entity(x.EntityId)
}

func (x *HistoricalEventEntityOverthrown) Html(c *context) string {
	return c.hf(x.InstigatorHfid) + " toppled the government of " + util.If(x.OverthrownHfid != -1, c.hfRelated(x.OverthrownHfid, x.InstigatorHfid)+" of ", "") + entity(x.EntityId) + " and " +
		util.If(x.PosTakerHfid == x.InstigatorHfid, "assumed control", "placed "+c.hfRelated(x.PosTakerHfid, x.InstigatorHfid)+" in power") + site(x.SiteId, " in") +
		util.If(len(x.ConspiratorHfid) > 0, ". The support of "+c.hfListRelated(x.ConspiratorHfid, x.InstigatorHfid)+" was crucial to the coup", "")
}

func (x *HistoricalEventEntityPersecuted) Html(c *context) string {
	var l []string
	if len(x.ExpelledHfid) > 0 {
		l = append(l, c.hfListRelated(x.ExpelledHfid, x.PersecutorHfid)+util.If(len(x.ExpelledHfid) > 1, " were", " was")+" expelled")
	}
	if len(x.PropertyConfiscatedFromHfid) > 0 {
		l = append(l, "most property was confiscated")
	}
	if x.DestroyedStructureId != -1 {
		l = append(l, structure(x.SiteId, x.DestroyedStructureId)+" was destroyed"+util.If(x.ShrineAmountDestroyed > 0, " along with several smaller sacred sites", ""))
	} else if x.ShrineAmountDestroyed > 0 {
		l = append(l, "some sacred sites were desecrated")
	}
	return c.hf(x.PersecutorHfid) + " of " + entity(x.PersecutorEnid) + " persecuted " + entity(x.TargetEnid) + " in " + site(x.SiteId, "") +
		util.If(len(l) > 0, ". "+util.Capitalize(andList(l)), "")
}

func (x *HistoricalEventEntityPrimaryCriminals) Html(c *context) string { // TODO structure
	switch x.Action {
	case HistoricalEventEntityPrimaryCriminalsAction_EntityPrimaryCriminals:
		return entity(x.EntityId) + " became the primary criminal organization in " + site(x.SiteId, "")
	}
	return "UNKNWON HistoricalEventEntityPrimaryCriminals"
}

func (x *HistoricalEventEntityRampagedInSite) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventEntityRampagedInSite"
}

func (x *HistoricalEventEntityRelocate) Html(c *context) string {
	switch x.Action {
	case HistoricalEventEntityRelocateAction_EntityRelocate:
		return entity(x.EntityId) + " moved" + siteStructure(x.SiteId, x.StructureId, "to")
	}
	return "UNKNWON HistoricalEventEntityRelocate"
}

func (x *HistoricalEventEntitySearchedSite) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventEntitySearchedSite"
}

func (x *HistoricalEventFailedFrameAttempt) Html(c *context) string {
	return c.hf(x.FramerHfid) + " attempted to frame " + c.hfRelated(x.TargetHfid, x.FramerHfid) + " for " + x.Crime.String() +
		util.If(x.PlotterHfid != -1, " at the behest of "+c.hfRelated(x.PlotterHfid, x.FramerHfid), "") +
		" by fooling " + c.hfRelated(x.FooledHfid, x.FramerHfid) + " and " + entity(x.ConvicterEnid) +
		" with fabricated evidence, but nothing came of it"
}

func (x *HistoricalEventFailedIntrigueCorruption) Html(c *context) string {
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
		method = "made a blackmail threat, due to embezzlement using the position " + position(x.RelevantEntityId, x.RelevantPositionProfileId, x.CorruptorHfid) + " of " + entity(x.RelevantEntityId)
	case HistoricalEventFailedIntrigueCorruptionMethod_Bribe:
		method = "offered a bribe"
	case HistoricalEventFailedIntrigueCorruptionMethod_Flatter:
		method = "made flattering remarks"
	case HistoricalEventFailedIntrigueCorruptionMethod_Intimidate:
		method = "made a threat"
	case HistoricalEventFailedIntrigueCorruptionMethod_OfferImmortality:
		method = "offered immortality"
	case HistoricalEventFailedIntrigueCorruptionMethod_Precedence:
		method = "pulled rank as " + position(x.RelevantEntityId, x.RelevantPositionProfileId, x.CorruptorHfid) + " of " + entity(x.RelevantEntityId)
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
		fail += ", despite being swayed by the emotional appeal" // TODO
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Vanity:
	case HistoricalEventFailedIntrigueCorruptionTopFacet_Vengeful:
	}
	return c.hf(x.CorruptorHfid) + " attempted to corrupt " + c.hfRelated(x.TargetHfid, x.CorruptorHfid) +
		" in order to " + action + location(x.SiteId, " in", x.SubregionId, " in") + ". " +
		util.Capitalize(util.If(x.LureHfid != -1,
			c.hfRelated(x.LureHfid, x.CorruptorHfid)+" lured "+c.hfShort(x.TargetHfid)+" to a meeting with "+c.hfShort(x.CorruptorHfid)+", where the latter",
			c.hfShort(x.CorruptorHfid)+" met with "+c.hfShort(x.TargetHfid))) +
		util.If(x.FailedJudgmentTest, ", while completely misreading the situation,", "") + " " + method + ". " + fail
}

func (x *HistoricalEventFieldBattle) Html(c *context) string {
	atk := entity(x.AttackerCivId)
	def := entity(x.DefenderCivId)
	generals := ""
	if x.AttackerGeneralHfid != -1 {
		generals += ". " + util.Capitalize(c.hf(x.AttackerGeneralHfid)) + " led the attack"
		if x.DefenderGeneralHfid != -1 {
			generals += ", and the defenders were led by " + c.hf(x.DefenderGeneralHfid)
		}
	}
	mercs := ""
	if x.AttackerMercEnid != -1 {
		mercs += fmt.Sprintf(". %s were hired by the attackers", entity(x.AttackerMercEnid))
	}
	if x.ASupportMercEnid != -1 {
		mercs += fmt.Sprintf(". %s were hired as scouts by the attackers", entity(x.ASupportMercEnid))
	}
	if x.DefenderMercEnid != -1 {
		mercs += fmt.Sprintf(". The defenders hired %s", entity(x.DefenderMercEnid))
	}
	if x.DSupportMercEnid != -1 {
		mercs += fmt.Sprintf(". The defenders hired %s as scouts", entity(x.DSupportMercEnid))
	}
	return fmt.Sprintf("%s attacked %s at %s%s%s", atk, def, region(x.SubregionId), generals, mercs)
}

func (x *HistoricalEventFirstContact) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventFirstContact"
}

func (x *HistoricalEventGamble) Html(c *context) string {
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
	return c.hf(x.GamblerHfid) + " " + outcome + " gambling" + siteStructure(x.SiteId, x.StructureId, " in") +
		util.If(x.OldAccount >= 0 && x.NewAccount < 0, " and went into debt", "")
}

func (x *HistoricalEventHfAbducted) Html(c *context) string {
	return c.hf(x.TargetHfid) + " was abducted " + location(x.SiteId, "from", x.SubregionId, "from") + " by " + c.hfRelated(x.SnatcherHfid, x.TargetHfid)
}

func (x *HistoricalEventHfAttackedSite) Html(c *context) string {
	return c.hf(x.AttackerHfid) + " attacked " + siteCiv(x.SiteCivId, x.DefenderCivId) + site(x.SiteId, " in")
}

func (x *HistoricalEventHfConfronted) Html(c *context) string {
	return c.hf(x.Hfid) + " aroused " + x.Situation.String() + location(x.SiteId, " in", x.SubregionId, " in") + " after " +
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
func (x *HistoricalEventHfConvicted) Html(c *context) string { // TODO no_prison_available, beating, hammerstrokes, interrogator_hfid
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
		util.If(x.ConvictIsContact, "as a go-between in a conspiracy to commit ", "of ") + x.Crime.String() + " by " + entity(x.ConvicterEnid)
	if x.FooledHfid != -1 {
		r += " after " + c.hfRelated(x.FramerHfid, x.ConvictedHfid) + " fooled " + c.hfRelated(x.FooledHfid, x.ConvictedHfid) + " with fabricated evidence" +
			util.If(x.PlotterHfid != -1, " at the behest of "+c.hfRelated(x.PlotterHfid, x.ConvictedHfid), "")
	}
	if x.CorruptConvicterHfid != -1 {
		r += " and the corrupt " + c.hfRelated(x.CorruptConvicterHfid, x.ConvictedHfid) + " through the machinations of " + c.hfRelated(x.PlotterHfid, x.ConvictedHfid)
	}
	switch {
	case x.DeathPenalty:
		r += " and sentenced to death"
	case x.Exiled:
		r += " and exiled"
	case x.PrisonMonths > 0:
		r += fmt.Sprintf(" and imprisoned for a term of %d years", x.PrisonMonths/12)
	}
	if x.HeldFirmInInterrogation {
		r += ". " + c.hfShort(x.ConvictedHfid) + " revealed nothing during interrogation"
	} else if len(x.ImplicatedHfid) > 0 {
		r += ". " + c.hfShort(x.ConvictedHfid) + " implicated " + c.hfList(x.ImplicatedHfid) + " during interrogation" +
			util.If(x.DidNotRevealAllInInterrogation, " but did not reveal eaverything", "")
	}
	return r
}

func (x *HistoricalEventHfDestroyedSite) Html(c *context) string {
	return c.hf(x.AttackerHfid) + " routed " + siteCiv(x.SiteCivId, x.DefenderCivId) + " and destroyed " + site(x.SiteId, "")
}

func (x *HistoricalEventHfDied) Html(c *context) string { // TODO force cause enum
	return "UNKNWON HistoricalEventHfDied"
}

func (x *HistoricalEventHfDisturbedStructure) Html(c *context) string {
	return c.hf(x.HistFigId) + " disturbed " + siteStructure(x.SiteId, x.StructureId, "")
}

func (x *HistoricalEventHfDoesInteraction) Html(c *context) string {
	i := strings.Index(x.InteractionAction, " ")
	if i > 0 {
		return c.hf(x.DoerHfid) + " " + x.InteractionAction[:i+1] + c.hfRelated(x.TargetHfid, x.DoerHfid) + x.InteractionAction[i:] + util.If(x.Site != -1, site(x.Site, " in"), "")
	} else {
		return c.hf(x.DoerHfid) + " UNKNOWN INTERACTION " + c.hfRelated(x.TargetHfid, x.DoerHfid) + util.If(x.Site != -1, site(x.Site, " in"), "")
	}
}

func (x *HistoricalEventHfEnslaved) Html(c *context) string {
	return c.hf(x.SellerHfid) + " sold " + c.hfRelated(x.EnslavedHfid, x.SellerHfid) + " to " + entity(x.PayerEntityId) + site(x.MovedToSiteId, " in")
}

func (x *HistoricalEventHfEquipmentPurchase) Html(c *context) string { // TODO site, structure, region
	return c.hf(x.GroupHfid) + " purchased " + equipmentLevel(x.Quality) + " equipment"
}

func (x *HistoricalEventHfFreed) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventHfFreed"
}

func (x *HistoricalEventHfGainsSecretGoal) Html(c *context) string {
	switch x.SecretGoal {
	case HistoricalEventHfGainsSecretGoalSecretGoal_Immortality:
		return c.hf(x.Hfid) + " became obsessed with " + posessivePronoun(x.Hfid) + " own mortality and sought to extend " + posessivePronoun(x.Hfid) + " life by any means"
	}
	return c.hf(x.Hfid) + " UNKNOWN SECRET GOAL"
}

func (x *HistoricalEventHfInterrogated) Html(c *context) string { // TODO wanted_and_recognized, held_firm_in_interrogation, implicated_hfid
	return c.hf(x.TargetHfid) + " was recognized and arrested by " + entity(x.ArrestingEnid) +
		". Despite the interrogation by " + c.hfRelated(x.InterrogatorHfid, x.TargetHfid) + ", " + c.hfShort(x.TargetHfid) + " refused to reveal anything and was released"
}

func (x *HistoricalEventHfLearnsSecret) Html(c *context) string {
	if x.ArtifactId != -1 {
		return c.hf(x.StudentHfid) + " learned " + x.SecretText.String() + " from " + artifact(x.ArtifactId)
	} else {
		return c.hf(x.TeacherHfid) + " taught " + c.hfRelated(x.StudentHfid, x.TeacherHfid) + " " + x.SecretText.String()
	}
}

func (x *HistoricalEventHfNewPet) Html(c *context) string {
	return c.hf(x.GroupHfid) + " tamed " + articled(x.Pets) + location(x.SiteId, " of", x.SubregionId, " of")
}
func (x *HistoricalEventHfPerformedHorribleExperiments) Html(c *context) string {
	return c.hf(x.GroupHfid) + " performed horrible experiments " + place(x.StructureId, x.SiteId, " in", x.SubregionId, " in")
}
func (x *HistoricalEventHfPrayedInsideStructure) Html(c *context) string {
	return c.hf(x.HistFigId) + " prayed " + siteStructure(x.SiteId, x.StructureId, "inside")
}

func (x *HistoricalEventHfPreach) Html(c *context) string { // relevant site
	topic := ""
	switch x.Topic {
	case HistoricalEventHfPreachTopic_Entity1ShouldLoveEntityTwo:
		topic = ", urging love to be shown to "
	case HistoricalEventHfPreachTopic_SetEntity1AgainstEntityTwo:
		topic = ", inveighing against "
	}
	return c.hf(x.SpeakerHfid) + " preached to " + entity(x.Entity1) + topic + entity(x.Entity2) + site(x.SiteHfid, " in")
}

func (x *HistoricalEventHfProfanedStructure) Html(c *context) string {
	return c.hf(x.HistFigId) + " profaned " + siteStructure(x.SiteId, x.StructureId, "")
}

func (x *HistoricalEventHfRansomed) Html(c *context) string {
	return c.hf(x.RansomerHfid) + " ransomed " + c.hfRelated(x.RansomedHfid, x.RansomerHfid) + " to " + util.If(x.PayerHfid != -1, c.hfRelated(x.PayerHfid, x.RansomerHfid), entity(x.PayerEntityId)) +
		". " + c.hfShort(x.RansomedHfid) + " was sent " + site(x.MovedToSiteId, "to")
}

func (x *HistoricalEventHfReachSummit) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventHfReachSummit"
}

func (x *HistoricalEventHfRecruitedUnitTypeForEntity) Html(c *context) string {
	return c.hf(x.Hfid) + " recruited " + x.UnitType.String() + "s into " + entity(x.EntityId) + location(x.SiteId, " in", x.SubregionId, " in")
}

func (x *HistoricalEventHfRelationshipDenied) Html(c *context) string {
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

func (x *HistoricalEventHfReunion) Html(c *context) string {
	return c.hf(x.Group1Hfid) + " was reunited with " + c.hfListRelated(x.Group2Hfid, x.Group1Hfid) + location(x.SiteId, " in", x.SubregionId, " in")
}

func (x *HistoricalEventHfRevived) Html(c *context) string {
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
		location(x.SiteId, " in", x.SubregionId, " in")
}

func (x *HistoricalEventHfSimpleBattleEvent) Html(c *context) string {
	group1 := c.hf(x.Group1Hfid)
	group2 := c.hfRelated(x.Group2Hfid, x.Group1Hfid)
	loc := location(x.SiteId, " in", x.SubregionId, " in")
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

func (x *HistoricalEventHfTravel) Html(c *context) string {
	return c.hfList(x.GroupHfid) + util.If(x.Return, " returned", " made a journey") + location(x.SiteId, " to", x.SubregionId, " to")
}

func (x *HistoricalEventHfViewedArtifact) Html(c *context) string {
	return c.hf(x.HistFigId) + " viewed " + artifact(x.ArtifactId) + siteStructure(x.SiteId, x.StructureId, " in")
}

func (x *HistoricalEventHfWounded) Html(c *context) string {
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

	return r + c.hfRelated(x.WounderHfid, x.WoundeeHfid) + location(x.SiteId, " in", x.SubregionId, " in") + util.If(x.WasTorture, " as a means of torture", "")
}

func (x *HistoricalEventHfsFormedIntrigueRelationship) Html(c *context) string {
	if x.Circumstance == HistoricalEventHfsFormedIntrigueRelationshipCircumstance_IsEntitySubordinate {
		return c.hf(x.CorruptorHfid) + " subordinated " + c.hfRelated(x.TargetHfid, x.CorruptorHfid) + " as a member of " + entity(x.CircumstanceId) +
			" toward the fullfillment of plots and schemes" + location(x.SiteId, " in", x.SubregionId, " in")
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
		method = "made a blackmail threat, due to embezzlement using the position " + position(x.RelevantEntityId, x.RelevantPositionProfileId, x.CorruptorHfid) + " of " + entity(x.RelevantEntityId)
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_Bribe:
		method = "offered a bribe"
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_Flatter:
		method = "made flattering remarks"
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_Intimidate:
		method = "made a threat"
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_OfferImmortality:
		method = "offered immortality"
	case HistoricalEventHfsFormedIntrigueRelationshipMethod_Precedence:
		method = "pulled rank as " + position(x.RelevantEntityId, x.RelevantPositionProfileId, x.CorruptorHfid) + " of " + entity(x.RelevantEntityId)
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
		// success += ", despite being swayed by the emotional appeal" // TODO
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Vanity:
	case HistoricalEventHfsFormedIntrigueRelationshipTopFacet_Vengeful:
	}
	return c.hf(x.CorruptorHfid) + " corrupted " + c.hfRelated(x.TargetHfid, x.CorruptorHfid) +
		" in order to " + action + location(x.SiteId, " in", x.SubregionId, " in") + ". " +
		util.Capitalize(util.If(x.LureHfid != -1,
			c.hfRelated(x.LureHfid, x.CorruptorHfid)+" lured "+c.hfShort(x.TargetHfid)+" to a meeting with "+c.hfShort(x.CorruptorHfid)+", where the latter",
			c.hfShort(x.CorruptorHfid)+" met with "+c.hfShort(x.TargetHfid))) +
		" and " + method + ". " + success
}

func (x *HistoricalEventHfsFormedReputationRelationship) Html(c *context) string {
	hf1 := c.hf(x.Hfid1) + util.If(x.IdentityId1 != -1, " as "+fullIdentity(x.IdentityId1), "")
	hf2 := c.hfRelated(x.Hfid2, x.Hfid1) + util.If(x.IdentityId2 != -1, " as "+fullIdentity(x.IdentityId2), "")
	loc := location(x.SiteId, " in", x.SubregionId, " in")
	switch x.HfRep2Of1 {
	case HistoricalEventHfsFormedReputationRelationshipHfRep2Of1_Friendly:
		return hf1 + " and " + hf2 + ", formed a false friendship where each used the other for information" + loc
	case HistoricalEventHfsFormedReputationRelationshipHfRep2Of1_InformationSource:
		return hf1 + ", formed a false friendship with " + hf2 + " in order to extract information" + loc
	}
	return hf1 + " and " + hf2 + ", formed an UNKNOWN RELATIONSHIP" + loc
}

func (x *HistoricalEventHolyCityDeclaration) Html(c *context) string {
	return entity(x.ReligionId) + " declared " + site(x.SiteId, "") + " to be a holy site"
}
func (x *HistoricalEventInsurrectionStarted) Html(c *context) string {
	e := util.If(x.TargetCivId != -1, entity(x.TargetCivId), "the local government")
	switch x.Outcome {
	case HistoricalEventInsurrectionStartedOutcome_LeadershipOverthrown:
		return "the insurrection " + site(x.SiteId, "in") + " concluded with " + e + " overthrown"
	case HistoricalEventInsurrectionStartedOutcome_PopulationGone:
		return "an insurrection " + site(x.SiteId, "in") + " against " + e + " ended with the disappearance of the rebelling population"
	default:
		return "an insurrection against " + e + " began " + site(x.SiteId, "in")
	}
}
func (x *HistoricalEventItemStolen) Html(c *context) string {
	i := util.If(x.Item != -1, artifact(x.Item), articled(x.Mat+" "+x.ItemType))
	circumstance := ""
	switch x.Circumstance.Type {
	case HistoricalEventItemStolenCircumstanceType_Defeated:
		circumstance = " after defeating " + c.hfRelated(x.Circumstance.Defeated, x.Histfig)
	case HistoricalEventItemStolenCircumstanceType_Histeventcollection:
	case HistoricalEventItemStolenCircumstanceType_Murdered:
	}

	switch x.TheftMethod {
	case HistoricalEventItemStolenTheftMethod_Confiscated: // TODO
	case HistoricalEventItemStolenTheftMethod_Looted:
	case HistoricalEventItemStolenTheftMethod_Recovered:
		return i + " was recovered by " + c.hf(x.Histfig) + circumstance
	}
	return i + " was stolen " + siteStructure(x.Site, x.Structure, "from") + " by " + c.hf(x.Histfig) + circumstance +
		util.If(x.StashSite != -1, " and brought "+site(x.StashSite, "to"), "")
}

func (x *HistoricalEventKnowledgeDiscovered) Html(c *context) string {
	return c.hf(x.Hfid) + util.If(x.First, " was the very first to discover ", " independently discovered ") + x.Knowledge
}

func (x *HistoricalEventMasterpieceArchConstructed) Html(c *context) string {
	return "UNKNWON HistoricalEventMasterpieceArchConstructed"
}
func (x *HistoricalEventMasterpieceEngraving) Html(c *context) string {
	return "UNKNWON HistoricalEventMasterpieceEngraving"
}
func (x *HistoricalEventMasterpieceFood) Html(c *context) string {
	return "UNKNWON HistoricalEventMasterpieceFood"
}
func (x *HistoricalEventMasterpieceItem) Html(c *context) string {
	return "UNKNWON HistoricalEventMasterpieceItem"
}
func (x *HistoricalEventMasterpieceItemImprovement) Html(c *context) string {
	return "UNKNWON HistoricalEventMasterpieceItemImprovement"
}
func (x *HistoricalEventMasterpieceLost) Html(c *context) string {
	return "UNKNWON HistoricalEventMasterpieceLost"
}
func (x *HistoricalEventMerchant) Html(c *context) string { return "UNKNWON HistoricalEventMerchant" }

func (x *HistoricalEventModifiedBuilding) Html(c *context) string {
	return c.hf(x.ModifierHfid) + " had " + articled(x.Modification.String()) + " added " + siteStructure(x.SiteId, x.StructureId, "to")
}

func (x *HistoricalEventMusicalFormCreated) Html(c *context) string {
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
	return musicalForm(x.FormId) + " was created by " + c.hf(x.HistFigureId) + site(x.SiteId, " in") + reason + circumstance
}

func (x *HistoricalEventNewSiteLeader) Html(c *context) string {
	return entity(x.AttackerCivId) + " defeated " + siteCiv(x.SiteCivId, x.DefenderCivId) + " and placed " + c.hf(x.NewLeaderHfid) + " in charge of" + site(x.SiteId, "") +
		". The new government was called " + entity(x.NewSiteCivId)
}

func (x *HistoricalEventPeaceAccepted) Html(c *context) string {
	return entity(x.Destination) + " accepted an offer of peace from " + entity(x.Source)
}

func (x *HistoricalEventPeaceRejected) Html(c *context) string {
	return entity(x.Destination) + " rejected an offer of peace from " + entity(x.Source)
}

func (x *HistoricalEventPerformance) Html(c *context) string {
	r := entity(x.CivId) + " held "
	if e, ok := world.Entities[x.CivId]; ok {
		o := e.Occasion[x.OccasionId]
		s := o.Schedule[x.ScheduleId]
		r += schedule(s)
		r += " as part of " + o.Name()
		r += site(x.SiteId, " in")
		r += string(util.Json(s))
	}
	return r
}

func (x *HistoricalEventPlunderedSite) Html(c *context) string { // TODO no_defeat_mention, took_items, took_livestock, was_raid
	return entity(x.AttackerCivId) + " defeated " + siteCiv(x.SiteCivId, x.DefenderCivId) + " and pillaged " + site(x.SiteId, "")
}

func (x *HistoricalEventPoeticFormCreated) Html(c *context) string {
	circumstance := ""
	switch x.Circumstance {
	case HistoricalEventPoeticFormCreatedCircumstance_Dream:
		circumstance = " after a dream"
	case HistoricalEventPoeticFormCreatedCircumstance_Nightmare:
		circumstance = " after a nightmare"
	}
	return poeticForm(x.FormId) + " was created by " + c.hf(x.HistFigureId) + site(x.SiteId, " in") + circumstance
}

func (x *HistoricalEventProcession) Html(c *context) string {
	r := entity(x.CivId) + " held a procession in " + site(x.SiteId, "")
	if e, ok := world.Entities[x.CivId]; ok {
		o := e.Occasion[x.OccasionId]
		r += " as part of " + o.Name()
		s := o.Schedule[x.ScheduleId]
		if s.Reference != -1 {
			r += ". It started at " + structure(x.SiteId, s.Reference)
			if s.Reference2 != -1 && s.Reference2 != s.Reference {
				r += " and ended at " + structure(x.SiteId, s.Reference2)
			} else {
				r += " and returned there after following its route"
			}
		}
		if len(s.Feature) > 0 {
			r += ". The event featured " + andList(util.Map(s.Feature, feature))
		}
		r += string(util.Json(s))
	}
	return r
}

func (x *HistoricalEventRazedStructure) Html(c *context) string {
	return entity(x.CivId) + " razed " + siteStructure(x.SiteId, x.StructureId, "")
}

func (x *HistoricalEventReclaimSite) Html(c *context) string { // TODO unretire
	return siteCiv(x.SiteCivId, x.CivId) + " launched an expedition to reclaim " + site(x.SiteId, "")
}

func (x *HistoricalEventRegionpopIncorporatedIntoEntity) Html(c *context) string { // TODO Race
	return strconv.Itoa(x.PopNumberMoved) + " of " + strconv.Itoa(x.PopRace) + " from " + region(x.PopSrid) + " joined with " + entity(x.JoinEntityId) + site(x.SiteId, " at")
}

func (x *HistoricalEventRemoveHfEntityLink) Html(c *context) string {
	hf := c.hf(x.Hfid)
	civ := entity(x.CivId)
	switch x.Link {
	case HistoricalEventRemoveHfEntityLinkLink_Member:
		return hf + " left " + civ
	case HistoricalEventRemoveHfEntityLinkLink_Position:
		return hf + " ceased to be the " + position(x.CivId, x.PositionId, x.Hfid) + " of " + civ
	case HistoricalEventRemoveHfEntityLinkLink_Prisoner:
		return hf + " escaped from the prisons of " + civ
	case HistoricalEventRemoveHfEntityLinkLink_Slave:
		return hf + " escaped from the slavery of " + civ
	}
	return hf + " left " + civ
}

func (x *HistoricalEventRemoveHfHfLink) Html(c *context) string { // divorced
	return c.hf(x.Hfid) + " and " + c.hfRelated(x.HfidTarget, x.Hfid) + " broke up"
}

func (x *HistoricalEventRemoveHfSiteLink) Html(c *context) string {
	switch x.LinkType {
	case HistoricalEventRemoveHfSiteLinkLinkType_HomeSiteAbstractBuilding:
		return c.hf(x.Histfig) + " moved out " + siteStructure(x.SiteId, x.Structure, "of")
	case HistoricalEventRemoveHfSiteLinkLinkType_Occupation:
		return c.hf(x.Histfig) + " stopped working " + siteStructure(x.SiteId, x.Structure, "at")
	case HistoricalEventRemoveHfSiteLinkLinkType_SeatOfPower:
		return c.hf(x.Histfig) + " stopped ruling " + siteStructure(x.SiteId, x.Structure, "from")
	}
	return c.hf(x.Histfig) + " stopped working " + siteStructure(x.SiteId, x.Structure, "at")
}

func (x *HistoricalEventReplacedStructure) Html(c *context) string {
	return siteCiv(x.SiteCivId, x.CivId) + " replaced " + siteStructure(x.SiteId, x.OldAbId, "") + " with " + structure(x.SiteId, x.NewAbId)
}

func (x *HistoricalEventSiteDied) Html(c *context) string { return "UNKNWON HistoricalEventSiteDied" } // TODO

func (x *HistoricalEventSiteDispute) Html(c *context) string {
	return entity(x.EntityId1) + " of " + site(x.SiteId1, "") + " and " + entity(x.EntityId2) + " of " + site(x.SiteId2, "") + " became embroiled in a dispute over " + x.Dispute.String()
}

func (x *HistoricalEventSiteRetired) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventSiteRetired"
}
func (x *HistoricalEventSiteSurrendered) Html(c *context) string { // TODO
	return "UNKNWON HistoricalEventSiteSurrendered"
}

func (x *HistoricalEventSiteTakenOver) Html(c *context) string {
	return entity(x.AttackerCivId) + " defeated " + siteCiv(x.SiteCivId, x.DefenderCivId) + " and took over " + site(x.SiteId, "") + ". The new government was called " + entity(x.NewSiteCivId)
}

func (x *HistoricalEventSiteTributeForced) Html(c *context) string {
	return "UNKNWON HistoricalEventSiteTributeForced"
}
func (x *HistoricalEventSneakIntoSite) Html(c *context) string {
	return "UNKNWON HistoricalEventSneakIntoSite"
}
func (x *HistoricalEventSpottedLeavingSite) Html(c *context) string {
	return "UNKNWON HistoricalEventSpottedLeavingSite"
}
func (x *HistoricalEventSquadVsSquad) Html(c *context) string {
	return "UNKNWON HistoricalEventSquadVsSquad"
}
func (x *HistoricalEventTacticalSituation) Html(c *context) string {
	return "UNKNWON HistoricalEventTacticalSituation"
}

func (x *HistoricalEventTrade) Html(c *context) string {
	outcome := ""
	switch d := x.AccountShift; {
	case d > 1000:
		outcome = " did well"
	case d < -1000:
		outcome = " did poorly"
	default:
		outcome = " broke even"
	}
	return c.hf(x.TraderHfid) + util.If(x.TraderEntityId != -1, " of "+entity(x.TraderEntityId), "") + outcome + " trading" + site(x.SourceSiteId, " from") + site(x.DestSiteId, " to")
}

func (x *HistoricalEventWrittenContentComposed) Html(c *context) string {
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
	return writtenContent(x.WcId) + " was authored by " + c.hf(x.HistFigureId) + location(x.SiteId, " in", x.SubregionId, " in") + reason + circumstance
}

func (x *HistoricalEventAgreementConcluded) Html(c *context) string {
	return "UNKNWON HistoricalEventAgreementConcluded"
}
func (x *HistoricalEventMasterpieceDye) Html(c *context) string {
	return "UNKNWON HistoricalEventMasterpieceDye"
}
