import { Component } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { UserServiceService } from 'src/app/user-service.service';

@Component({
  selector: 'app-register-codes',
  templateUrl: './register-codes.component.html',
  styleUrls: ['./register-codes.component.sass']
})
export class RegisterCodesComponent {
  form: FormGroup;
  constructor(public userService: UserServiceService, private router: Router, private formBuilder: FormBuilder) {
    this.form = this.formBuilder.group({
      code: ""
    })
   }
  ngOnInit(): void {
    if(this.userService.userData?.role != "admin") {
      this.router.navigate(["/dashboard"]);
    }
  }
  submit(): void {
    this.userService.registerCode(this.form);
  }
  disable(): boolean {
    const raw = this.form.getRawValue();
    if(raw.code == "" || raw.code.length != 4) {
      return true;
    } else return false;
  }
}
