import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CivilizationsComponent } from './pages/civilizations/civilizations.component';
import { EntityComponent } from './pages/entity/entity.component';
import { HfComponent } from './pages/hf/hf.component';
import { EventListComponent } from './components/event-list/event-list.component';
import { InlineEntityComponent } from './components/inline-entity/inline-entity.component';
import { InlineHfComponent } from './components/inline-hf/inline-hf.component';

@NgModule({
  declarations: [
    AppComponent,
    CivilizationsComponent,
    EntityComponent,
    HfComponent,
    EventListComponent,
    InlineEntityComponent,
    InlineHfComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
