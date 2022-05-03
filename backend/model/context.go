package model

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

type Context struct {
	World *DfWorld
	HfId  int
	Story bool
}

func NewContext(w *DfWorld, ref any) *Context {
	c := &Context{World: w}
	switch r := ref.(type) {
	case *WrittenContent:
		c.HfId = r.AuthorHfid
	default:
		fmt.Printf("unknown type for context %T\n", ref)
	}
	return c
}

func (c *Context) hf(id int) string {
	if c.HfId != -1 {
		if c.HfId == id {
			return c.hfShort(id)
		} else {
			return c.hfRelated(id, c.HfId)
		}
	}
	if x, ok := c.World.HistoricalFigures[id]; ok {
		return fmt.Sprintf(`the %s <a class="hf" href="/hf/%d">%s</a>`, x.Race+util.If(x.Deity, " deity", "")+util.If(x.Force, " force", ""), x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func (c *Context) hfShort(id int) string {
	if x, ok := c.World.HistoricalFigures[id]; ok {
		return fmt.Sprintf(`<a class="hf" href="/hf/%d">%s</a>`, x.Id(), util.Title(x.FirstName()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func (c *Context) hfRelated(id, to int) string {
	if c.HfId != -1 && to != c.HfId {
		if c.HfId == id {
			return c.hfShort(id)
		} else {
			return c.hfRelated(id, c.HfId)
		}
	}
	if x, ok := c.World.HistoricalFigures[id]; ok {
		if t, ok := c.World.HistoricalFigures[to]; ok {
			if y, ok := util.Find(t.HfLink, func(l *HfLink) bool { return l.Hfid == id }); ok {
				return fmt.Sprintf(`%s %s <a class="hf" href="/hf/%d">%s</a>`, t.PossesivePronoun(), y.LinkType, x.Id(), util.Title(x.Name()))
			}
		}
		return fmt.Sprintf(`the %s <a class="hf" href="/hf/%d">%s</a>`, x.Race+util.If(x.Deity, " deity", "")+util.If(x.Force, " force", ""), x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func (c *Context) hfList(ids []int) string {
	return andList(util.Map(ids, func(id int) string { return c.hf(id) }))
}

func (c *Context) hfListRelated(ids []int, to int) string {
	return andList(util.Map(ids, func(id int) string { return c.hfRelated(id, to) }))
}

func (c *Context) artifact(id int) string {
	if x, ok := c.World.Artifacts[id]; ok {
		return fmt.Sprintf(`<a class="artifact" href="/artifact/%d"><i class="%s fa-xs"></i> %s</a>`, x.Id(), x.Icon(), util.Title(x.Name()))
	}
	return "UNKNOWN ARTIFACT"
}

func (c *Context) entity(id int) string {
	if x, ok := c.World.Entities[id]; ok {
		return fmt.Sprintf(`<a class="entity" href="/entity/%d"><i class="%s fa-xs"></i> %s</a>`, x.Id(), x.Icon(), util.Title(x.Name()))
	}
	return "UNKNOWN ENTITY"
}

func (c *Context) entityList(ids []int) string {
	return andList(util.Map(ids, func(id int) string { return c.entity(id) }))
}

func (c *Context) position(entityId, positionId, hfId int) string {
	if e, ok := c.World.Entities[entityId]; ok {
		if h, ok := c.World.HistoricalFigures[hfId]; ok {
			return e.Position(positionId).GenderName(h)
		}
	}
	return "UNKNOWN POSITION"
}

func (c *Context) siteCiv(siteCivId, civId int) string {
	if siteCivId == civId {
		return c.entity(civId)
	}
	return util.If(siteCivId != -1, c.entity(siteCivId), "") + util.If(civId != -1 && siteCivId != -1, " of ", "") + util.If(civId != -1, c.entity(civId), "")
}

func (c *Context) siteStructure(siteId, structureId int, prefix string) string {
	if siteId == -1 {
		return ""
	}
	return " " + prefix + " " + util.If(structureId != -1, c.structure(siteId, structureId)+" in ", "") + c.site(siteId, "")
}

func (c *Context) site(id int, prefix string) string {
	if x, ok := c.World.Sites[id]; ok {
		return fmt.Sprintf(`%s <a class="site" href="/site/%d"><i class="%s fa-xs"></i> %s</a>`, prefix, x.Id(), x.Icon(), util.Title(x.Name()))
	}
	return "UNKNOWN SITE"
}

func (c *Context) structure(siteId, structureId int) string {
	if x, ok := c.World.Sites[siteId]; ok {
		if y, ok := x.Structures[structureId]; ok {
			return fmt.Sprintf(`<a class="structure" href="/site/%d/structure/%d"><i class="%s fa-xs"></i> %s</a>`, siteId, structureId, y.Icon(), util.Title(y.Name()))
		}
	}
	return "UNKNOWN STRUCTURE"
}

func (c *Context) property(siteId, propertyId int) string {
	if x, ok := c.World.Sites[siteId]; ok {
		if y, ok := x.SiteProperties[propertyId]; ok {
			if y.StructureId != -1 {
				return c.structure(siteId, y.StructureId)
			}
			return articled(y.Type_.String())
		}
	}
	return "UNKNOWN PROPERTY"
}

func (c *Context) region(id int) string {
	if x, ok := c.World.Regions[id]; ok {
		return fmt.Sprintf(`<a class="region" href="/region/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN REGION"
}

func (c *Context) location(siteId int, sitePrefix string, regionId int, regionPrefix string) string {
	if siteId != -1 {
		return c.site(siteId, sitePrefix)
	}
	if regionId != -1 {
		return regionPrefix + " " + c.region(regionId)
	}
	return ""
}

func (c *Context) place(structureId, siteId int, sitePrefix string, regionId int, regionPrefix string) string {
	if siteId != -1 {
		return c.siteStructure(siteId, structureId, sitePrefix)
	}
	if regionId != -1 {
		return regionPrefix + " " + c.region(regionId)
	}
	return ""
}

func (c *Context) mountain(id int) string {
	if x, ok := c.World.MountainPeaks[id]; ok {
		return fmt.Sprintf(`<a class="mountain" href="/site/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN MOUNTAIN"
}

func (c *Context) identity(id int) string {
	if x, ok := c.World.Identities[id]; ok {
		return fmt.Sprintf(`<a class="identity" href="/region/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN IDENTITY"
}

func (c *Context) fullIdentity(id int) string {
	if x, ok := c.World.Identities[id]; ok {
		return fmt.Sprintf(`&quot;the %s <a class="identity" href="/region/%d">%s</a> of %s&quot;`, x.Profession.String(), x.Id(), util.Title(x.Name()), c.entity(x.EntityId))
	}
	return "UNKNOWN IDENTITY"
}

func (c *Context) danceForm(id int) string {
	if x, ok := c.World.DanceForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/danceform/%d"><i class="fa-solid fa-shoe-prints fa-xs"></i> %s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN DANCE FORM"
}

func (c *Context) musicalForm(id int) string {
	if x, ok := c.World.MusicalForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/musicalform/%d"><i class="fa-solid fa-music fa-xs"></i> %s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN MUSICAL FORM"
}

func (c *Context) poeticForm(id int) string {
	if x, ok := c.World.PoeticForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/poeticform/%d"><i class="fa-solid fa-comment-dots fa-xs"></i> %s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN POETIC FORM"
}

func (c *Context) worldConstruction(id int) string {
	if x, ok := c.World.WorldConstructions[id]; ok {
		return fmt.Sprintf(`<a class="worldconstruction" href="/worldconstruction/%d"><i class="%s fa-xs"></i> %s</a>`, id, x.Icon(), util.Title(x.Name()))
	}
	return "UNKNOWN WORLD CONSTRUCTION"
}

func (c *Context) writtenContent(id int) string {
	if x, ok := c.World.WrittenContents[id]; ok {
		return fmt.Sprintf(`<a class="writtencontent" href="/writtencontent/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN WORLD CONSTRUCTION"
}

func (c *Context) collection(id int) string {
	if x, ok := c.World.HistoricalEventCollections[id]; ok {
		return x.Details.Html(x, c)
	}
	return "UNKNOWN EVENT COLLECTION"
}

func (c *Context) feature(x *Feature) string {
	switch x.Type_ {
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
			if e, ok := c.World.HistoricalEvents[x.Reference]; ok {
				return "a telling of the story of " + e.Details.Html(&Context{World: c.World, Story: true}) + " in " + Time(e.Year, e.Seconds72)
			}
		}
		return "a story recital"
	default:
		return strcase.ToDelimited(x.Type_.String(), ' ')
	}
}

func (c *Context) schedule(x *Schedule) string {
	switch x.Type_ {
	case ScheduleType_DancePerformance:
		return "a perfomance of " + c.danceForm(x.Reference)
	case ScheduleType_MusicalPerformance:
		return "a perfomance of " + c.musicalForm(x.Reference)
	case ScheduleType_PoetryRecital:
		return "a recital of " + c.poeticForm(x.Reference)
	case ScheduleType_Storytelling:
		if x.Reference != -1 {
			if e, ok := c.World.HistoricalEvents[x.Reference]; ok {
				return "the story of " + e.Details.Html(&Context{World: c.World, Story: true}) + " in " + Time(e.Year, e.Seconds72)
			}
		}
		return "a story recital"
	default:
		return strcase.ToDelimited(x.Type_.String(), ' ')
	}
}

func (c *Context) pronoun(id int) string {
	if x, ok := c.World.HistoricalFigures[id]; ok {
		return x.Pronoun()
	}
	return "he"
}

func (c *Context) posessivePronoun(id int) string {
	if x, ok := c.World.HistoricalFigures[id]; ok {
		return x.PossesivePronoun()
	}
	return "his"
}
