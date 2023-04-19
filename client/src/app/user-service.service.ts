import { HttpClient, HttpErrorResponse, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { DomSanitizer } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

export interface Social {
  name: string;
  link: string;
  active: boolean;
}

interface User {
  first_name: string;
  last_name: string;
  email: string;
  id: string;
  bio: string;
  code: string;
  pictureURL: string;
  active: boolean;
  role: string;
  socials: Social[];
}
export const serverURL = "http://localhost:8000/api/";

@Injectable({
  providedIn: 'root'
})
export class UserServiceService {
  profileImg: any;
  userData: User | undefined;
  isLoggedIn: boolean = false;
  loading: boolean = true;
  reader: FileReader = new FileReader();
  constructor(private http: HttpClient, private router: Router, private toastr: ToastrService, private sanitizer : DomSanitizer) { 
    this.fetchUser();
  }
  fetchUser(): void {
    this.http.get(serverURL + "get-user", { withCredentials: true}).subscribe({
      next: (v) => {
        this.userData = v as User;
        this.isLoggedIn = true;
        this.fetchProfilePicture();
      },
      error: (err) => {
        this.userData = undefined;
        this.isLoggedIn = false;
        this.loading = false;
      }
    })
  }
  fetchProfilePicture(): void {
    this.loading = true;
    this.profileImg = undefined;
    let httpHeaders = new HttpHeaders()
         .set('Accept', "image/jpeg,*/*");
    this.http.get(serverURL + "profile-picture/" + this.userData!.id, {
      headers: httpHeaders,
      responseType: 'blob'
    }).subscribe({
      next: (a) => {
        this.reader.readAsDataURL(a);
        this.reader.onload = (_) => {
          this.profileImg = this.reader.result;
          this.loading = false;
        }
        
      },
      error: (err) => {
        this.toastr.error("Error fetching profile picture", "Error");
        this.loading = false;
      }
    })
  }
  login(form: FormGroup): void {
    this.http.post("http://localhost:8000/api/login", form.getRawValue(), {
      withCredentials: true, responseType: "json"
    }).subscribe({ 
      next: (v: any) => {
        if(v.message === "success") {
          this.fetchUser();
          this.toastr.success("Login Successful", "Success");
          this.isLoggedIn = true;
          this.router.navigate(["/dashboard"]);
          
        }
      },
      error: (err: any) => {
        this.toastr.error(err.error.message, "Error");
      }
    })
  }
  logout(): void {
    if(this.isLoggedIn) {
      this.loading = true;
    this.http.post("http://localhost:8000/api/logout", {}, { withCredentials: true }).subscribe((res: any) => {
      if(res.message === "success") {
        this.isLoggedIn = false;
        this.router.navigate(["/"]);
        this.loading = false;
      }
    })
  }
  }
  register(form: FormGroup): void {
    this.loading = true;
    this.http.post(serverURL + "register", form.getRawValue()).subscribe({
      next: (v: any) => {
          if(v.message === "success") {
            this.loading = false;
            this.router.navigate(["/login"]);
            this.toastr.success("Registered Successfully", "Success");
          }
      },
      error: (err: HttpErrorResponse) => {
        this.loading = false;
        this.toastr.error(err.error.Message, "Error");
      }
    })
  }
  updateName(form: FormGroup): void {
    this.loading = true;
    this.http.post(serverURL + "update-name", form.getRawValue(), { withCredentials: true }).subscribe({
    next: (v: any) => {
      if(v.message === "success") {
        if(this.userData != undefined) {
        this.userData.first_name = form.getRawValue().first_name;
        this.userData.last_name = form.getRawValue().last_name;
      }
        this.router.navigate(["/dashboard"]);
        this.loading = false;
        this.toastr.success("Name updated successfully", "Success");
      }
    },
    error: (err) => {
      console.log(err);
    }
    })
  }
  updateEmail(form: FormGroup): void {
    this.loading = true;
    this.http.post(serverURL + "update-email", form.getRawValue(), { withCredentials: true }).subscribe({
      next: (v: any) => {
        if(v.message === "success") {
          if(this.userData != undefined) {
            this.userData.email = form.getRawValue().email;
          }
          this.router.navigate(["/dashboard"]);
          this.loading = false;
          this.toastr.success("Email updated successfully", "Success");
        }
      },
      error: (err: HttpErrorResponse) => {
        this.loading = false;
        this.toastr.error(err.error.message, "Error");
        console.log(err);
      }
    })
  }
  updatePassword(form: FormGroup): void {
    this.loading = true;
    this.http.post(serverURL + "update-password", form.getRawValue(), { withCredentials: true }).subscribe({
      next: (v: any) => {
        if(v.message === "success") {
          this.router.navigate(["/dashboard"]);
          this.loading = false;
          this.toastr.success("Password updated successfully", "Success");
        }
      },
      error: (err) => {
        this.loading = false;
        this.toastr.error(err.error.message, "Error");
      }
    })
  }
  registerCode(form: FormGroup): void {
    this.loading = true;
    this.http.post(serverURL + "register-code", form.getRawValue(), { withCredentials: true }).subscribe({
      next: (v: any) => {
        if(v.message === "success") {
          this.loading = false;
          this.toastr.success("Code Registered Successfully", "Success");
        }
      },
      error: (err) => {
        this.loading = false;
        this.toastr.error(err.error.message, "Error");
      }
    })
  }
  updateBio(form: FormGroup): void {
    this.loading = true;
    this.http.post(serverURL + "update-bio", form.getRawValue(), { withCredentials: true }).subscribe({
      next: (v: any) => {
        if(v.message === "success") {
          if(this.userData != undefined) {
            this.userData.bio = form.getRawValue().bio;
          }
          this.router.navigate(["/dashboard"]);
          this.loading = false;
          this.toastr.success("Bio updated successfully", "Success");
        }
      },
      error: (err) => {
        this.toastr.error(err.error.message, "Error");
      }
    })
  }
  gotSocialCopy(): Social[] {
    return this.userData!.socials.map(s => Object.assign({}, s));
  }
  updateSocials(socials: Social[]): void {
    this.loading = true;
    this.http.post(serverURL + "update-socials", socials, { withCredentials: true }).subscribe({
      next: (v: any) => {
        if(v.message === "success") {
          if(this.userData != undefined) {
            this.userData.socials = socials;
          }
          this.router.navigate(["/dashboard"]);
          this.loading = false;
          this.toastr.success("Socials updated successfully", "Success");
        }
      },
      error: (err) => {
        this.loading = false;
        this.toastr.error(err.error.message, "Error");
      }
    })
  }
  updateProfilePicture(form: FormData): void {
    this.loading = true;
    this.http.post(serverURL + "update-picture", form, { withCredentials: true }).subscribe({
      next: (v: any) => {
        this.fetchProfilePicture();
        this.toastr.success(v.message, "Success");
       this.loading = false;
      },
      error: (err) => {
        this.loading = false;
        this.toastr.error(err.error.message, "Error");
      }
    })
  }
}
