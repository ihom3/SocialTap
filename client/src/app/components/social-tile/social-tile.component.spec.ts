import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SocialTileComponent } from './social-tile.component';

import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faFacebook, faInstagram, faLinkedin, faSnapchat, faTwitter, faTiktok, faYoutube, faDiscord, faTwitch, faGithub} from '@fortawesome/free-brands-svg-icons';
import { faFilePdf, faTimesCircle, faEnvelope } from '@fortawesome/free-regular-svg-icons';

const { fetchIcon } = require('./fetchIcon'); // import the fetchIcon function

describe('SocialTileComponent', () => {
  let component: SocialTileComponent;
  let fixture: ComponentFixture<SocialTileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SocialTileComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SocialTileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

// describe('fetchIcon', () => {
//   it('returns the Facebook icon for name "facebook"', () => {
//     expect(fetchIcon('facebook')).toEqual(faFacebook);
//   });

//   it('returns the Snapchat icon for name "snapchat"', () => {
//     expect(fetchIcon('snapchat')).toEqual(faSnapchat);
//   });

//   it('returns the LinkedIn icon for name "linkedin"', () => {
//     expect(fetchIcon('linkedin')).toEqual(faLinkedin);
//   });

//   it('returns the Instagram icon for name "instagram"', () => {
//     expect(fetchIcon('instagram')).toEqual(faInstagram);
//   });

//   it('returns the Twitter icon for name "twitter"', () => {
//     expect(fetchIcon('twitter')).toEqual(faTwitter);
//   });

//   it('returns the TikTok icon for name "tiktok"', () => {
//     expect(fetchIcon('tiktok')).toEqual(faTiktok);
//   });

//   it('returns the YouTube icon for name "youtube"', () => {
//     expect(fetchIcon('youtube')).toEqual(faYoutube);
//   });

//   it('returns the Discord icon for name "discord"', () => {
//     expect(fetchIcon('discord')).toEqual(faDiscord);
//   });

//   it('returns the Twitch icon for name "twitch"', () => {
//     expect(fetchIcon('twitch')).toEqual(faTwitch);
//   });

//   it('returns the GitHub icon for name "github"', () => {
//     expect(fetchIcon('github')).toEqual(faGithub);
//   });

//   it('returns the PDF file icon for name "resume"', () => {
//     expect(fetchIcon('resume')).toEqual(faFilePdf);
//   });

//   it('returns the envelope icon for name "email"', () => {
//     expect(fetchIcon('email')).toEqual(faEnvelope);
//   });

//   it('returns the times circle icon for an unknown name', () => {
//     expect(fetchIcon('unknown')).toEqual(faTimesCircle);
//   });
// });
