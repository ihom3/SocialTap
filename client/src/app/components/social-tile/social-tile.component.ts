import { Component, Input } from '@angular/core';
import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faFacebook, faInstagram, faLinkedin, faSnapchat, faTwitter } from '@fortawesome/free-brands-svg-icons';
import { faTimesCircle } from '@fortawesome/free-regular-svg-icons';
import { Social } from 'src/app/models/Social';

@Component({
  selector: 'social-tile',
  templateUrl: './social-tile.component.html',
  styleUrls: ['./social-tile.component.sass']
})
export class SocialTileComponent {
  @Input() social: Social;

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
      default:
        return faTimesCircle;
    }
  }
  redirectPage(site: string, url: string): void {
    var main: string = "https://";
    switch(site) {
      case "snapchat":
        main += "snapchat.com/";
        break;
      case "facebook":
        main += "facebook.com/";
        break;
      case "linkedin":
        main += "linkedin.com/";
        break;
      case "instagram":
        main += "instagram.com/";
        break;
      case "twitter":
        main += "twitter.com/";
        break;
      default:
        return;
    }
    main += url;
    window.location.href = main; 
  }
}
