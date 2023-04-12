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
  header: Header = {
    bio: "",
    imageUrl: "",
    name: ""
  };
  socials: Social[] = [];
  id: string;
  user: any;
  constructor(private router: Router) {
    const state = this.router.getCurrentNavigation()?.extras.state;
    this.id = state ? state['id'] : "";
    this.user = state ? state['user'] : "";
    if(this.id == "") {
    }
  }
  ngOnInit(): void {
    //fetch user data from DB
    console.log(this.user.user.social_list);
    Object.values(this.user.user.social_list).forEach((key: any, index) => {
      if(key.status) {
        var social: Social = { name: key.name, url: key.url, active: key.status};
        this.socials.push(social);
      }
    });
    
    this.header.bio = this.user.user.bio_text;
    this.header.imageUrl = this.user.user.profile_picture;
    this.header.name = this.user.user.first_name + " " + this.user.user.last_name;
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
