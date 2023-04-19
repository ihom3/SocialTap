import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { Location } from '@angular/common';
import { UserServiceService } from '../user-service.service';
@Component({
  selector: 'app-dashboard-routing',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.sass'],
  host: {'class': 'container'}
})
export class DashboardComponent {
  constructor(private http: HttpClient, private router: Router, public location: Location, public userService: UserServiceService) { 
  }
  updatePicture(): void {
    let input = document.createElement('input');
  input.type = 'file';
  input.accept="image/jpeg";
  input.click();
  const formData = new FormData();
  input.onchange = (event: any) => {
    const file: File = event.target.files[0];
    formData.append("profile_picture" , file);
    this.userService.updateProfilePicture(formData);
  }
  }
  
}
