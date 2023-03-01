import { Component, OnInit } from '@angular/core';
import {AuthService } from "@auth0/auth0-angular";
@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.sass']
})
export class DashboardComponent {
  constructor(public auth: AuthService) {}
  code: string;

  ngOnInit(): void {

    this.auth.user$.subscribe(
      (profile)=> this.code = profile ? profile["get/sticker_code"] : ""
    )
  }
}
