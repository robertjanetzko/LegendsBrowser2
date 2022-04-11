package model

import "encoding/xml"

type Region struct {
	XMLName xml.Name `xml:"region" json:"-"`
	NamedObject
	Type string `xml:"type" json:"type"`
	OtherElements
	EventObject
}

type UndergroundRegion struct {
	XMLName xml.Name `xml:"underground_region" json:"-"`
	NamedObject
	Type  string `xml:"type" json:"type"`
	Depth *int   `xml:"depth" json:"depth,omitempty"`
	OtherElements
	EventObject
}

type Landmass struct {
	XMLName xml.Name `xml:"landmass" json:"-"`
	NamedObject
	OtherElements
	EventObject
}

type Site struct {
	XMLName xml.Name `xml:"site" json:"-"`
	NamedObject
	Type       string      `xml:"type" json:"type"`
	Coords     string      `xml:"coords" json:"coords"`
	Rectangle  string      `xml:"rectangle" json:"rectangle"`
	Structures []Structure `xml:"structures>structure" json:"structures"`

	OtherElements
	EventObject
}

// func (obj Site) id() int      { return obj.Id }
// func (obj Site) name() string { return obj.Name }

type Structure struct {
	XMLName xml.Name `xml:"structure" json:"-"`
	LocalId int      `xml:"local_id" json:"localId"`
	Name    string   `xml:"name" json:"name"`
	Type    string   `xml:"type" json:"type"`
	OtherElements
	EventObject
}

type WorldConstruction struct {
	XMLName xml.Name `xml:"world_construction" json:"-"`
	NamedObject
	OtherElements
	EventObject
}

type Artifact struct {
	XMLName xml.Name `xml:"artifact" json:"-"`
	NamedObject
	SiteId     int  `xml:"site_id" json:"siteId"`
	AbsTileX   *int `xml:"abs_tile_x" json:"absTileX,omitempty"`
	AbsTileY   *int `xml:"abs_tile_y" json:"absTileY,omitempty"`
	AbsTileZ   *int `xml:"abs_tile_z" json:"absTileZ,omitempty"`
	HolderHfid *int `xml:"holder_hfid" json:"holderHfid,omitempty" legend:"hf"`
	// item object
	StructureLocalId *int `xml:"structure_local_id" json:"structureLocalId,omitempty"`
	SubregionId      *int `xml:"subregion_id" json:"subregionId,omitempty"`

	OtherElements

	EventObject
}

type HistoricalFigure struct {
	XMLName xml.Name `xml:"historical_figure" json:"-"`
	NamedObject

	Race *string `xml:"race" json:"race"`
	// Caste             *string   `xml:"caste" json:"caste"`
	// ActiveInteraction *[]string `xml:"active_interaction" json:"activeInteraction,omitempty"`
	// Animated          *string   `xml:"animated" json:"animated,omitempty"`
	// AnimatedString    *string   `xml:"animated_string" json:"animatedString,omitempty"`
	// Appeared          *int      `xml:"appeared" json:"appeared,omitempty"`
	// AssociatedType    *string   `xml:"associated_type" json:"associatedType,omitempty"`
	// BirthSeconds72    *int      `xml:"birth_seconds72" json:"birthSeconds72,omitempty"`
	// BirthYear         *int      `xml:"birth_year" json:"birthYear,omitempty"`
	// CurrentIdentityId *int      `xml:"current_identity_id" json:"currentIdentityId,omitempty" legend:"entity"`
	// DeathSeconds72    *int      `xml:"death_seconds72" json:"deathSeconds72,omitempty"`
	// DeathYear         *int      `xml:"death_year" json:"deathYear,omitempty"`
	// Deity             *string   `xml:"deity" json:"deity,omitempty"`
	// EntPopId          *int      `xml:"ent_pop_id" json:"entPopId,omitempty"`
	// // entity_former_position_link object
	// // entity_link object
	// // entity_position_link object
	// // entity_reputation object
	// // entity_squad_link object
	// Force *string   `xml:"force" json:"force,omitempty"`
	// Goal  *[]string `xml:"goal" json:"goal,omitempty"`
	// // hf_link object
	// // hf_skill object
	// HoldsArtifact *[]int `xml:"holds_artifact" json:"holdsArtifact,omitempty"`
	// // honor_entity object
	// InteractionKnowledge *[]string `xml:"interaction_knowledge" json:"interactionKnowledge,omitempty"`
	// // intrigue_actor object
	// // intrigue_plot object
	// JourneyPet *[]string `xml:"journey_pet" json:"journeyPet,omitempty"`
	// // relationship_profile_hf_historical object
	// // relationship_profile_hf_visual object
	// // site_link object
	// // site_property object
	// Sphere         *[]string `xml:"sphere" json:"sphere,omitempty"`
	// UsedIdentityId *[]int    `xml:"used_identity_id" json:"usedIdentityId,omitempty" legend:"entity"`
	// // vague_relationship object

	OtherElements

	EventObject
}

func (r *HistoricalFigure) Type() string { return "hf" }

type HistoricalEventCollection struct {
	XMLName xml.Name `xml:"historical_event_collection" json:"-"`
	NamedObject
	Year       int    `xml:"year"`
	Seconds    int    `xml:"seconds72"`
	EndYear    int    `xml:"end_year"`
	EndSeconds int    `xml:"end_seconds72"`
	Type       string `xml:"type" json:"type"`

	EventIds []int `xml:"event" json:"eventIds"`

	AggressorEntId          *int      `xml:"aggressor_ent_id" json:"aggressorEntId,omitempty"`
	AttackingEnid           *int      `xml:"attacking_enid" json:"attackingEnid,omitempty" legend:"entity"`
	AttackingHfid           *[]int    `xml:"attacking_hfid" json:"attackingHfid,omitempty" legend:"hf"`
	AttackingSquadDeaths    *[]int    `xml:"attacking_squad_deaths" json:"attackingSquadDeaths,omitempty"`
	AttackingSquadEntityPop *[]int    `xml:"attacking_squad_entity_pop" json:"attackingSquadEntityPop,omitempty"`
	AttackingSquadNumber    *[]int    `xml:"attacking_squad_number" json:"attackingSquadNumber,omitempty"`
	AttackingSquadRace      *[]string `xml:"attacking_squad_race" json:"attackingSquadRace,omitempty"`
	AttackingSquadSite      *[]int    `xml:"attacking_squad_site" json:"attackingSquadSite,omitempty"`
	CivId                   *int      `xml:"civ_id" json:"civId,omitempty" legend:"entity"`
	Coords                  *string   `xml:"coords" json:"coords,omitempty"`
	DefenderEntId           *int      `xml:"defender_ent_id" json:"defenderEntId,omitempty"`
	DefendingEnid           *int      `xml:"defending_enid" json:"defendingEnid,omitempty" legend:"entity"`
	DefendingHfid           *[]int    `xml:"defending_hfid" json:"defendingHfid,omitempty" legend:"hf"`
	DefendingSquadDeaths    *[]int    `xml:"defending_squad_deaths" json:"defendingSquadDeaths,omitempty"`
	DefendingSquadEntityPop *[]int    `xml:"defending_squad_entity_pop" json:"defendingSquadEntityPop,omitempty"`
	DefendingSquadNumber    *[]int    `xml:"defending_squad_number" json:"defendingSquadNumber,omitempty"`
	DefendingSquadRace      *[]string `xml:"defending_squad_race" json:"defendingSquadRace,omitempty"`
	DefendingSquadSite      *[]int    `xml:"defending_squad_site" json:"defendingSquadSite,omitempty"`
	Eventcol                *[]int    `xml:"eventcol" json:"eventcol,omitempty"`
	FeatureLayerId          *int      `xml:"feature_layer_id" json:"featureLayerId,omitempty"`
	IndividualMerc          *[]string `xml:"individual_merc" json:"individualMerc,omitempty"`
	NoncomHfid              *[]int    `xml:"noncom_hfid" json:"noncomHfid,omitempty" legend:"hf"`
	OccasionId              *int      `xml:"occasion_id" json:"occasionId,omitempty"`
	Ordinal                 *int      `xml:"ordinal" json:"ordinal,omitempty"`
	Outcome                 *string   `xml:"outcome" json:"outcome,omitempty"`
	ParentEventcol          *int      `xml:"parent_eventcol" json:"parentEventcol,omitempty"`
	SiteId                  *int      `xml:"site_id" json:"siteId,omitempty" legend:"site"`
	StartSeconds72          *int      `xml:"start_seconds72" json:"startSeconds72,omitempty"`
	StartYear               *int      `xml:"start_year" json:"startYear,omitempty"`
	SubregionId             *int      `xml:"subregion_id" json:"subregionId,omitempty"`
	WarEventcol             *int      `xml:"war_eventcol" json:"warEventcol,omitempty"`
	ASupportMercEnid        *int      `xml:"a_support_merc_enid" json:"aSupportMercEnid,omitempty" legend:"entity"`
	ASupportMercHfid        *[]int    `xml:"a_support_merc_hfid" json:"aSupportMercHfid,omitempty" legend:"hf"`
	AttackingMercEnid       *int      `xml:"attacking_merc_enid" json:"attackingMercEnid,omitempty" legend:"entity"`
	AttackingSquadAnimated  *[]string `xml:"attacking_squad_animated" json:"attackingSquadAnimated,omitempty"`
	CompanyMerc             *[]string `xml:"company_merc" json:"companyMerc,omitempty"`
	DSupportMercEnid        *int      `xml:"d_support_merc_enid" json:"dSupportMercEnid,omitempty" legend:"entity"`
	DSupportMercHfid        *[]int    `xml:"d_support_merc_hfid" json:"dSupportMercHfid,omitempty" legend:"hf"`
	DefendingMercEnid       *int      `xml:"defending_merc_enid" json:"defendingMercEnid,omitempty" legend:"entity"`
	DefendingSquadAnimated  *[]string `xml:"defending_squad_animated" json:"defendingSquadAnimated,omitempty"`
	TargetEntityId          *int      `xml:"target_entity_id" json:"targetEntityId,omitempty" legend:"entity"`

	OtherElements
}

type Entity struct {
	XMLName xml.Name `xml:"entity" json:"-"`
	NamedObject
	OtherElements
	EventObject
}

type EntityPopulation struct {
	XMLName xml.Name `xml:"entity_population" json:"-"`
	OtherElements
}

type HistoricalEra struct {
	XMLName xml.Name `xml:"historical_era" json:"-"`
	NamedObject
	StartYear *int `xml:"start_year" json:"startYear,omitempty"`
	OtherElements
}

type DanceForm struct {
	XMLName xml.Name `xml:"dance_form" json:"-"`
	ArtForm
}

type MusicalForm struct {
	XMLName xml.Name `xml:"musical_form" json:"-"`
	ArtForm
}

type PoeticForm struct {
	XMLName xml.Name `xml:"poetic_form" json:"-"`
	ArtForm
}

type WrittenContent struct {
	XMLName xml.Name `xml:"written_content" json:"-"`
	NamedObject

	AuthorHfid *int      `xml:"author_hfid" json:"authorHfid,omitempty" legend:"hf"`
	AuthorRoll *int      `xml:"author_roll" json:"authorRoll,omitempty"`
	Form       *string   `xml:"form" json:"form,omitempty"`
	FormId     *int      `xml:"form_id" json:"formId,omitempty"`
	Style      *[]string `xml:"style" json:"style,omitempty"`
	Title      *string   `xml:"title" json:"title,omitempty"`

	OtherElements
	EventObject
}

type ArtForm struct {
	NamedObject

	Description *string `xml:"description" json:"description,omitempty"`

	OtherElements
	EventObject
}
