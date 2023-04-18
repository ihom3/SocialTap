import { Component } from '@angular/core';
import { faDiscord, faFacebook, faGithub, faInstagram, faLinkedin, faSnapchat, faTiktok, faTwitch, faTwitter, faYoutube, IconDefinition } from '@fortawesome/free-brands-svg-icons';
import { faEnvelope, faFilePdf, faTimesCircle } from '@fortawesome/free-regular-svg-icons';
import { SOCIALS } from 'src/app/models/mockData';
import { Social } from 'src/app/models/Social';

@Component({
  selector: 'app-update-socials',
  templateUrl: './update-socials.component.html',
  styleUrls: ['./update-socials.component.sass']
})
export class UpdateSocialsComponent {
  socials: Social[] = SOCIALS;
  fetchIcon(name: string): IconDefinition {
    switch(name) {
      case "facebook":
        return faFacebook;
      case "snapchat":
        return faSnapchat;
      case "linkedin":
        return faLinkedin;
      case "instagram":
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
}
