import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from '@auth0/auth0-angular';
import { UserRequest } from 'src/app/models/UserReq';
import { GetUserService } from 'src/app/services/get-user.service';

@Component({
  selector: 'app-id-discovery',
  templateUrl: './id-discovery.component.html',
  styleUrls: ['./id-discovery.component.sass']
})
export class IdDiscoveryComponent implements OnInit {

  constructor(private route: ActivatedRoute, private router: Router, public auth: AuthService, private service: GetUserService) {}
  
  id: string;
  user: UserRequest;

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.id = params['id'];
    })
    this.service.getUserByCode(this.id).subscribe(res => { 
      this.user = res;
      if(this.id == "unregistered-test") {
        this.auth.loginWithRedirect({authorizationParams: {screen_hint: "signup"}, appState: {id: this.id}})
        // this.router.navigate(['register'], { state: {id: this.id}})
      } else if(this.user.first_name != "" || this.id == "registered-test") {
         //redirect to social-page with userData
        this.router.navigate(['user'], { state: { id: this.id, user: this.user}});
      } else {
        this.router.navigate(['error']);
      }
    })

    //query backend for id in unregistered and registered
    
    
    //redirect to page not found


}
}
