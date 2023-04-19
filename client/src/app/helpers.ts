import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import { faDiscord, faFacebook, faGithub, faInstagram, faLinkedin, faReddit, faSnapchat, faSoundcloud, faSpotify, faTiktok, faTwitch, faTwitter, faYoutube } from '@fortawesome/free-brands-svg-icons';
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
    "GitHub",
    "Spotify",
    "SoundCloud",
    "Discord",
    "Email",
    "Reddit"
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
      case "TikTok":
        return faTiktok;
      case "YouTube":
        return faYoutube;
      case "Discord":
        return faDiscord;
      case "Twitch":
        return faTwitch;
      case "GitHub":
        return faGithub;
      case "Email":
        return faEnvelope;
      case "Spotify":
        return faSpotify;
      case "SoundCloud":
        return faSoundcloud;
      case "Discord":
        return faDiscord;
      case "Reddit":
        return faReddit;
      default:
        return faTimesCircle;
    }
}



export function fetchURL(social: string, username: string): string {
  var baseURL = "https://";
  switch(social) {
    case "Facebook":
      baseURL += "facebook.com/"
      break;
    case "Twitter":
      baseURL += "twitter.com/"
      break;
    case "Instagram":
      baseURL += "instagram.com/"
      break;
    case "LinkedIn":
      baseURL += "linkedin.com/in/"
      break;
    case "Snapchat":
      baseURL += "snapchat.com/add/"
      break;
    case "TikTok":
      baseURL += "tiktok.com/@"
      break;
    case "YouTube":
      baseURL += "youtube.com/@"
      break;
    case "Twitch":
      baseURL += "twitch.tv/"
      break;
    case "GitHub":
      baseURL += "github.com/"
      break;
    case "Spotify":
      baseURL += "open.spotify.com/user/"
      break;
    case "SoundCloud":
      baseURL += "soundcloud.com/"
      break;
    case "Email":
      baseURL = "mailto:";
      break;
    case "Reddit":
      baseURL += "reddit.com/user/";
      break;
    default:
      break;

  }
  baseURL += username;
  return baseURL;

}