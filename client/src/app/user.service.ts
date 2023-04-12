import { Injectable } from '@angular/core';
import { GetUserService } from './services/get-user.service';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private getUser: GetUserService) { }

  dashboardRoute = () => {
    return this.getUser.getDashboardProps();
  }
  
}
