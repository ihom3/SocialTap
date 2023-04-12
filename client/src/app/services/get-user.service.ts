import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment.development';
import { UserRequest } from '../models/UserReq';
import {Observable } from "rxjs"
import { Social } from '../models/Social';
@Injectable({
  providedIn: 'root'
})
export class GetUserService {
  private readonly url = environment.ApiURL;
  constructor(private httpClient: HttpClient) { }


  getUserByCode = (code: string) => {
    return this.httpClient.get(this.url + code);
  }
  getDashboardProps = () => {
    return this.httpClient.get(this.url + "dashboard");
  }
  updateName = (name: string) => {
    return this.httpClient.put(this.url + "update-name", { name });
  }
  updateBio = (bio: string) => {
    return this.httpClient.put(this.url + "update-bio", { bio });
  }
  updateProfilePic = (profilePic: string) => {
    return this.httpClient.put(this.url + "update-profile-pic", { profilePic });
  }
  updateSocials = (socials: Social[]) => {
    return this.httpClient.put(this.url + "update-socials", { socials });
  }
}
