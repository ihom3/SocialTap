import { Component } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { UserServiceService } from 'src/app/user-service.service';

@Component({
  selector: 'app-update-email',
  templateUrl: './update-email.component.html',
  styleUrls: ['./update-email.component.sass']
})
export class UpdateEmailComponent {
  form: FormGroup;
  constructor(public userService: UserServiceService, private formBuilder: FormBuilder) {
    this.form = this.formBuilder.group({
      email: "",
      emailConfirmation: ""
    })
  }
  submit(): void {
    this.form.removeControl("emailConfirmation");
    this.userService.updateEmail(this.form);
  }
  disable(): boolean {
    const raw = this.form.getRawValue();
    if(raw.email != raw.emailConfirmation || raw.email == "" || raw.emailConfirmation == this.userService.userData?.email) {
      return true;
    } else return false;
  }
}
