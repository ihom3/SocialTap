import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-update-name',
  templateUrl: './update-name.component.html',
  styleUrls: ['./update-name.component.sass']
})
export class UpdateNameComponent {
  constructor(private router: Router) {}
  Name: string = "";
  handleNameChange(name: string) {
    this.Name = name;

  }

}
