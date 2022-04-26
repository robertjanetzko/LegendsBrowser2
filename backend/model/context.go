package model

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

type context struct {
	world *DfWorld
	hfId  int
	story bool
}

func (c *context) hf(id int) string {
	if c.hfId != -1 {
		if c.hfId == id {
			return c.hfShort(id)
		} else {
			return c.hfRelated(id, c.hfId)
		}
	}
	if x, ok := c.world.HistoricalFigures[id]; ok {
		return fmt.Sprintf(`the %s <a class="hf" href="/hf/%d">%s</a>`, x.Race+util.If(x.Deity, " deity", "")+util.If(x.Force, " force", ""), x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func (c *context) hfShort(id int) string {
	if x, ok := c.world.HistoricalFigures[id]; ok {
		return fmt.Sprintf(`<a class="hf" href="/hf/%d">%s</a>`, x.Id(), util.Title(x.FirstName()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func (c *context) hfRelated(id, to int) string {
	if c.hfId != -1 {
		if c.hfId == id {
			return c.hfShort(id)
		} else {
			return c.hfRelated(id, c.hfId)
		}
	}
	if x, ok := c.world.HistoricalFigures[id]; ok {
		if t, ok := c.world.HistoricalFigures[to]; ok {
			if y, ok := util.Find(t.HfLink, func(l *HfLink) bool { return l.Hfid == id }); ok {
				return fmt.Sprintf(`%s %s <a class="hf" href="/hf/%d">%s</a>`, t.PossesivePronoun(), y.LinkType, x.Id(), util.Title(x.Name()))
			}
		}
		return fmt.Sprintf(`the %s <a class="hf" href="/hf/%d">%s</a>`, x.Race+util.If(x.Deity, " deity", "")+util.If(x.Force, " force", ""), x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func (c *context) hfList(ids []int) string {
	return andList(util.Map(ids, func(id int) string { return c.hf(id) }))
}

func (c *context) hfListRelated(ids []int, to int) string {
	return andList(util.Map(ids, func(id int) string { return c.hfRelated(id, to) }))
}

func (c *context) artifact(id int) string {
	if x, ok := c.world.Artifacts[id]; ok {
		return fmt.Sprintf(`<a class="artifact" href="/artifact/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN ARTIFACT"
}

func (c *context) entity(id int) string {
	if x, ok := c.world.Entities[id]; ok {
		return fmt.Sprintf(`<a class="entity" href="/entity/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN ENTITY"
}

func (c *context) entityList(ids []int) string {
	return andList(util.Map(ids, func(id int) string { return c.entity(id) }))
}

func (c *context) position(entityId, positionId, hfId int) string {
	if e, ok := c.world.Entities[entityId]; ok {
		if h, ok := c.world.HistoricalFigures[hfId]; ok {
			return e.Position(positionId).GenderName(h)
		}
	}
	return "UNKNOWN POSITION"
}

func (c *context) siteCiv(siteCivId, civId int) string {
	if siteCivId == civId {
		return c.entity(civId)
	}
	return util.If(siteCivId != -1, c.entity(siteCivId), "") + util.If(civId != -1 && siteCivId != -1, " of ", "") + util.If(civId != -1, c.entity(civId), "")
}

func (c *context) siteStructure(siteId, structureId int, prefix string) string {
	if siteId == -1 {
		return ""
	}
	return " " + prefix + " " + util.If(structureId != -1, c.structure(siteId, structureId)+" in ", "") + c.site(siteId, "")
}

func (c *context) site(id int, prefix string) string {
	if x, ok := c.world.Sites[id]; ok {
		return fmt.Sprintf(`%s <a class="site" href="/site/%d">%s</a>`, prefix, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN SITE"
}

func (c *context) structure(siteId, structureId int) string {
	if x, ok := c.world.Sites[siteId]; ok {
		if y, ok := x.Structures[structureId]; ok {
			return fmt.Sprintf(`<a class="structure" href="/site/%d/structure/%d">%s</a>`, siteId, structureId, util.Title(y.Name()))
		}
	}
	return "UNKNOWN STRUCTURE"
}

func (c *context) property(siteId, propertyId int) string {
	if x, ok := c.world.Sites[siteId]; ok {
		if y, ok := x.SiteProperties[propertyId]; ok {
			if y.StructureId != -1 {
				return c.structure(siteId, y.StructureId)
			}
			return articled(y.Type.String())
		}
	}
	return "UNKNOWN PROPERTY"
}

func (c *context) region(id int) string {
	if x, ok := c.world.Regions[id]; ok {
		return fmt.Sprintf(`<a class="region" href="/region/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN REGION"
}

func (c *context) location(siteId int, sitePrefix string, regionId int, regionPrefix string) string {
	if siteId != -1 {
		return c.site(siteId, sitePrefix)
	}
	if regionId != -1 {
		return regionPrefix + " " + c.region(regionId)
	}
	return ""
}

func (c *context) place(structureId, siteId int, sitePrefix string, regionId int, regionPrefix string) string {
	if siteId != -1 {
		return c.siteStructure(siteId, structureId, sitePrefix)
	}
	if regionId != -1 {
		return regionPrefix + " " + c.region(regionId)
	}
	return ""
}

func (c *context) mountain(id int) string {
	if x, ok := c.world.MountainPeaks[id]; ok {
		return fmt.Sprintf(`<a class="mountain" href="/site/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN MOUNTAIN"
}

func (c *context) identity(id int) string {
	if x, ok := c.world.Identities[id]; ok {
		return fmt.Sprintf(`<a class="identity" href="/region/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN IDENTITY"
}

func (c *context) fullIdentity(id int) string {
	if x, ok := c.world.Identities[id]; ok {
		return fmt.Sprintf(`&quot;the %s <a class="identity" href="/region/%d">%s</a> of %s&quot;`, x.Profession.String(), x.Id(), util.Title(x.Name()), c.entity(x.EntityId))
	}
	return "UNKNOWN IDENTITY"
}

func (c *context) danceForm(id int) string {
	if x, ok := c.world.DanceForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/danceForm/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN DANCE FORM"
}

func (c *context) musicalForm(id int) string {
	if x, ok := c.world.MusicalForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/musicalForm/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN MUSICAL FORM"
}

func (c *context) poeticForm(id int) string {
	if x, ok := c.world.PoeticForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/poeticForm/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN POETIC FORM"
}

func (c *context) worldConstruction(id int) string {
	if x, ok := c.world.WorldConstructions[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/wc/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN WORLD CONSTRUCTION"
}

func (c *context) writtenContent(id int) string {
	if x, ok := c.world.WrittenContents[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/writtenContent/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN WORLD CONSTRUCTION"
}

func (c *context) feature(x *Feature) string {
	switch x.Type {
	case FeatureType_DancePerformance:
		return "a perfomance of " + c.danceForm(x.Reference)
	case FeatureType_Images:
		if x.Reference != -1 {
			return "images of " + c.hf(x.Reference)
		}
		return "images"
	case FeatureType_MusicalPerformance:
		return "a perfomance of " + c.musicalForm(x.Reference)
	case FeatureType_PoetryRecital:
		return "a recital of " + c.poeticForm(x.Reference)
	case FeatureType_Storytelling:
		if x.Reference != -1 {
			if e, ok := c.world.HistoricalEvents[x.Reference]; ok {
				return "a telling of the story of " + e.Details.Html(&context{story: true}) + " in " + Time(e.Year, e.Seconds72)
			}
		}
		return "a story recital"
	default:
		return strcase.ToDelimited(x.Type.String(), ' ')
	}
}

func (c *context) schedule(x *Schedule) string {
	switch x.Type {
	case ScheduleType_DancePerformance:
		return "a perfomance of " + c.danceForm(x.Reference)
	case ScheduleType_MusicalPerformance:
		return "a perfomance of " + c.musicalForm(x.Reference)
	case ScheduleType_PoetryRecital:
		return "a recital of " + c.poeticForm(x.Reference)
	case ScheduleType_Storytelling:
		if x.Reference != -1 {
			if e, ok := c.world.HistoricalEvents[x.Reference]; ok {
				return "the story of " + e.Details.Html(&context{story: true}) + " in " + Time(e.Year, e.Seconds72)
			}
		}
		return "a story recital"
	default:
		return strcase.ToDelimited(x.Type.String(), ' ')
	}
}

func (c *context) pronoun(id int) string {
	if x, ok := c.world.HistoricalFigures[id]; ok {
		return x.Pronoun()
	}
	return "he"
}

func (c *context) posessivePronoun(id int) string {
	if x, ok := c.world.HistoricalFigures[id]; ok {
		return x.PossesivePronoun()
	}
	return "his"
}
