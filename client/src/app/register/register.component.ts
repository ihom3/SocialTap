import { HttpClient } from '@angular/common/http';
import { Component, Input } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { UserServiceService } from '../user-service.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.sass'],
  host: {'class': 'container'}
})
export class RegisterComponent {
  @Input() code: string = "";
  loading: boolean;
  form: FormGroup;
  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router, private toastr: ToastrService, public userServie: UserServiceService) {
    this.form = this.formBuilder.group({
      first_name: '',
      last_name: '',
      email: '',
      password: '',
      code: '',
      passwordConfirmation: ''
    });
    this.loading = this.userServie.loading;
  }
  disableButton(): boolean {
    const raw = this.form.getRawValue();
    if(raw.email === "" || raw.firstName === "" || raw.lastName === "" || raw.password === "" || raw.passwordConfirmation === "" || raw.password != raw.passwordConfirmation) {
      return true;
    } else return false;
  }
  submit(): void {
    this.userServie.register(this.form);
}
 ngOnInit(): void {
  if(this.userServie.isLoggedIn) {
    this.router.navigate(["/dashboard"]);
  }
  this.form.patchValue({ code: this.code});
 }
}
