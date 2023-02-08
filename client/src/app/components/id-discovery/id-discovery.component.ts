import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-id-discovery',
  templateUrl: './id-discovery.component.html',
  styleUrls: ['./id-discovery.component.sass']
})
export class IdDiscoveryComponent implements OnInit {

  constructor(private route: ActivatedRoute, private router: Router) {}
  
  id: string;

  ngOnInit(): void {
    this.route.params.subscribe(params => {
      this.id = params['id'];
    })

    //query backend for id in unregistered and registered
    if(this.id == "unregistered-test") {
      this.router.navigate(['register'], { state: {id: this.id}})
    } else if(this.id == "registered-test") {
       //redirect to social-page with userData
      this.router.navigate(['user'], { state: { id: this.id}});
    } else {
      this.router.navigate(['error']);
    }
    
    //redirect to page not found


}
}
