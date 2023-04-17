import { Component, Input } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { UserServiceService } from 'src/app/user-service.service';

@Component({
  selector: 'app-update-name',
  templateUrl: './update-name.component.html',
  styleUrls: ['./update-name.component.sass']
})
export class UpdateNameComponent {
  form: FormGroup;
  constructor(private formBuilder: FormBuilder, public userService: UserServiceService) {
    this.form = this.formBuilder.group({
      first_name: "",
      last_name: ""
         });
  }
  disable(): boolean {
    const raw = this.form.getRawValue();
    if((raw.first_name === "" && raw.last_name === "") || (raw.first_name === this.userService.userData?.first_name && raw.last_name === this.userService.userData?.last_name)) {
      return true;
    } else return false;
  }
}
