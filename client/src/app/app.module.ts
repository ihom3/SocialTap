import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { environment } from 'src/environments/environment.development';
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
import { MatToolbarModule } from "@angular/material/toolbar";
import { MatIconModule } from "@angular/material/icon";
import { MatButtonModule} from "@angular/material/button";
import { ActivateCodeComponent } from './components/activate-code/activate-code.component';
import { MatCardModule } from "@angular/material/card";
import {MatFormFieldModule} from "@angular/material/form-field"
import {MatInputModule} from "@angular/material/input";
import { CardComponent } from './components/card/card.component';
import { FormsModule } from '@angular/forms';
import { AuthModule } from '@auth0/auth0-angular';
import { AuthButtonComponent } from './components/auth-button/auth-button.component';
import { HttpClientModule } from '@angular/common/http';
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
    ActivateCodeComponent,
    CardComponent,
    AuthButtonComponent
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FontAwesomeModule,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    AuthModule.forRoot({
      domain: environment.Auth0Domain,
      clientId: environment.Auth0ClientID,
      authorizationParams: {
        redirect_uri: window.location.origin
      }
    })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
