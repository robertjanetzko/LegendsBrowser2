package model

func (x *Entity) Icon() string {
	switch x.Type_ {
	case EntityType_Civilization:
		return "fa-solid fa-star"
	case EntityType_Guild:
		return "fa-solid fa-wrench"
	case EntityType_Merchantcompany:
		return "fa-solid fa-coins"
	case EntityType_Migratinggroup:
		return "fa-solid fa-person-hiking"
	case EntityType_Militaryunit:
		return "fa-solid fa-chess-knight"
	case EntityType_Nomadicgroup:
		return "fa-solid fa-tent"
	case EntityType_Outcast:
		return "fa-solid fa-campground"
	case EntityType_Performancetroupe:
		return "fa-solid fa-guitar"
	case EntityType_Religion:
		return "fa-solid fa-building-columns"
	case EntityType_Sitegovernment:
		return "fa-solid fa-scale-balanced"
	}
	return "fa-solid fa-star"
}

func (x *Site) Icon() string {
	switch x.Type_ {
	case SiteType_Camp:
		return "fa-solid fa-campground"
	case SiteType_Castle, SiteType_DarkFortress, SiteType_Fort, SiteType_Fortress:
		return "fa-brands fa-fort-awesome"
	case SiteType_Cave:
		return "fa-solid fa-mound"
	case SiteType_DarkPits:
		return "fa-solid fa-square"
	case SiteType_ForestRetreat:
		return "fa-solid fa-tree-city"
	case SiteType_Hamlet, SiteType_Hillocks:
		return "fa-solid fa-home"
	case SiteType_ImportantLocation:
		return "fa-solid fa-monument"
	case SiteType_Labyrinth:
		return "fa-solid fa-border-all"
	case SiteType_Lair:
		return "fa-solid fa-paw"
	case SiteType_Monastery:
		return "fa-solid fa-building-columns"
	case SiteType_MountainHalls:
		return "fa-solid fa-mountain-city"
	case SiteType_Shrine:
		return "fa-solid fa-landmark-dome"
	case SiteType_Tomb:
		return "fa-solid fa-circle-stop"
	case SiteType_Town:
		return "fa-solid fa-city"
	case SiteType_Tower:
		return "fa-solid fa-chess-rook"
	case SiteType_Vault:
		return "fa-solid fa-vault"
	}
	return ""
}

func (x *WorldConstruction) Icon() string {
	switch x.Type_ {
	case WorldConstructionType_Bridge:
		return "fa-solid fa-bridge"
	case WorldConstructionType_Road:
		return "fa-solid fa-road"
	case WorldConstructionType_Tunnel:
		return "fa-solid fa-archway"
	}
	return ""
}

func (e *Artifact) Icon() string {
	switch e.ItemSubtype {
	case "scroll":
		return "fa-solid fa-scroll"
	}
	switch e.ItemType {
	case "weapon":
		return "fa-solid fa-baseball-bat-ball"
	case "tool":
		return "fa-solid fa-wrench"
	case "book":
		return "fa-solid fa-book"
	case "slab":
		return "fa-solid fa-square"
	case "armor", "shoe", "gloves", "helm", "pants", "shield":
		return "fa-solid fa-shield"
	default:
		return "fa-solid fa-circle"
	}
}
