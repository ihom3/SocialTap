import { Component, Input, OnInit } from '@angular/core';
import { FaIconComponent } from '@fortawesome/angular-fontawesome/public_api';
import { Social } from 'src/app/models/Social';
import { faFacebook, IconDefinition } from "@fortawesome/free-brands-svg-icons";
import { faTimesCircle } from '@fortawesome/free-regular-svg-icons';
@Component({
  selector: 'social-grid',
  templateUrl: './social-grid.component.html',
  styleUrls: ['./social-grid.component.sass']
})
export class SocialGridComponent implements OnInit {
  @Input() socialList: Social[];
  rowString: string = "";

  ngOnInit(): void {
    console.log(this.socialList);
    var count: number = 0;
    this.socialList.forEach(li => {if(li.active) count++})
    count = Math.ceil(count / 2);
    while(count != 0) {
      this.rowString += "1fr ";
      count--;
    }
  }

  
}
