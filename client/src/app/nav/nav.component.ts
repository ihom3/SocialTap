import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';
import {Location} from "@angular/common"
import { HttpClient } from '@angular/common/http';
import { UserServiceService } from '../user-service.service';
@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.sass']
})
export class NavComponent {
  constructor(public location: Location, private http: HttpClient, private router: Router, public userService: UserServiceService) {

  }

  logout(): void {
    this.userService.logout();
  }

}
