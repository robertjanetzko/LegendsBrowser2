// Code generated by legendsbrowser; DO NOT EDIT.
export interface Artifact {
	absTileX: number;
	absTileY: number;
	absTileZ: number;
	holderHfid: number;
	id: number;
	item: Item;
	itemDescription: string;
	itemSubtype: string;
	itemType: string;
	mat: string;
	name: string;
	pageCount: number;
	siteId: number;
	structureLocalId: number;
	subregionId: number;
	writing: number;
}
export interface Creature {
	allCastesAlive: boolean;
	artificialHiveable: boolean;
	biomeDesertBadland: boolean;
	biomeDesertRock: boolean;
	biomeDesertSand: boolean;
	biomeForestTaiga: boolean;
	biomeForestTemperateBroadleaf: boolean;
	biomeForestTemperateConifer: boolean;
	biomeForestTropicalConifer: boolean;
	biomeForestTropicalDryBroadleaf: boolean;
	biomeForestTropicalMoistBroadleaf: boolean;
	biomeGlacier: boolean;
	biomeGrasslandTemperate: boolean;
	biomeGrasslandTropical: boolean;
	biomeLakeTemperateBrackishwater: boolean;
	biomeLakeTemperateFreshwater: boolean;
	biomeLakeTemperateSaltwater: boolean;
	biomeLakeTropicalBrackishwater: boolean;
	biomeLakeTropicalFreshwater: boolean;
	biomeLakeTropicalSaltwater: boolean;
	biomeMarshTemperateFreshwater: boolean;
	biomeMarshTemperateSaltwater: boolean;
	biomeMarshTropicalFreshwater: boolean;
	biomeMarshTropicalSaltwater: boolean;
	biomeMountain: boolean;
	biomeOceanArctic: boolean;
	biomeOceanTemperate: boolean;
	biomeOceanTropical: boolean;
	biomePoolTemperateBrackishwater: boolean;
	biomePoolTemperateFreshwater: boolean;
	biomePoolTemperateSaltwater: boolean;
	biomePoolTropicalBrackishwater: boolean;
	biomePoolTropicalFreshwater: boolean;
	biomePoolTropicalSaltwater: boolean;
	biomeRiverTemperateBrackishwater: boolean;
	biomeRiverTemperateFreshwater: boolean;
	biomeRiverTemperateSaltwater: boolean;
	biomeRiverTropicalBrackishwater: boolean;
	biomeRiverTropicalFreshwater: boolean;
	biomeRiverTropicalSaltwater: boolean;
	biomeSavannaTemperate: boolean;
	biomeSavannaTropical: boolean;
	biomeShrublandTemperate: boolean;
	biomeShrublandTropical: boolean;
	biomeSubterraneanChasm: boolean;
	biomeSubterraneanLava: boolean;
	biomeSubterraneanWater: boolean;
	biomeSwampMangrove: boolean;
	biomeSwampTemperateFreshwater: boolean;
	biomeSwampTemperateSaltwater: boolean;
	biomeSwampTropicalFreshwater: boolean;
	biomeSwampTropicalSaltwater: boolean;
	biomeTundra: boolean;
	creatureId: string;
	doesNotExist: boolean;
	equipment: boolean;
	equipmentWagon: boolean;
	evil: boolean;
	fanciful: boolean;
	generated: boolean;
	good: boolean;
	hasAnyBenign: boolean;
	hasAnyCanSwim: boolean;
	hasAnyCannotBreatheAir: boolean;
	hasAnyCannotBreatheWater: boolean;
	hasAnyCarnivore: boolean;
	hasAnyCommonDomestic: boolean;
	hasAnyCuriousBeast: boolean;
	hasAnyDemon: boolean;
	hasAnyFeatureBeast: boolean;
	hasAnyFlier: boolean;
	hasAnyFlyRaceGait: boolean;
	hasAnyGrasp: boolean;
	hasAnyGrazer: boolean;
	hasAnyHasBlood: boolean;
	hasAnyImmobile: boolean;
	hasAnyIntelligentLearns: boolean;
	hasAnyIntelligentSpeaks: boolean;
	hasAnyLargePredator: boolean;
	hasAnyLocalPopsControllable: boolean;
	hasAnyLocalPopsProduceHeroes: boolean;
	hasAnyMegabeast: boolean;
	hasAnyMischievous: boolean;
	hasAnyNaturalAnimal: boolean;
	hasAnyNightCreature: boolean;
	hasAnyNightCreatureBogeyman: boolean;
	hasAnyNightCreatureHunter: boolean;
	hasAnyNightCreatureNightmare: boolean;
	hasAnyNotFireimmune: boolean;
	hasAnyNotLiving: boolean;
	hasAnyOutsiderControllable: boolean;
	hasAnyRaceGait: boolean;
	hasAnySemimegabeast: boolean;
	hasAnySlowLearner: boolean;
	hasAnySupernatural: boolean;
	hasAnyTitan: boolean;
	hasAnyUniqueDemon: boolean;
	hasAnyUtterances: boolean;
	hasAnyVerminHateable: boolean;
	hasAnyVerminMicro: boolean;
	hasFemale: boolean;
	hasMale: boolean;
	largeRoaming: boolean;
	looseClusters: boolean;
	matesToBreed: boolean;
	mundane: boolean;
	namePlural: string;
	nameSingular: string;
	occursAsEntityRace: boolean;
	savage: boolean;
	smallRace: boolean;
	twoGenders: boolean;
	ubiquitous: boolean;
	verminEater: boolean;
	verminFish: boolean;
	verminGrounder: boolean;
	verminRotter: boolean;
	verminSoil: boolean;
	verminSoilColony: boolean;
}
export interface DanceForm {
	description: string;
	id: number;
	name: string;
}
export interface DfWorld {
	altname: string;
	artifacts: { [key:number]:Artifact; };
	creatureRaw: Creature[];
	danceForms: { [key:number]:DanceForm; };
	entities: { [key:number]:Entity; };
	entityPopulations: { [key:number]:EntityPopulation; };
	historicalEras: HistoricalEra[];
	historicalEventCollections: { [key:number]:HistoricalEventCollection; };
	historicalEventRelationshipSupplements: HistoricalEventRelationshipSupplement[];
	historicalEventRelationships: HistoricalEventRelationship[];
	historicalEvents: { [key:number]:HistoricalEvent; };
	historicalFigures: { [key:number]:HistoricalFigure; };
	identities: { [key:number]:Identity; };
	landmasses: { [key:number]:Landmass; };
	mountainPeaks: { [key:number]:MountainPeak; };
	musicalForms: { [key:number]:MusicalForm; };
	name: string;
	poeticForms: { [key:number]:PoeticForm; };
	regions: { [key:number]:Region; };
	rivers: River[];
	sites: { [key:number]:Site; };
	undergroundRegions: { [key:number]:UndergroundRegion; };
	worldConstructions: { [key:number]:WorldConstruction; };
	writtenContents: { [key:number]:WrittenContent; };
}
export interface Entity {
	child: number[];
	claims: string;
	entityLink: EntityEntityLink[];
	entityPosition: EntityPosition[];
	entityPositionAssignment: EntityPositionAssignment[];
	histfigId: number[];
	honor: Honor[];
	id: number;
	name: string;
	occasion: Occasion[];
	profession: string;
	race: string;
	type: string;
	weapon: string[];
	worshipId: number[];
}
export interface EntityEntityLink {
	strength: number;
	target: number;
	type: string;
}
export interface EntityFormerPositionLink {
	endYear: number;
	entityId: number;
	positionProfileId: number;
	startYear: number;
}
export interface EntityPopulation {
	civId: number;
	id: number;
	race: string;
}
export interface EntityPosition {
	id: number;
	name: string;
	nameFemale: string;
	nameMale: string;
	spouse: string;
	spouseFemale: string;
	spouseMale: string;
}
export interface EntityPositionAssignment {
	histfig: number;
	id: number;
	positionId: number;
	squadId: number;
}
export interface EntityPositionLink {
	entityId: number;
	positionProfileId: number;
	startYear: number;
}
export interface EntityReputation {
	entityId: number;
	firstAgelessSeasonCount: number;
	firstAgelessYear: number;
	unsolvedMurders: number;
}
export interface EntitySquadLink {
	entityId: number;
	squadId: number;
	squadPosition: number;
	startYear: number;
}
export interface Feature {
	reference: number;
	type: string;
}
export interface HfLink {
	hfid: number;
	linkStrength: number;
	linkType: string;
}
export interface HfSkill {
	skill: string;
	totalIp: number;
}
export interface HistoricalEra {
	name: string;
	startYear: number;
}
export interface HistoricalEvent {
	id: number;
	seconds72: number;
	year: number;
	Details: any;
}
export interface HistoricalEventAddHfEntityHonor {
	entityId: number;
	hfid: number;
	honorId: number;
}
export interface HistoricalEventAddHfEntityLink {
	appointerHfid: number;
	civ: number;
	civId: number;
	hfid: number;
	histfig: number;
	link: string;
	linkType: string;
	position: string;
	positionId: number;
	promiseToHfid: number;
}
export interface HistoricalEventAddHfHfLink {
	hf: number;
	hfTarget: number;
	hfid: number;
	hfidTarget: number;
	linkType: string;
}
export interface HistoricalEventAddHfSiteLink {
	civ: number;
	histfig: number;
	linkType: string;
	site: number;
	siteId: number;
	structure: number;
}
export interface HistoricalEventAgreementFormed {
	action: string;
	agreementId: number;
	allyDefenseBonus: number;
	coconspiratorBonus: number;
	delegated: boolean;
	failedJudgmentTest: boolean;
	method: string;
	relevantEntityId: number;
	relevantIdForMethod: number;
	relevantPositionProfileId: number;
	successful: boolean;
	topFacet: string;
	topFacetModifier: number;
	topFacetRating: number;
	topRelationshipFactor: string;
	topRelationshipModifier: number;
	topRelationshipRating: number;
	topValue: string;
	topValueModifier: number;
	topValueRating: number;
}
export interface HistoricalEventArtifactClaimFormed {
	artifactId: number;
	circumstance: string;
	claim: string;
	entityId: number;
	histFigureId: number;
	positionProfileId: number;
}
export interface HistoricalEventArtifactCopied {
	artifactId: number;
	destEntityId: number;
	destSiteId: number;
	destStructureId: number;
	fromOriginal: boolean;
	sourceEntityId: number;
	sourceSiteId: number;
	sourceStructureId: number;
}
export interface HistoricalEventArtifactCreated {
	artifactId: number;
	circumstance: HistoricalEventArtifactCreatedCircumstance;
	creatorHfid: number;
	creatorUnitId: number;
	histFigureId: number;
	nameOnly: boolean;
	reason: string;
	sanctifyHf: number;
	site: number;
	siteId: number;
	unitId: number;
}
export interface HistoricalEventArtifactCreatedCircumstance {
	defeated: number;
	type: string;
}
export interface HistoricalEventArtifactDestroyed {
	artifactId: number;
	destroyerEnid: number;
	siteId: number;
}
export interface HistoricalEventArtifactFound {
	artifactId: number;
	histFigureId: number;
	siteId: number;
	unitId: number;
}
export interface HistoricalEventArtifactGiven {
	artifactId: number;
	giverEntityId: number;
	giverHistFigureId: number;
	receiverEntityId: number;
	receiverHistFigureId: number;
}
export interface HistoricalEventArtifactLost {
	artifactId: number;
	siteId: number;
	sitePropertyId: number;
	subregionId: number;
}
export interface HistoricalEventArtifactPossessed {
	artifactId: number;
	circumstance: string;
	circumstanceId: number;
	featureLayerId: number;
	histFigureId: number;
	reason: string;
	reasonId: number;
	siteId: number;
	subregionId: number;
	unitId: number;
}
export interface HistoricalEventArtifactRecovered {
	artifactId: number;
	featureLayerId: number;
	histFigureId: number;
	siteId: number;
	structureId: number;
	subregionId: number;
	unitId: number;
}
export interface HistoricalEventArtifactStored {
	artifactId: number;
	histFigureId: number;
	siteId: number;
	unitId: number;
}
export interface HistoricalEventAssumeIdentity {
	identityCaste: string;
	identityHistfigId: number;
	identityId: number;
	identityName: string;
	identityNemesisId: number;
	identityRace: string;
	target: number;
	targetEnid: number;
	trickster: number;
	tricksterHfid: number;
}
export interface HistoricalEventAttackedSite {
	aSupportMercEnid: number;
	attackerCivId: number;
	attackerGeneralHfid: number;
	attackerMercEnid: number;
	dSupportMercEnid: number;
	defenderCivId: number;
	defenderGeneralHfid: number;
	defenderMercEnid: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventBodyAbused {
	abuseType: string;
	bodies: number[];
	civ: number;
	coords: string;
	featureLayerId: number;
	histfig: number;
	interaction: number;
	itemMat: string;
	itemSubtype: string;
	itemType: string;
	pileType: string;
	site: number;
	siteId: number;
	structure: number;
	subregionId: number;
	tree: number;
	victimEntity: number;
}
export interface HistoricalEventBuildingProfileAcquired {
	acquirerEnid: number;
	acquirerHfid: number;
	buildingProfileId: number;
	inherited: boolean;
	lastOwnerHfid: number;
	purchasedUnowned: boolean;
	rebuiltRuined: boolean;
	siteId: number;
}
export interface HistoricalEventCeremony {
	civId: number;
	featureLayerId: number;
	occasionId: number;
	scheduleId: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventChangeHfBodyState {
	bodyState: string;
	coords: string;
	featureLayerId: number;
	hfid: number;
	siteId: number;
	structureId: number;
	subregionId: number;
}
export interface HistoricalEventChangeHfJob {
	featureLayerId: number;
	hfid: number;
	newJob: string;
	oldJob: string;
	site: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventChangeHfState {
	coords: string;
	featureLayerId: number;
	hfid: number;
	mood: string;
	reason: string;
	site: number;
	siteId: number;
	state: string;
	subregionId: number;
}
export interface HistoricalEventChangedCreatureType {
	changee: number;
	changeeHfid: number;
	changer: number;
	changerHfid: number;
	newCaste: string;
	newRace: string;
	oldCaste: string;
	oldRace: string;
}
export interface HistoricalEventCollection {
	endSeconds72: number;
	endYear: number;
	event: number[];
	eventcol: number[];
	id: number;
	startSeconds72: number;
	startYear: number;
	Details: any;
}
export interface HistoricalEventCollectionAbduction {
	attackingEnid: number;
	coords: string;
	defendingEnid: number;
	featureLayerId: number;
	ordinal: number;
	parentEventcol: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventCollectionBattle {
	aSupportMercEnid: number;
	aSupportMercHfid: number[];
	attackingHfid: number[];
	attackingMercEnid: number;
	attackingSquadAnimated: boolean[];
	attackingSquadDeaths: number[];
	attackingSquadEntityPop: number[];
	attackingSquadNumber: number[];
	attackingSquadRace: string[];
	attackingSquadSite: number[];
	companyMerc: boolean[];
	coords: string;
	dSupportMercEnid: number;
	dSupportMercHfid: number[];
	defendingHfid: number[];
	defendingMercEnid: number;
	defendingSquadAnimated: boolean[];
	defendingSquadDeaths: number[];
	defendingSquadEntityPop: number[];
	defendingSquadNumber: number[];
	defendingSquadRace: string[];
	defendingSquadSite: number[];
	featureLayerId: number;
	individualMerc: boolean[];
	name: string;
	noncomHfid: number[];
	outcome: string;
	siteId: number;
	subregionId: number;
	warEventcol: number;
}
export interface HistoricalEventCollectionBeastAttack {
	coords: string;
	defendingEnid: number;
	featureLayerId: number;
	ordinal: number;
	parentEventcol: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventCollectionCeremony {
	ordinal: number;
}
export interface HistoricalEventCollectionCompetition {
	ordinal: number;
}
export interface HistoricalEventCollectionDuel {
	attackingHfid: number;
	coords: string;
	defendingHfid: number;
	featureLayerId: number;
	ordinal: number;
	parentEventcol: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventCollectionEntityOverthrown {
	ordinal: number;
	siteId: number;
	targetEntityId: number;
}
export interface HistoricalEventCollectionJourney {
	ordinal: number;
}
export interface HistoricalEventCollectionOccasion {
	civId: number;
	occasionId: number;
	ordinal: number;
}
export interface HistoricalEventCollectionPerformance {
	ordinal: number;
}
export interface HistoricalEventCollectionPersecution {
	ordinal: number;
	siteId: number;
	targetEntityId: number;
}
export interface HistoricalEventCollectionProcession {
	ordinal: number;
}
export interface HistoricalEventCollectionPurge {
	adjective: string;
	ordinal: number;
	siteId: number;
}
export interface HistoricalEventCollectionSiteConquered {
	attackingEnid: number;
	defendingEnid: number;
	ordinal: number;
	siteId: number;
	warEventcol: number;
}
export interface HistoricalEventCollectionTheft {
	attackingEnid: number;
	coords: string;
	defendingEnid: number;
	featureLayerId: number;
	ordinal: number;
	parentEventcol: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventCollectionWar {
	aggressorEntId: number;
	defenderEntId: number;
	name: string;
}
export interface HistoricalEventCompetition {
	civId: number;
	competitorHfid: number[];
	featureLayerId: number;
	occasionId: number;
	scheduleId: number;
	siteId: number;
	subregionId: number;
	winnerHfid: number;
}
export interface HistoricalEventCreateEntityPosition {
	civ: number;
	histfig: number;
	position: string;
	reason: string;
	siteCiv: number;
}
export interface HistoricalEventCreatedSite {
	builderHfid: number;
	civId: number;
	residentCivId: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventCreatedStructure {
	builderHf: number;
	builderHfid: number;
	civ: number;
	civId: number;
	rebuild: string;
	rebuilt: boolean;
	site: number;
	siteCiv: number;
	siteCivId: number;
	siteId: number;
	structure: number;
	structureId: number;
}
export interface HistoricalEventCreatedWorldConstruction {
	civId: number;
	masterWcid: number;
	siteCivId: number;
	siteId1: number;
	siteId2: number;
	wcid: number;
}
export interface HistoricalEventCreatureDevoured {
	caste: string;
	eater: number;
	entity: number;
	featureLayerId: number;
	race: string;
	site: number;
	siteId: number;
	subregionId: number;
	victim: number;
}
export interface HistoricalEventDanceFormCreated {
	circumstance: string;
	circumstanceId: number;
	formId: number;
	histFigureId: number;
	reason: string;
	reasonId: number;
	siteId: number;
}
export interface HistoricalEventDestroyedSite {
	attackerCivId: number;
	defenderCivId: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventEntityAllianceFormed {
	initiatingEnid: number;
	joiningEnid: number[];
}
export interface HistoricalEventEntityBreachFeatureLayer {
	civEntityId: number;
	featureLayerId: number;
	siteEntityId: number;
	siteId: number;
}
export interface HistoricalEventEntityCreated {
	creatorHfid: number;
	entityId: number;
	siteId: number;
	structureId: number;
}
export interface HistoricalEventEntityDissolved {
	entityId: number;
	reason: string;
}
export interface HistoricalEventEntityEquipmentPurchase {
	entityId: number;
	hfid: number[];
	newEquipmentLevel: number;
}
export interface HistoricalEventEntityIncorporated {
	joinedEntityId: number;
	joinerEntityId: number;
	leaderHfid: number;
	partialIncorporation: boolean;
	siteId: number;
}
export interface HistoricalEventEntityLaw {
	entityId: number;
	histFigureId: number;
	lawAdd: string;
	lawRemove: string;
}
export interface HistoricalEventEntityOverthrown {
	conspiratorHfid: number[];
	entityId: number;
	instigatorHfid: number;
	overthrownHfid: number;
	posTakerHfid: number;
	positionProfileId: number;
	siteId: number;
}
export interface HistoricalEventEntityPersecuted {
	destroyedStructureId: number;
	expelledCreature: number[];
	expelledHfid: number[];
	expelledNumber: number[];
	expelledPopId: number[];
	persecutorEnid: number;
	persecutorHfid: number;
	propertyConfiscatedFromHfid: number[];
	shrineAmountDestroyed: number;
	siteId: number;
	targetEnid: number;
}
export interface HistoricalEventEntityPrimaryCriminals {
	action: string;
	entity: number;
	entityId: number;
	site: number;
	siteId: number;
	structure: number;
	structureId: number;
}
export interface HistoricalEventEntityRelocate {
	action: string;
	entity: number;
	entityId: number;
	site: number;
	siteId: number;
	structure: number;
	structureId: number;
}
export interface HistoricalEventFailedFrameAttempt {
	convicterEnid: number;
	crime: string;
	fooledHfid: number;
	framerHfid: number;
	plotterHfid: number;
	targetHfid: number;
}
export interface HistoricalEventFailedIntrigueCorruption {
	action: string;
	allyDefenseBonus: number;
	coconspiratorBonus: number;
	corruptorHfid: number;
	corruptorIdentity: number;
	failedJudgmentTest: boolean;
	featureLayerId: number;
	lureHfid: number;
	method: string;
	relevantEntityId: number;
	relevantIdForMethod: number;
	relevantPositionProfileId: number;
	siteId: number;
	subregionId: number;
	targetHfid: number;
	targetIdentity: number;
	topFacet: string;
	topFacetModifier: number;
	topFacetRating: number;
	topRelationshipFactor: string;
	topRelationshipModifier: number;
	topRelationshipRating: number;
	topValue: string;
	topValueModifier: number;
	topValueRating: number;
}
export interface HistoricalEventFieldBattle {
	aSupportMercEnid: number;
	attackerCivId: number;
	attackerGeneralHfid: number;
	attackerMercEnid: number;
	coords: string;
	dSupportMercEnid: number;
	defenderCivId: number;
	defenderGeneralHfid: number;
	defenderMercEnid: number;
	featureLayerId: number;
	subregionId: number;
}
export interface HistoricalEventGamble {
	gamblerHfid: number;
	newAccount: number;
	oldAccount: number;
	siteId: number;
	structureId: number;
}
export interface HistoricalEventHfAbducted {
	featureLayerId: number;
	siteId: number;
	snatcherHfid: number;
	subregionId: number;
	targetHfid: number;
}
export interface HistoricalEventHfAttackedSite {
	attackerHfid: number;
	defenderCivId: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventHfConfronted {
	coords: string;
	featureLayerId: number;
	hfid: number;
	reason: string[];
	siteId: number;
	situation: string;
	subregionId: number;
}
export interface HistoricalEventHfConvicted {
	coconspiratorHfid: number;
	confessedAfterApbArrestEnid: number;
	contactHfid: number;
	convictIsContact: boolean;
	convictedHfid: number;
	convicterEnid: number;
	corruptConvicterHfid: number;
	crime: string;
	deathPenalty: boolean;
	didNotRevealAllInInterrogation: boolean;
	exiled: boolean;
	fooledHfid: number;
	framerHfid: number;
	heldFirmInInterrogation: boolean;
	implicatedHfid: number[];
	interrogatorHfid: number;
	plotterHfid: number;
	prisonMonths: number;
	surveiledCoconspirator: boolean;
	surveiledContact: boolean;
	surveiledConvicted: boolean;
	surveiledTarget: boolean;
	targetHfid: number;
	wrongfulConviction: boolean;
}
export interface HistoricalEventHfDestroyedSite {
	attackerHfid: number;
	defenderCivId: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventHfDied {
	artifactId: number;
	cause: string;
	deathCause: string;
	featureLayerId: number;
	hfid: number;
	item: number;
	itemSubtype: string;
	itemType: string;
	mat: string;
	site: number;
	siteId: number;
	slayerCaste: string;
	slayerHf: number;
	slayerHfid: number;
	slayerItemId: number;
	slayerRace: string;
	slayerShooterItemId: number;
	subregionId: number;
	victimHf: number;
}
export interface HistoricalEventHfDisturbedStructure {
	action: string;
	histFigId: number;
	histfig: number;
	site: number;
	siteId: number;
	structure: number;
	structureId: number;
}
export interface HistoricalEventHfDoesInteraction {
	doer: number;
	doerHfid: number;
	interaction: string;
	interactionAction: string;
	region: number;
	site: number;
	source: number;
	target: number;
	targetHfid: number;
}
export interface HistoricalEventHfEnslaved {
	enslavedHfid: number;
	movedToSiteId: number;
	payerEntityId: number;
	sellerHfid: number;
}
export interface HistoricalEventHfEquipmentPurchase {
	featureLayerId: number;
	groupHfid: number;
	quality: number;
	siteId: number;
	structureId: number;
	subregionId: number;
}
export interface HistoricalEventHfGainsSecretGoal {
	hfid: number;
	secretGoal: string;
}
export interface HistoricalEventHfInterrogated {
	arrestingEnid: number;
	heldFirmInInterrogation: boolean;
	interrogatorHfid: number;
	targetHfid: number;
	wantedAndRecognized: boolean;
}
export interface HistoricalEventHfLearnsSecret {
	artifact: number;
	artifactId: number;
	interaction: string;
	secretText: string;
	student: number;
	studentHfid: number;
	teacher: number;
	teacherHfid: number;
}
export interface HistoricalEventHfNewPet {
	coords: string;
	featureLayerId: number;
	group: number;
	groupHfid: number;
	pets: string;
	site: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventHfPerformedHorribleExperiments {
	featureLayerId: number;
	groupHfid: number;
	siteId: number;
	structureId: number;
	subregionId: number;
}
export interface HistoricalEventHfPrayedInsideStructure {
	action: string;
	histFigId: number;
	histfig: number;
	site: number;
	siteId: number;
	structure: number;
	structureId: number;
}
export interface HistoricalEventHfPreach {
	entity1: number;
	entity2: number;
	siteHfid: number;
	speakerHfid: number;
	topic: string;
}
export interface HistoricalEventHfProfanedStructure {
	action: string;
	histFigId: number;
	histfig: number;
	site: number;
	siteId: number;
	structure: number;
	structureId: number;
}
export interface HistoricalEventHfRecruitedUnitTypeForEntity {
	entityId: number;
	featureLayerId: number;
	hfid: number;
	siteId: number;
	subregionId: number;
	unitType: string;
}
export interface HistoricalEventHfRelationshipDenied {
	featureLayerId: number;
	reason: string;
	reasonId: number;
	relationship: string;
	seekerHfid: number;
	siteId: number;
	subregionId: number;
	targetHfid: number;
}
export interface HistoricalEventHfReunion {
	featureLayerId: number;
	group1Hfid: number;
	group2Hfid: number[];
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventHfRevived {
	actorHfid: number;
	disturbance: boolean;
	featureLayerId: number;
	hfid: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventHfSimpleBattleEvent {
	featureLayerId: number;
	group1Hfid: number;
	group2Hfid: number;
	siteId: number;
	subregionId: number;
	subtype: string;
}
export interface HistoricalEventHfTravel {
	coords: string;
	featureLayerId: number;
	groupHfid: number[];
	return: boolean;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventHfViewedArtifact {
	artifactId: number;
	histFigId: number;
	siteId: number;
	structureId: number;
}
export interface HistoricalEventHfWounded {
	bodyPart: number;
	featureLayerId: number;
	injuryType: string;
	partLost: string;
	site: number;
	siteId: number;
	subregionId: number;
	woundee: number;
	woundeeCaste: number;
	woundeeHfid: number;
	woundeeRace: number;
	wounder: number;
	wounderHfid: number;
}
export interface HistoricalEventHfsFormedIntrigueRelationship {
	action: string;
	allyDefenseBonus: number;
	circumstance: string;
	circumstanceId: number;
	coconspiratorBonus: number;
	corruptorHfid: number;
	corruptorIdentity: number;
	corruptorSeenAs: string;
	failedJudgmentTest: boolean;
	featureLayerId: number;
	lureHfid: number;
	method: string;
	relevantEntityId: number;
	relevantIdForMethod: number;
	relevantPositionProfileId: number;
	siteId: number;
	subregionId: number;
	successful: boolean;
	targetHfid: number;
	targetIdentity: number;
	targetSeenAs: string;
	topFacet: string;
	topFacetModifier: number;
	topFacetRating: number;
	topRelationshipFactor: string;
	topRelationshipModifier: number;
	topRelationshipRating: number;
	topValue: string;
	topValueModifier: number;
	topValueRating: number;
}
export interface HistoricalEventHfsFormedReputationRelationship {
	featureLayerId: number;
	hfRep1Of2: string;
	hfRep2Of1: string;
	hfid1: number;
	hfid2: number;
	identityId1: number;
	identityId2: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventHolyCityDeclaration {
	religionId: number;
	siteId: number;
}
export interface HistoricalEventItemStolen {
	circumstance: HistoricalEventItemStolenCircumstance;
	circumstanceId: number;
	entity: number;
	histfig: number;
	item: number;
	itemSubtype: string;
	itemType: string;
	mat: string;
	matindex: number;
	mattype: number;
	site: number;
	stashSite: number;
	structure: number;
	theftMethod: string;
}
export interface HistoricalEventItemStolenCircumstance {
	defeated: number;
	histEventCollection: number;
	murdered: number;
	type: string;
}
export interface HistoricalEventKnowledgeDiscovered {
	first: boolean;
	hfid: number;
	knowledge: string;
}
export interface HistoricalEventMasterpieceItem {
	entityId: number;
	hfid: number;
	itemId: number;
	itemType: string;
	maker: number;
	makerEntity: number;
	mat: string;
	site: number;
	siteId: number;
	skillAtTime: string;
}
export interface HistoricalEventMerchant {
	depotEntityId: number;
	destination: number;
	site: number;
	siteId: number;
	source: number;
	traderEntityId: number;
}
export interface HistoricalEventModifiedBuilding {
	modification: string;
	modifierHfid: number;
	siteId: number;
	structureId: number;
}
export interface HistoricalEventMusicalFormCreated {
	circumstance: string;
	circumstanceId: number;
	formId: number;
	histFigureId: number;
	reason: string;
	reasonId: number;
	siteId: number;
}
export interface HistoricalEventNewSiteLeader {
	attackerCivId: number;
	defenderCivId: number;
	newLeaderHfid: number;
	newSiteCivId: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventPeaceAccepted {
	destination: number;
	site: number;
	siteId: number;
	source: number;
	topic: string;
}
export interface HistoricalEventPeaceRejected {
	destination: number;
	site: number;
	siteId: number;
	source: number;
	topic: string;
}
export interface HistoricalEventPerformance {
	civId: number;
	featureLayerId: number;
	occasionId: number;
	scheduleId: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventPlunderedSite {
	attackerCivId: number;
	defenderCivId: number;
	detected: boolean;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventPoeticFormCreated {
	circumstance: string;
	formId: number;
	histFigureId: number;
	siteId: number;
}
export interface HistoricalEventProcession {
	civId: number;
	featureLayerId: number;
	occasionId: number;
	scheduleId: number;
	siteId: number;
	subregionId: number;
}
export interface HistoricalEventRazedStructure {
	civId: number;
	siteId: number;
	structureId: number;
}
export interface HistoricalEventReclaimSite {
	civId: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventRegionpopIncorporatedIntoEntity {
	joinEntityId: number;
	popFlid: number;
	popNumberMoved: number;
	popRace: number;
	popSrid: number;
	siteId: number;
}
export interface HistoricalEventRelationship {
	event: number;
	relationship: string;
	sourceHf: number;
	targetHf: number;
	year: number;
}
export interface HistoricalEventRelationshipSupplement {
	event: number;
	occasionType: number;
	site: number;
	unk1: number;
}
export interface HistoricalEventRemoveHfEntityLink {
	civ: number;
	civId: number;
	hfid: number;
	histfig: number;
	link: string;
	linkType: string;
	position: string;
	positionId: number;
}
export interface HistoricalEventRemoveHfHfLink {
	hfid: number;
	hfidTarget: number;
}
export interface HistoricalEventRemoveHfSiteLink {
	civ: number;
	histfig: number;
	linkType: string;
	site: number;
	siteId: number;
	structure: number;
}
export interface HistoricalEventReplacedStructure {
	civ: number;
	civId: number;
	newAbId: number;
	newStructure: number;
	oldAbId: number;
	oldStructure: number;
	site: number;
	siteCiv: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventSiteDispute {
	dispute: string;
	entityId1: number;
	entityId2: number;
	siteId1: number;
	siteId2: number;
}
export interface HistoricalEventSiteTakenOver {
	attackerCivId: number;
	defenderCivId: number;
	newSiteCivId: number;
	siteCivId: number;
	siteId: number;
}
export interface HistoricalEventSquadVsSquad {
	aHfid: number;
	aSquadId: number;
	dEffect: number;
	dInteraction: number;
	dNumber: number;
	dRace: number;
	dSlain: number;
	dSquadId: number;
	featureLayerId: number;
	siteId: number;
	structureId: number;
	subregionId: number;
}
export interface HistoricalEventTacticalSituation {
	aTacticianHfid: number;
	aTacticsRoll: number;
	dTacticianHfid: number;
	dTacticsRoll: number;
	featureLayerId: number;
	siteId: number;
	situation: string;
	start: boolean;
	structureId: number;
	subregionId: number;
}
export interface HistoricalEventTrade {
	accountShift: number;
	allotment: number;
	allotmentIndex: number;
	destSiteId: number;
	productionZoneId: number;
	sourceSiteId: number;
	traderEntityId: number;
	traderHfid: number;
}
export interface HistoricalEventWrittenContentComposed {
	circumstance: string;
	circumstanceId: number;
	histFigureId: number;
	reason: string;
	reasonId: number;
	siteId: number;
	subregionId: number;
	wcId: number;
}
export interface HistoricalFigure {
	activeInteraction: string[];
	animated: boolean;
	animatedString: string;
	appeared: number;
	associatedType: string;
	birthSeconds72: number;
	birthYear: number;
	caste: string;
	currentIdentityId: number;
	deathSeconds72: number;
	deathYear: number;
	deity: boolean;
	entPopId: number;
	entityFormerPositionLink: EntityFormerPositionLink[];
	entityLink: HistoricalFigureEntityLink[];
	entityPositionLink: EntityPositionLink[];
	entityReputation: EntityReputation[];
	entitySquadLink: EntitySquadLink;
	force: boolean;
	goal: string[];
	hfLink: HfLink[];
	hfSkill: HfSkill[];
	holdsArtifact: number[];
	honorEntity: HonorEntity[];
	id: number;
	interactionKnowledge: string[];
	intrigueActor: IntrigueActor[];
	intriguePlot: IntriguePlot[];
	journeyPet: string[];
	name: string;
	race: string;
	relationshipProfileHfHistorical: RelationshipProfileHfHistorical[];
	relationshipProfileHfVisual: RelationshipProfileHfVisual[];
	sex: number;
	siteLink: SiteLink[];
	siteProperty: HistoricalFigureSiteProperty[];
	sphere: string[];
	usedIdentityId: number[];
	vagueRelationship: VagueRelationship[];
}
export interface HistoricalFigureEntityLink {
	entityId: number;
	linkStrength: number;
	linkType: string;
}
export interface HistoricalFigureSiteProperty {
	propertyId: number;
	siteId: number;
}
export interface Honor {
	exemptEpid: number;
	exemptFormerEpid: number;
	givesPrecedence: number;
	grantedToEverybody: boolean;
	id: number;
	name: string;
	requiredBattles: number;
	requiredKills: number;
	requiredSkill: string;
	requiredSkillIpTotal: number;
	requiredYears: number;
	requiresAnyMeleeOrRangedSkill: boolean;
}
export interface HonorEntity {
	battles: number;
	entity: number;
	honorId: number[];
	kills: number;
}
export interface Identity {
	birthSecond: number;
	birthYear: number;
	caste: string;
	entityId: number;
	histfigId: number;
	id: number;
	name: string;
	nemesisId: number;
	profession: string;
	race: string;
}
export interface IntrigueActor {
	entityId: number;
	handleActorId: number;
	hfid: number;
	localId: number;
	promisedActorImmortality: boolean;
	promisedMeImmortality: boolean;
	role: string;
	strategy: string;
	strategyEnid: number;
	strategyEppid: number;
}
export interface IntriguePlot {
	actorId: number;
	artifactId: number;
	delegatedPlotHfid: number;
	delegatedPlotId: number;
	entityId: number;
	localId: number;
	onHold: boolean;
	parentPlotHfid: number;
	parentPlotId: number;
	plotActor: PlotActor[];
	type: string;
}
export interface Item {
	nameString: string;
	pageNumber: number;
	pageWrittenContentId: number;
	writingWrittenContentId: number;
}
export interface Landmass {
	coord1: string;
	coord2: string;
	id: number;
	name: string;
}
export interface MountainPeak {
	coords: string;
	height: number;
	id: number;
	isVolcano: boolean;
	name: string;
}
export interface MusicalForm {
	description: string;
	id: number;
	name: string;
}
export interface Occasion {
	event: number;
	id: number;
	name: string;
	schedule: Schedule[];
}
export interface PlotActor {
	actorId: number;
	agreementHasMessenger: boolean;
	agreementId: number;
	plotRole: string;
}
export interface PoeticForm {
	description: string;
	id: number;
	name: string;
}
export interface Reference {
	id: number;
	type: string;
}
export interface Region {
	coords: string;
	evilness: string;
	forceId: number;
	id: number;
	name: string;
	type: string;
}
export interface RelationshipProfileHfHistorical {
	fear: number;
	hfId: number;
	love: number;
	loyalty: number;
	respect: number;
	trust: number;
}
export interface RelationshipProfileHfVisual {
	fear: number;
	hfId: number;
	knownIdentityId: number;
	lastMeetSeconds72: number;
	lastMeetYear: number;
	love: number;
	loyalty: number;
	meetCount: number;
	repFriendly: number;
	repInformationSource: number;
	respect: number;
	trust: number;
}
export interface River {
	endPos: string;
	name: string;
	path: string;
}
export interface Schedule {
	feature: Feature[];
	id: number;
	itemSubtype: string;
	itemType: string;
	reference: number;
	reference2: number;
	type: string;
}
export interface Site {
	civId: number;
	coords: string;
	curOwnerId: number;
	id: number;
	name: string;
	rectangle: string;
	siteProperties: { [key:number]:SiteSiteProperty; };
	structures: { [key:number]:Structure; };
	type: string;
}
export interface SiteLink {
	entityId: number;
	linkType: string;
	occupationId: number;
	siteId: number;
	subId: number;
}
export interface SiteSiteProperty {
	id: number;
	ownerHfid: number;
	structureId: number;
	type: string;
}
export interface Structure {
	copiedArtifactId: number[];
	deity: number;
	deityType: number;
	dungeonType: number;
	entityId: number;
	id: number;
	inhabitant: number[];
	localId: number;
	name: string;
	name2: string;
	religion: number;
	subtype: string;
	type: string;
	worshipHfid: number;
}
export interface UndergroundRegion {
	coords: string;
	depth: number;
	id: number;
	type: string;
}
export interface VagueRelationship {
	artisticBuddy: boolean;
	atheleticRival: boolean;
	athleteBuddy: boolean;
	businessRival: boolean;
	childhoodFriend: boolean;
	grudge: boolean;
	hfid: number;
	jealousObsession: boolean;
	jealousRelationshipGrudge: boolean;
	persecutionGrudge: boolean;
	religiousPersecutionGrudge: boolean;
	scholarBuddy: boolean;
	supernaturalGrudge: boolean;
	warBuddy: boolean;
}
export interface WorldConstruction {
	coords: string;
	id: number;
	name: string;
	type: string;
}
export interface WrittenContent {
	author: number;
	authorHfid: number;
	authorRoll: number;
	form: string;
	formId: number;
	id: number;
	pageEnd: number;
	pageStart: number;
	reference: Reference[];
	style: string[];
	title: string;
	type: string;
}
