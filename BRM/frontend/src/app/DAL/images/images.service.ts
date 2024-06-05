import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../../environments/environment';
import {Observable} from 'rxjs';
import {IdResponse} from './model/IdResponse';

@Injectable({
  providedIn: 'root',
})
export class ImagesService {
  constructor(private http: HttpClient) {
  }

  saveImage(image: FormData): Observable<IdResponse> {
    return this.http.post<IdResponse>(`${environment.imageUrl}/images`, image);
  }

  loadImage(id: string): Observable<Blob> {
    return this.http.get<Blob>(`${environment.imageUrl}/images/${id}`);
  }
}
