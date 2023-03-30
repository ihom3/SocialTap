import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Header } from 'src/app/models/Header';
import { HEADER, SOCIALS } from 'src/app/models/mockData';
import { Social } from 'src/app/models/Social';
import { UserRequest, SocialReq } from 'src/app/models/UserReq';

@Component({
  selector: 'social-page',
  templateUrl: './social-page.component.html',
  styleUrls: ['./social-page.component.sass']
})
export class SocialPageComponent implements OnInit {
  header: Header;
  socials: Social[] = [];
  id: string;
  user: UserRequest;
  constructor(private router: Router) {
    const state = this.router.getCurrentNavigation()?.extras.state;
    this.id = state ? state['id'] : "";
    this.user = state ? state['user'] : "";
    if(this.id == "") {
    }
  }
  ngOnInit(): void {
    //fetch user data from DB
    this.socials = SOCIALS;
    this.header = HEADER;
    // Object.entries(this.user.user.social_list).forEach(([key, value]: [string, SocialReq], index) => {
    //   if(value.status) {
    //   var social: Social = { name: key, url: value.url, active: value.status};
    //   this.socials.push(social);
    // }
    // })
    // this.header = { bio: this.user.user.bio_text, imageUrl: this.user.user.profile_picture, name: this.user.user.first_name + " " + this.user.user.last_name};
      //set header and socials
  }



}
