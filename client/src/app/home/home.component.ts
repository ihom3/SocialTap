import { Component } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.sass'],
  host: {'class': 'container'}
})
export class HomeComponent {
  gitHub = (): void  => {
    window.location.href = 'https://github.com/ihom3/SocialTap';
  }
}
