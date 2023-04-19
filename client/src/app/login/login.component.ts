import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { UserServiceService, serverURL } from '../user-service.service';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.sass'],
  host: {'class': 'container-sm'}
})
export class LoginComponent {
  loading: boolean = true;
  form: FormGroup;
  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router, public userService: UserServiceService, private toast: ToastrService) {
    this.form = this.formBuilder.group({
      email: "",
      password: ""
    });
  }
  ngOnInit(): void {
    this.http.get(serverURL + "is-logged-in", { withCredentials: true}).subscribe({
      next: (v: any) => {
        if(v.status === true) {
          this.router.navigate(["/dashboard"]);
        } else {
          this.loading = false;
        }
      },
      error: (err: HttpErrorResponse) => {
        this.toast.error(err.message, "Error");
      }
    })
  }
}
