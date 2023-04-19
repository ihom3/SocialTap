import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, Input } from '@angular/core';
import { serverURL } from '../user-service.service';
import { ToastrService } from 'ngx-toastr';
import { fetchIcon, fetchURL } from '../helpers';
@Component({
  selector: 'app-social-page',
  templateUrl: './social-page.component.html',
  styleUrls: ['./social-page.component.sass']
})
export class SocialPageComponent {
  @Input() userData: any;
  profileImg: string | ArrayBuffer | null = null;
  loading: boolean = true;
  reader: FileReader = new FileReader();
  constructor(private http: HttpClient, private toastr: ToastrService) {}
  ngOnInit(): void {
    this.fetchUerImage();
  }
  fetchIcon(name: string) {
    return fetchIcon(name);
  }
  fetchUerImage(): void {
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
        this.toastr.error(err.error.message, "Error");
        this.loading = false;
      }
    })
  }
  redirect(social: string, url: string) {
    window.location.href = fetchURL(social,url);
  }
}
