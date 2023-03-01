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
const routes: Routes = [{ path: '', component: HomePageComponent},
{ path: "dashboard", component: DashboardComponent, canActivate: [AuthGuard]},
{path: "register", component: RegisterComponent},
{path: "error", component: PageNotFoundComponent},
{path: "user", component: SocialPageComponent},
{path: "activate-code", component: ActivateCodeComponent},
{path: ":id", component: IdDiscoveryComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
  constructor() {}
  
 
 
  

 }
