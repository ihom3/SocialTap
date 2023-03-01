import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Header } from 'src/app/models/Header';
import { Social } from 'src/app/models/Social';
import { HEADER, SOCIALS } from 'src/app/models/mockData';


@Component({
  selector: 'social-page',
  templateUrl: './social-page.component.html',
  styleUrls: ['./social-page.component.sass']
})
export class SocialPageComponent implements OnInit {
  header: Header;
  socials: Social[];
  id: string;
  constructor(private router: Router) {
    const state = this.router.getCurrentNavigation()?.extras.state;
    this.id = state ? state['id'] : "";
    if(this.id == "") {
      this.router.navigate(["error"]);
    }
  }
  ngOnInit(): void {
    //fetch user data from DB
    this.socials = SOCIALS;
    this.header = HEADER;
      //set header and socials
  }



}
