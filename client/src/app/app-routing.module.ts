import { NgModule, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterModule, Routes } from '@angular/router';
import { Subscription } from 'rxjs';
import { ActivateCodeComponent } from './components/activate-code/activate-code.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { HomePageComponent } from './components/home-page/home-page.component';
import { IdDiscoveryComponent } from './components/id-discovery/id-discovery.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { RegisterComponent } from './components/register/register.component';
import { SocialPageComponent } from './components/social-page/social-page.component';
import { AuthGuard} from "@auth0/auth0-angular";
import { UpdateNameComponent } from './component/dashboard/update-name/update-name.component';
import { UpdateBioComponent } from './component/dashboard/update-bio/update-bio.component';
import { UpdateSocialsComponent } from './component/dashboard/update-socials/update-socials.component';
const routes: Routes = [{ path: '', component: HomePageComponent},
{ path: "dashboard", component: DashboardComponent, canActivate: [AuthGuard]},
{path: "register", component: RegisterComponent},
{path: "error", component: PageNotFoundComponent},
{path: "user", component: SocialPageComponent},
{path: "activate-code", component: ActivateCodeComponent},
{path: ":id", component: IdDiscoveryComponent},
{path: "dashboard/update-name", component: UpdateNameComponent, canActivate: [AuthGuard]},
{path: "dashboard/update-bio", component: UpdateBioComponent, canActivate: [AuthGuard]},
{path: "dashboard/update-socials", component: UpdateSocialsComponent, canActivate: [AuthGuard]},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
  constructor() {}

 }
