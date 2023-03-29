import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import {AuthService } from "@auth0/auth0-angular";
@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.sass']
})
export class DashboardComponent {
  constructor(public auth: AuthService, private router: Router) {}
  code: string;

  visitUpdateName = () =>
    this.router.navigate(['/dashboard/update-name']);
  visitUpdateBio = () => this.router.navigate(['dashboard/update-bio']);
  visitPage = () => this.router.navigate(['dashboard/page']);
  visitUpdateSocials = () => this.router.navigate(['dashboard/update-socials']);
  ngOnInit(): void {

    this.auth.user$.subscribe(
      (profile)=> this.code = profile ? profile["get/sticker_code"] : ""
    )
  }
}
