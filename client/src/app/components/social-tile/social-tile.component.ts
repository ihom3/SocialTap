import { Component, Input } from '@angular/core';
import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faFacebook, faInstagram, faLinkedin, faSnapchat, faTwitter, faTiktok, faYoutube, faDiscord, faTwitch, faGithub} from '@fortawesome/free-brands-svg-icons';
import { faFilePdf, faTimesCircle, faEnvelope } from '@fortawesome/free-regular-svg-icons';
import { faD } from '@fortawesome/free-solid-svg-icons';
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
      case "tiktok":
        main += "tiktok.com/";
        break;
      case "youtube":
        main += "youtube.com/";
        break;
      case "discord":
        main += "discord.com/";
        break;
      case "twitch":
        main += "twitch.com/";
        break;
      case "github":
        main += "github.com/";
        break;
      case "resume":
        //main = "../../../../assets/resume.pdf"
        main = "/assets/resume.pdf";
        break;
      case "email":
        main += "mailto:";
        break;
      default:
        return;
    }
    main += url;
    window.location.href = main; 
  }
}


