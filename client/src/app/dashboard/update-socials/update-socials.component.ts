import { Component } from '@angular/core';
import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faDiscord, faFacebook, faGithub, faInstagram, faLinkedin, faSnapchat, faTiktok, faTwitch, faTwitter, faYoutube } from '@fortawesome/free-brands-svg-icons';
import { faEnvelope, faFilePdf, faTimesCircle } from '@fortawesome/free-solid-svg-icons';
import { UserServiceService, Social } from 'src/app/user-service.service';

@Component({
  selector: 'app-update-socials',
  templateUrl: './update-socials.component.html',
  styleUrls: ['./update-socials.component.sass']
})

export class UpdateSocialsComponent {
  readonly oldSocials: Social[] = [];
  socials: Social[] = [];
  constructor(public userService: UserServiceService) {
    this.socials = this.userService.gotSocialCopy();
    this.oldSocials = this.userService.gotSocialCopy();
    
  }
  updateURL(name: string, event: any) {
    this.socials.forEach((social) => {
      if (social.name === name) {
        social.link = event.target.value;
      }
    })
  }
  toggleSocial(name: string) {
    var found: boolean = false;
    this.socials.forEach((social) => {
      if (social.name === name) {
        social.active = !social.active;
        found = true;
      }
    })
    if(!found) {
      this.socials.push({
        name: name,
        link: "",
        active: true
      })
    }
  }
  isChecked(name: string) {
    var c = this.socials.filter(val => val.name === name);
    if(c.length === 0) {
      this.socials.push({
        name,
        link: "",
        active: false
      })
      return false;
    } else return c[0].active;
  }
  SocialList: string[] = [
    "Facebook",
    "Twitter",
    "Instagram",
    "LinkedIn",
    "Snapchat",
    "TikTok",
    "YouTube",
    "Twitch",
    "Reddit",
    "GitHub",
    "Discord",
    "Spotify",
    "SoundCloud",
    "Apple Music",
  ]
    fetchIcon(name: string): IconDefinition {
      switch(name) {
        case "Facebook":
          return faFacebook;
        case "Snapchat":
          return faSnapchat;
        case "LinkedIn":
          return faLinkedin;
        case "Instagram":
          return faInstagram;
        case "twitter":
          return faTwitter;
        case "tiktok":
          return faTiktok;
        case "youtube":
          return faYoutube;
        case "discord":
          return faDiscord;
        case "twitch":
          return faTwitch;
        case "github":
          return faGithub;
        case "resume":
          return faFilePdf;
        case "email":
          return faEnvelope;
        default:
          return faTimesCircle;
      }
  }
  buttonDisable(): boolean {

    if(JSON.stringify(this.oldSocials) === JSON.stringify(this.socials.filter(val => val.active || val.link != ""))) {
      return true;
    } else return false;
  

  }
  submit(): void {
    this.userService.updateSocials(this.socials.filter(val => val.active || val.link != ""));
  }
  getURL(name: string): string {
    var f = this.socials.filter(val => val.name === name);
    if(f.length === 0) {
      return "";
    }
    else return f[0].link;
  }
}
 