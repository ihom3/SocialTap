import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatFormFieldModule} from '@angular/material/form-field'
import {MatInputModule} from "@angular/material/input"
import {MatButtonModule} from "@angular/material/button"
import {MatToolbarModule} from "@angular/material/toolbar"
import {MatIconModule} from "@angular/material/icon";
import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { NavComponent } from './nav/nav.component';
import { HomeComponent } from './home/home.component';
import { IdRouteComponent } from './id-route/id-route.component'
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";
import { UserNotFoundComponent } from './user-not-found/user-not-found.component';
import { SocialPageComponent } from './social-page/social-page.component';
import {MatCardModule} from "@angular/material/card";
import {MatDividerModule } from "@angular/material/divider"
import {MatListModule} from "@angular/material/list";
import { UpdateNameComponent } from './dashboard/update-name/update-name.component';
import { UpdateEmailComponent } from './dashboard/update-email/update-email.component';
import { UpdatePasswordComponent } from './dashboard/update-password/update-password.component'
import { DashboardComponent } from './dashboard/dashboard.component';
import { ToastrModule } from 'ngx-toastr';
import { RegisterCodesComponent } from './dashboard/register-codes/register-codes.component';
import { UpdateSocialsComponent } from './dashboard/update-socials/update-socials.component';
import { UpdateBioComponent } from './dashboard/update-bio/update-bio.component';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
@NgModule({
  declarations: [
    AppComponent,
    RegisterComponent,
    LoginComponent,
    NavComponent,
    HomeComponent,
    IdRouteComponent,
    UserNotFoundComponent,
    SocialPageComponent,
    DashboardComponent,
    UpdateNameComponent,
    UpdateEmailComponent,
    UpdatePasswordComponent,
    RegisterCodesComponent,
    UpdateSocialsComponent,
    UpdateBioComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgbModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatToolbarModule,
    MatIconModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatProgressSpinnerModule,
    MatCardModule,
    MatDividerModule,
    MatListModule,
    ToastrModule.forRoot({
      positionClass: 'toast-bottom-right',
      preventDuplicates: true
    }
    ),
    MatSlideToggleModule,
    FontAwesomeModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
