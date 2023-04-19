import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { serverURL } from '../user-service.service';

@Component({
  selector: 'app-id-route',
  templateUrl: './id-route.component.html',
  styleUrls: ['./id-route.component.sass'],
  host: {'class': 'w-full h-full grid'}
})
export class IdRouteComponent {
  id: string = "";
  unregistered: boolean = false;
  found: boolean = false;
  userData = {};
  constructor( private http: HttpClient, private route: ActivatedRoute, private router: Router) {route.params.subscribe(params => {
    this.id = params['id'];
  })
  }
  ngOnInit(): void {
    this.http.get(serverURL + this.id).subscribe((res: any) => {
      if(res.message === "User Not Found") {
        this.router.navigate(["/user-not-found"]);
      } else if(res.message === "User Not Registered") {
        this.unregistered = true;
      } else {
        this.userData = res;
        this.found = true;
      }
    })
  }
}
