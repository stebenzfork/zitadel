import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, switchMap, throwError } from 'rxjs';
import { ExhaustedService } from '../exhausted.service';

/**
 * ExhaustedHttpInterceptor shows the exhausted dialog before sending the request if the exhausted cookie is there.
 * Also, it shows the exhausted dialog after receiving an HTTP response status 429.
 */
@Injectable()
export class ExhaustedHttpInterceptor implements HttpInterceptor {
  constructor(private exhaustedSvc: ExhaustedService) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return this.exhaustedSvc.checkCookie().pipe(
      switchMap(() => next.handle(req)),
      catchError((error: HttpErrorResponse) => {
        if (error.status === 429) {
          return this.exhaustedSvc.showExhaustedDialog().pipe(switchMap(() => throwError(() => error)));
        }
        return throwError(() => error);
      }),
    );
  }
}
