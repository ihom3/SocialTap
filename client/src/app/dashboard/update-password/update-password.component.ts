import { Component } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { UserServiceService } from 'src/app/user-service.service';

@Component({
  selector: 'app-update-password',
  templateUrl: './update-password.component.html',
  styleUrls: ['./update-password.component.sass']
})
export class UpdatePasswordComponent {
  form: FormGroup;
  constructor(public userService: UserServiceService, private formBuilder: FormBuilder) {
    this.form = this.formBuilder.group({
      oldPassword: "",
      newPassword: "",
      newPasswordConfirmation: ""
    })
  }
  submit(): void {
    this.form.removeControl("newPasswordConfirmation");
    this.userService.updatePassword(this.form);
  }
  disable(): boolean {
    const raw = this.form.getRawValue();
    if(raw.newPassword != raw.newPasswordConfirmation || raw.oldPassword == "" || raw.oldPassword == raw.newPassword || raw.newPassword == "") {
      return true;
    } else return false;
  }
}
