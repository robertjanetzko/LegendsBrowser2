package model

import "encoding/xml"

type HistoricalEvent struct {
	XMLName xml.Name `xml:"historical_event" json:"-"`
	Id_     int      `xml:"id" json:"id"`
	Year    int      `xml:"year" json:"year"`
	Seconds int      `xml:"seconds72" json:"seconds72"`
	TypedObject

	ASupportMercEnid               *int    `xml:"a_support_merc_enid" json:"aSupportMercEnid,omitempty" legend:"entity"`
	AccountShift                   *int    `xml:"account_shift" json:"accountShift,omitempty"`
	AcquirerEnid                   *int    `xml:"acquirer_enid" json:"acquirerEnid,omitempty" legend:"entity"`
	AcquirerHfid                   *int    `xml:"acquirer_hfid" json:"acquirerHfid,omitempty" legend:"hf"`
	Action                         *string `xml:"action" json:"action,omitempty"`
	ActorHfid                      *int    `xml:"actor_hfid" json:"actorHfid,omitempty" legend:"hf"`
	AgreementId                    *int    `xml:"agreement_id" json:"agreementId,omitempty"`
	Allotment                      *int    `xml:"allotment" json:"allotment,omitempty"`
	AllotmentIndex                 *int    `xml:"allotment_index" json:"allotmentIndex,omitempty"`
	AllyDefenseBonus               *int    `xml:"ally_defense_bonus" json:"allyDefenseBonus,omitempty"`
	AppointerHfid                  *int    `xml:"appointer_hfid" json:"appointerHfid,omitempty" legend:"hf"`
	ArrestingEnid                  *int    `xml:"arresting_enid" json:"arrestingEnid,omitempty" legend:"entity"`
	ArtifactId                     *int    `xml:"artifact_id" json:"artifactId,omitempty" legend:"artifact"`
	AttackerCivId                  *int    `xml:"attacker_civ_id" json:"attackerCivId,omitempty" legend:"entity"`
	AttackerGeneralHfid            *int    `xml:"attacker_general_hfid" json:"attackerGeneralHfid,omitempty" legend:"hf"`
	AttackerHfid                   *int    `xml:"attacker_hfid" json:"attackerHfid,omitempty" legend:"hf"`
	AttackerMercEnid               *int    `xml:"attacker_merc_enid" json:"attackerMercEnid,omitempty" legend:"entity"`
	BodyState                      *string `xml:"body_state" json:"bodyState,omitempty"`
	BuilderHfid                    *int    `xml:"builder_hfid" json:"builderHfid,omitempty" legend:"hf"`
	BuildingProfileId              *int    `xml:"building_profile_id" json:"buildingProfileId,omitempty"`
	Cause                          *string `xml:"cause" json:"cause,omitempty"`
	ChangeeHfid                    *int    `xml:"changee_hfid" json:"changeeHfid,omitempty" legend:"hf"`
	ChangerHfid                    *int    `xml:"changer_hfid" json:"changerHfid,omitempty" legend:"hf"`
	Circumstance                   *string `xml:"circumstance" json:"circumstance,omitempty"`
	CircumstanceId                 *int    `xml:"circumstance_id" json:"circumstanceId,omitempty"`
	CivEntityId                    *int    `xml:"civ_entity_id" json:"civEntityId,omitempty" legend:"entity"`
	CivId                          *int    `xml:"civ_id" json:"civId,omitempty" legend:"entity"`
	Claim                          *string `xml:"claim" json:"claim,omitempty"`
	CoconspiratorBonus             *int    `xml:"coconspirator_bonus" json:"coconspiratorBonus,omitempty"`
	CompetitorHfid                 *[]int  `xml:"competitor_hfid" json:"competitorHfid,omitempty" legend:"hf"`
	ConfessedAfterApbArrestEnid    *int    `xml:"confessed_after_apb_arrest_enid" json:"confessedAfterApbArrestEnid,omitempty" legend:"entity"`
	ConspiratorHfid                *[]int  `xml:"conspirator_hfid" json:"conspiratorHfid,omitempty" legend:"hf"`
	ContactHfid                    *int    `xml:"contact_hfid" json:"contactHfid,omitempty" legend:"hf"`
	ConvictIsContact               *string `xml:"convict_is_contact" json:"convictIsContact,omitempty"`
	ConvictedHfid                  *int    `xml:"convicted_hfid" json:"convictedHfid,omitempty" legend:"hf"`
	ConvicterEnid                  *int    `xml:"convicter_enid" json:"convicterEnid,omitempty" legend:"entity"`
	Coords                         *string `xml:"coords" json:"coords,omitempty"`
	CorruptConvicterHfid           *int    `xml:"corrupt_convicter_hfid" json:"corruptConvicterHfid,omitempty" legend:"hf"`
	CorruptorHfid                  *int    `xml:"corruptor_hfid" json:"corruptorHfid,omitempty" legend:"hf"`
	CorruptorIdentity              *int    `xml:"corruptor_identity" json:"corruptorIdentity,omitempty"`
	CorruptorSeenAs                *string `xml:"corruptor_seen_as" json:"corruptorSeenAs,omitempty"`
	CreatorHfid                    *int    `xml:"creator_hfid" json:"creatorHfid,omitempty" legend:"hf"`
	Crime                          *string `xml:"crime" json:"crime,omitempty"`
	DSupportMercEnid               *int    `xml:"d_support_merc_enid" json:"dSupportMercEnid,omitempty" legend:"entity"`
	DeathPenalty                   *string `xml:"death_penalty" json:"deathPenalty,omitempty"`
	DefenderCivId                  *int    `xml:"defender_civ_id" json:"defenderCivId,omitempty" legend:"entity"`
	DefenderGeneralHfid            *int    `xml:"defender_general_hfid" json:"defenderGeneralHfid,omitempty" legend:"hf"`
	DefenderMercEnid               *int    `xml:"defender_merc_enid" json:"defenderMercEnid,omitempty" legend:"entity"`
	Delegated                      *string `xml:"delegated" json:"delegated,omitempty"`
	DestEntityId                   *int    `xml:"dest_entity_id" json:"destEntityId,omitempty" legend:"entity"`
	DestSiteId                     *int    `xml:"dest_site_id" json:"destSiteId,omitempty" legend:"site"`
	DestStructureId                *int    `xml:"dest_structure_id" json:"destStructureId,omitempty" legend:"structure"`
	DestroyedStructureId           *int    `xml:"destroyed_structure_id" json:"destroyedStructureId,omitempty" legend:"structure"`
	DestroyerEnid                  *int    `xml:"destroyer_enid" json:"destroyerEnid,omitempty" legend:"entity"`
	Detected                       *string `xml:"detected" json:"detected,omitempty"`
	DidNotRevealAllInInterrogation *string `xml:"did_not_reveal_all_in_interrogation" json:"didNotRevealAllInInterrogation,omitempty"`
	Dispute                        *string `xml:"dispute" json:"dispute,omitempty"`
	DoerHfid                       *int    `xml:"doer_hfid" json:"doerHfid,omitempty" legend:"hf"`
	Entity1                        *int    `xml:"entity_1" json:"entity1,omitempty" legend:"entity"`
	Entity2                        *int    `xml:"entity_2" json:"entity2,omitempty" legend:"entity"`
	EntityId                       *int    `xml:"entity_id" json:"entityId,omitempty" legend:"entity"`
	EntityId1                      *int    `xml:"entity_id_1" json:"entityId1,omitempty" legend:"entity"`
	EntityId2                      *int    `xml:"entity_id_2" json:"entityId2,omitempty" legend:"entity"`
	Exiled                         *string `xml:"exiled" json:"exiled,omitempty"`
	ExpelledCreature               *[]int  `xml:"expelled_creature" json:"expelledCreature,omitempty"`
	ExpelledHfid                   *[]int  `xml:"expelled_hfid" json:"expelledHfid,omitempty" legend:"hf"`
	ExpelledNumber                 *[]int  `xml:"expelled_number" json:"expelledNumber,omitempty"`
	ExpelledPopId                  *[]int  `xml:"expelled_pop_id" json:"expelledPopId,omitempty"`
	FailedJudgmentTest             *string `xml:"failed_judgment_test" json:"failedJudgmentTest,omitempty"`
	FeatureLayerId                 *int    `xml:"feature_layer_id" json:"featureLayerId,omitempty"`
	First                          *string `xml:"first" json:"first,omitempty"`
	FooledHfid                     *int    `xml:"fooled_hfid" json:"fooledHfid,omitempty" legend:"hf"`
	FormId                         *int    `xml:"form_id" json:"formId,omitempty"`
	FramerHfid                     *int    `xml:"framer_hfid" json:"framerHfid,omitempty" legend:"hf"`
	FromOriginal                   *string `xml:"from_original" json:"fromOriginal,omitempty"`
	GamblerHfid                    *int    `xml:"gambler_hfid" json:"gamblerHfid,omitempty" legend:"hf"`
	GiverEntityId                  *int    `xml:"giver_entity_id" json:"giverEntityId,omitempty" legend:"entity"`
	GiverHistFigureId              *int    `xml:"giver_hist_figure_id" json:"giverHistFigureId,omitempty" legend:"hf"`
	Group1Hfid                     *int    `xml:"group_1_hfid" json:"group1Hfid,omitempty" legend:"hf"`
	Group2Hfid                     *[]int  `xml:"group_2_hfid" json:"group2Hfid,omitempty" legend:"hf"`
	GroupHfid                      *[]int  `xml:"group_hfid" json:"groupHfid,omitempty" legend:"hf"`
	HeldFirmInInterrogation        *string `xml:"held_firm_in_interrogation" json:"heldFirmInInterrogation,omitempty"`
	HfRep1Of2                      *string `xml:"hf_rep_1_of_2" json:"hfRep1Of2,omitempty"`
	HfRep2Of1                      *string `xml:"hf_rep_2_of_1" json:"hfRep2Of1,omitempty"`
	Hfid                           *[]int  `xml:"hfid" json:"hfid,omitempty" legend:"hf"`
	Hfid1                          *int    `xml:"hfid1" json:"hfid1,omitempty" legend:"hf"`
	Hfid2                          *int    `xml:"hfid2" json:"hfid2,omitempty" legend:"hf"`
	HfidTarget                     *int    `xml:"hfid_target" json:"hfidTarget,omitempty" legend:"hf"`
	HistFigId                      *int    `xml:"hist_fig_id" json:"histFigId,omitempty" legend:"hf"`
	HistFigureId                   *int    `xml:"hist_figure_id" json:"histFigureId,omitempty" legend:"hf"`
	HonorId                        *int    `xml:"honor_id" json:"honorId,omitempty"`
	IdentityId                     *int    `xml:"identity_id" json:"identityId,omitempty" legend:"entity"`
	IdentityId1                    *int    `xml:"identity_id1" json:"identityId1,omitempty" legend:"entity"`
	IdentityId2                    *int    `xml:"identity_id2" json:"identityId2,omitempty" legend:"entity"`
	ImplicatedHfid                 *[]int  `xml:"implicated_hfid" json:"implicatedHfid,omitempty" legend:"hf"`
	Inherited                      *string `xml:"inherited" json:"inherited,omitempty"`
	InitiatingEnid                 *int    `xml:"initiating_enid" json:"initiatingEnid,omitempty" legend:"entity"`
	InstigatorHfid                 *int    `xml:"instigator_hfid" json:"instigatorHfid,omitempty" legend:"hf"`
	Interaction                    *string `xml:"interaction" json:"interaction,omitempty"`
	InterrogatorHfid               *int    `xml:"interrogator_hfid" json:"interrogatorHfid,omitempty" legend:"hf"`
	JoinEntityId                   *int    `xml:"join_entity_id" json:"joinEntityId,omitempty" legend:"entity"`
	JoinedEntityId                 *int    `xml:"joined_entity_id" json:"joinedEntityId,omitempty" legend:"entity"`
	JoinerEntityId                 *int    `xml:"joiner_entity_id" json:"joinerEntityId,omitempty" legend:"entity"`
	JoiningEnid                    *[]int  `xml:"joining_enid" json:"joiningEnid,omitempty" legend:"entity"`
	Knowledge                      *string `xml:"knowledge" json:"knowledge,omitempty"`
	LastOwnerHfid                  *int    `xml:"last_owner_hfid" json:"lastOwnerHfid,omitempty" legend:"hf"`
	LeaderHfid                     *int    `xml:"leader_hfid" json:"leaderHfid,omitempty" legend:"hf"`
	Link                           *string `xml:"link" json:"link,omitempty"`
	LureHfid                       *int    `xml:"lure_hfid" json:"lureHfid,omitempty" legend:"hf"`
	MasterWcid                     *int    `xml:"master_wcid" json:"masterWcid,omitempty" legend:"wc"`
	Method                         *string `xml:"method" json:"method,omitempty"`
	Modification                   *string `xml:"modification" json:"modification,omitempty"`
	ModifierHfid                   *int    `xml:"modifier_hfid" json:"modifierHfid,omitempty" legend:"hf"`
	Mood                           *string `xml:"mood" json:"mood,omitempty"`
	NameOnly                       *string `xml:"name_only" json:"nameOnly,omitempty"`
	NewAbId                        *int    `xml:"new_ab_id" json:"newAbId,omitempty"`
	NewAccount                     *int    `xml:"new_account" json:"newAccount,omitempty"`
	NewCaste                       *string `xml:"new_caste" json:"newCaste,omitempty"`
	NewEquipmentLevel              *int    `xml:"new_equipment_level" json:"newEquipmentLevel,omitempty"`
	NewLeaderHfid                  *int    `xml:"new_leader_hfid" json:"newLeaderHfid,omitempty" legend:"hf"`
	NewRace                        *string `xml:"new_race" json:"newRace,omitempty"`
	NewSiteCivId                   *int    `xml:"new_site_civ_id" json:"newSiteCivId,omitempty" legend:"entity"`
	OccasionId                     *int    `xml:"occasion_id" json:"occasionId,omitempty"`
	OldAbId                        *int    `xml:"old_ab_id" json:"oldAbId,omitempty"`
	OldAccount                     *int    `xml:"old_account" json:"oldAccount,omitempty"`
	OldCaste                       *string `xml:"old_caste" json:"oldCaste,omitempty"`
	OldRace                        *string `xml:"old_race" json:"oldRace,omitempty"`
	OverthrownHfid                 *int    `xml:"overthrown_hfid" json:"overthrownHfid,omitempty" legend:"hf"`
	PartialIncorporation           *string `xml:"partial_incorporation" json:"partialIncorporation,omitempty"`
	PersecutorEnid                 *int    `xml:"persecutor_enid" json:"persecutorEnid,omitempty" legend:"entity"`
	PersecutorHfid                 *int    `xml:"persecutor_hfid" json:"persecutorHfid,omitempty" legend:"hf"`
	PlotterHfid                    *int    `xml:"plotter_hfid" json:"plotterHfid,omitempty" legend:"hf"`
	PopFlid                        *int    `xml:"pop_flid" json:"popFlid,omitempty"`
	PopNumberMoved                 *int    `xml:"pop_number_moved" json:"popNumberMoved,omitempty"`
	PopRace                        *int    `xml:"pop_race" json:"popRace,omitempty"`
	PopSrid                        *int    `xml:"pop_srid" json:"popSrid,omitempty"`
	PosTakerHfid                   *int    `xml:"pos_taker_hfid" json:"posTakerHfid,omitempty" legend:"hf"`
	PositionId                     *int    `xml:"position_id" json:"positionId,omitempty"`
	PositionProfileId              *int    `xml:"position_profile_id" json:"positionProfileId,omitempty"`
	PrisonMonths                   *int    `xml:"prison_months" json:"prisonMonths,omitempty"`
	ProductionZoneId               *int    `xml:"production_zone_id" json:"productionZoneId,omitempty"`
	PromiseToHfid                  *int    `xml:"promise_to_hfid" json:"promiseToHfid,omitempty" legend:"hf"`
	PropertyConfiscatedFromHfid    *[]int  `xml:"property_confiscated_from_hfid" json:"propertyConfiscatedFromHfid,omitempty" legend:"hf"`
	PurchasedUnowned               *string `xml:"purchased_unowned" json:"purchasedUnowned,omitempty"`
	Quality                        *int    `xml:"quality" json:"quality,omitempty"`
	Reason                         *string `xml:"reason" json:"reason,omitempty"`
	ReasonId                       *int    `xml:"reason_id" json:"reasonId,omitempty"`
	RebuiltRuined                  *string `xml:"rebuilt_ruined" json:"rebuiltRuined,omitempty"`
	ReceiverEntityId               *int    `xml:"receiver_entity_id" json:"receiverEntityId,omitempty" legend:"entity"`
	ReceiverHistFigureId           *int    `xml:"receiver_hist_figure_id" json:"receiverHistFigureId,omitempty" legend:"hf"`
	Relationship                   *string `xml:"relationship" json:"relationship,omitempty"`
	RelevantEntityId               *int    `xml:"relevant_entity_id" json:"relevantEntityId,omitempty" legend:"entity"`
	RelevantIdForMethod            *int    `xml:"relevant_id_for_method" json:"relevantIdForMethod,omitempty"`
	RelevantPositionProfileId      *int    `xml:"relevant_position_profile_id" json:"relevantPositionProfileId,omitempty"`
	ReligionId                     *int    `xml:"religion_id" json:"religionId,omitempty"`
	ResidentCivId                  *int    `xml:"resident_civ_id" json:"residentCivId,omitempty" legend:"entity"`
	Return                         *string `xml:"return" json:"return,omitempty"`
	ScheduleId                     *int    `xml:"schedule_id" json:"scheduleId,omitempty"`
	SecretGoal                     *string `xml:"secret_goal" json:"secretGoal,omitempty"`
	SeekerHfid                     *int    `xml:"seeker_hfid" json:"seekerHfid,omitempty" legend:"hf"`
	ShrineAmountDestroyed          *int    `xml:"shrine_amount_destroyed" json:"shrineAmountDestroyed,omitempty"`
	SiteCivId                      *int    `xml:"site_civ_id" json:"siteCivId,omitempty" legend:"entity"`
	SiteEntityId                   *int    `xml:"site_entity_id" json:"siteEntityId,omitempty" legend:"entity"`
	SiteHfid                       *int    `xml:"site_hfid" json:"siteHfid,omitempty" legend:"hf"`
	SiteId                         *int    `xml:"site_id" json:"siteId,omitempty" legend:"site"`
	SiteId1                        *int    `xml:"site_id1" json:"siteId1,omitempty" legend:"site"`
	SiteId2                        *int    `xml:"site_id2" json:"siteId2,omitempty" legend:"site"`
	SiteId_1                       *int    `xml:"site_id_1" json:"siteId_1,omitempty" legend:"site"`
	SiteId_2                       *int    `xml:"site_id_2" json:"siteId_2,omitempty" legend:"site"`
	SitePropertyId                 *int    `xml:"site_property_id" json:"sitePropertyId,omitempty"`
	Situation                      *string `xml:"situation" json:"situation,omitempty"`
	SlayerCaste                    *string `xml:"slayer_caste" json:"slayerCaste,omitempty"`
	SlayerHfid                     *int    `xml:"slayer_hfid" json:"slayerHfid,omitempty" legend:"hf"`
	SlayerItemId                   *int    `xml:"slayer_item_id" json:"slayerItemId,omitempty"`
	SlayerRace                     *string `xml:"slayer_race" json:"slayerRace,omitempty"`
	SlayerShooterItemId            *int    `xml:"slayer_shooter_item_id" json:"slayerShooterItemId,omitempty"`
	SnatcherHfid                   *int    `xml:"snatcher_hfid" json:"snatcherHfid,omitempty" legend:"hf"`
	SourceEntityId                 *int    `xml:"source_entity_id" json:"sourceEntityId,omitempty" legend:"entity"`
	SourceSiteId                   *int    `xml:"source_site_id" json:"sourceSiteId,omitempty" legend:"site"`
	SourceStructureId              *int    `xml:"source_structure_id" json:"sourceStructureId,omitempty" legend:"structure"`
	SpeakerHfid                    *int    `xml:"speaker_hfid" json:"speakerHfid,omitempty" legend:"hf"`
	State                          *string `xml:"state" json:"state,omitempty"`
	StructureId                    *int    `xml:"structure_id" json:"structureId,omitempty" legend:"structure"`
	StudentHfid                    *int    `xml:"student_hfid" json:"studentHfid,omitempty" legend:"hf"`
	SubregionId                    *int    `xml:"subregion_id" json:"subregionId,omitempty"`
	Subtype                        *string `xml:"subtype" json:"subtype,omitempty"`
	Successful                     *string `xml:"successful" json:"successful,omitempty"`
	SurveiledContact               *string `xml:"surveiled_contact" json:"surveiledContact,omitempty"`
	SurveiledConvicted             *string `xml:"surveiled_convicted" json:"surveiledConvicted,omitempty"`
	TargetEnid                     *int    `xml:"target_enid" json:"targetEnid,omitempty" legend:"entity"`
	TargetHfid                     *int    `xml:"target_hfid" json:"targetHfid,omitempty" legend:"hf"`
	TargetIdentity                 *int    `xml:"target_identity" json:"targetIdentity,omitempty"`
	TargetSeenAs                   *string `xml:"target_seen_as" json:"targetSeenAs,omitempty"`
	TeacherHfid                    *int    `xml:"teacher_hfid" json:"teacherHfid,omitempty" legend:"hf"`
	TopFacet                       *string `xml:"top_facet" json:"topFacet,omitempty"`
	TopFacetModifier               *int    `xml:"top_facet_modifier" json:"topFacetModifier,omitempty"`
	TopFacetRating                 *int    `xml:"top_facet_rating" json:"topFacetRating,omitempty"`
	TopRelationshipFactor          *string `xml:"top_relationship_factor" json:"topRelationshipFactor,omitempty"`
	TopRelationshipModifier        *int    `xml:"top_relationship_modifier" json:"topRelationshipModifier,omitempty"`
	TopRelationshipRating          *int    `xml:"top_relationship_rating" json:"topRelationshipRating,omitempty"`
	TopValue                       *string `xml:"top_value" json:"topValue,omitempty"`
	TopValueModifier               *int    `xml:"top_value_modifier" json:"topValueModifier,omitempty"`
	TopValueRating                 *int    `xml:"top_value_rating" json:"topValueRating,omitempty"`
	Topic                          *string `xml:"topic" json:"topic,omitempty"`
	TraderEntityId                 *int    `xml:"trader_entity_id" json:"traderEntityId,omitempty" legend:"entity"`
	TraderHfid                     *int    `xml:"trader_hfid" json:"traderHfid,omitempty" legend:"hf"`
	TricksterHfid                  *int    `xml:"trickster_hfid" json:"tricksterHfid,omitempty" legend:"hf"`
	UnitId                         *int    `xml:"unit_id" json:"unitId,omitempty"`
	UnitType                       *string `xml:"unit_type" json:"unitType,omitempty"`
	WantedAndRecognized            *string `xml:"wanted_and_recognized" json:"wantedAndRecognized,omitempty"`
	WcId                           *int    `xml:"wc_id" json:"wcId,omitempty" legend:"wc"`
	Wcid                           *int    `xml:"wcid" json:"wcid,omitempty" legend:"wc"`
	WinnerHfid                     *int    `xml:"winner_hfid" json:"winnerHfid,omitempty" legend:"hf"`
	WoundeeHfid                    *int    `xml:"woundee_hfid" json:"woundeeHfid,omitempty" legend:"hf"`
	WounderHfid                    *int    `xml:"wounder_hfid" json:"wounderHfid,omitempty" legend:"hf"`
	WrongfulConviction             *string `xml:"wrongful_conviction" json:"wrongfulConviction,omitempty"`

	OtherElements
}

func (r *HistoricalEvent) Id() int      { return r.Id_ }
func (r *HistoricalEvent) Name() string { return r.Type() }

type EventObject struct {
	Events []*HistoricalEvent `json:"events"`
}

func (r *EventObject) GetEvents() []*HistoricalEvent       { return r.Events }
func (r *EventObject) SetEvents(events []*HistoricalEvent) { r.Events = events }

type HasEvents interface {
	GetEvents() []*HistoricalEvent
	SetEvents([]*HistoricalEvent)
}
