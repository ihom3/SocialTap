import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { IdRouteComponent } from './id-route/id-route.component';
import { UserNotFoundComponent } from './user-not-found/user-not-found.component';
import { UpdateNameComponent } from './dashboard/update-name/update-name.component';
import { UpdateEmailComponent } from './dashboard/update-email/update-email.component';
import { UpdatePasswordComponent } from './dashboard/update-password/update-password.component';
import {DashboardComponent} from "./dashboard/dashboard.component";
import { RegisterCodesComponent } from './dashboard/register-codes/register-codes.component';
import { AuthGuard } from './auth/auth.guard';
import { UpdateBioComponent } from './dashboard/update-bio/update-bio.component';
import { UpdateSocialsComponent } from './dashboard/update-socials/update-socials.component';
const routes: Routes = [{
  path: "", component: HomeComponent,
  
}, {
  path: "login", component: LoginComponent
}, {
  path: "user-not-found", component: UserNotFoundComponent
},{
  path: "dashboard", component: DashboardComponent, canActivate:[AuthGuard], children: [
    {
      path: "update-name", component: UpdateNameComponent
    }, {
      path: "update-email", component: UpdateEmailComponent
    },
    {
      path: "update-password", component: UpdatePasswordComponent
    }, {
      path: "update-bio", component: UpdateBioComponent
    },
    {
      path: "update-socials", component: UpdateSocialsComponent
    },
    {
      path: "register-codes", component: RegisterCodesComponent
    }
  ]
},
 {
  path: ":id", component: IdRouteComponent
},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
