import {HttpInterceptorFn} from '@angular/common/http';
import {environment} from '../../environments/environment';

export const authInterceptor: HttpInterceptorFn = (request, next) => {
  const token = localStorage.getItem('token') ?? '';

  if (
    token &&
    token != '' &&
    !request.url.includes(`${environment.registerUrl}`) &&
    !request.url.includes(`${environment.imageUrl}`)
  )
    request = request.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`,
      },
    });

  return next(request);
};
