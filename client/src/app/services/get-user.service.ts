import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment.development';
import { UserRequest } from '../models/UserReq';
import {Observable } from "rxjs"

@Injectable({
  providedIn: 'root'
})
export class GetUserService {
  private url = environment.ApiURL + "/users/";
  constructor(private httpClient: HttpClient) { }


  getUserByCode = (code: string): Observable<UserRequest> => {
    return this.httpClient.get<UserRequest>(this.url + code);
  }
}
