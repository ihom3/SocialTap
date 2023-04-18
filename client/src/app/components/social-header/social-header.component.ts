import { Component, Input } from '@angular/core';
import { Header } from 'src/app/models/Header';

@Component({
  selector: 'social-header',
  templateUrl: './social-header.component.html',
  styleUrls: ['./social-header.component.sass']
})
export class SocialHeaderComponent {
  @Input() headerData: Header;


}
