import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SocialTileComponent } from './components/social-tile/social-tile.component';
import { SocialGridComponent } from './components/social-grid/social-grid.component';
import { SocialHeaderComponent } from './components/social-header/social-header.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { SocialPageComponent } from './components/social-page/social-page.component';
import { RegisterComponent } from './components/register/register.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomePageComponent } from './components/home-page/home-page.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { IdDiscoveryComponent } from './components/id-discovery/id-discovery.component';
import { ErrorPageComponent } from './components/error-page/error-page.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { PageHeaderComponent } from './components/page-header/page-header.component';
import { ButtonComponent } from './components/button/button.component';

@NgModule({
  declarations: [
    AppComponent,
    SocialTileComponent,
    SocialGridComponent,
    SocialHeaderComponent,
    DashboardComponent,
    SocialPageComponent,
    RegisterComponent,
    HomePageComponent,
    PageNotFoundComponent,
    IdDiscoveryComponent,
    ErrorPageComponent,
    PageHeaderComponent,
    ButtonComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FontAwesomeModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
