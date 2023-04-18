import { Component, Inject,OnInit } from '@angular/core';
import {Router} from "@angular/router";
import { AuthService } from '@auth0/auth0-angular';
import { DOCUMENT } from "@angular/common";
@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.sass']
})
export class HomePageComponent {
  constructor(private router: Router, public auth: AuthService, @Inject(DOCUMENT) public document: Document) {}

  visitExample = () =>
    this.router.navigate(['/registered-test']);

  async ngOnInit() {
    console.log(this.auth.isAuthenticated$)
  }
}
