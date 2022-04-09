import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CivilizationsComponent } from './pages/civilizations/civilizations.component';
import { EntityComponent } from './pages/entity/entity.component';

@NgModule({
  declarations: [
    AppComponent,
    CivilizationsComponent,
    EntityComponent
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
