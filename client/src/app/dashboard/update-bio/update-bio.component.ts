import { Component } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { UserServiceService } from 'src/app/user-service.service';

@Component({
  selector: 'app-update-bio',
  templateUrl: './update-bio.component.html',
  styleUrls: ['./update-bio.component.sass']
})
export class UpdateBioComponent {
  form: FormGroup;
  constructor(public userService: UserServiceService, private formBuilder: FormBuilder) {
    this.form = this.formBuilder.group({
      bio: ""
    });
  }
  submit(): void {
    this.userService.updateBio(this.form);
  }
  disable(): boolean {
    const raw = this.form.getRawValue();
    if(raw.bio === "" || raw.bio === this.userService.userData!.bio) {
      return true;
    } else return false;
  }

}
