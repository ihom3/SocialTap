import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import { faDiscord, faFacebook, faGithub, faInstagram, faLinkedin, faSnapchat, faTiktok, faTwitch, faTwitter, faYoutube } from '@fortawesome/free-brands-svg-icons';
import { faEnvelope, faFilePdf, faTimesCircle } from "@fortawesome/free-solid-svg-icons";

export const SocialList: string[] = [
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

export function fetchIcon(name: string): IconDefinition {
    switch(name) {
      case "Facebook":
        return faFacebook;
      case "Snapchat":
        return faSnapchat;
      case "LinkedIn":
        return faLinkedin;
      case "Instagram":
        return faInstagram;
      case "Twitter":
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

export function fetchURL(social: string, username: string): string {
  let tempUrl: string = '';
  switch(social) {
    case "Facebook":
      tempUrl = "https://www.facebook.com/";
      break;
    default:
      break;
  }
  tempUrl += username;
  return tempUrl;

}