import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EntitiesResolver, EntityResolver } from './entity.service';
import { CivilizationsComponent } from './pages/civilizations/civilizations.component';
import { EntityComponent } from './pages/entity/entity.component';
import { HfComponent } from './pages/hf/hf.component';

const routes: Routes = [
  { path: '', component: CivilizationsComponent, resolve: { civilizations: EntitiesResolver } },
  { path: 'entity/:id', component: EntityComponent, resolve: { entity: EntityResolver }, data: { resource: "entity" } },
  { path: 'hf/:id', component: HfComponent, resolve: { data: EntityResolver }, data: { resource: "hf" } },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
