import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.sass']
})
export class RegisterComponent {
  id: string;
  constructor(private router: Router) {
    const state = this.router.getCurrentNavigation()?.extras.state;
    this.id = state ? state['id'] : "";
    if(this.id == "") {
      this.router.navigate(["error"]);
    }
  }
}
